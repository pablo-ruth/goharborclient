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

// GetReplicationPoliciesReader is a Reader for the GetReplicationPolicies structure.
type GetReplicationPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReplicationPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetReplicationPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetReplicationPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetReplicationPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetReplicationPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReplicationPoliciesOK creates a GetReplicationPoliciesOK with default headers values
func NewGetReplicationPoliciesOK() *GetReplicationPoliciesOK {
	return &GetReplicationPoliciesOK{}
}

/*GetReplicationPoliciesOK handles this case with default header values.

Get policy successfully.
*/
type GetReplicationPoliciesOK struct {
	Payload []*models.ReplicationPolicy
}

func (o *GetReplicationPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] getReplicationPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetReplicationPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationPoliciesBadRequest creates a GetReplicationPoliciesBadRequest with default headers values
func NewGetReplicationPoliciesBadRequest() *GetReplicationPoliciesBadRequest {
	return &GetReplicationPoliciesBadRequest{}
}

/*GetReplicationPoliciesBadRequest handles this case with default header values.

Bad Request
*/
type GetReplicationPoliciesBadRequest struct {
}

func (o *GetReplicationPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] getReplicationPoliciesBadRequest ", 400)
}

func (o *GetReplicationPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesUnauthorized creates a GetReplicationPoliciesUnauthorized with default headers values
func NewGetReplicationPoliciesUnauthorized() *GetReplicationPoliciesUnauthorized {
	return &GetReplicationPoliciesUnauthorized{}
}

/*GetReplicationPoliciesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetReplicationPoliciesUnauthorized struct {
}

func (o *GetReplicationPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] getReplicationPoliciesUnauthorized ", 401)
}

func (o *GetReplicationPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesForbidden creates a GetReplicationPoliciesForbidden with default headers values
func NewGetReplicationPoliciesForbidden() *GetReplicationPoliciesForbidden {
	return &GetReplicationPoliciesForbidden{}
}

/*GetReplicationPoliciesForbidden handles this case with default header values.

Forbidden
*/
type GetReplicationPoliciesForbidden struct {
}

func (o *GetReplicationPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] getReplicationPoliciesForbidden ", 403)
}

func (o *GetReplicationPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReplicationPoliciesInternalServerError creates a GetReplicationPoliciesInternalServerError with default headers values
func NewGetReplicationPoliciesInternalServerError() *GetReplicationPoliciesInternalServerError {
	return &GetReplicationPoliciesInternalServerError{}
}

/*GetReplicationPoliciesInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetReplicationPoliciesInternalServerError struct {
}

func (o *GetReplicationPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/policies][%d] getReplicationPoliciesInternalServerError ", 500)
}

func (o *GetReplicationPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
