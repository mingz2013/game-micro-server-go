package web_proxy

import (
	"github.com/labstack/echo"
	"io"
	"net/http"
	"os"
)

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func saveUser(c echo.Context) error {
	//id:=c.Param("id")
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)

}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// /show?team=x-men&member=wolverine
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// $ curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/saveFile
// => <b>Thank you! Joe Smith</b>
func saveFile(c echo.Context) error {
	name := c.FormValue("name")
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank u"+name+"</b>")

}
