package controllers

import (
	"github.com/labstack/echo"
)

func RegisterRouters(e *echo.Echo) {

	user := &UserController{}
	user.registerRouter(e.Group("/user"))

	post := &PostController{}
	post.registerRouter(e.Group("/post"))

}

//func registerFilter(e *echo.Echo) {
//	// middleware
//	e.Pre(filters.SetLogTrace())
//	e.Use(filters.GetLogMiddleware())
//	e.Use(middleware.Recover())
//}

//type customValidator struct {
//}
//
//func (cv *customValidator) Validate(i interface{}) error {
//	if _, err := govalidator.ValidateStruct(i); err != nil {
//		return err
//	}
//	return nil
//}

//SetupRouter 启动路由，添加中间件
//func SetupRouter(e*echo.Echo) *echo.Echo {
//	//e := echo.New()
//	//govalidator.TagMap["alphaversion"] = govalidator.Validator(func(str string) bool {
//	//	return str == "alpha" || str == "beta" || str == "rc" || str == "release"
//	//})
//	//e.Validator = &customValidator{}
//	//registerFilter(e)
//	registerRouters(e)
//	return e
//}
