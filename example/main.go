package main

import (
	"fmt"

	"github.com/AmyangXYZ/sweetygo"
)

func main() {
	app := sweetygo.New()

	app.USE(sweetygo.Logger())
	// app.GET("/static/*files", staticServer)
	app.GET("/", home)
	app.POST("/api", home)
	app.GET("/usr/:user/:sex/:age", hello)

	app.RunServer(":8080")
}

func home(c *sweetygo.Context) {
	c.Resp.WriteHeader(200)
	fmt.Fprintf(c.Resp, "Welcome \n")
}

func hello(c *sweetygo.Context) {
	c.Resp.WriteHeader(200)
	fmt.Fprintf(c.Resp, "Hello %s\n", c.Params()["user"][0])
}
