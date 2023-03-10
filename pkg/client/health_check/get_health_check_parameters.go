// Code generated by go-swagger; DO NOT EDIT.

package health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetHealthCheckParams creates a new GetHealthCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHealthCheckParams() *GetHealthCheckParams {
	return &GetHealthCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHealthCheckParamsWithTimeout creates a new GetHealthCheckParams object
// with the ability to set a timeout on a request.
func NewGetHealthCheckParamsWithTimeout(timeout time.Duration) *GetHealthCheckParams {
	return &GetHealthCheckParams{
		timeout: timeout,
	}
}

// NewGetHealthCheckParamsWithContext creates a new GetHealthCheckParams object
// with the ability to set a context for a request.
func NewGetHealthCheckParamsWithContext(ctx context.Context) *GetHealthCheckParams {
	return &GetHealthCheckParams{
		Context: ctx,
	}
}

// NewGetHealthCheckParamsWithHTTPClient creates a new GetHealthCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHealthCheckParamsWithHTTPClient(client *http.Client) *GetHealthCheckParams {
	return &GetHealthCheckParams{
		HTTPClient: client,
	}
}

/*
GetHealthCheckParams contains all the parameters to send to the API endpoint

	for the get health check operation.

	Typically these are written to a http.Request.
*/
type GetHealthCheckParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHealthCheckParams) WithDefaults() *GetHealthCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get health check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHealthCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get health check params
func (o *GetHealthCheckParams) WithTimeout(timeout time.Duration) *GetHealthCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get health check params
func (o *GetHealthCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get health check params
func (o *GetHealthCheckParams) WithContext(ctx context.Context) *GetHealthCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get health check params
func (o *GetHealthCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get health check params
func (o *GetHealthCheckParams) WithHTTPClient(client *http.Client) *GetHealthCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get health check params
func (o *GetHealthCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHealthCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
