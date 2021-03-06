// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePetReader is a Reader for the DeletePet structure.
type DeletePetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeletePetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePetOK creates a DeletePetOK with default headers values
func NewDeletePetOK() *DeletePetOK {
	return &DeletePetOK{}
}

/*DeletePetOK handles this case with default header values.

success
*/
type DeletePetOK struct {
}

func (o *DeletePetOK) Error() string {
	return fmt.Sprintf("[DELETE /pet/{petId}][%d] deletePetOK ", 200)
}

func (o *DeletePetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePetBadRequest creates a DeletePetBadRequest with default headers values
func NewDeletePetBadRequest() *DeletePetBadRequest {
	return &DeletePetBadRequest{}
}

/*DeletePetBadRequest handles this case with default header values.

Invalid ID supplied
*/
type DeletePetBadRequest struct {
}

func (o *DeletePetBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pet/{petId}][%d] deletePetBadRequest ", 400)
}

func (o *DeletePetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePetNotFound creates a DeletePetNotFound with default headers values
func NewDeletePetNotFound() *DeletePetNotFound {
	return &DeletePetNotFound{}
}

/*DeletePetNotFound handles this case with default header values.

Pet not found
*/
type DeletePetNotFound struct {
}

func (o *DeletePetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pet/{petId}][%d] deletePetNotFound ", 404)
}

func (o *DeletePetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
