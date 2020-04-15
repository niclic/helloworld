package routing

import (
	"net/http"

	"github.com/niclic/helloworld/internal/handlers"
)

func init() {
	router.RegisterRoutes(Routes)
}

var mux = http.NewServeMux()

var router = &HttpRouter{ServeMux: mux}

var DefaultRouter = router

type HttpRouter struct {
	*http.ServeMux
}

func (r *HttpRouter) RegisterRoutes(routes map[string]handlers.HttpHandler) {
	for p, h := range routes {
		r.HandleFunc(p, h)
	}
}
