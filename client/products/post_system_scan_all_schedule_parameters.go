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

	models "github.com/pablo-ruth/goharborclient/models"
)

// NewPostSystemScanAllScheduleParams creates a new PostSystemScanAllScheduleParams object
// with the default values initialized.
func NewPostSystemScanAllScheduleParams() *PostSystemScanAllScheduleParams {
	var ()
	return &PostSystemScanAllScheduleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSystemScanAllScheduleParamsWithTimeout creates a new PostSystemScanAllScheduleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSystemScanAllScheduleParamsWithTimeout(timeout time.Duration) *PostSystemScanAllScheduleParams {
	var ()
	return &PostSystemScanAllScheduleParams{

		timeout: timeout,
	}
}

// NewPostSystemScanAllScheduleParamsWithContext creates a new PostSystemScanAllScheduleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSystemScanAllScheduleParamsWithContext(ctx context.Context) *PostSystemScanAllScheduleParams {
	var ()
	return &PostSystemScanAllScheduleParams{

		Context: ctx,
	}
}

// NewPostSystemScanAllScheduleParamsWithHTTPClient creates a new PostSystemScanAllScheduleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSystemScanAllScheduleParamsWithHTTPClient(client *http.Client) *PostSystemScanAllScheduleParams {
	var ()
	return &PostSystemScanAllScheduleParams{
		HTTPClient: client,
	}
}

/*PostSystemScanAllScheduleParams contains all the parameters to send to the API endpoint
for the post system scan all schedule operation typically these are written to a http.Request
*/
type PostSystemScanAllScheduleParams struct {

	/*Schedule
	  Create a schedule or a manual trigger for the scan all job.

	*/
	Schedule *models.AdminJobSchedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) WithTimeout(timeout time.Duration) *PostSystemScanAllScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) WithContext(ctx context.Context) *PostSystemScanAllScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) WithHTTPClient(client *http.Client) *PostSystemScanAllScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchedule adds the schedule to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) WithSchedule(schedule *models.AdminJobSchedule) *PostSystemScanAllScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the post system scan all schedule params
func (o *PostSystemScanAllScheduleParams) SetSchedule(schedule *models.AdminJobSchedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *PostSystemScanAllScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
