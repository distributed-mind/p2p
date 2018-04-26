// Code generated by go-swagger; DO NOT EDIT.

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DaemonOptionsHandlerFunc turns a function with the right signature into a daemon options handler
type DaemonOptionsHandlerFunc func(DaemonOptionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DaemonOptionsHandlerFunc) Handle(params DaemonOptionsParams) middleware.Responder {
	return fn(params)
}

// DaemonOptionsHandler interface for that can handle valid daemon options params
type DaemonOptionsHandler interface {
	Handle(DaemonOptionsParams) middleware.Responder
}

// NewDaemonOptions creates a new http.Handler for the daemon options operation
func NewDaemonOptions(ctx *middleware.Context, handler DaemonOptionsHandler) *DaemonOptions {
	return &DaemonOptions{Context: ctx, Handler: handler}
}

/*DaemonOptions swagger:route POST /daemon daemon daemonOptions

Modify daemon

Modify daemon options on runtime

*/
type DaemonOptions struct {
	Context *middleware.Context
	Handler DaemonOptionsHandler
}

func (o *DaemonOptions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDaemonOptionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
