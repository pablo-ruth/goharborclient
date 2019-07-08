// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetChartrepoHealthParams creates a new GetChartrepoHealthParams object
// with the default values initialized.
func NewGetChartrepoHealthParams() *GetChartrepoHealthParams {

	return &GetChartrepoHealthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChartrepoHealthParamsWithTimeout creates a new GetChartrepoHealthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChartrepoHealthParamsWithTimeout(timeout time.Duration) *GetChartrepoHealthParams {

	return &GetChartrepoHealthParams{

		timeout: timeout,
	}
}

// NewGetChartrepoHealthParamsWithContext creates a new GetChartrepoHealthParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChartrepoHealthParamsWithContext(ctx context.Context) *GetChartrepoHealthParams {

	return &GetChartrepoHealthParams{

		Context: ctx,
	}
}

// NewGetChartrepoHealthParamsWithHTTPClient creates a new GetChartrepoHealthParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChartrepoHealthParamsWithHTTPClient(client *http.Client) *GetChartrepoHealthParams {

	return &GetChartrepoHealthParams{
		HTTPClient: client,
	}
}

/*GetChartrepoHealthParams contains all the parameters to send to the API endpoint
for the get chartrepo health operation typically these are written to a http.Request
*/
type GetChartrepoHealthParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get chartrepo health params
func (o *GetChartrepoHealthParams) WithTimeout(timeout time.Duration) *GetChartrepoHealthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chartrepo health params
func (o *GetChartrepoHealthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chartrepo health params
func (o *GetChartrepoHealthParams) WithContext(ctx context.Context) *GetChartrepoHealthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chartrepo health params
func (o *GetChartrepoHealthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chartrepo health params
func (o *GetChartrepoHealthParams) WithHTTPClient(client *http.Client) *GetChartrepoHealthParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chartrepo health params
func (o *GetChartrepoHealthParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetChartrepoHealthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
