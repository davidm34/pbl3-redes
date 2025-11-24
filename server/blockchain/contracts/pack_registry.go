// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// PackRegistryMetaData contains all meta data concerning the PackRegistry contract.
var PackRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialStock\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cardId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"ownerId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CardAssigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cardId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CardTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"matchId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"winner\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"loser\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"MatchRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"StockUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_playerId\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"_cardIds\",\"type\":\"string[]\"}],\"name\":\"assignCards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"cardOwner\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decrementStock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatchCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_ownerId\",\"type\":\"string\"}],\"name\":\"getUserCards\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"matches\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"matchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"winnerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"loserId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerCards\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_matchId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_winnerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_loserId\",\"type\":\"string\"}],\"name\":\"recordMatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPacks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fromId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_toId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cardId\",\"type\":\"string\"}],\"name\":\"transferCard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001b9738038062001b97833981810160405281019062000037919062000085565b8060008190555050620000b7565b600080fd5b6000819050919050565b6200005f816200004a565b81146200006b57600080fd5b50565b6000815190506200007f8162000054565b92915050565b6000602082840312156200009e576200009d62000045565b5b6000620000ae848285016200006e565b91505092915050565b611ad080620000c76000396000f3fe608060405234801561001057600080fd5b50600436106100a85760003560e01c80636d34550e116100715780636d34550e146101525780637281df121461016e578063888ad8981461018c5780638c4f7dae146101bc5780638f2ce142146101da578063e38ffcd81461020a576100a8565b8062375a44146100ad578063118c54ec146100b757806312636a62146100d3578063194b13b6146100ef5780634768d4ef1461011f575b600080fd5b6100b5610228565b005b6100d160048036038101906100cc9190610d78565b6102c0565b005b6100ed60048036038101906100e89190610d78565b610404565b005b61010960048036038101906101049190610e1f565b6104e0565b6040516101169190610fa9565b60405180910390f35b61013960048036038101906101349190611001565b6105d7565b6040516101499493929190611087565b60405180910390f35b61016c600480360381019061016791906111c7565b6107af565b005b610176610912565b604051610183919061123f565b60405180910390f35b6101a660048036038101906101a19190610e1f565b610918565b6040516101b3919061125a565b60405180910390f35b6101c46109ce565b6040516101d1919061123f565b60405180910390f35b6101f460048036038101906101ef919061127c565b6109db565b604051610201919061125a565b60405180910390f35b610212610aaa565b60405161021f919061123f565b60405180910390f35b600080541161026c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161026390611324565b60405180910390fd5b600160008082825461027e9190611373565b925050819055507f2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a06000546040516102b691906113f3565b60405180910390a1565b82805190602001206002826040516102d8919061145d565b90815260200160405180910390206040516102f39190611577565b60405180910390201461033b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610332906115da565b60405180910390fd5b8160028260405161034c919061145d565b9081526020016040518091039020908161036691906117a6565b506103718382610ab3565b600382604051610381919061145d565b9081526020016040518091039020819080600181540180825580915050600190039060005260206000200160009091909190915090816103c191906117a6565b507f6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d818484426040516103f79493929190611087565b60405180910390a1505050565b60016040518060800160405280858152602001848152602001838152602001428152509080600181540180825580915050600190039060005260206000209060040201600090919091909150600082015181600001908161046591906117a6565b50602082015181600101908161047b91906117a6565b50604082015181600201908161049191906117a6565b506060820151816003015550507fe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37838383426040516104d39493929190611087565b60405180910390a1505050565b60606003826040516104f2919061145d565b9081526020016040518091039020805480602002602001604051908101604052809291908181526020016000905b828210156105cc57838290600052602060002001805461053f906114a3565b80601f016020809104026020016040519081016040528092919081815260200182805461056b906114a3565b80156105b85780601f1061058d576101008083540402835291602001916105b8565b820191906000526020600020905b81548152906001019060200180831161059b57829003601f168201915b505050505081526020019060010190610520565b505050509050919050565b600181815481106105e757600080fd5b906000526020600020906004020160009150905080600001805461060a906114a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610636906114a3565b80156106835780601f1061065857610100808354040283529160200191610683565b820191906000526020600020905b81548152906001019060200180831161066657829003601f168201915b505050505090806001018054610698906114a3565b80601f01602080910402602001604051908101604052809291908181526020018280546106c4906114a3565b80156107115780601f106106e657610100808354040283529160200191610711565b820191906000526020600020905b8154815290600101906020018083116106f457829003601f168201915b505050505090806002018054610726906114a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610752906114a3565b801561079f5780601f106107745761010080835404028352916020019161079f565b820191906000526020600020905b81548152906001019060200180831161078257829003601f168201915b5050505050908060030154905084565b60005b815181101561090d5760008282815181106107d0576107cf611878565b5b6020026020010151905060006002826040516107ec919061145d565b90815260200160405180910390208054610805906114a3565b905014610847576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083e906118f3565b60405180910390fd5b83600282604051610858919061145d565b9081526020016040518091039020908161087291906117a6565b50600384604051610883919061145d565b9081526020016040518091039020819080600181540180825580915050600190039060005260206000200160009091909190915090816108c391906117a6565b507f3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e8185426040516108f793929190611913565b60405180910390a15080806001019150506107b2565b505050565b60005481565b600281805160208101820180518482526020830160208501208183528095505050505050600091509050805461094d906114a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610979906114a3565b80156109c65780601f1061099b576101008083540402835291602001916109c6565b820191906000526020600020905b8154815290600101906020018083116109a957829003601f168201915b505050505081565b6000600180549050905090565b6003828051602081018201805184825260208301602085012081835280955050505050508181548110610a0d57600080fd5b90600052602060002001600091509150508054610a29906114a3565b80601f0160208091040260200160405190810160405280929190818152602001828054610a55906114a3565b8015610aa25780601f10610a7757610100808354040283529160200191610aa2565b820191906000526020600020905b815481529060010190602001808311610a8557829003601f168201915b505050505081565b60008054905090565b6000600383604051610ac5919061145d565b9081526020016040518091039020905060005b8180549050811015610bbb578280519060200120828281548110610aff57610afe611878565b5b90600052602060002001604051610b169190611577565b604051809103902003610bae578160018380549050610b359190611373565b81548110610b4657610b45611878565b5b90600052602060002001828281548110610b6357610b62611878565b5b906000526020600020019081610b799190611983565b5081805480610b8b57610b8a611a6b565b5b600190038181906000526020600020016000610ba79190610bc1565b9055610bbb565b8080600101915050610ad8565b50505050565b508054610bcd906114a3565b6000825580601f10610bdf5750610bfe565b601f016020900490600052602060002090810190610bfd9190610c01565b5b50565b5b80821115610c1a576000816000905550600101610c02565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610c8582610c3c565b810181811067ffffffffffffffff82111715610ca457610ca3610c4d565b5b80604052505050565b6000610cb7610c1e565b9050610cc38282610c7c565b919050565b600067ffffffffffffffff821115610ce357610ce2610c4d565b5b610cec82610c3c565b9050602081019050919050565b82818337600083830152505050565b6000610d1b610d1684610cc8565b610cad565b905082815260208101848484011115610d3757610d36610c37565b5b610d42848285610cf9565b509392505050565b600082601f830112610d5f57610d5e610c32565b5b8135610d6f848260208601610d08565b91505092915050565b600080600060608486031215610d9157610d90610c28565b5b600084013567ffffffffffffffff811115610daf57610dae610c2d565b5b610dbb86828701610d4a565b935050602084013567ffffffffffffffff811115610ddc57610ddb610c2d565b5b610de886828701610d4a565b925050604084013567ffffffffffffffff811115610e0957610e08610c2d565b5b610e1586828701610d4a565b9150509250925092565b600060208284031215610e3557610e34610c28565b5b600082013567ffffffffffffffff811115610e5357610e52610c2d565b5b610e5f84828501610d4a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ece578082015181840152602081019050610eb3565b60008484015250505050565b6000610ee582610e94565b610eef8185610e9f565b9350610eff818560208601610eb0565b610f0881610c3c565b840191505092915050565b6000610f1f8383610eda565b905092915050565b6000602082019050919050565b6000610f3f82610e68565b610f498185610e73565b935083602082028501610f5b85610e84565b8060005b85811015610f975784840389528151610f788582610f13565b9450610f8383610f27565b925060208a01995050600181019050610f5f565b50829750879550505050505092915050565b60006020820190508181036000830152610fc38184610f34565b905092915050565b6000819050919050565b610fde81610fcb565b8114610fe957600080fd5b50565b600081359050610ffb81610fd5565b92915050565b60006020828403121561101757611016610c28565b5b600061102584828501610fec565b91505092915050565b600082825260208201905092915050565b600061104a82610e94565b611054818561102e565b9350611064818560208601610eb0565b61106d81610c3c565b840191505092915050565b61108181610fcb565b82525050565b600060808201905081810360008301526110a1818761103f565b905081810360208301526110b5818661103f565b905081810360408301526110c9818561103f565b90506110d86060830184611078565b95945050505050565b600067ffffffffffffffff8211156110fc576110fb610c4d565b5b602082029050602081019050919050565b600080fd5b6000611125611120846110e1565b610cad565b905080838252602082019050602084028301858111156111485761114761110d565b5b835b8181101561118f57803567ffffffffffffffff81111561116d5761116c610c32565b5b80860161117a8982610d4a565b8552602085019450505060208101905061114a565b5050509392505050565b600082601f8301126111ae576111ad610c32565b5b81356111be848260208601611112565b91505092915050565b600080604083850312156111de576111dd610c28565b5b600083013567ffffffffffffffff8111156111fc576111fb610c2d565b5b61120885828601610d4a565b925050602083013567ffffffffffffffff81111561122957611228610c2d565b5b61123585828601611199565b9150509250929050565b60006020820190506112546000830184611078565b92915050565b60006020820190508181036000830152611274818461103f565b905092915050565b6000806040838503121561129357611292610c28565b5b600083013567ffffffffffffffff8111156112b1576112b0610c2d565b5b6112bd85828601610d4a565b92505060206112ce85828601610fec565b9150509250929050565b7f4573746f717565206573676f7461646f00000000000000000000000000000000600082015250565b600061130e60108361102e565b9150611319826112d8565b602082019050919050565b6000602082019050818103600083015261133d81611301565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061137e82610fcb565b915061138983610fcb565b92508282039050818111156113a1576113a0611344565b5b92915050565b7f5061636f74652061626572746f00000000000000000000000000000000000000600082015250565b60006113dd600d8361102e565b91506113e8826113a7565b602082019050919050565b60006040820190506114086000830184611078565b8181036020830152611419816113d0565b905092915050565b600081905092915050565b600061143782610e94565b6114418185611421565b9350611451818560208601610eb0565b80840191505092915050565b6000611469828461142c565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114bb57607f821691505b6020821081036114ce576114cd611474565b5b50919050565b600081905092915050565b60008190508160005260206000209050919050565b60008154611501816114a3565b61150b81866114d4565b94506001821660008114611526576001811461153b5761156e565b60ff198316865281151582028601935061156e565b611544856114df565b60005b8381101561156657815481890152600182019150602081019050611547565b838801955050505b50505092915050565b600061158382846114f4565b915081905092915050565b7f4e616f2065206f20646f6e6f0000000000000000000000000000000000000000600082015250565b60006115c4600c8361102e565b91506115cf8261158e565b602082019050919050565b600060208201905081810360008301526115f3816115b7565b9050919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b60006008830261165c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261161f565b611666868361161f565b95508019841693508086168417925050509392505050565b6000819050919050565b60006116a361169e61169984610fcb565b61167e565b610fcb565b9050919050565b6000819050919050565b6116bd83611688565b6116d16116c9826116aa565b84845461162c565b825550505050565b600090565b6116e66116d9565b6116f18184846116b4565b505050565b5b818110156117155761170a6000826116de565b6001810190506116f7565b5050565b601f82111561175a5761172b816115fa565b6117348461160f565b81016020851015611743578190505b61175761174f8561160f565b8301826116f6565b50505b505050565b600082821c905092915050565b600061177d6000198460080261175f565b1980831691505092915050565b6000611796838361176c565b9150826002028217905092915050565b6117af82610e94565b67ffffffffffffffff8111156117c8576117c7610c4d565b5b6117d282546114a3565b6117dd828285611719565b600060209050601f83116001811461181057600084156117fe578287015190505b611808858261178a565b865550611870565b601f19841661181e866115fa565b60005b8281101561184657848901518255600182019150602085019450602081019050611821565b86831015611863578489015161185f601f89168261176c565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4361727461206a612074656d20646f6e6f000000000000000000000000000000600082015250565b60006118dd60118361102e565b91506118e8826118a7565b602082019050919050565b6000602082019050818103600083015261190c816118d0565b9050919050565b6000606082019050818103600083015261192d818661103f565b90508181036020830152611941818561103f565b90506119506040830184611078565b949350505050565b600081549050611967816114a3565b9050919050565b60008190508160005260206000209050919050565b818103611991575050611a69565b61199a82611958565b67ffffffffffffffff8111156119b3576119b2610c4d565b5b6119bd82546114a3565b6119c8828285611719565b6000601f8311600181146119f757600084156119e5578287015490505b6119ef858261178a565b865550611a62565b601f198416611a058761196e565b9650611a10866115fa565b60005b82811015611a3857848901548255600182019150600185019450602081019050611a13565b86831015611a555784890154611a51601f89168261176c565b8355505b6001600288020188555050505b5050505050505b565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea2646970667358221220a5316be042296550b54be55cea7bf3d5bf191b873af7084fc26105c74140c52e64736f6c63430008180033",
}

// PackRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PackRegistryMetaData.ABI instead.
var PackRegistryABI = PackRegistryMetaData.ABI

// PackRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PackRegistryMetaData.Bin instead.
var PackRegistryBin = PackRegistryMetaData.Bin

// DeployPackRegistry deploys a new Ethereum contract, binding an instance of PackRegistry to it.
func DeployPackRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _initialStock *big.Int) (common.Address, *types.Transaction, *PackRegistry, error) {
	parsed, err := PackRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PackRegistryBin), backend, _initialStock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PackRegistry{PackRegistryCaller: PackRegistryCaller{contract: contract}, PackRegistryTransactor: PackRegistryTransactor{contract: contract}, PackRegistryFilterer: PackRegistryFilterer{contract: contract}}, nil
}

// PackRegistry is an auto generated Go binding around an Ethereum contract.
type PackRegistry struct {
	PackRegistryCaller     // Read-only binding to the contract
	PackRegistryTransactor // Write-only binding to the contract
	PackRegistryFilterer   // Log filterer for contract events
}

// PackRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PackRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PackRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackRegistrySession struct {
	Contract     *PackRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackRegistryCallerSession struct {
	Contract *PackRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PackRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackRegistryTransactorSession struct {
	Contract     *PackRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PackRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PackRegistryRaw struct {
	Contract *PackRegistry // Generic contract binding to access the raw methods on
}

// PackRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackRegistryCallerRaw struct {
	Contract *PackRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PackRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackRegistryTransactorRaw struct {
	Contract *PackRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPackRegistry creates a new instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistry(address common.Address, backend bind.ContractBackend) (*PackRegistry, error) {
	contract, err := bindPackRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PackRegistry{PackRegistryCaller: PackRegistryCaller{contract: contract}, PackRegistryTransactor: PackRegistryTransactor{contract: contract}, PackRegistryFilterer: PackRegistryFilterer{contract: contract}}, nil
}

// NewPackRegistryCaller creates a new read-only instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistryCaller(address common.Address, caller bind.ContractCaller) (*PackRegistryCaller, error) {
	contract, err := bindPackRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackRegistryCaller{contract: contract}, nil
}

// NewPackRegistryTransactor creates a new write-only instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PackRegistryTransactor, error) {
	contract, err := bindPackRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackRegistryTransactor{contract: contract}, nil
}

// NewPackRegistryFilterer creates a new log filterer instance of PackRegistry, bound to a specific deployed contract.
func NewPackRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PackRegistryFilterer, error) {
	contract, err := bindPackRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackRegistryFilterer{contract: contract}, nil
}

