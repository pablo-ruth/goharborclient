// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRepositoriesRepoNameLabelsLabelIDReader is a Reader for the DeleteRepositoriesRepoNameLabelsLabelID structure.
type DeleteRepositoriesRepoNameLabelsLabelIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoriesRepoNameLabelsLabelIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteRepositoriesRepoNameLabelsLabelIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteRepositoriesRepoNameLabelsLabelIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteRepositoriesRepoNameLabelsLabelIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteRepositoriesRepoNameLabelsLabelIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRepositoriesRepoNameLabelsLabelIDOK creates a DeleteRepositoriesRepoNameLabelsLabelIDOK with default headers values
func NewDeleteRepositoriesRepoNameLabelsLabelIDOK() *DeleteRepositoriesRepoNameLabelsLabelIDOK {
	return &DeleteRepositoriesRepoNameLabelsLabelIDOK{}
}

/*DeleteRepositoriesRepoNameLabelsLabelIDOK handles this case with default header values.

Successfully.
*/
type DeleteRepositoriesRepoNameLabelsLabelIDOK struct {
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDOK) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}/labels/{label_id}][%d] deleteRepositoriesRepoNameLabelsLabelIdOK ", 200)
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameLabelsLabelIDUnauthorized creates a DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized with default headers values
func NewDeleteRepositoriesRepoNameLabelsLabelIDUnauthorized() *DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized {
	return &DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized{}
}

/*DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized handles this case with default header values.

Unauthorized.
*/
type DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized struct {
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}/labels/{label_id}][%d] deleteRepositoriesRepoNameLabelsLabelIdUnauthorized ", 401)
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameLabelsLabelIDForbidden creates a DeleteRepositoriesRepoNameLabelsLabelIDForbidden with default headers values
func NewDeleteRepositoriesRepoNameLabelsLabelIDForbidden() *DeleteRepositoriesRepoNameLabelsLabelIDForbidden {
	return &DeleteRepositoriesRepoNameLabelsLabelIDForbidden{}
}

/*DeleteRepositoriesRepoNameLabelsLabelIDForbidden handles this case with default header values.

Forbidden. User should have write permisson for the repository to perform the action.
*/
type DeleteRepositoriesRepoNameLabelsLabelIDForbidden struct {
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}/labels/{label_id}][%d] deleteRepositoriesRepoNameLabelsLabelIdForbidden ", 403)
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRepositoriesRepoNameLabelsLabelIDNotFound creates a DeleteRepositoriesRepoNameLabelsLabelIDNotFound with default headers values
func NewDeleteRepositoriesRepoNameLabelsLabelIDNotFound() *DeleteRepositoriesRepoNameLabelsLabelIDNotFound {
	return &DeleteRepositoriesRepoNameLabelsLabelIDNotFound{}
}

/*DeleteRepositoriesRepoNameLabelsLabelIDNotFound handles this case with default header values.

Resource not found.
*/
type DeleteRepositoriesRepoNameLabelsLabelIDNotFound struct {
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /repositories/{repo_name}/labels/{label_id}][%d] deleteRepositoriesRepoNameLabelsLabelIdNotFound ", 404)
}

func (o *DeleteRepositoriesRepoNameLabelsLabelIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
