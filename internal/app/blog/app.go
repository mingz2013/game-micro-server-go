package blog

import (
	"github.com/labstack/echo"
	"github.com/mingz2013/game-micro-server-go/internal/app/blog/controllers"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
)

func notFound(c echo.Context) error {
	return c.String(http.StatusNotFound, "not found")
}

func Start() {
	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New(nil)}

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "hello, world!")
	})

	controllers.RegisterRouters(e)

	e.Any("/", notFound)

	e.Static("/static", "./static")

	e.Logger.Fatal(e.Start(":8001"))
}
