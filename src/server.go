// flusk/server.go
package flusk

import (
	"net/http"
)

func RunServer(router *Router, addr string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler, ok := router.Find(r.Method, r.URL.Path)
		if !ok {
			http.NotFound(w, r)
			return
		}
		ctx := NewContext(w, r)
		handler(ctx)
	})
	return http.ListenAndServe(addr, nil)
}