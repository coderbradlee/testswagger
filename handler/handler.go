package handler

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lzxm160/testswagger/restapi/operations/update"
)

//type UpdateResponse struct {
//
//	// In: body
//	Payload *Response `json:"response,omitempty"`
//}
//
//// Handler to the client
//func (o UpdateResponse) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
//
//	rw.WriteHeader(200)
//	if o.Payload != nil {
//		if err := producer.Produce(rw, o.Payload); err != nil {
//			panic(err) // let the recovery middleware deal with this
//		}
//	}
//}

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
