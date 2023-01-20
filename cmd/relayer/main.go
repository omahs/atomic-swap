package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/apd/v3"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"

	p2pnet "github.com/athanorlabs/go-p2p-net"
	rcommon "github.com/athanorlabs/go-relayer/common"
	rcontracts "github.com/athanorlabs/go-relayer/impls/gsnforwarder"
	net "github.com/athanorlabs/go-relayer/net"
	"github.com/athanorlabs/go-relayer/relayer"
	rrpc "github.com/athanorlabs/go-relayer/rpc"

	"github.com/athanorlabs/atomic-swap/cliutil"
	"github.com/athanorlabs/atomic-swap/coins"
	contracts "github.com/athanorlabs/atomic-swap/ethereum"
	swapnet "github.com/athanorlabs/atomic-swap/net"

	logging "github.com/ipfs/go-log"
)

const (
	flagEndpoint          = "endpoint"
	flagForwarderAddress  = "forwarder-address"
	flagKey               = "key"
	flagRPC               = "rpc"
	flagRPCPort           = "rpc-port"
	flagDeploy            = "deploy"
	flagLog               = "log-level"
	flagWithNetwork       = "with-network" // TODO: do we need this?
	flagLibp2pKey         = "libp2p-key"
	flagLibp2pPort        = "libp2p-port"
	flagBootnodes         = "bootnodes"
	flagRelayerCommission = "relayer-commission"

	defaultLibp2pPort = 10900
)

var (
	log = logging.Logger("main")

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:  flagEndpoint,
			Value: "http://localhost:8545",
			Usage: "Ethereum RPC endpoint",
		},
		&cli.StringFlag{
			Name:  flagKey,
			Value: "eth.key",
			Usage: "Path to file containing Ethereum private key",
		},
		&cli.UintFlag{
			Name:  flagRPCPort,
			Value: 7799,
			Usage: "Relayer RPC server port",
		},
		&cli.BoolFlag{
			Name:  flagRPC,
			Value: false,
			Usage: "Run the relayer RPC server. Defaults false",
		},
		&cli.BoolFlag{
			Name:  flagDeploy,
			Usage: "Deploy an instance of the forwarder contract",
		},
		&cli.StringFlag{
			Name:  flagLog,
			Value: "info",
			Usage: "Set log level: one of [error|warn|info|debug]",
		},
		&cli.BoolFlag{
			Name:  flagWithNetwork,
			Value: true,
			Usage: "Run the relayer with p2p network capabilities",
		},
		&cli.StringFlag{
			Name:  flagLibp2pKey,
			Usage: "libp2p private key",
			Value: fmt.Sprintf("%s/node.key", os.TempDir()),
		},
		&cli.UintFlag{
			Name:  flagLibp2pPort,
			Usage: "libp2p port to listen on",
			Value: defaultLibp2pPort,
		},
		&cli.StringSliceFlag{
			Name:    flagBootnodes,
			Aliases: []string{"bn"},
			Usage:   "libp2p bootnode, comma separated if passing multiple to a single flag",
			EnvVars: []string{"SWAPD_BOOTNODES"},
		},
		&cli.StringFlag{
			Name: flagRelayerCommission,
			Usage: "Minimum commission percentage (of the swap value) to receive:" +
				" eg. --relayer-commission=0.01 for 1% commission",
		},
	}

	errInvalidAddress       = errors.New("invalid forwarder address")
	errNoEthereumPrivateKey = errors.New("must provide ethereum private key with --key")
)

