// flusk/examples/main.go
package main

import (
	"github.com/AlexanderXinarxZenDev/flusk"
)

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

	app.Run(":8080")
}