// bindPackRegistry binds a generic wrapper to an already deployed contract.
func bindPackRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PackRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackRegistry *PackRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackRegistry.Contract.PackRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackRegistry *PackRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackRegistry.Contract.PackRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackRegistry *PackRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackRegistry.Contract.PackRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PackRegistry *PackRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PackRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PackRegistry *PackRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PackRegistry *PackRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PackRegistry.Contract.contract.Transact(opts, method, params...)
}

// CardOwner is a free data retrieval call binding the contract method 0x888ad898.
//
// Solidity: function cardOwner(string ) view returns(string)
func (_PackRegistry *PackRegistryCaller) CardOwner(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "cardOwner", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CardOwner is a free data retrieval call binding the contract method 0x888ad898.
//
// Solidity: function cardOwner(string ) view returns(string)
func (_PackRegistry *PackRegistrySession) CardOwner(arg0 string) (string, error) {
	return _PackRegistry.Contract.CardOwner(&_PackRegistry.CallOpts, arg0)
}

// CardOwner is a free data retrieval call binding the contract method 0x888ad898.
//
// Solidity: function cardOwner(string ) view returns(string)
func (_PackRegistry *PackRegistryCallerSession) CardOwner(arg0 string) (string, error) {
	return _PackRegistry.Contract.CardOwner(&_PackRegistry.CallOpts, arg0)
}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_PackRegistry *PackRegistryCaller) GetMatchCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "getMatchCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_PackRegistry *PackRegistrySession) GetMatchCount() (*big.Int, error) {
	return _PackRegistry.Contract.GetMatchCount(&_PackRegistry.CallOpts)
}

