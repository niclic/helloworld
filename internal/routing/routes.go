package routing

import "github.com/niclic/helloworld/internal/handlers"

var Routes = map[string]handlers.HttpHandler{
	"/hello": handlers.Hello,
}
