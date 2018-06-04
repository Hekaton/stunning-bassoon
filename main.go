package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	app.RegisterView(iris.Django("./views", ".html"))

	// Method: GET
	// Resource http://localhost:8080

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("message", "Hello World!")

		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"))
}
