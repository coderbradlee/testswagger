package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/lzxm160/testswagger/contract"

	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"

	"github.com/lzxm160/testswagger/restapi/operations/get"
	"github.com/lzxm160/testswagger/restapi/operations/update"
)

const (
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	Chainpoint            = "api.testnet.iotex.one:443"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
)

func UpdateHandler(params update.UpdateParams) *Response {
	fmt.Println("UpdateHandler:", *params.Body)
	chainpoint := os.Getenv("CHAINPOINT")
	if chainpoint == "" {
		chainpoint = Chainpoint
	}
	DIDAddress := os.Getenv("IoTeXDIDPROXYADDRESS")
	if DIDAddress == "" {
		DIDAddress = IoTeXDIDProxy_address
	}
	d, err := NewDID(chainpoint, privateKey, DIDAddress, contract.IoTeXDIDProxyABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	if err != nil {
		ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
		return ret
	}
	var result string

	switch *params.Body.Method {
	case createDID:
		if len(params.Body.Params) != 3 {
			ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
			return ret
		}
		result, err = d.CreateDID(params.Body.Params[0], params.Body.Params[1], params.Body.Params[2])
	case deleteDID:
		if len(params.Body.Params) != 1 {
			ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
			return ret
		}
		result, err = d.DeleteDID(params.Body.Params[0])
	case updateHash:
		if len(params.Body.Params) != 2 {
			ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
			return ret
		}
		result, err = d.UpdateHash(params.Body.Params[0], params.Body.Params[1])
	case updateURI:
		if len(params.Body.Params) != 2 {
			ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
			return ret
		}
		result, err = d.UpdateUri(params.Body.Params[0], params.Body.Params[1])
	default:
		err = errors.New("request invalid method")
	}
	if err != nil {
		ret, _ := NewResponse(nil, nil, ErrRPCInternal)
		return ret
	}
	marshalledResult, err := json.Marshal(result)
	if err != nil {
		return nil
	}
	ret, _ := NewResponse(*params.Body.ID, marshalledResult, nil)
	return ret
}

func GetHandler(params get.GetParams) *Response {
	fmt.Println("GetHandler:", *params.Body)
	chainpoint := os.Getenv("CHAINPOINT")
	if chainpoint == "" {
		chainpoint = Chainpoint
	}
	DIDAddress := os.Getenv("IoTeXDIDPROXYADDRESS")
	if DIDAddress == "" {
		DIDAddress = IoTeXDIDProxy_address
	}
	if len(params.Body.Params) != 1 {
		ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
		return ret
	}
	d, err := NewDID(chainpoint, privateKey, DIDAddress, contract.IoTeXDIDProxyABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
	if err != nil {
		ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
		return ret
	}
	var result string
	switch *params.Body.Method {
	case getHash:
		result, err = d.GetHash(params.Body.Params[0])
	case getURI:
		result, err = d.GetUri(params.Body.Params[0])
	default:
		err = errors.New("request invalid method")
	}
	if err != nil {
		ret, _ := NewResponse(nil, nil, ErrRPCMethodNotFound)
		return ret
	}
	marshalledResult, err := json.Marshal(result)
	if err != nil {
		return nil
	}
	ret, _ := NewResponse(*params.Body.ID, marshalledResult, nil)
	return ret
}
