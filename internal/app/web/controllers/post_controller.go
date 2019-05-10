package controllers

import (
	"github.com/labstack/echo"
	"github.com/mingz2013/echo-demo-go/internal/app/blog/models"
	"net/http"
)

type PostController struct {
}

func (c *PostController) registerRouter(g *echo.Group) {
	g.GET("/", c.list)
	g.POST("/", c.save)
	g.GET("/:id", c.get)
	g.PUT("/:id", c.update)
	g.DELETE("/:id", c.delete)
}

func (c *PostController) list(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func (c *PostController) get(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func (c *PostController) save(ctx echo.Context) error {
	//id:=c.Param("id")
	u := new(models.Post)
	if err := ctx.Bind(u); err != nil {
		return err
	}

	if err := ctx.Validate(u); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, u)

}

func (c *PostController) update(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}

func (c *PostController) delete(ctx echo.Context) error {
	id := ctx.Param("id")
	return ctx.String(http.StatusOK, id)
}
