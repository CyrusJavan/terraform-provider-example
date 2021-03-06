// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdatePetHandlerFunc turns a function with the right signature into a update pet handler
type UpdatePetHandlerFunc func(UpdatePetParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdatePetHandlerFunc) Handle(params UpdatePetParams) middleware.Responder {
	return fn(params)
}

// UpdatePetHandler interface for that can handle valid update pet params
type UpdatePetHandler interface {
	Handle(UpdatePetParams) middleware.Responder
}

// NewUpdatePet creates a new http.Handler for the update pet operation
func NewUpdatePet(ctx *middleware.Context, handler UpdatePetHandler) *UpdatePet {
	return &UpdatePet{Context: ctx, Handler: handler}
}

/*UpdatePet swagger:route PUT /pet pet updatePet

Update an existing pet

*/
type UpdatePet struct {
	Context *middleware.Context
	Handler UpdatePetHandler
}

func (o *UpdatePet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdatePetParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
