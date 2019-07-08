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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetConfigurationsParams creates a new GetConfigurationsParams object
// with the default values initialized.
func NewGetConfigurationsParams() *GetConfigurationsParams {

	return &GetConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConfigurationsParamsWithTimeout creates a new GetConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConfigurationsParamsWithTimeout(timeout time.Duration) *GetConfigurationsParams {

	return &GetConfigurationsParams{

		timeout: timeout,
	}
}

// NewGetConfigurationsParamsWithContext creates a new GetConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConfigurationsParamsWithContext(ctx context.Context) *GetConfigurationsParams {

	return &GetConfigurationsParams{

		Context: ctx,
	}
}

// NewGetConfigurationsParamsWithHTTPClient creates a new GetConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConfigurationsParamsWithHTTPClient(client *http.Client) *GetConfigurationsParams {

	return &GetConfigurationsParams{
		HTTPClient: client,
	}
}

/*GetConfigurationsParams contains all the parameters to send to the API endpoint
for the get configurations operation typically these are written to a http.Request
*/
type GetConfigurationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get configurations params
func (o *GetConfigurationsParams) WithTimeout(timeout time.Duration) *GetConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get configurations params
func (o *GetConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get configurations params
func (o *GetConfigurationsParams) WithContext(ctx context.Context) *GetConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get configurations params
func (o *GetConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get configurations params
func (o *GetConfigurationsParams) WithHTTPClient(client *http.Client) *GetConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get configurations params
func (o *GetConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
