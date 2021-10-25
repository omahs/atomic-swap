// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swapOnChain

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SwapOnChainMetaData contains all meta data concerning the SwapOnChain contract.
var SwapOnChainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_pubKeyClaim\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_pubKeyRefund\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"Claimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"p\",\"type\":\"bytes32\"}],\"name\":\"Constructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"b\",\"type\":\"bool\"}],\"name\":\"IsReady\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_s\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubKeyClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubKeyRefund\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_s\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"set_ready\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout_0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040526000600160006101000a81548160ff02191690831515021790555060405162001c7738038062001c77833981810160405281019062000045919062000190565b3373ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508160c081815250508060e0818152505062015180426200009a919062000210565b6101008181525050604051620000b09062000142565b604051809103906000f080158015620000cd573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250507f1ba2159dcdf5aa313440d6540b9acb108e5f7907737e884db04579a584275fbb816040516200013291906200027e565b60405180910390a150506200029b565b610df78062000e8083390190565b600080fd5b6000819050919050565b6200016a8162000155565b81146200017657600080fd5b50565b6000815190506200018a816200015f565b92915050565b60008060408385031215620001aa57620001a962000150565b5b6000620001ba8582860162000179565b9250506020620001cd8582860162000179565b9150509250929050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200021d82620001d7565b91506200022a83620001d7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620002625762000261620001e1565b5b828201905092915050565b620002788162000155565b82525050565b60006020820190506200029560008301846200026d565b92915050565b60805160a05160c05160e05161010051610b766200030a600039600081816101c80152818161032a0152818161040e01526104c501526000818161013e015261022f015260008181610392015261043201526000818161028c015261046e0152600061055e0152610b766000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806345bb8e091161005b57806345bb8e09146100d85780634ded8d52146100f6578063736290f81461011457806374d7c138146101325761007d565b806303f7e24614610082578063278ecde1146100a0578063379607f5146100bc575b600080fd5b61008a61013c565b6040516100979190610684565b60405180910390f35b6100ba60048036038101906100b591906106da565b610160565b005b6100d660048036038101906100d191906106da565b6102c3565b005b6100e0610406565b6040516100ed9190610716565b60405180910390f35b6100fe61040c565b60405161010b9190610716565b60405180910390f35b61011c610430565b6040516101299190610684565b60405180910390f35b61013a610454565b005b7f000000000000000000000000000000000000000000000000000000000000000081565b60011515600160009054906101000a900460ff16151514156101c6576000544210156101c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101b8906107b4565b60405180910390fd5b610229565b7f00000000000000000000000000000000000000000000000000000000000000004210610228576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161021f90610820565b60405180910390fd5b5b610253817f0000000000000000000000000000000000000000000000000000000000000000610559565b7f3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c816040516102829190610716565b60405180910390a17f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16ff5b60011515600160009054906101000a900460ff1615151415610328576000544210610323576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031a9061088c565b60405180910390fd5b61038c565b7f000000000000000000000000000000000000000000000000000000000000000042101561038b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103829061091e565b60405180910390fd5b5b6103b6817f0000000000000000000000000000000000000000000000000000000000000000610559565b7f7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8816040516103e59190610716565b60405180910390a13373ffffffffffffffffffffffffffffffffffffffff16ff5b60005481565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600160009054906101000a900460ff161580156104bc57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80156104e757507f000000000000000000000000000000000000000000000000000000000000000042105b6104f057600080fd5b60018060006101000a81548160ff0219169083151502179055506201518042610519919061096d565b6000819055507f2724cf6c3ad6a3399ad72482e4013d0171794f3ef4c462b7e24790c658cb3cd4600160405161054f91906109de565b60405180910390a1565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663c4f4912b856040518263ffffffff1660e01b81526004016105b59190610716565b604080518083038186803b1580156105cc57600080fd5b505afa1580156105e0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106049190610a0e565b91509150600060ff6002846106199190610a7d565b901b82179050838160001b14610664576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065b90610b20565b60405180910390fd5b5050505050565b6000819050919050565b61067e8161066b565b82525050565b60006020820190506106996000830184610675565b92915050565b600080fd5b6000819050919050565b6106b7816106a4565b81146106c257600080fd5b50565b6000813590506106d4816106ae565b92915050565b6000602082840312156106f0576106ef61069f565b5b60006106fe848285016106c5565b91505092915050565b610710816106a4565b82525050565b600060208201905061072b6000830184610707565b92915050565b600082825260208201905092915050565b7f4974277320426f622773207475726e206e6f772c20706c65617365207761697460008201527f2100000000000000000000000000000000000000000000000000000000000000602082015250565b600061079e602183610731565b91506107a982610742565b604082019050919050565b600060208201905081810360008301526107cd81610791565b9050919050565b7f4d697373656420796f7572206368616e63652100000000000000000000000000600082015250565b600061080a601383610731565b9150610815826107d4565b602082019050919050565b60006020820190508181036000830152610839816107fd565b9050919050565b7f546f6f206c61746520746f20636c61696d210000000000000000000000000000600082015250565b6000610876601283610731565b915061088182610840565b602082019050919050565b600060208201905081810360008301526108a581610869565b9050919050565b7f2769735265616479203d3d2066616c7365272063616e6e6f7420636c61696d2060008201527f7965742100000000000000000000000000000000000000000000000000000000602082015250565b6000610908602483610731565b9150610913826108ac565b604082019050919050565b60006020820190508181036000830152610937816108fb565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610978826106a4565b9150610983836106a4565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156109b8576109b761093e565b5b828201905092915050565b60008115159050919050565b6109d8816109c3565b82525050565b60006020820190506109f360008301846109cf565b92915050565b600081519050610a08816106ae565b92915050565b60008060408385031215610a2557610a2461069f565b5b6000610a33858286016109f9565b9250506020610a44858286016109f9565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000610a88826106a4565b9150610a93836106a4565b925082610aa357610aa2610a4e565b5b828206905092915050565b7f70726f76696465642073656372657420646f6573206e6f74206d61746368207460008201527f6865206578706563746564207075624b65790000000000000000000000000000602082015250565b6000610b0a603283610731565b9150610b1582610aae565b604082019050919050565b60006020820190508181036000830152610b3981610afd565b905091905056fea26469706673582212200094259fccf6352712768e50e970ba52232a8591f9a94c62135d10808705691364736f6c63430008090033608060405234801561001057600080fd5b50610dd7806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c4f4912b14610030575b600080fd5b61004a60048036038101906100459190610caa565b610061565b604051610058929190610ce6565b60405180910390f35b60008061006c610c09565b610074610c09565b7f216936d3cd6e53fec0a4e231fdd6dc5c692cc7609525a7b2c9562d608f25d51a8260000181815250507f666666666666666666666666666666666666666666666666666666666666665882602001818152505060018260400181815250506000816000018181525050600181602001818152505060018160400181815250505b600085111561012d57600180861614156101165761011381836101d2565b90505b600185901c945061012682610701565b91506100f5565b600061013c8260400151610b6b565b90507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061016d5761016c610d0f565b5b818360000151098260000181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806101ac576101ab610d0f565b5b818360200151098260200181815250508160000151826020015194509450505050915091565b6101da610c09565b6101e2610c2a565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061021157610210610d0f565b5b83604001518560400151098160000181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061025457610253610d0f565b5b81600001518260000151098160200181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061029757610296610d0f565b5b83600001518560000151098160400181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806102da576102d9610d0f565b5b83602001518560200151098160600181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061031d5761031c610d0f565b5b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061034c5761034b610d0f565b5b82606001518360400151097f52036cee2b6ffe738cc740797779e89800700a4d4141d8ab75eb4dca135978a3098160800181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806103b1576103b0610d0f565b5b81608001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed6103e19190610d6d565b8260200151088160a00181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061041f5761041e610d0f565b5b81608001518260200151088160c00181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061046257610461610d0f565b5b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061049157610490610d0f565b5b82606001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed6104c19190610d6d565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806104f0576104ef610d0f565b5b84604001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed6105209190610d6d565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061054f5761054e610d0f565b5b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061057e5761057d610d0f565b5b89602001518a60000151087f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806105b8576105b7610d0f565b5b8b602001518c60000151080908087f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806105f5576105f4610d0f565b5b8360a00151846000015109098260000181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061063957610638610d0f565b5b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061066857610667610d0f565b5b82604001518360600151087f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806106a2576106a1610d0f565b5b8360c00151846000015109098260200181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806106e6576106e5610d0f565b5b8160c001518260a00151098260400181815250505092915050565b610709610c09565b610711610c2a565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806107405761073f610d0f565b5b83602001518460000151088160000181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061078357610782610d0f565b5b81600001518260000151098160200181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806107c6576107c5610d0f565b5b83600001518460000151098160400181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061080957610808610d0f565b5b836020015184602001510981606001818152505080604001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed61084d9190610d6d565b8160800181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061088557610884610d0f565b5b81606001518260800151088160a00181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806108c8576108c7610d0f565b5b83604001518460400151098160e00181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061090b5761090a610d0f565b5b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed8061093a57610939610d0f565b5b8260e001516002097f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed61096d9190610d6d565b8260a00151088160c00181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806109ab576109aa610d0f565b5b8160c001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed806109df576109de610d0f565b5b83606001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed610a0f9190610d6d565b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed80610a3e57610a3d610d0f565b5b85604001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed610a6e9190610d6d565b86602001510808098260000181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed80610aae57610aad610d0f565b5b7f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed80610add57610adc610d0f565b5b82606001517f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed610b0d9190610d6d565b8360800151088260a00151098260200181815250507f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed80610b5157610b50610d0f565b5b8160c001518260a001510982604001818152505050919050565b60008060027f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed610b9b9190610d6d565b905060007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffed905060405160208152602080820152602060408201528460608201528260808201528160a082015260208160c0836005600019fa610bfd57600080fd5b80519350505050919050565b60405180606001604052806000815260200160008152602001600081525090565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b600080fd5b6000819050919050565b610c8781610c74565b8114610c9257600080fd5b50565b600081359050610ca481610c7e565b92915050565b600060208284031215610cc057610cbf610c6f565b5b6000610cce84828501610c95565b91505092915050565b610ce081610c74565b82525050565b6000604082019050610cfb6000830185610cd7565b610d086020830184610cd7565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d7882610c74565b9150610d8383610c74565b925082821015610d9657610d95610d3e565b5b82820390509291505056fea264697066735822122026763aba14e28dd6f9cbc211a6db8bcee0af925f8e7aef10e72e4164c3feaf8c64736f6c63430008090033",
}

// SwapOnChainABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapOnChainMetaData.ABI instead.
var SwapOnChainABI = SwapOnChainMetaData.ABI

// SwapOnChainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapOnChainMetaData.Bin instead.
var SwapOnChainBin = SwapOnChainMetaData.Bin

// DeploySwapOnChain deploys a new Ethereum contract, binding an instance of SwapOnChain to it.
func DeploySwapOnChain(auth *bind.TransactOpts, backend bind.ContractBackend, _pubKeyClaim [32]byte, _pubKeyRefund [32]byte) (common.Address, *types.Transaction, *SwapOnChain, error) {
	parsed, err := SwapOnChainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapOnChainBin), backend, _pubKeyClaim, _pubKeyRefund)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapOnChain{SwapOnChainCaller: SwapOnChainCaller{contract: contract}, SwapOnChainTransactor: SwapOnChainTransactor{contract: contract}, SwapOnChainFilterer: SwapOnChainFilterer{contract: contract}}, nil
}

// SwapOnChain is an auto generated Go binding around an Ethereum contract.
type SwapOnChain struct {
	SwapOnChainCaller     // Read-only binding to the contract
	SwapOnChainTransactor // Write-only binding to the contract
	SwapOnChainFilterer   // Log filterer for contract events
}

// SwapOnChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapOnChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapOnChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapOnChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapOnChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapOnChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapOnChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapOnChainSession struct {
	Contract     *SwapOnChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapOnChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapOnChainCallerSession struct {
	Contract *SwapOnChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwapOnChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapOnChainTransactorSession struct {
	Contract     *SwapOnChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwapOnChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapOnChainRaw struct {
	Contract *SwapOnChain // Generic contract binding to access the raw methods on
}

// SwapOnChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapOnChainCallerRaw struct {
	Contract *SwapOnChainCaller // Generic read-only contract binding to access the raw methods on
}

// SwapOnChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapOnChainTransactorRaw struct {
	Contract *SwapOnChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapOnChain creates a new instance of SwapOnChain, bound to a specific deployed contract.
func NewSwapOnChain(address common.Address, backend bind.ContractBackend) (*SwapOnChain, error) {
	contract, err := bindSwapOnChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapOnChain{SwapOnChainCaller: SwapOnChainCaller{contract: contract}, SwapOnChainTransactor: SwapOnChainTransactor{contract: contract}, SwapOnChainFilterer: SwapOnChainFilterer{contract: contract}}, nil
}

// NewSwapOnChainCaller creates a new read-only instance of SwapOnChain, bound to a specific deployed contract.
func NewSwapOnChainCaller(address common.Address, caller bind.ContractCaller) (*SwapOnChainCaller, error) {
	contract, err := bindSwapOnChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapOnChainCaller{contract: contract}, nil
}

// NewSwapOnChainTransactor creates a new write-only instance of SwapOnChain, bound to a specific deployed contract.
func NewSwapOnChainTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapOnChainTransactor, error) {
	contract, err := bindSwapOnChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapOnChainTransactor{contract: contract}, nil
}

