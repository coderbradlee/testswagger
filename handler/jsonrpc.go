// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package handler

import (
	"encoding/json"
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

// NewRequest returns a new JSON-RPC 1.0 request object given the provided id,
// method, and parameters.  The parameters are marshalled into a json.RawMessage
// for the Params field of the returned request object.  This function is only
// provided in case the caller wants to construct raw requests for some reason.
//
// Typically callers will instead want to create a registered concrete command
// type with the NewCmd or New<Foo>Cmd functions and call the MarshalCmd
// function with that command to generate the marshalled JSON-RPC request.
//func NewRequest(id interface{}, method string, params []interface{}) (*Request, error) {
//	if !IsValidIDType(id) {
//		str := fmt.Sprintf("the id of type '%T' is invalid", id)
//		return nil, makeError(ErrInvalidType, str)
//	}
//
//	rawParams := make([]json.RawMessage, 0, len(params))
//	for _, param := range params {
//		marshalledParam, err := json.Marshal(param)
//		if err != nil {
//			return nil, err
//		}
//		rawMessage := json.RawMessage(marshalledParam)
//		rawParams = append(rawParams, rawMessage)
//	}
//
//	return &Request{
//		Jsonrpc: "1.0",
//		ID:      id,
//		Method:  method,
//		Params:  rawParams,
//	}, nil
//}

// Response is the general form of a JSON-RPC response.  The type of the Result
// field varies from one command to the next, so it is implemented as an
// interface.  The ID field has to be a pointer for Go to put a null in it when
// empty.
type Response struct {
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
	ID     *interface{}    `json:"id"`
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

func UpdateHandler(params update.UpdateParams) []byte {
	fmt.Println("UpdateHandler:", *params.Body)
	var ret []byte
	switch *params.Body.Method {
	case "createDID":
		result := "createDID"
		ret, _ = MarshalResponse(*params.Body.ID, &result, nil)
	case "deleteDID":
		result := "deleteDID"
		ret, _ = MarshalResponse(*params.Body.ID, &result, nil)
	case "updateHash":
		result := "updateHash"
		ret, _ = MarshalResponse(*params.Body.ID, &result, nil)
	case "updateURI":
		result := "updateURI"
		ret, _ = MarshalResponse(*params.Body.ID, &result, nil)
	default:
		ret, _ = MarshalResponse(nil, nil, &RPCError{ErrInvalidMethod, "request invalid method"})
	}
	return ret
}