// GetMatchCount is a free data retrieval call binding the contract method 0x8c4f7dae.
//
// Solidity: function getMatchCount() view returns(uint256)
func (_PackRegistry *PackRegistryCallerSession) GetMatchCount() (*big.Int, error) {
	return _PackRegistry.Contract.GetMatchCount(&_PackRegistry.CallOpts)
}

// GetStock is a free data retrieval call binding the contract method 0xe38ffcd8.
//
// Solidity: function getStock() view returns(uint256)
func (_PackRegistry *PackRegistryCaller) GetStock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "getStock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStock is a free data retrieval call binding the contract method 0xe38ffcd8.
//
// Solidity: function getStock() view returns(uint256)
func (_PackRegistry *PackRegistrySession) GetStock() (*big.Int, error) {
	return _PackRegistry.Contract.GetStock(&_PackRegistry.CallOpts)
}

// GetStock is a free data retrieval call binding the contract method 0xe38ffcd8.
//
// Solidity: function getStock() view returns(uint256)
func (_PackRegistry *PackRegistryCallerSession) GetStock() (*big.Int, error) {
	return _PackRegistry.Contract.GetStock(&_PackRegistry.CallOpts)
}

// GetUserCards is a free data retrieval call binding the contract method 0x194b13b6.
//
// Solidity: function getUserCards(string _ownerId) view returns(string[])
func (_PackRegistry *PackRegistryCaller) GetUserCards(opts *bind.CallOpts, _ownerId string) ([]string, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "getUserCards", _ownerId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetUserCards is a free data retrieval call binding the contract method 0x194b13b6.
//
// Solidity: function getUserCards(string _ownerId) view returns(string[])
func (_PackRegistry *PackRegistrySession) GetUserCards(_ownerId string) ([]string, error) {
	return _PackRegistry.Contract.GetUserCards(&_PackRegistry.CallOpts, _ownerId)
}

// GetUserCards is a free data retrieval call binding the contract method 0x194b13b6.
//
// Solidity: function getUserCards(string _ownerId) view returns(string[])
func (_PackRegistry *PackRegistryCallerSession) GetUserCards(_ownerId string) ([]string, error) {
	return _PackRegistry.Contract.GetUserCards(&_PackRegistry.CallOpts, _ownerId)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(string matchId, string winnerId, string loserId, uint256 timestamp)
func (_PackRegistry *PackRegistryCaller) Matches(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MatchId   string
	WinnerId  string
	LoserId   string
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "matches", arg0)

	outstruct := new(struct {
		MatchId   string
		WinnerId  string
		LoserId   string
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MatchId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.WinnerId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.LoserId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(string matchId, string winnerId, string loserId, uint256 timestamp)
func (_PackRegistry *PackRegistrySession) Matches(arg0 *big.Int) (struct {
	MatchId   string
	WinnerId  string
	LoserId   string
	Timestamp *big.Int
}, error) {
	return _PackRegistry.Contract.Matches(&_PackRegistry.CallOpts, arg0)
}

// Matches is a free data retrieval call binding the contract method 0x4768d4ef.
//
// Solidity: function matches(uint256 ) view returns(string matchId, string winnerId, string loserId, uint256 timestamp)
func (_PackRegistry *PackRegistryCallerSession) Matches(arg0 *big.Int) (struct {
	MatchId   string
	WinnerId  string
	LoserId   string
	Timestamp *big.Int
}, error) {
	return _PackRegistry.Contract.Matches(&_PackRegistry.CallOpts, arg0)
}

// OwnerCards is a free data retrieval call binding the contract method 0x8f2ce142.
//
// Solidity: function ownerCards(string , uint256 ) view returns(string)
func (_PackRegistry *PackRegistryCaller) OwnerCards(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "ownerCards", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// OwnerCards is a free data retrieval call binding the contract method 0x8f2ce142.
//
// Solidity: function ownerCards(string , uint256 ) view returns(string)
func (_PackRegistry *PackRegistrySession) OwnerCards(arg0 string, arg1 *big.Int) (string, error) {
	return _PackRegistry.Contract.OwnerCards(&_PackRegistry.CallOpts, arg0, arg1)
}

// OwnerCards is a free data retrieval call binding the contract method 0x8f2ce142.
//
// Solidity: function ownerCards(string , uint256 ) view returns(string)
func (_PackRegistry *PackRegistryCallerSession) OwnerCards(arg0 string, arg1 *big.Int) (string, error) {
	return _PackRegistry.Contract.OwnerCards(&_PackRegistry.CallOpts, arg0, arg1)
}

// TotalPacks is a free data retrieval call binding the contract method 0x7281df12.
//
// Solidity: function totalPacks() view returns(uint256)
func (_PackRegistry *PackRegistryCaller) TotalPacks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PackRegistry.contract.Call(opts, &out, "totalPacks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPacks is a free data retrieval call binding the contract method 0x7281df12.
//
// Solidity: function totalPacks() view returns(uint256)
func (_PackRegistry *PackRegistrySession) TotalPacks() (*big.Int, error) {
	return _PackRegistry.Contract.TotalPacks(&_PackRegistry.CallOpts)
}

// TotalPacks is a free data retrieval call binding the contract method 0x7281df12.
//
// Solidity: function totalPacks() view returns(uint256)
func (_PackRegistry *PackRegistryCallerSession) TotalPacks() (*big.Int, error) {
	return _PackRegistry.Contract.TotalPacks(&_PackRegistry.CallOpts)
}

// AssignCards is a paid mutator transaction binding the contract method 0x6d34550e.
//
// Solidity: function assignCards(string _playerId, string[] _cardIds) returns()
func (_PackRegistry *PackRegistryTransactor) AssignCards(opts *bind.TransactOpts, _playerId string, _cardIds []string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "assignCards", _playerId, _cardIds)
}

// AssignCards is a paid mutator transaction binding the contract method 0x6d34550e.
//
// Solidity: function assignCards(string _playerId, string[] _cardIds) returns()
func (_PackRegistry *PackRegistrySession) AssignCards(_playerId string, _cardIds []string) (*types.Transaction, error) {
	return _PackRegistry.Contract.AssignCards(&_PackRegistry.TransactOpts, _playerId, _cardIds)
}

// AssignCards is a paid mutator transaction binding the contract method 0x6d34550e.
//
// Solidity: function assignCards(string _playerId, string[] _cardIds) returns()
func (_PackRegistry *PackRegistryTransactorSession) AssignCards(_playerId string, _cardIds []string) (*types.Transaction, error) {
	return _PackRegistry.Contract.AssignCards(&_PackRegistry.TransactOpts, _playerId, _cardIds)
}

// DecrementStock is a paid mutator transaction binding the contract method 0x00375a44.
//
// Solidity: function decrementStock() returns()
func (_PackRegistry *PackRegistryTransactor) DecrementStock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "decrementStock")
}

// DecrementStock is a paid mutator transaction binding the contract method 0x00375a44.
//
// Solidity: function decrementStock() returns()
func (_PackRegistry *PackRegistrySession) DecrementStock() (*types.Transaction, error) {
	return _PackRegistry.Contract.DecrementStock(&_PackRegistry.TransactOpts)
}

// DecrementStock is a paid mutator transaction binding the contract method 0x00375a44.
//
// Solidity: function decrementStock() returns()
func (_PackRegistry *PackRegistryTransactorSession) DecrementStock() (*types.Transaction, error) {
	return _PackRegistry.Contract.DecrementStock(&_PackRegistry.TransactOpts)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x12636a62.
//
// Solidity: function recordMatch(string _matchId, string _winnerId, string _loserId) returns()
func (_PackRegistry *PackRegistryTransactor) RecordMatch(opts *bind.TransactOpts, _matchId string, _winnerId string, _loserId string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "recordMatch", _matchId, _winnerId, _loserId)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x12636a62.
//
// Solidity: function recordMatch(string _matchId, string _winnerId, string _loserId) returns()
func (_PackRegistry *PackRegistrySession) RecordMatch(_matchId string, _winnerId string, _loserId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.RecordMatch(&_PackRegistry.TransactOpts, _matchId, _winnerId, _loserId)
}

// RecordMatch is a paid mutator transaction binding the contract method 0x12636a62.
//
// Solidity: function recordMatch(string _matchId, string _winnerId, string _loserId) returns()
func (_PackRegistry *PackRegistryTransactorSession) RecordMatch(_matchId string, _winnerId string, _loserId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.RecordMatch(&_PackRegistry.TransactOpts, _matchId, _winnerId, _loserId)
}

// TransferCard is a paid mutator transaction binding the contract method 0x118c54ec.
//
// Solidity: function transferCard(string _fromId, string _toId, string _cardId) returns()
func (_PackRegistry *PackRegistryTransactor) TransferCard(opts *bind.TransactOpts, _fromId string, _toId string, _cardId string) (*types.Transaction, error) {
	return _PackRegistry.contract.Transact(opts, "transferCard", _fromId, _toId, _cardId)
}

// TransferCard is a paid mutator transaction binding the contract method 0x118c54ec.
//
// Solidity: function transferCard(string _fromId, string _toId, string _cardId) returns()
func (_PackRegistry *PackRegistrySession) TransferCard(_fromId string, _toId string, _cardId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.TransferCard(&_PackRegistry.TransactOpts, _fromId, _toId, _cardId)
}

// TransferCard is a paid mutator transaction binding the contract method 0x118c54ec.
//
// Solidity: function transferCard(string _fromId, string _toId, string _cardId) returns()
func (_PackRegistry *PackRegistryTransactorSession) TransferCard(_fromId string, _toId string, _cardId string) (*types.Transaction, error) {
	return _PackRegistry.Contract.TransferCard(&_PackRegistry.TransactOpts, _fromId, _toId, _cardId)
}

// PackRegistryCardAssignedIterator is returned from FilterCardAssigned and is used to iterate over the raw logs and unpacked data for CardAssigned events raised by the PackRegistry contract.
type PackRegistryCardAssignedIterator struct {
	Event *PackRegistryCardAssigned // Event containing the contract specifics and raw log

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
func (it *PackRegistryCardAssignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryCardAssigned)
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
		it.Event = new(PackRegistryCardAssigned)
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
func (it *PackRegistryCardAssignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryCardAssignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryCardAssigned represents a CardAssigned event raised by the PackRegistry contract.
type PackRegistryCardAssigned struct {
	CardId    string
	OwnerId   string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCardAssigned is a free log retrieval operation binding the contract event 0x3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e.
//
// Solidity: event CardAssigned(string cardId, string ownerId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) FilterCardAssigned(opts *bind.FilterOpts) (*PackRegistryCardAssignedIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "CardAssigned")
	if err != nil {
		return nil, err
	}
	return &PackRegistryCardAssignedIterator{contract: _PackRegistry.contract, event: "CardAssigned", logs: logs, sub: sub}, nil
}

// WatchCardAssigned is a free log subscription operation binding the contract event 0x3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e.
//
// Solidity: event CardAssigned(string cardId, string ownerId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) WatchCardAssigned(opts *bind.WatchOpts, sink chan<- *PackRegistryCardAssigned) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "CardAssigned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryCardAssigned)
				if err := _PackRegistry.contract.UnpackLog(event, "CardAssigned", log); err != nil {
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

// ParseCardAssigned is a log parse operation binding the contract event 0x3c3087eb611c2be646224be60e6856109eeb94b7d90e6e707fed5c096a82703e.
//
// Solidity: event CardAssigned(string cardId, string ownerId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) ParseCardAssigned(log types.Log) (*PackRegistryCardAssigned, error) {
	event := new(PackRegistryCardAssigned)
	if err := _PackRegistry.contract.UnpackLog(event, "CardAssigned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRegistryCardTransferredIterator is returned from FilterCardTransferred and is used to iterate over the raw logs and unpacked data for CardTransferred events raised by the PackRegistry contract.
type PackRegistryCardTransferredIterator struct {
	Event *PackRegistryCardTransferred // Event containing the contract specifics and raw log

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
func (it *PackRegistryCardTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryCardTransferred)
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
		it.Event = new(PackRegistryCardTransferred)
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
func (it *PackRegistryCardTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryCardTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryCardTransferred represents a CardTransferred event raised by the PackRegistry contract.
type PackRegistryCardTransferred struct {
	CardId    string
	FromId    string
	ToId      string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCardTransferred is a free log retrieval operation binding the contract event 0x6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d.
//
// Solidity: event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) FilterCardTransferred(opts *bind.FilterOpts) (*PackRegistryCardTransferredIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "CardTransferred")
	if err != nil {
		return nil, err
	}
	return &PackRegistryCardTransferredIterator{contract: _PackRegistry.contract, event: "CardTransferred", logs: logs, sub: sub}, nil
}

// WatchCardTransferred is a free log subscription operation binding the contract event 0x6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d.
//
// Solidity: event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) WatchCardTransferred(opts *bind.WatchOpts, sink chan<- *PackRegistryCardTransferred) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "CardTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryCardTransferred)
				if err := _PackRegistry.contract.UnpackLog(event, "CardTransferred", log); err != nil {
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

// ParseCardTransferred is a log parse operation binding the contract event 0x6b557a50618841e0cdca1b7d2e68c490faf639331b28e17efa4930ea03d88e9d.
//
// Solidity: event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) ParseCardTransferred(log types.Log) (*PackRegistryCardTransferred, error) {
	event := new(PackRegistryCardTransferred)
	if err := _PackRegistry.contract.UnpackLog(event, "CardTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRegistryMatchRecordedIterator is returned from FilterMatchRecorded and is used to iterate over the raw logs and unpacked data for MatchRecorded events raised by the PackRegistry contract.
type PackRegistryMatchRecordedIterator struct {
	Event *PackRegistryMatchRecorded // Event containing the contract specifics and raw log

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
func (it *PackRegistryMatchRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryMatchRecorded)
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
		it.Event = new(PackRegistryMatchRecorded)
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
func (it *PackRegistryMatchRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryMatchRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryMatchRecorded represents a MatchRecorded event raised by the PackRegistry contract.
type PackRegistryMatchRecorded struct {
	MatchId   string
	Winner    string
	Loser     string
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMatchRecorded is a free log retrieval operation binding the contract event 0xe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37.
//
// Solidity: event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) FilterMatchRecorded(opts *bind.FilterOpts) (*PackRegistryMatchRecordedIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "MatchRecorded")
	if err != nil {
		return nil, err
	}
	return &PackRegistryMatchRecordedIterator{contract: _PackRegistry.contract, event: "MatchRecorded", logs: logs, sub: sub}, nil
}

// WatchMatchRecorded is a free log subscription operation binding the contract event 0xe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37.
//
// Solidity: event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) WatchMatchRecorded(opts *bind.WatchOpts, sink chan<- *PackRegistryMatchRecorded) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "MatchRecorded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryMatchRecorded)
				if err := _PackRegistry.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
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

// ParseMatchRecorded is a log parse operation binding the contract event 0xe6bfbd050def41eb3a8f15a21087f802b211548d57828abcb3741ec236bc1b37.
//
// Solidity: event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp)
func (_PackRegistry *PackRegistryFilterer) ParseMatchRecorded(log types.Log) (*PackRegistryMatchRecorded, error) {
	event := new(PackRegistryMatchRecorded)
	if err := _PackRegistry.contract.UnpackLog(event, "MatchRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRegistryStockUpdatedIterator is returned from FilterStockUpdated and is used to iterate over the raw logs and unpacked data for StockUpdated events raised by the PackRegistry contract.
type PackRegistryStockUpdatedIterator struct {
	Event *PackRegistryStockUpdated // Event containing the contract specifics and raw log

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
func (it *PackRegistryStockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRegistryStockUpdated)
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
		it.Event = new(PackRegistryStockUpdated)
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
func (it *PackRegistryStockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRegistryStockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRegistryStockUpdated represents a StockUpdated event raised by the PackRegistry contract.
type PackRegistryStockUpdated struct {
	NewStock *big.Int
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStockUpdated is a free log retrieval operation binding the contract event 0x2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0.
//
// Solidity: event StockUpdated(uint256 newStock, string reason)
func (_PackRegistry *PackRegistryFilterer) FilterStockUpdated(opts *bind.FilterOpts) (*PackRegistryStockUpdatedIterator, error) {

	logs, sub, err := _PackRegistry.contract.FilterLogs(opts, "StockUpdated")
	if err != nil {
		return nil, err
	}
	return &PackRegistryStockUpdatedIterator{contract: _PackRegistry.contract, event: "StockUpdated", logs: logs, sub: sub}, nil
}

// WatchStockUpdated is a free log subscription operation binding the contract event 0x2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0.
//
// Solidity: event StockUpdated(uint256 newStock, string reason)
func (_PackRegistry *PackRegistryFilterer) WatchStockUpdated(opts *bind.WatchOpts, sink chan<- *PackRegistryStockUpdated) (event.Subscription, error) {

	logs, sub, err := _PackRegistry.contract.WatchLogs(opts, "StockUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRegistryStockUpdated)
				if err := _PackRegistry.contract.UnpackLog(event, "StockUpdated", log); err != nil {
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

// ParseStockUpdated is a log parse operation binding the contract event 0x2926060a4c291371fea4258fa171860291f0144cd93a2accc8a0da15fc05c7a0.
//
// Solidity: event StockUpdated(uint256 newStock, string reason)
func (_PackRegistry *PackRegistryFilterer) ParseStockUpdated(log types.Log) (*PackRegistryStockUpdated, error) {
	event := new(PackRegistryStockUpdated)
	if err := _PackRegistry.contract.UnpackLog(event, "StockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
