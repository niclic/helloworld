package routing

import (
	"net/http"

	"github.com/niclic/helloworld/internal/handlers"
)

var mux = http.NewServeMux()

var DefaultRouter = &HttpRouter{ServeMux: mux}

type HttpRouter struct {
	*http.ServeMux
}

func (r *HttpRouter) RegisterRoutes(routes map[string]handlers.HttpHandler) {
	for p, h := range routes {
		r.HandleFunc(p, h)
	}
}
