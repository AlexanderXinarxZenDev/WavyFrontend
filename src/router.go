// flusk/router.go
package flusk

type Handler func(*Context)

type Router struct {
	routes map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]Handler),
	}
}

func (r *Router) Add(method, path string, handler Handler) {
	key := method + " " + path
	r.routes[key] = handler
}

func (r *Router) Find(method, path string) (Handler, bool) {
	key := method + " " + path
	h, ok := r.routes[key]
	return h, ok
}