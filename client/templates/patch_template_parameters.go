package templates

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

	"github.com/galthaus/swagger-test/models"
)

// NewPatchTemplateParams creates a new PatchTemplateParams object
// with the default values initialized.
func NewPatchTemplateParams() *PatchTemplateParams {
	var ()
	return &PatchTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTemplateParamsWithTimeout creates a new PatchTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchTemplateParamsWithTimeout(timeout time.Duration) *PatchTemplateParams {
	var ()
	return &PatchTemplateParams{

		timeout: timeout,
	}
}

// NewPatchTemplateParamsWithContext creates a new PatchTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchTemplateParamsWithContext(ctx context.Context) *PatchTemplateParams {
	var ()
	return &PatchTemplateParams{

		Context: ctx,
	}
}

/*PatchTemplateParams contains all the parameters to send to the API endpoint
for the patch template operation typically these are written to a http.Request
*/
type PatchTemplateParams struct {

	/*Body*/
	Body models.Patchinput
	/*UUID*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch template params
func (o *PatchTemplateParams) WithTimeout(timeout time.Duration) *PatchTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch template params
func (o *PatchTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch template params
func (o *PatchTemplateParams) WithContext(ctx context.Context) *PatchTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch template params
func (o *PatchTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the patch template params
func (o *PatchTemplateParams) WithBody(body models.Patchinput) *PatchTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch template params
func (o *PatchTemplateParams) SetBody(body models.Patchinput) {
	o.Body = body
}

// WithUUID adds the uuid to the patch template params
func (o *PatchTemplateParams) WithUUID(uuid string) *PatchTemplateParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the patch template params
func (o *PatchTemplateParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}