// NewSwapOnChainFilterer creates a new log filterer instance of SwapOnChain, bound to a specific deployed contract.
func NewSwapOnChainFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapOnChainFilterer, error) {
	contract, err := bindSwapOnChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapOnChainFilterer{contract: contract}, nil
}

// bindSwapOnChain binds a generic wrapper to an already deployed contract.
func bindSwapOnChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapOnChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapOnChain *SwapOnChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapOnChain.Contract.SwapOnChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapOnChain *SwapOnChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapOnChain.Contract.SwapOnChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapOnChain *SwapOnChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapOnChain.Contract.SwapOnChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapOnChain *SwapOnChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapOnChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapOnChain *SwapOnChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapOnChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapOnChain *SwapOnChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapOnChain.Contract.contract.Transact(opts, method, params...)
}

// PubKeyClaim is a free data retrieval call binding the contract method 0x736290f8.
//
// Solidity: function pubKeyClaim() view returns(bytes32)
func (_SwapOnChain *SwapOnChainCaller) PubKeyClaim(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapOnChain.contract.Call(opts, &out, "pubKeyClaim")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PubKeyClaim is a free data retrieval call binding the contract method 0x736290f8.
//
// Solidity: function pubKeyClaim() view returns(bytes32)
func (_SwapOnChain *SwapOnChainSession) PubKeyClaim() ([32]byte, error) {
	return _SwapOnChain.Contract.PubKeyClaim(&_SwapOnChain.CallOpts)
}

// PubKeyClaim is a free data retrieval call binding the contract method 0x736290f8.
//
// Solidity: function pubKeyClaim() view returns(bytes32)
func (_SwapOnChain *SwapOnChainCallerSession) PubKeyClaim() ([32]byte, error) {
	return _SwapOnChain.Contract.PubKeyClaim(&_SwapOnChain.CallOpts)
}

// PubKeyRefund is a free data retrieval call binding the contract method 0x03f7e246.
//
// Solidity: function pubKeyRefund() view returns(bytes32)
func (_SwapOnChain *SwapOnChainCaller) PubKeyRefund(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapOnChain.contract.Call(opts, &out, "pubKeyRefund")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PubKeyRefund is a free data retrieval call binding the contract method 0x03f7e246.
//
// Solidity: function pubKeyRefund() view returns(bytes32)
func (_SwapOnChain *SwapOnChainSession) PubKeyRefund() ([32]byte, error) {
	return _SwapOnChain.Contract.PubKeyRefund(&_SwapOnChain.CallOpts)
}

// PubKeyRefund is a free data retrieval call binding the contract method 0x03f7e246.
//
// Solidity: function pubKeyRefund() view returns(bytes32)
func (_SwapOnChain *SwapOnChainCallerSession) PubKeyRefund() ([32]byte, error) {
	return _SwapOnChain.Contract.PubKeyRefund(&_SwapOnChain.CallOpts)
}

// Timeout0 is a free data retrieval call binding the contract method 0x4ded8d52.
//
// Solidity: function timeout_0() view returns(uint256)
func (_SwapOnChain *SwapOnChainCaller) Timeout0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapOnChain.contract.Call(opts, &out, "timeout_0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timeout0 is a free data retrieval call binding the contract method 0x4ded8d52.
//
// Solidity: function timeout_0() view returns(uint256)
func (_SwapOnChain *SwapOnChainSession) Timeout0() (*big.Int, error) {
	return _SwapOnChain.Contract.Timeout0(&_SwapOnChain.CallOpts)
}

// Timeout0 is a free data retrieval call binding the contract method 0x4ded8d52.
//
// Solidity: function timeout_0() view returns(uint256)
func (_SwapOnChain *SwapOnChainCallerSession) Timeout0() (*big.Int, error) {
	return _SwapOnChain.Contract.Timeout0(&_SwapOnChain.CallOpts)
}

// Timeout1 is a free data retrieval call binding the contract method 0x45bb8e09.
//
// Solidity: function timeout_1() view returns(uint256)
func (_SwapOnChain *SwapOnChainCaller) Timeout1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapOnChain.contract.Call(opts, &out, "timeout_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timeout1 is a free data retrieval call binding the contract method 0x45bb8e09.
//
// Solidity: function timeout_1() view returns(uint256)
func (_SwapOnChain *SwapOnChainSession) Timeout1() (*big.Int, error) {
	return _SwapOnChain.Contract.Timeout1(&_SwapOnChain.CallOpts)
}

// Timeout1 is a free data retrieval call binding the contract method 0x45bb8e09.
//
// Solidity: function timeout_1() view returns(uint256)
func (_SwapOnChain *SwapOnChainCallerSession) Timeout1() (*big.Int, error) {
	return _SwapOnChain.Contract.Timeout1(&_SwapOnChain.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _s) returns()
func (_SwapOnChain *SwapOnChainTransactor) Claim(opts *bind.TransactOpts, _s *big.Int) (*types.Transaction, error) {
	return _SwapOnChain.contract.Transact(opts, "claim", _s)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _s) returns()
func (_SwapOnChain *SwapOnChainSession) Claim(_s *big.Int) (*types.Transaction, error) {
	return _SwapOnChain.Contract.Claim(&_SwapOnChain.TransactOpts, _s)
}

// Claim is a paid mutator transaction binding the contract method 0x379607f5.
//
// Solidity: function claim(uint256 _s) returns()
func (_SwapOnChain *SwapOnChainTransactorSession) Claim(_s *big.Int) (*types.Transaction, error) {
	return _SwapOnChain.Contract.Claim(&_SwapOnChain.TransactOpts, _s)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _s) returns()
func (_SwapOnChain *SwapOnChainTransactor) Refund(opts *bind.TransactOpts, _s *big.Int) (*types.Transaction, error) {
	return _SwapOnChain.contract.Transact(opts, "refund", _s)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _s) returns()
func (_SwapOnChain *SwapOnChainSession) Refund(_s *big.Int) (*types.Transaction, error) {
	return _SwapOnChain.Contract.Refund(&_SwapOnChain.TransactOpts, _s)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 _s) returns()
func (_SwapOnChain *SwapOnChainTransactorSession) Refund(_s *big.Int) (*types.Transaction, error) {
	return _SwapOnChain.Contract.Refund(&_SwapOnChain.TransactOpts, _s)
}

// SetReady is a paid mutator transaction binding the contract method 0x74d7c138.
//
// Solidity: function set_ready() returns()
func (_SwapOnChain *SwapOnChainTransactor) SetReady(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapOnChain.contract.Transact(opts, "set_ready")
}

// SetReady is a paid mutator transaction binding the contract method 0x74d7c138.
//
// Solidity: function set_ready() returns()
func (_SwapOnChain *SwapOnChainSession) SetReady() (*types.Transaction, error) {
	return _SwapOnChain.Contract.SetReady(&_SwapOnChain.TransactOpts)
}

// SetReady is a paid mutator transaction binding the contract method 0x74d7c138.
//
// Solidity: function set_ready() returns()
func (_SwapOnChain *SwapOnChainTransactorSession) SetReady() (*types.Transaction, error) {
	return _SwapOnChain.Contract.SetReady(&_SwapOnChain.TransactOpts)
}

// SwapOnChainClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the SwapOnChain contract.
type SwapOnChainClaimedIterator struct {
	Event *SwapOnChainClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapOnChainClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOnChainClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapOnChainClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapOnChainClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOnChainClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOnChainClaimed represents a Claimed event raised by the SwapOnChain contract.
type SwapOnChainClaimed struct {
	S   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 s)
func (_SwapOnChain *SwapOnChainFilterer) FilterClaimed(opts *bind.FilterOpts) (*SwapOnChainClaimedIterator, error) {

	logs, sub, err := _SwapOnChain.contract.FilterLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return &SwapOnChainClaimedIterator{contract: _SwapOnChain.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 s)
func (_SwapOnChain *SwapOnChainFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *SwapOnChainClaimed) (event.Subscription, error) {

	logs, sub, err := _SwapOnChain.contract.WatchLogs(opts, "Claimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOnChainClaimed)
				if err := _SwapOnChain.contract.UnpackLog(event, "Claimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaimed is a log parse operation binding the contract event 0x7a355715549cfe7c1cba26304350343fbddc4b4f72d3ce3e7c27117dd20b5cb8.
//
// Solidity: event Claimed(uint256 s)
func (_SwapOnChain *SwapOnChainFilterer) ParseClaimed(log types.Log) (*SwapOnChainClaimed, error) {
	event := new(SwapOnChainClaimed)
	if err := _SwapOnChain.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapOnChainConstructedIterator is returned from FilterConstructed and is used to iterate over the raw logs and unpacked data for Constructed events raised by the SwapOnChain contract.
type SwapOnChainConstructedIterator struct {
	Event *SwapOnChainConstructed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapOnChainConstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOnChainConstructed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapOnChainConstructed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapOnChainConstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOnChainConstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOnChainConstructed represents a Constructed event raised by the SwapOnChain contract.
type SwapOnChainConstructed struct {
	P   [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterConstructed is a free log retrieval operation binding the contract event 0x1ba2159dcdf5aa313440d6540b9acb108e5f7907737e884db04579a584275fbb.
//
// Solidity: event Constructed(bytes32 p)
func (_SwapOnChain *SwapOnChainFilterer) FilterConstructed(opts *bind.FilterOpts) (*SwapOnChainConstructedIterator, error) {

	logs, sub, err := _SwapOnChain.contract.FilterLogs(opts, "Constructed")
	if err != nil {
		return nil, err
	}
	return &SwapOnChainConstructedIterator{contract: _SwapOnChain.contract, event: "Constructed", logs: logs, sub: sub}, nil
}

// WatchConstructed is a free log subscription operation binding the contract event 0x1ba2159dcdf5aa313440d6540b9acb108e5f7907737e884db04579a584275fbb.
//
// Solidity: event Constructed(bytes32 p)
func (_SwapOnChain *SwapOnChainFilterer) WatchConstructed(opts *bind.WatchOpts, sink chan<- *SwapOnChainConstructed) (event.Subscription, error) {

	logs, sub, err := _SwapOnChain.contract.WatchLogs(opts, "Constructed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOnChainConstructed)
				if err := _SwapOnChain.contract.UnpackLog(event, "Constructed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConstructed is a log parse operation binding the contract event 0x1ba2159dcdf5aa313440d6540b9acb108e5f7907737e884db04579a584275fbb.
//
// Solidity: event Constructed(bytes32 p)
func (_SwapOnChain *SwapOnChainFilterer) ParseConstructed(log types.Log) (*SwapOnChainConstructed, error) {
	event := new(SwapOnChainConstructed)
	if err := _SwapOnChain.contract.UnpackLog(event, "Constructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapOnChainIsReadyIterator is returned from FilterIsReady and is used to iterate over the raw logs and unpacked data for IsReady events raised by the SwapOnChain contract.
type SwapOnChainIsReadyIterator struct {
	Event *SwapOnChainIsReady // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapOnChainIsReadyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOnChainIsReady)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapOnChainIsReady)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapOnChainIsReadyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOnChainIsReadyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOnChainIsReady represents a IsReady event raised by the SwapOnChain contract.
type SwapOnChainIsReady struct {
	B   bool
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIsReady is a free log retrieval operation binding the contract event 0x2724cf6c3ad6a3399ad72482e4013d0171794f3ef4c462b7e24790c658cb3cd4.
//
// Solidity: event IsReady(bool b)
func (_SwapOnChain *SwapOnChainFilterer) FilterIsReady(opts *bind.FilterOpts) (*SwapOnChainIsReadyIterator, error) {

	logs, sub, err := _SwapOnChain.contract.FilterLogs(opts, "IsReady")
	if err != nil {
		return nil, err
	}
	return &SwapOnChainIsReadyIterator{contract: _SwapOnChain.contract, event: "IsReady", logs: logs, sub: sub}, nil
}

// WatchIsReady is a free log subscription operation binding the contract event 0x2724cf6c3ad6a3399ad72482e4013d0171794f3ef4c462b7e24790c658cb3cd4.
//
// Solidity: event IsReady(bool b)
func (_SwapOnChain *SwapOnChainFilterer) WatchIsReady(opts *bind.WatchOpts, sink chan<- *SwapOnChainIsReady) (event.Subscription, error) {

	logs, sub, err := _SwapOnChain.contract.WatchLogs(opts, "IsReady")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOnChainIsReady)
				if err := _SwapOnChain.contract.UnpackLog(event, "IsReady", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIsReady is a log parse operation binding the contract event 0x2724cf6c3ad6a3399ad72482e4013d0171794f3ef4c462b7e24790c658cb3cd4.
//
// Solidity: event IsReady(bool b)
func (_SwapOnChain *SwapOnChainFilterer) ParseIsReady(log types.Log) (*SwapOnChainIsReady, error) {
	event := new(SwapOnChainIsReady)
	if err := _SwapOnChain.contract.UnpackLog(event, "IsReady", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapOnChainRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the SwapOnChain contract.
type SwapOnChainRefundedIterator struct {
	Event *SwapOnChainRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapOnChainRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOnChainRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapOnChainRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapOnChainRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOnChainRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOnChainRefunded represents a Refunded event raised by the SwapOnChain contract.
type SwapOnChainRefunded struct {
	S   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 s)
func (_SwapOnChain *SwapOnChainFilterer) FilterRefunded(opts *bind.FilterOpts) (*SwapOnChainRefundedIterator, error) {

	logs, sub, err := _SwapOnChain.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &SwapOnChainRefundedIterator{contract: _SwapOnChain.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 s)
func (_SwapOnChain *SwapOnChainFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *SwapOnChainRefunded) (event.Subscription, error) {

	logs, sub, err := _SwapOnChain.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOnChainRefunded)
				if err := _SwapOnChain.contract.UnpackLog(event, "Refunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRefunded is a log parse operation binding the contract event 0x3d2a04f53164bedf9a8a46353305d6b2d2261410406df3b41f99ce6489dc003c.
//
// Solidity: event Refunded(uint256 s)
func (_SwapOnChain *SwapOnChainFilterer) ParseRefunded(log types.Log) (*SwapOnChainRefunded, error) {
	event := new(SwapOnChainRefunded)
	if err := _SwapOnChain.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}