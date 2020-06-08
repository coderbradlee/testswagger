package handler

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	Chainpoint            = "api.testnet.iotex.one:443"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
	GasPrice              = "1000000000000"
	GasLimit              = "1000000"
)

var (
	chainpoint string
	DIDAddress string
	gasPrice   *big.Int
	gasLimit   uint64
)

func init() {
	chainpoint = os.Getenv("CHAINPOINT")
	if chainpoint == "" {
		chainpoint = Chainpoint
	}
	DIDAddress = os.Getenv("IoTeXDIDPROXYADDRESS")
	if DIDAddress == "" {
		DIDAddress = IoTeXDIDProxy_address
	}
	gasPriceString := os.Getenv("GASPRICE")
	if gasPriceString == "" {
		gasPriceString = GasPrice
	}
	var ok bool
	gasPrice, ok = new(big.Int).SetString(gasPriceString, 10)
	if !ok {
		fmt.Println("gas price convert error")
	}
	gasLimitString := os.Getenv("GASLIMIT")
	if gasLimitString == "" {
		gasLimitString = GasLimit
	}
	var err error
	gasLimit, err = strconv.ParseUint(gasLimitString, 10, 64)
	if err != nil {
		fmt.Println("gas limit convert error", err)
	}
}

func GetHandler(id string) *Response {
	//fmt.Println("GetHandler:", *params.Body)
	fmt.Println(chainpoint, DIDAddress, gasPrice, gasLimit,id)
	//if len(params.Body.Params) != 1 {
	//	ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
	//	return ret
	//}
	//d, err := NewDID(chainpoint, privateKey, DIDAddress, IoTeXDID.IoTeXDIDABI, gasPrice, gasLimit)
	//if err != nil {
	//	ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
	//	return ret
	//}
	//fmt.Println("121")
	//var result string
	//switch *params.Body.Method {
	//case getHash:
	//	result, err = d.GetHash(params.Body.Params[0])
	//case getURI:
	//	result, err = d.GetUri(params.Body.Params[0])
	//default:
	//	err = errors.New("request invalid method")
	//}
	//if err != nil {
	//	ret, _ := NewResponse(nil, nil, ErrRPCMethodNotFound)
	//	return ret
	//}
	//marshalledResult, err := json.Marshal(result)
	//if err != nil {
	//	return nil
	//}
	//ret, _ := NewResponse(*params.Body.ID, marshalledResult, nil)
	//return ret
	return nil
}
