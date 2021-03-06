// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/CyrusJavan/terraform-provider-example/example-server/models"
)

// AddPetOKCode is the HTTP code returned for type AddPetOK
const AddPetOKCode int = 200

/*AddPetOK successfully added

swagger:response addPetOK
*/
type AddPetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pet `json:"body,omitempty"`
}

// NewAddPetOK creates AddPetOK with default headers values
func NewAddPetOK() *AddPetOK {

	return &AddPetOK{}
}

// WithPayload adds the payload to the add pet o k response
func (o *AddPetOK) WithPayload(payload *models.Pet) *AddPetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add pet o k response
func (o *AddPetOK) SetPayload(payload *models.Pet) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddPetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddPetMethodNotAllowedCode is the HTTP code returned for type AddPetMethodNotAllowed
const AddPetMethodNotAllowedCode int = 405

/*AddPetMethodNotAllowed Invalid input

swagger:response addPetMethodNotAllowed
*/
type AddPetMethodNotAllowed struct {
}

// NewAddPetMethodNotAllowed creates AddPetMethodNotAllowed with default headers values
func NewAddPetMethodNotAllowed() *AddPetMethodNotAllowed {

	return &AddPetMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddPetMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}
