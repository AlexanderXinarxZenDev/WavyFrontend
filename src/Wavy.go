// Wavy/Wavy.go
package Wavy

type App struct {
	router *Router
}

func New() *App {
	return &App{
		router: NewRouter(),
	}
}

func (a *App) GET(path string, handler Handler) {
	a.router.Add("GET", path, handler)
}

func (a *App) POST(path string, handler Handler) {
	a.router.Add("POST", path, handler)
}

func (a *App) Run(addr string) error {
	return RunServer(a.router, addr)
}