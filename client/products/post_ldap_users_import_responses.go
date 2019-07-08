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

// PostLdapUsersImportReader is a Reader for the PostLdapUsersImport structure.
type PostLdapUsersImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLdapUsersImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostLdapUsersImportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostLdapUsersImportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostLdapUsersImportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostLdapUsersImportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 415:
		result := NewPostLdapUsersImportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLdapUsersImportOK creates a PostLdapUsersImportOK with default headers values
func NewPostLdapUsersImportOK() *PostLdapUsersImportOK {
	return &PostLdapUsersImportOK{}
}

/*PostLdapUsersImportOK handles this case with default header values.

Add ldap users successfully.
*/
type PostLdapUsersImportOK struct {
}

func (o *PostLdapUsersImportOK) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] postLdapUsersImportOK ", 200)
}

func (o *PostLdapUsersImportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapUsersImportUnauthorized creates a PostLdapUsersImportUnauthorized with default headers values
func NewPostLdapUsersImportUnauthorized() *PostLdapUsersImportUnauthorized {
	return &PostLdapUsersImportUnauthorized{}
}

/*PostLdapUsersImportUnauthorized handles this case with default header values.

User need to login first.
*/
type PostLdapUsersImportUnauthorized struct {
}

func (o *PostLdapUsersImportUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] postLdapUsersImportUnauthorized ", 401)
}

func (o *PostLdapUsersImportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapUsersImportForbidden creates a PostLdapUsersImportForbidden with default headers values
func NewPostLdapUsersImportForbidden() *PostLdapUsersImportForbidden {
	return &PostLdapUsersImportForbidden{}
}

/*PostLdapUsersImportForbidden handles this case with default header values.

Only admin has this authority.
*/
type PostLdapUsersImportForbidden struct {
}

func (o *PostLdapUsersImportForbidden) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] postLdapUsersImportForbidden ", 403)
}

func (o *PostLdapUsersImportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostLdapUsersImportNotFound creates a PostLdapUsersImportNotFound with default headers values
func NewPostLdapUsersImportNotFound() *PostLdapUsersImportNotFound {
	return &PostLdapUsersImportNotFound{}
}

/*PostLdapUsersImportNotFound handles this case with default header values.

Failed import some users.
*/
type PostLdapUsersImportNotFound struct {
	Payload []*models.LdapFailedImportUsers
}

func (o *PostLdapUsersImportNotFound) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] postLdapUsersImportNotFound  %+v", 404, o.Payload)
}

func (o *PostLdapUsersImportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLdapUsersImportUnsupportedMediaType creates a PostLdapUsersImportUnsupportedMediaType with default headers values
func NewPostLdapUsersImportUnsupportedMediaType() *PostLdapUsersImportUnsupportedMediaType {
	return &PostLdapUsersImportUnsupportedMediaType{}
}

/*PostLdapUsersImportUnsupportedMediaType handles this case with default header values.

The Media Type of the request is not supported, it has to be "application/json"
*/
type PostLdapUsersImportUnsupportedMediaType struct {
}

func (o *PostLdapUsersImportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /ldap/users/import][%d] postLdapUsersImportUnsupportedMediaType ", 415)
}

func (o *PostLdapUsersImportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
