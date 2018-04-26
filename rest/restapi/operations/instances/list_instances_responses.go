// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/subutai-io/p2p/rest/models"
)

// ListInstancesOKCode is the HTTP code returned for type ListInstancesOK
const ListInstancesOKCode int = 200

/*ListInstancesOK Sucessful operation

swagger:response listInstancesOK
*/
type ListInstancesOK struct {

	/*
	  In: Body
	*/
	Payload models.Instances `json:"body,omitempty"`
}

// NewListInstancesOK creates ListInstancesOK with default headers values
func NewListInstancesOK() *ListInstancesOK {

	return &ListInstancesOK{}
}

// WithPayload adds the payload to the list instances o k response
func (o *ListInstancesOK) WithPayload(payload models.Instances) *ListInstancesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list instances o k response
func (o *ListInstancesOK) SetPayload(payload models.Instances) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListInstancesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.Instances, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ListInstancesServiceUnavailableCode is the HTTP code returned for type ListInstancesServiceUnavailable
const ListInstancesServiceUnavailableCode int = 503

/*ListInstancesServiceUnavailable Service unavailable

swagger:response listInstancesServiceUnavailable
*/
type ListInstancesServiceUnavailable struct {
}

// NewListInstancesServiceUnavailable creates ListInstancesServiceUnavailable with default headers values
func NewListInstancesServiceUnavailable() *ListInstancesServiceUnavailable {

	return &ListInstancesServiceUnavailable{}
}

// WriteResponse to the client
func (o *ListInstancesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(503)
}
