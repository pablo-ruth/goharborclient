// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pablo-ruth/goharborclient/models"
)

// GetLabelsIDReader is a Reader for the GetLabelsID structure.
type GetLabelsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLabelsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLabelsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetLabelsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLabelsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetLabelsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLabelsIDOK creates a GetLabelsIDOK with default headers values
func NewGetLabelsIDOK() *GetLabelsIDOK {
	return &GetLabelsIDOK{}
}

/*GetLabelsIDOK handles this case with default header values.

Get successfully.
*/
type GetLabelsIDOK struct {
	Payload *models.Label
}

func (o *GetLabelsIDOK) Error() string {
	return fmt.Sprintf("[GET /labels/{id}][%d] getLabelsIdOK  %+v", 200, o.Payload)
}

func (o *GetLabelsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Label)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLabelsIDUnauthorized creates a GetLabelsIDUnauthorized with default headers values
func NewGetLabelsIDUnauthorized() *GetLabelsIDUnauthorized {
	return &GetLabelsIDUnauthorized{}
}

/*GetLabelsIDUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetLabelsIDUnauthorized struct {
}

func (o *GetLabelsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /labels/{id}][%d] getLabelsIdUnauthorized ", 401)
}

func (o *GetLabelsIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsIDNotFound creates a GetLabelsIDNotFound with default headers values
func NewGetLabelsIDNotFound() *GetLabelsIDNotFound {
	return &GetLabelsIDNotFound{}
}

/*GetLabelsIDNotFound handles this case with default header values.

The resource does not exist.
*/
type GetLabelsIDNotFound struct {
}

func (o *GetLabelsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /labels/{id}][%d] getLabelsIdNotFound ", 404)
}

func (o *GetLabelsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLabelsIDInternalServerError creates a GetLabelsIDInternalServerError with default headers values
func NewGetLabelsIDInternalServerError() *GetLabelsIDInternalServerError {
	return &GetLabelsIDInternalServerError{}
}

/*GetLabelsIDInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetLabelsIDInternalServerError struct {
}

func (o *GetLabelsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /labels/{id}][%d] getLabelsIdInternalServerError ", 500)
}

func (o *GetLabelsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
