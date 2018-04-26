// Code generated by go-swagger; DO NOT EDIT.

package instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CloseInstanceHandlerFunc turns a function with the right signature into a close instance handler
type CloseInstanceHandlerFunc func(CloseInstanceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CloseInstanceHandlerFunc) Handle(params CloseInstanceParams) middleware.Responder {
	return fn(params)
}

// CloseInstanceHandler interface for that can handle valid close instance params
type CloseInstanceHandler interface {
	Handle(CloseInstanceParams) middleware.Responder
}

// NewCloseInstance creates a new http.Handler for the close instance operation
func NewCloseInstance(ctx *middleware.Context, handler CloseInstanceHandler) *CloseInstance {
	return &CloseInstance{Context: ctx, Handler: handler}
}

/*CloseInstance swagger:route DELETE /instance instances closeInstance

Destroy P2P instance

This command will shutdown P2P instance

*/
type CloseInstance struct {
	Context *middleware.Context
	Handler CloseInstanceHandler
}

func (o *CloseInstance) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCloseInstanceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
