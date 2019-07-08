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

// NewPostChartrepoChartsParams creates a new PostChartrepoChartsParams object
// with the default values initialized.
func NewPostChartrepoChartsParams() *PostChartrepoChartsParams {
	var ()
	return &PostChartrepoChartsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostChartrepoChartsParamsWithTimeout creates a new PostChartrepoChartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostChartrepoChartsParamsWithTimeout(timeout time.Duration) *PostChartrepoChartsParams {
	var ()
	return &PostChartrepoChartsParams{

		timeout: timeout,
	}
}

// NewPostChartrepoChartsParamsWithContext creates a new PostChartrepoChartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostChartrepoChartsParamsWithContext(ctx context.Context) *PostChartrepoChartsParams {
	var ()
	return &PostChartrepoChartsParams{

		Context: ctx,
	}
}

// NewPostChartrepoChartsParamsWithHTTPClient creates a new PostChartrepoChartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostChartrepoChartsParamsWithHTTPClient(client *http.Client) *PostChartrepoChartsParams {
	var ()
	return &PostChartrepoChartsParams{
		HTTPClient: client,
	}
}

/*PostChartrepoChartsParams contains all the parameters to send to the API endpoint
for the post chartrepo charts operation typically these are written to a http.Request
*/
type PostChartrepoChartsParams struct {

	/*Chart
	  The chart file

	*/
	Chart runtime.NamedReadCloser
	/*Prov
	  The provance file

	*/
	Prov runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post chartrepo charts params
func (o *PostChartrepoChartsParams) WithTimeout(timeout time.Duration) *PostChartrepoChartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post chartrepo charts params
func (o *PostChartrepoChartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post chartrepo charts params
func (o *PostChartrepoChartsParams) WithContext(ctx context.Context) *PostChartrepoChartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post chartrepo charts params
func (o *PostChartrepoChartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post chartrepo charts params
func (o *PostChartrepoChartsParams) WithHTTPClient(client *http.Client) *PostChartrepoChartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post chartrepo charts params
func (o *PostChartrepoChartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChart adds the chart to the post chartrepo charts params
func (o *PostChartrepoChartsParams) WithChart(chart runtime.NamedReadCloser) *PostChartrepoChartsParams {
	o.SetChart(chart)
	return o
}

// SetChart adds the chart to the post chartrepo charts params
func (o *PostChartrepoChartsParams) SetChart(chart runtime.NamedReadCloser) {
	o.Chart = chart
}

// WithProv adds the prov to the post chartrepo charts params
func (o *PostChartrepoChartsParams) WithProv(prov runtime.NamedReadCloser) *PostChartrepoChartsParams {
	o.SetProv(prov)
	return o
}

// SetProv adds the prov to the post chartrepo charts params
func (o *PostChartrepoChartsParams) SetProv(prov runtime.NamedReadCloser) {
	o.Prov = prov
}

// WriteToRequest writes these params to a swagger request
func (o *PostChartrepoChartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param chart
	if err := r.SetFileParam("chart", o.Chart); err != nil {
		return err
	}

	if o.Prov != nil {

		if o.Prov != nil {

			// form file param prov
			if err := r.SetFileParam("prov", o.Prov); err != nil {
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}