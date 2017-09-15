// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cilium/cilium/api/v1/models"
)

// GetEndpointIDConfigOKCode is the HTTP code returned for type GetEndpointIDConfigOK
const GetEndpointIDConfigOKCode int = 200

/*GetEndpointIDConfigOK Success

swagger:response getEndpointIdConfigOK
*/
type GetEndpointIDConfigOK struct {

	/*
	  In: Body
	*/
	Payload *models.Configuration `json:"body,omitempty"`
}

// NewGetEndpointIDConfigOK creates GetEndpointIDConfigOK with default headers values
func NewGetEndpointIDConfigOK() *GetEndpointIDConfigOK {
	return &GetEndpointIDConfigOK{}
}

// WithPayload adds the payload to the get endpoint Id config o k response
func (o *GetEndpointIDConfigOK) WithPayload(payload *models.Configuration) *GetEndpointIDConfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get endpoint Id config o k response
func (o *GetEndpointIDConfigOK) SetPayload(payload *models.Configuration) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEndpointIDConfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEndpointIDConfigNotFoundCode is the HTTP code returned for type GetEndpointIDConfigNotFound
const GetEndpointIDConfigNotFoundCode int = 404

/*GetEndpointIDConfigNotFound Endpoint not found

swagger:response getEndpointIdConfigNotFound
*/
type GetEndpointIDConfigNotFound struct {
}

// NewGetEndpointIDConfigNotFound creates GetEndpointIDConfigNotFound with default headers values
func NewGetEndpointIDConfigNotFound() *GetEndpointIDConfigNotFound {
	return &GetEndpointIDConfigNotFound{}
}

// WriteResponse to the client
func (o *GetEndpointIDConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}
