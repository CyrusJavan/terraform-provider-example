// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CyrusJavan/terraform-provider-example/models"
)

// GetPetByIDReader is a Reader for the GetPetByID structure.
type GetPetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPetByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPetByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPetByIDOK creates a GetPetByIDOK with default headers values
func NewGetPetByIDOK() *GetPetByIDOK {
	return &GetPetByIDOK{}
}

/*GetPetByIDOK handles this case with default header values.

successful operation
*/
type GetPetByIDOK struct {
	Payload *models.Pet
}

func (o *GetPetByIDOK) Error() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] getPetByIdOK  %+v", 200, o.Payload)
}

func (o *GetPetByIDOK) GetPayload() *models.Pet {
	return o.Payload
}

func (o *GetPetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPetByIDBadRequest creates a GetPetByIDBadRequest with default headers values
func NewGetPetByIDBadRequest() *GetPetByIDBadRequest {
	return &GetPetByIDBadRequest{}
}

/*GetPetByIDBadRequest handles this case with default header values.

Invalid ID supplied
*/
type GetPetByIDBadRequest struct {
}

func (o *GetPetByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] getPetByIdBadRequest ", 400)
}

func (o *GetPetByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPetByIDNotFound creates a GetPetByIDNotFound with default headers values
func NewGetPetByIDNotFound() *GetPetByIDNotFound {
	return &GetPetByIDNotFound{}
}

/*GetPetByIDNotFound handles this case with default header values.

Pet not found
*/
type GetPetByIDNotFound struct {
}

func (o *GetPetByIDNotFound) Error() string {
	return fmt.Sprintf("[GET /pet/{petId}][%d] getPetByIdNotFound ", 404)
}

func (o *GetPetByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
