package j_tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJTagFetchLastInteractorsIDParams creates a new PostRemoteAPIJTagFetchLastInteractorsIDParams object
// with the default values initialized.
func NewPostRemoteAPIJTagFetchLastInteractorsIDParams() *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	var ()
	return &PostRemoteAPIJTagFetchLastInteractorsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJTagFetchLastInteractorsIDParamsWithTimeout creates a new PostRemoteAPIJTagFetchLastInteractorsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJTagFetchLastInteractorsIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	var ()
	return &PostRemoteAPIJTagFetchLastInteractorsIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJTagFetchLastInteractorsIDParamsWithContext creates a new PostRemoteAPIJTagFetchLastInteractorsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJTagFetchLastInteractorsIDParamsWithContext(ctx context.Context) *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	var ()
	return &PostRemoteAPIJTagFetchLastInteractorsIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJTagFetchLastInteractorsIDParams contains all the parameters to send to the API endpoint
for the post remote API j tag fetch last interactors ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJTagFetchLastInteractorsIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) WithContext(ctx context.Context) *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) WithID(id string) *PostRemoteAPIJTagFetchLastInteractorsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j tag fetch last interactors ID params
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJTagFetchLastInteractorsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
