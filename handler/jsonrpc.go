// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package handler

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lzxm160/testswagger/restapi/operations/update"
)

type RPCErrorCode int

const (
	ErrDuplicateMethod RPCErrorCode = iota
	ErrInvalidMethod
	ErrInvalidType
	ErrEmbeddedType
	ErrUnexportedField
	ErrUnsupportedFieldType
	ErrNonOptionalField
	ErrNonOptionalDefault
	ErrMismatchedDefault
	ErrUnregisteredMethod
	ErrUnmarshal
	ErrNumParams
)

type RPCError struct {
	Code    RPCErrorCode `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
}

var _, _ error = RPCError{}, (*RPCError)(nil)

func (e RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewRPCError(code RPCErrorCode, message string) *RPCError {
	return &RPCError{
		Code:    code,
		Message: message,
	}
}

func IsValidIDType(id interface{}) bool {
	switch id.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		string,
		nil:
		return true
	default:
		return false
	}
}

type Request struct {
	Jsonrpc string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
	ID      interface{}       `json:"id"`
}

type Response struct {
	ID     *interface{}    `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
}

func NewResponse(id interface{}, marshalledResult []byte, rpcErr *RPCError) (*Response, error) {
	if !IsValidIDType(id) {
		str := fmt.Sprintf("the id of type '%T' is invalid", id)
		return nil, RPCError{ErrInvalidType, str}
	}

	pid := &id
	return &Response{
		Result: marshalledResult,
		Error:  rpcErr,
		ID:     pid,
	}, nil
}

func MarshalResponse(id interface{}, result interface{}, rpcErr *RPCError) ([]byte, error) {
	marshalledResult, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}
	response, err := NewResponse(id, marshalledResult, rpcErr)
	if err != nil {
		return nil, err
	}
	return json.Marshal(&response)
}

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
		ret, _ := NewResponse(nil, nil, &RPCError{ErrInvalidMethod, err.Error()})
		return ret
	}

	marshalledResult, err := json.Marshal(result)
	if err != nil {
		return nil
	}
	ret, _ := NewResponse(*params.Body.ID, marshalledResult, nil)

	return ret
}
