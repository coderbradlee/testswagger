package handler

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/lzxm160/testswagger/models"
)

type UpdateResponse struct {

	// In: body
	Payload *models.Jsonrpc `json:"body,omitempty"`
}

// Handler to the client
func (o UpdateResponse) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
