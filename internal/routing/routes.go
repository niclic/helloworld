package routing

import "github.com/niclic/helloworld/internal/handlers"

// Routes is a map of URL path patterns and request handler functions.
var Routes = map[string]handlers.HTTPHandler{
	"/hello": handlers.Hello,
}
