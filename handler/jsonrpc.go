// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.
package handler

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

type Response struct {
	DIDDocument string `json:"DIDDocument"`
}

func NewResponse(marshalledResult string) (*Response, error) {
	return &Response{
		DIDDocument: marshalledResult,
	}, nil
}

func (o Response) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(200)
	if err := producer.Produce(rw, o); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
