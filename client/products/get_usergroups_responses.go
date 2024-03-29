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

// GetUsergroupsReader is a Reader for the GetUsergroups structure.
type GetUsergroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsergroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsergroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUsergroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetUsergroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUsergroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsergroupsOK creates a GetUsergroupsOK with default headers values
func NewGetUsergroupsOK() *GetUsergroupsOK {
	return &GetUsergroupsOK{}
}

/*GetUsergroupsOK handles this case with default header values.

Get user group successfully.
*/
type GetUsergroupsOK struct {
	Payload []*models.UserGroup
}

func (o *GetUsergroupsOK) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] getUsergroupsOK  %+v", 200, o.Payload)
}

func (o *GetUsergroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsergroupsUnauthorized creates a GetUsergroupsUnauthorized with default headers values
func NewGetUsergroupsUnauthorized() *GetUsergroupsUnauthorized {
	return &GetUsergroupsUnauthorized{}
}

/*GetUsergroupsUnauthorized handles this case with default header values.

User need to log in first.
*/
type GetUsergroupsUnauthorized struct {
}

func (o *GetUsergroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] getUsergroupsUnauthorized ", 401)
}

func (o *GetUsergroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsergroupsForbidden creates a GetUsergroupsForbidden with default headers values
func NewGetUsergroupsForbidden() *GetUsergroupsForbidden {
	return &GetUsergroupsForbidden{}
}

/*GetUsergroupsForbidden handles this case with default header values.

User in session does not have permission to the user group.
*/
type GetUsergroupsForbidden struct {
}

func (o *GetUsergroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] getUsergroupsForbidden ", 403)
}

func (o *GetUsergroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsergroupsInternalServerError creates a GetUsergroupsInternalServerError with default headers values
func NewGetUsergroupsInternalServerError() *GetUsergroupsInternalServerError {
	return &GetUsergroupsInternalServerError{}
}

/*GetUsergroupsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetUsergroupsInternalServerError struct {
}

func (o *GetUsergroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usergroups][%d] getUsergroupsInternalServerError ", 500)
}

func (o *GetUsergroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