func main() {
	app := &cli.App{
		Name:                 "relayer",
		Usage:                "Ethereum transaction relayer",
		Version:              getVersion(),
		Flags:                flags,
		Action:               run,
		EnableBashCompletion: true,
		Suggest:              true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func setLogLevels(c *cli.Context) error {
	const (
		levelError = "error"
		levelWarn  = "warn"
		levelInfo  = "info"
		levelDebug = "debug"
	)

	level := c.String(flagLog)
	switch level {
	case levelError, levelWarn, levelInfo, levelDebug:
	default:
		return fmt.Errorf("invalid log level %q", level)
	}

	_ = logging.SetLogLevel("main", level)
	_ = logging.SetLogLevel("relayer", level)
	_ = logging.SetLogLevel("rpc", level)
	return nil
}

func run(c *cli.Context) error {
	err := setLogLevels(c)
	if err != nil {
		return err
	}

	port := uint16(c.Uint(flagRPCPort))
	endpoint := c.String(flagEndpoint)
	ec, err := ethclient.Dial(endpoint)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(c.Context)
	defer cancel()

	chainID, err := ec.ChainID(ctx)
	if err != nil {
		return err
	}

	log.Infof("starting relayer with ethereum endpoint %s and chain ID %s", endpoint, chainID)

	key, err := getPrivateKey(c.String(flagKey))
	if err != nil {
		return err
	}

	contractAddr := c.String(flagForwarderAddress)
	deploy := c.Bool(flagDeploy)
	contractSet := contractAddr != ""
	if deploy && contractSet {
		return fmt.Errorf("flags --%s and --%s are mutually exclusive", flagDeploy, flagForwarderAddress)
	}
	if !deploy && !contractSet {
		return fmt.Errorf("either --%s or --%s is required", flagDeploy, flagForwarderAddress)
	}

	forwarder, err := deployOrGetForwarder(
		contractAddr,
		ec,
		key,
		chainID,
	)
	if err != nil {
		return err
	}

	relayerCommission, err := cliutil.ReadUnsignedDecimalFlag(c, flagRelayerCommission)
	if err != nil {
		return err
	}

	// TODO: do we need to restrict potential commission values? eg. 1%, 1.25%, 1.5%, etc
	v := &validator{
		ctx:               ctx,
		ec:                ec,
		relayerCommission: relayerCommission,
	}

	// TODO: the forwarder contract is fixed here; thus it needs to be the same
	// as what's hardcoded in the swap contract addr for that network.
	rcfg := &relayer.Config{
		Ctx:                     ctx,
		EthClient:               ec,
		Forwarder:               rcontracts.NewIForwarderWrapped(forwarder),
		Key:                     key,
		ValidateTransactionFunc: v.validateTransactionFunc,
	}

	r, err := relayer.NewRelayer(rcfg)
	if err != nil {
		return err
	}

	if c.Bool(flagWithNetwork) {
		h, err := setupNetwork(ctx, c, ec, r) //nolint:govet
		if err != nil {
			return err
		}

		defer func() {
			_ = h.Stop()
		}()
	}

	if c.Bool(flagRPC) {
		go signalHandler(ctx, cancel)
		rpcCfg := &rrpc.Config{
			Ctx:     ctx,
			Address: fmt.Sprintf("127.0.0.1:%d", port),
			Relayer: r,
		}

		server, err := rrpc.NewServer(rpcCfg) //nolint:govet
		if err != nil {
			return err
		}

		err = server.Start()
		if errors.Is(err, context.Canceled) || errors.Is(err, http.ErrServerClosed) {
			return nil
		}
	} else {
		signalHandler(ctx, cancel)
	}

	return err
}

type validator struct {
	ctx               context.Context
	ec                *ethclient.Client
	relayerCommission *apd.Decimal
}

func (v *validator) validateTransactionFunc(req *rcommon.SubmitTransactionRequest) error {
	// validate that:
	// 1. the `to` address is a swap contract;
	// 2. the function being called is `claimRelayer`;
	// 3. the fee passed to `claimRelayer` is equal to or greater
	// than our desired commission percentage.

	_, err := contracts.CheckSwapFactoryContractCode(
		v.ctx, v.ec, req.To,
	)
	if err != nil {
		return err
	}

	// hardcoded, from swap_factory.go bindings
	claimRelayerSig := ethcommon.Hex2Bytes("0x73e4771c")
	if !bytes.Equal(claimRelayerSig, req.Data[:4]) {
		return errors.New("call must be to claimRelayer()")
	}

	err = validateRelayerFee(req.Data[4:], v.relayerCommission)
	if err != nil {
		return err
	}

	return nil
}

func validateRelayerFee(data []byte, minFeePercentage *apd.Decimal) error {
	uint256Ty, err := abi.NewType("uint256", "", nil)
	if err != nil {
		return fmt.Errorf("failed to create uint256 type: %w", err)
	}

	bytes32Ty, err := abi.NewType("bytes32", "", nil)
	if err != nil {
		return fmt.Errorf("failed to create bytes32 type: %w", err)
	}

	addressTy, err := abi.NewType("address", "", nil)
	if err != nil {
		return fmt.Errorf("failed to create address type: %w", err)
	}

	arguments := abi.Arguments{
		// Swap
		{
			Type: addressTy,
		},
		{
			Type: addressTy,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: bytes32Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: addressTy,
		},
		{
			Type: uint256Ty, // value
		},
		{
			Type: uint256Ty,
		},
		// _s
		{
			Type: bytes32Ty,
		},
		// _fee
		{
			Type: uint256Ty,
		},
	}
	args, err := arguments.Unpack(data)
	if err != nil {
		return err
	}

	value, ok := args[7].(*big.Int)
	if !ok {
		// this shouldn't happen afaik
		return errors.New("value argument was not marshalled into a *big.Int")
	}

	fee, ok := args[10].(*big.Int)
	if !ok {
		// this shouldn't happen afaik
		return errors.New("fee argument was not marshalled into a *big.Int")
	}

	valueD := apd.NewWithBigInt(
		new(apd.BigInt).SetMathBigInt(value), // swap value, in wei
		0,
	)
	feeD := apd.NewWithBigInt(
		new(apd.BigInt).SetMathBigInt(fee), // fee, in wei
		0,
	)

	percentage := new(apd.Decimal)
	_, err = coins.DecimalCtx().Quo(percentage, feeD, valueD)
	if err != nil {
		return err
	}

	if percentage.Cmp(minFeePercentage) == -1 {
		return fmt.Errorf("fee too low: percentage is %s, expected minimum %s",
			percentage,
			minFeePercentage,
		)
	}

	return nil
}

func setupNetwork(
	ctx context.Context,
	c *cli.Context,
	ec *ethclient.Client,
	r *relayer.Relayer,
) (*net.Host, error) {
	chainID, err := ec.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	var bootnodes []string
	if c.IsSet(flagBootnodes) {
		bootnodes = expandBootnodes(c.StringSlice(flagBootnodes))
	}

	listenIP := "0.0.0.0"
	netCfg := &p2pnet.Config{
		Ctx:        ctx,
		DataDir:    os.TempDir(),
		Port:       uint16(c.Uint(flagLibp2pPort)),
		KeyFile:    c.String(flagLibp2pKey),
		Bootnodes:  bootnodes,
		ProtocolID: fmt.Sprintf("/%s/%s/%d", swapnet.ProtocolID, net.ProtocolID, chainID.Int64()),
		ListenIP:   listenIP,
	}

	cfg := &net.Config{
		Context:              ctx,
		P2pConfig:            netCfg,
		TransactionSubmitter: r,
		IsRelayer:            true,
	}

	h, err := net.NewHost(cfg)
	if err != nil {
		return nil, err
	}

	err = h.Start()
	if err != nil {
		return nil, err
	}

	return h, nil
}

func deployOrGetForwarder(
	addressString string,
	ec *ethclient.Client,
	key *rcommon.Key,
	chainID *big.Int,
) (*rcontracts.IForwarder, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey(), chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to make transactor: %w", err)
	}

	if addressString == "" {
		address, tx, _, err := rcontracts.DeployForwarder(txOpts, ec)
		if err != nil {
			return nil, err
		}

		_, err = bind.WaitMined(context.Background(), ec, tx)
		if err != nil {
			return nil, err
		}

		log.Infof("deployed Forwarder.sol to %s", address)
		return rcontracts.NewIForwarder(address, ec)
	}

	ok := ethcommon.IsHexAddress(addressString)
	if !ok {
		return nil, errInvalidAddress
	}

	log.Infof("loaded Forwarder.sol at %s", ethcommon.HexToAddress(addressString))
	return rcontracts.NewIForwarder(ethcommon.HexToAddress(addressString), ec)
}

func getPrivateKey(keyFile string) (*rcommon.Key, error) {
	if keyFile != "" {
		fileData, err := os.ReadFile(filepath.Clean(keyFile))
		if err != nil {
			return nil, fmt.Errorf("failed to read private key file: %w", err)
		}
		keyHex := strings.TrimSpace(string(fileData))
		return rcommon.NewKeyFromPrivateKeyString(keyHex)
	}
	return nil, errNoEthereumPrivateKey
}

// expandBootnodes expands the boot nodes passed on the command line that
// can be specified individually with multiple flags, but can also contain
// multiple boot nodes passed to single flag separated by commas.
func expandBootnodes(nodesCLI []string) []string {
	var nodes []string // nodes from all flag values combined
	for _, flagVal := range nodesCLI {
		splitNodes := strings.Split(flagVal, ",")
		for _, n := range splitNodes {
			n = strings.TrimSpace(n)
			// Handle the empty string to not use default bootnodes. Doing it here after
			// the split has the arguably positive side effect of skipping empty entries.
			if len(n) > 0 {
				nodes = append(nodes, strings.TrimSpace(n))
			}
		}
	}
	return nodes
}