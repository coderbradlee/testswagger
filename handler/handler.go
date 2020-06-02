package handler

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lzxm160/testswagger/restapi/operations/get"
	"github.com/lzxm160/testswagger/restapi/operations/update"
)

func UpdateHandler(params update.UpdateParams) *Response {
	fmt.Println("UpdateHandler:", *params.Body)
	var (
		result string
		err    error
	)

	switch *params.Body.Method {
	case "createDID":
		result = "createDID"
	case "deleteDID":
		result = "deleteDID"
	case "updateHash":
		result = "updateHash"
	case "updateURI":
		result = "updateURI"
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
func GetHandler(params get.GetParams) *Response {
	fmt.Println("GetHandler:", *params.Body)
	var (
		result string
		err    error
	)

	switch *params.Body.Method {
	case "getHash":
		result = "getHash"
	case "getURI":
		result = "getURI"
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
