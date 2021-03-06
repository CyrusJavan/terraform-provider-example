// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pet API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pet API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddPet(params *AddPetParams) (*AddPetOK, error)

	DeletePet(params *DeletePetParams) (*DeletePetOK, error)

	GetPetByID(params *GetPetByIDParams) (*GetPetByIDOK, error)

	UpdatePet(params *UpdatePetParams) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddPet adds a new pet to the store
*/
func (a *Client) AddPet(params *AddPetParams) (*AddPetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addPet",
		Method:             "POST",
		PathPattern:        "/pet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddPetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddPetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addPet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePet deletes a pet
*/
func (a *Client) DeletePet(params *DeletePetParams) (*DeletePetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePet",
		Method:             "DELETE",
		PathPattern:        "/pet/{petId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePet: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPetByID finds pet by ID

  Returns a single pet
*/
func (a *Client) GetPetByID(params *GetPetByIDParams) (*GetPetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPetById",
		Method:             "GET",
		PathPattern:        "/pet/{petId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPetByIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPetByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPetById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePet updates an existing pet
*/
func (a *Client) UpdatePet(params *UpdatePetParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePetParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePet",
		Method:             "PUT",
		PathPattern:        "/pet",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePetReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
