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

// NewPutUsersUserIDPasswordParams creates a new PutUsersUserIDPasswordParams object
// with the default values initialized.
func NewPutUsersUserIDPasswordParams() *PutUsersUserIDPasswordParams {
	var ()
	return &PutUsersUserIDPasswordParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutUsersUserIDPasswordParamsWithTimeout creates a new PutUsersUserIDPasswordParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutUsersUserIDPasswordParamsWithTimeout(timeout time.Duration) *PutUsersUserIDPasswordParams {
	var ()
	return &PutUsersUserIDPasswordParams{

		timeout: timeout,
	}
}

// NewPutUsersUserIDPasswordParamsWithContext creates a new PutUsersUserIDPasswordParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutUsersUserIDPasswordParamsWithContext(ctx context.Context) *PutUsersUserIDPasswordParams {
	var ()
	return &PutUsersUserIDPasswordParams{

		Context: ctx,
	}
}

// NewPutUsersUserIDPasswordParamsWithHTTPClient creates a new PutUsersUserIDPasswordParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutUsersUserIDPasswordParamsWithHTTPClient(client *http.Client) *PutUsersUserIDPasswordParams {
	var ()
	return &PutUsersUserIDPasswordParams{
		HTTPClient: client,
	}
}

/*PutUsersUserIDPasswordParams contains all the parameters to send to the API endpoint
for the put users user ID password operation typically these are written to a http.Request
*/
type PutUsersUserIDPasswordParams struct {

	/*Password
	  Password to be updated, the attribute 'old_password' is optional when the API is called by the system administrator.

	*/
	Password *models.Password
	/*UserID
	  Registered user ID.

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) WithTimeout(timeout time.Duration) *PutUsersUserIDPasswordParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) WithContext(ctx context.Context) *PutUsersUserIDPasswordParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) WithHTTPClient(client *http.Client) *PutUsersUserIDPasswordParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPassword adds the password to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) WithPassword(password *models.Password) *PutUsersUserIDPasswordParams {
	o.SetPassword(password)
	return o
}

// SetPassword adds the password to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) SetPassword(password *models.Password) {
	o.Password = password
}

// WithUserID adds the userID to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) WithUserID(userID int64) *PutUsersUserIDPasswordParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put users user ID password params
func (o *PutUsersUserIDPasswordParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutUsersUserIDPasswordParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Password != nil {
		if err := r.SetBodyParam(o.Password); err != nil {
			return err
		}
	}

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
