# 🔥 Flusk v1.5

**A minimal Flask‑like web framework in Go** – lightweight, zero dependencies, and dead simple.

Flusk gives you just enough to build HTTP APIs without the clutter. Inspired by Flask's elegance, built with Go's performance.

---

## ✨ Features

- 🚀 **Simple routing** – `GET` and `POST` with path matching
- 📦 **No external dependencies** – only the Go standard library
- 🎯 **Context wrapper** – easy access to request/response
- 📄 **Text & JSON responses** – one line each
- 🏃 **Single‑line server startup** – `app.Run(":8080")`
- 🔍 **404 handling** – automatic for missing routes
- 🧹 **Under 300 lines** – tiny, readable, maintainable

---

## 📦 Installation

```bash
go get github.com/AlexanderXinarxZenDev/flusk/src
```

Or if you're using the local development structure:

```bash
git clone https://github.com/yourname/flusk
cd flusk
go work init ./src ./examples   # if using go.work
```

---

## 🚀 Quick start

```go
package main

import "github.com/yourname/flusk"

func main() {
    app := flusk.New()

    app.GET("/", func(c *flusk.Context) {
        c.Text(200, "Hello Flusk v1.5")
    })

    app.GET("/json", func(c *flusk.Context) {
        c.JSON(200, map[string]string{
            "message": "working",
        })
    })

    app.POST("/submit", func(c *flusk.Context) {
        c.Text(201, "Created")
    })

    app.Run(":8080")
}
```

Run it:

```bash
go run main.go
# Output: 🚀 Flusk v1.5 running on http://localhost:8080
```

Test with `curl`:

```bash
curl http://localhost:8080/
# Hello Flusk v1.5

curl http://localhost:8080/json
# {"message":"working"}
```

---

## 📚 API Reference

### `flusk.New() *App`
Creates a new Flusk application instance.

### `app.GET(path string, handler Handler)`
Registers a handler for HTTP `GET` requests at the given path.

### `app.POST(path string, handler Handler)`
Registers a handler for HTTP `POST` requests.

### `app.Run(addr string) error`
Starts the HTTP server on the specified address (e.g., `":8080"`).

### `type Context`
Wraps `http.ResponseWriter` and `*http.Request`. Provides two response methods:

#### `c.Text(status int, msg string)`
Sends a plain text response with the given status code.

#### `c.JSON(status int, data any)`
Sends a JSON response. Automatically sets `Content-Type: application/json`.

---

## 📁 Project structure

```
flusk/
├── flusk.go        # App core
├── router.go       # Simple route map
├── context.go      # Request/response wrapper
├── server.go       # HTTP server bootstrap
├── examples/
│   └── main.go     # Demo usage
└── go.mod
```

---

## 🧠 Design philosophy

- **Minimalism** – no middleware system, no templates, no complex routing trees.
- **Simplicity over features** – you can always add more, but you can't take away.
- **Clean idiomatic Go** – each file has a single responsibility.
- **Production‑ready** – uses standard `http.ListenAndServe` under the hood.

---

## 🚫 What Flusk is NOT

- A full‑featured framework (use Gin, Echo, or Fiber for that)
- A template engine (use `html/template` directly)
- A middleware powerhouse (but you can easily wrap handlers yourself)

---

## 🤝 Contributing

Flusk is intentionally small. If you have an idea that fits the minimalist spirit, open an issue or PR.

Guidelines:
- Keep it under 400 lines total
- No external dependencies
- Preserve the simple API surface

---

## 📄 License

BCD 3 Clause License – use it anywhere, for anything.

---

**Built with ❤️ for Go developers who miss Flask's simplicity.**
