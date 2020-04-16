package routing

import (
	"net/http"

	"github.com/niclic/helloworld/internal/handlers"
)

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{http.NewServeMux()}
}

type HttpRouter struct {
	mux *http.ServeMux
}

func (r *HttpRouter) Mux() http.Handler {
	return r.mux
}

func (r *HttpRouter) RegisterRoutes(routes map[string]handlers.HttpHandler) {
	for p, h := range routes {
		r.mux.HandleFunc(p, h)
	}
}
