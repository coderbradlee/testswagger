package handler

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/lzxm160/testswagger/contract/IoTeXDID"
)

const (
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	Chainpoint            = "api.testnet.iotex.one:443"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
	GasPrice              = "1000000000000"
	GasLimit              = "1000000"
	testDID               = `{"@context": "https://w3id.org/did/v1","id": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6","authentication": [{"type": "Secp256k1SignatureAuthentication2018","publicKey": ["did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6#owner"]}],"publicKey": [{"id": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6#owner","type": "Secp256k1VerificationKey2018","ethereumAddress": "0xe6fe788d8ca214a080b0f6ac7f48480b2aefa9a6","owner": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6"},{"id": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6#delegate-1","type": "Secp256k1VerificationKey2018","owner": "did:ethr:0xE6Fe788d8ca214A080b0f6aC7F48480b2AEfa9a6","publicKeyHex": "0295dda1dca7f80e308ef60155ddeac00e46b797fd40ef407f422e88d2467a27eb"}]}`
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

func GetHandler(did string) *Response {
	ret, _ := NewResponse("")
	d, err := NewDID(chainpoint, "", DIDAddress, IoTeXDID.IoTeXDIDABI, gasPrice, gasLimit)
	if err != nil {
		return ret
	}
	result, err := d.GetUri(did)
	if err != nil {
		return ret
	}
	fmt.Println(result)
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
	ret, _ = NewResponse(testDID)
	return ret
}
