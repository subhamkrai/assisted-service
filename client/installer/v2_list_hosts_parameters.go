// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewV2ListHostsParams creates a new V2ListHostsParams object
// with the default values initialized.
func NewV2ListHostsParams() *V2ListHostsParams {
	var ()
	return &V2ListHostsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewV2ListHostsParamsWithTimeout creates a new V2ListHostsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewV2ListHostsParamsWithTimeout(timeout time.Duration) *V2ListHostsParams {
	var ()
	return &V2ListHostsParams{

		timeout: timeout,
	}
}

// NewV2ListHostsParamsWithContext creates a new V2ListHostsParams object
// with the default values initialized, and the ability to set a context for a request
func NewV2ListHostsParamsWithContext(ctx context.Context) *V2ListHostsParams {
	var ()
	return &V2ListHostsParams{

		Context: ctx,
	}
}

// NewV2ListHostsParamsWithHTTPClient creates a new V2ListHostsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewV2ListHostsParamsWithHTTPClient(client *http.Client) *V2ListHostsParams {
	var ()
	return &V2ListHostsParams{
		HTTPClient: client,
	}
}

/*V2ListHostsParams contains all the parameters to send to the API endpoint
for the v2 list hosts operation typically these are written to a http.Request
*/
type V2ListHostsParams struct {

	/*InfraEnvID
	  The InfraEnv that the hosts are asociated with.

	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the v2 list hosts params
func (o *V2ListHostsParams) WithTimeout(timeout time.Duration) *V2ListHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 list hosts params
func (o *V2ListHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 list hosts params
func (o *V2ListHostsParams) WithContext(ctx context.Context) *V2ListHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 list hosts params
func (o *V2ListHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 list hosts params
func (o *V2ListHostsParams) WithHTTPClient(client *http.Client) *V2ListHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 list hosts params
func (o *V2ListHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraEnvID adds the infraEnvID to the v2 list hosts params
func (o *V2ListHostsParams) WithInfraEnvID(infraEnvID strfmt.UUID) *V2ListHostsParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the v2 list hosts params
func (o *V2ListHostsParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *V2ListHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}