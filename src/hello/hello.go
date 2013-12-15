package main

import "github.com/codegangsta/martini"
import "github.com/codegangsta/martini-contrib/render"

type MessageStruct struct {
        Message string `json:"message"`
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/plaintext", func() string {
		return "Hello World!"
		})

	m.Get("/json", func(r render.Render) {
		r.JSON(300, MessageStruct{"Hello, World!"})
		})

	m.Run()
}