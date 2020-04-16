package routing

import (
	"net/http"

	"github.com/niclic/helloworld/internal/handlers"
)

// NewHTTPRouter returns a new HTTPRouter struct with a http.ServeMux.
func NewHTTPRouter() *HTTPRouter {
	return &HTTPRouter{http.NewServeMux()}
}

// HTTPRouter is a wrapper around http.ServeMux.
type HTTPRouter struct {
	mux *http.ServeMux
}

// Mux returns the http.ServeMux associated with this HTTPRouter.
func (r *HTTPRouter) Mux() http.Handler {
	return r.mux
}

// RegisterRoutes iterates over a map of route patterns and handlers,
// and registers each pattern/handler pair with the http.ServeMux.
func (r *HTTPRouter) RegisterRoutes(routes map[string]handlers.HTTPHandler) {
	for p, h := range routes {
		r.mux.HandleFunc(p, h)
	}
}
