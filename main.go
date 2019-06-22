package main

import (
	r "github.com/marceloagmelo/cursogomvc/routers"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))
}
