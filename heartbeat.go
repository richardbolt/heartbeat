/*
Package heartbeat is a small module to use with a Goa app providing a heartbeat to ping for health check purposes.

By adding Heartbeat your app will respond with a 200 OK and a JSON object like so: `{"ENV":"production"}` on the url you specify or /health if no url is specified. Heartbeat returns a simple JSON object with one key, ENV, having the value of the ENV environment variable or "dev" if ENV has no value.

Heartbeat is not a middleware, instead it adds itself to service.Mux as an http handler.
*/
package heartbeat

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goadesign/goa"
	"golang.org/x/net/context"
)

// Heartbeat adds a standard response on a given url,
// which defaults to /health if url is an empty string.
func Heartbeat(service *goa.Service, url string) {
	if url == "" {
		url = "/health"
	}
	h := newHeartbeatController(service)
	mountHeartbeatController(service, h, url)
}

// heartbeat runs the heartbeat action.
func (c *actualHeartbeatController) heartbeat(ctx *heartbeatContext) error {
	// Heartbeat returns basic information.
	env := "dev"
	if os.Getenv("ENV") != "" {
		env = os.Getenv("ENV")
	}
	return ctx.OK([]byte(fmt.Sprintf(`{"ENV":"%s"}`, env)))
}

//-----------------------------
// Below is autowiring for Goa.
// We do all this stuff to get logging and other niceties Goa provides

// heartbeatContext provides the Heartbeat heartbeat action context.
type heartbeatContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewHeartbeatContext parses the incoming request URL and body, performs validations and creates the
// context used by the Heartbeat controller heartbeat action.
func newHeartbeatContext(ctx context.Context) (*heartbeatContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := heartbeatContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *heartbeatContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/json")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// heartbeatController is the controller interface for the Heartbeat actions.
type heartbeatController interface {
	goa.Muxer
	heartbeat(*heartbeatContext) error
}

// mountHeartbeatController "mounts" a Heartbeat resource controller on the given service.
func mountHeartbeatController(service *goa.Service, ctrl heartbeatController, url string) {
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := newHeartbeatContext(ctx)
		if err != nil {
			return err
		}
		return ctrl.heartbeat(rctx)
	}
	service.Mux.Handle("GET", url, ctrl.MuxHandler("Heartbeat", h, nil))
	service.LogInfo("mount", "ctrl", "Heartbeat", "action", "Heartbeat", "route", fmt.Sprintf("GET %s", url))
}

// actualHeartbeatController implements the Heartbeat resource.
type actualHeartbeatController struct {
	*goa.Controller
}

// newHeartbeatController creates a Heartbeat controller.
func newHeartbeatController(service *goa.Service) *actualHeartbeatController {
	return &actualHeartbeatController{Controller: service.NewController("Heartbeat")}
}
