package controllers

import (
	"github.com/labstack/echo"
	"github.com/mingz2013/echo-demo-go/internal/app/login/models"
	"net/http"
)

type UserController struct {
}

func (c *UserController) registerRouter(g *echo.Group) {
	g.GET("/", c.list)
	g.POST("/", c.save)
	g.GET("/:id", c.get)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.delete)
}

func (c *UserController) list(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func (c *UserController) get(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func (c *UserController) save(ctx echo.Context) error {
	//id:=c.Param("id")
	u := new(models.User)
	if err := ctx.Bind(u); err != nil {
		return err
	}

	if err := ctx.Validate(u); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)

}

func (c *UserController) update(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func (c *UserController) delete(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}
