package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/lzxm160/testswagger/contract"

	"github.com/iotexproject/iotex-antenna-go/v2/utils/unit"

	"github.com/lzxm160/testswagger/restapi/operations/get"
	"github.com/lzxm160/testswagger/restapi/operations/update"
)

const (
	sender                = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	privateKey            = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	endpoint              = "api.testnet.iotex.one:443"
	IoTeXDID_address      = "io1eurq3lx4lzx9wdj56plw5rm59f5qzanacr3raz"
	IoTeXDIDProxy_address = "io1zgs5gqjl679qlj4gqqpa9t329r8f5gr8xc9lr0"
)

func UpdateHandler(params update.UpdateParams) *Response {
	fmt.Println("UpdateHandler:", *params.Body)
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDProxyABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
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
	if len(params.Body.Params) != 1 {
		ret, _ := NewResponse(nil, nil, ErrRPCInvalidParams)
		return ret
	}
	d, err := NewDID(endpoint, privateKey, IoTeXDIDProxy_address, contract.IoTeXDIDProxyABI, big.NewInt(int64(unit.Qev)), uint64(1000000))
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
