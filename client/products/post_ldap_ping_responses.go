// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLdapPingReader is a Reader for the PostLdapPing structure.
type PostLdapPingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLdapPingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLdapPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostLdapPingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostLdapPingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostLdapPingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewPostLdapPingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostLdapPingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLdapPingOK creates a PostLdapPingOK with default headers values
func NewPostLdapPingOK() *PostLdapPingOK {
	return &PostLdapPingOK{}
}

/*PostLdapPingOK handles this case with default header values.

Ping ldap service successfully.
*/
type PostLdapPingOK struct {
}

func (o *PostLdapPingOK) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] postLdapPingOK ", 200)
}

func (o *PostLdapPingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapPingBadRequest creates a PostLdapPingBadRequest with default headers values
func NewPostLdapPingBadRequest() *PostLdapPingBadRequest {
	return &PostLdapPingBadRequest{}
}

/*PostLdapPingBadRequest handles this case with default header values.

Inviald ldap configuration parameters.
*/
type PostLdapPingBadRequest struct {
}

func (o *PostLdapPingBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] postLdapPingBadRequest ", 400)
}

func (o *PostLdapPingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapPingUnauthorized creates a PostLdapPingUnauthorized with default headers values
func NewPostLdapPingUnauthorized() *PostLdapPingUnauthorized {
	return &PostLdapPingUnauthorized{}
}

/*PostLdapPingUnauthorized handles this case with default header values.

User need to login first.
*/
type PostLdapPingUnauthorized struct {
}

func (o *PostLdapPingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] postLdapPingUnauthorized ", 401)
}

func (o *PostLdapPingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapPingForbidden creates a PostLdapPingForbidden with default headers values
func NewPostLdapPingForbidden() *PostLdapPingForbidden {
	return &PostLdapPingForbidden{}
}

/*PostLdapPingForbidden handles this case with default header values.

Only admin has this authority.
*/
type PostLdapPingForbidden struct {
}

func (o *PostLdapPingForbidden) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] postLdapPingForbidden ", 403)
}

func (o *PostLdapPingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapPingUnsupportedMediaType creates a PostLdapPingUnsupportedMediaType with default headers values
func NewPostLdapPingUnsupportedMediaType() *PostLdapPingUnsupportedMediaType {
	return &PostLdapPingUnsupportedMediaType{}
}

/*PostLdapPingUnsupportedMediaType handles this case with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type PostLdapPingUnsupportedMediaType struct {
}

func (o *PostLdapPingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] postLdapPingUnsupportedMediaType ", 415)
}

func (o *PostLdapPingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapPingInternalServerError creates a PostLdapPingInternalServerError with default headers values
func NewPostLdapPingInternalServerError() *PostLdapPingInternalServerError {
	return &PostLdapPingInternalServerError{}
}

/*PostLdapPingInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PostLdapPingInternalServerError struct {
}

func (o *PostLdapPingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldap/ping][%d] postLdapPingInternalServerError ", 500)
}

func (o *PostLdapPingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
