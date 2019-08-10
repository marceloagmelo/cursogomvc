package main

import (
	"github.com/labstack/echo/middleware"
	r "github.com/marceloagmelo/cursogomvc/routers"
	"github.com/marceloagmelo/pongor-echo"
)

func main() {
	e := r.App

	p := pongor.GetRenderer()
	p.Directory = "views"

	e.Renderer = p
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))
}
