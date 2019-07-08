// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pablo-ruth/goharborclient/models"
)

// NewPostProjectsProjectIDMembersParams creates a new PostProjectsProjectIDMembersParams object
// with the default values initialized.
func NewPostProjectsProjectIDMembersParams() *PostProjectsProjectIDMembersParams {
	var ()
	return &PostProjectsProjectIDMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostProjectsProjectIDMembersParamsWithTimeout creates a new PostProjectsProjectIDMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostProjectsProjectIDMembersParamsWithTimeout(timeout time.Duration) *PostProjectsProjectIDMembersParams {
	var ()
	return &PostProjectsProjectIDMembersParams{

		timeout: timeout,
	}
}

// NewPostProjectsProjectIDMembersParamsWithContext creates a new PostProjectsProjectIDMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostProjectsProjectIDMembersParamsWithContext(ctx context.Context) *PostProjectsProjectIDMembersParams {
	var ()
	return &PostProjectsProjectIDMembersParams{

		Context: ctx,
	}
}

// NewPostProjectsProjectIDMembersParamsWithHTTPClient creates a new PostProjectsProjectIDMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostProjectsProjectIDMembersParamsWithHTTPClient(client *http.Client) *PostProjectsProjectIDMembersParams {
	var ()
	return &PostProjectsProjectIDMembersParams{
		HTTPClient: client,
	}
}

/*PostProjectsProjectIDMembersParams contains all the parameters to send to the API endpoint
for the post projects project ID members operation typically these are written to a http.Request
*/
type PostProjectsProjectIDMembersParams struct {

	/*ProjectID
	  Relevant project ID.

	*/
	ProjectID int64
	/*ProjectMember*/
	ProjectMember *models.ProjectMember

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) WithTimeout(timeout time.Duration) *PostProjectsProjectIDMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) WithContext(ctx context.Context) *PostProjectsProjectIDMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) WithHTTPClient(client *http.Client) *PostProjectsProjectIDMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) WithProjectID(projectID int64) *PostProjectsProjectIDMembersParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WithProjectMember adds the projectMember to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) WithProjectMember(projectMember *models.ProjectMember) *PostProjectsProjectIDMembersParams {
	o.SetProjectMember(projectMember)
	return o
}

// SetProjectMember adds the projectMember to the post projects project ID members params
func (o *PostProjectsProjectIDMembersParams) SetProjectMember(projectMember *models.ProjectMember) {
	o.ProjectMember = projectMember
}

// WriteToRequest writes these params to a swagger request
func (o *PostProjectsProjectIDMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if o.ProjectMember != nil {
		if err := r.SetBodyParam(o.ProjectMember); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
