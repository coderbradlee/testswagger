// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"
)

// Standard JSON-RPC 2.0 errors.
var (
	ErrRPCInvalidRequest = &RPCError{
		Code:    -32600,
		Message: "Invalid request",
	}
	ErrRPCMethodNotFound = &RPCError{
		Code:    -32601,
		Message: "Method not found",
	}
	ErrRPCInvalidParams = &RPCError{
		Code:    -32602,
		Message: "Invalid parameters",
	}
	ErrRPCInternal = &RPCError{
		Code:    -32603,
		Message: "Internal error",
	}
	ErrRPCParse = &RPCError{
		Code:    -32700,
		Message: "Parse error",
	}
)

type RPCError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e RPCError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
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
	ID      interface{}       `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

type Response struct {
	ID     *interface{}    `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  *RPCError       `json:"error"`
}

func NewResponse(id interface{}, marshalledResult []byte, rpcErr *RPCError) (*Response, error) {
	if !IsValidIDType(id) {
		return nil, ErrRPCInvalidParams
	}

	pid := &id
	return &Response{
		Result: marshalledResult,
		Error:  rpcErr,
		ID:     pid,
	}, nil
}

func (o Response) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	if err := producer.Produce(rw, o); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
