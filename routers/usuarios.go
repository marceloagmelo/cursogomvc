package routers

import (
	"github.com/labstack/echo"
	"github.com/marceloagmelo/cursogomvc/controllers"
)

//App é uma instancia de App
var App *echo.Echo

func init() {
	App = echo.New()

	// A página inicial da aplicação
	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/atualizar/:id", controllers.Atualizar)

	api := App.Group("/v1")
	api.POST("/insert", controllers.Inserir)
	api.DELETE("/delete/:id", controllers.Deletar)
	api.PUT("/update/:id", controllers.Update)
}
