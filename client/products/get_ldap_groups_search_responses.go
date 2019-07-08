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

// GetLdapGroupsSearchReader is a Reader for the GetLdapGroupsSearch structure.
type GetLdapGroupsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapGroupsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLdapGroupsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetLdapGroupsSearchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetLdapGroupsSearchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetLdapGroupsSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLdapGroupsSearchOK creates a GetLdapGroupsSearchOK with default headers values
func NewGetLdapGroupsSearchOK() *GetLdapGroupsSearchOK {
	return &GetLdapGroupsSearchOK{}
}

/*GetLdapGroupsSearchOK handles this case with default header values.

Search ldap group successfully.
*/
type GetLdapGroupsSearchOK struct {
	Payload []*models.UserGroup
}

func (o *GetLdapGroupsSearchOK) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] getLdapGroupsSearchOK  %+v", 200, o.Payload)
}

func (o *GetLdapGroupsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLdapGroupsSearchBadRequest creates a GetLdapGroupsSearchBadRequest with default headers values
func NewGetLdapGroupsSearchBadRequest() *GetLdapGroupsSearchBadRequest {
	return &GetLdapGroupsSearchBadRequest{}
}

/*GetLdapGroupsSearchBadRequest handles this case with default header values.

The Ldap group DN is invalid.
*/
type GetLdapGroupsSearchBadRequest struct {
}

func (o *GetLdapGroupsSearchBadRequest) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] getLdapGroupsSearchBadRequest ", 400)
}

func (o *GetLdapGroupsSearchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapGroupsSearchNotFound creates a GetLdapGroupsSearchNotFound with default headers values
func NewGetLdapGroupsSearchNotFound() *GetLdapGroupsSearchNotFound {
	return &GetLdapGroupsSearchNotFound{}
}

/*GetLdapGroupsSearchNotFound handles this case with default header values.

No ldap group found.
*/
type GetLdapGroupsSearchNotFound struct {
}

func (o *GetLdapGroupsSearchNotFound) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] getLdapGroupsSearchNotFound ", 404)
}

func (o *GetLdapGroupsSearchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLdapGroupsSearchInternalServerError creates a GetLdapGroupsSearchInternalServerError with default headers values
func NewGetLdapGroupsSearchInternalServerError() *GetLdapGroupsSearchInternalServerError {
	return &GetLdapGroupsSearchInternalServerError{}
}

/*GetLdapGroupsSearchInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type GetLdapGroupsSearchInternalServerError struct {
}

func (o *GetLdapGroupsSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] getLdapGroupsSearchInternalServerError ", 500)
}

func (o *GetLdapGroupsSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}