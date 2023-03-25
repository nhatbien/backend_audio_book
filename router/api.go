package router

import (
	"backend/controller"
	"backend/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type API struct {
	Echo                   *echo.Echo
	UserController         controller.UserController
	CategoryBookController controller.CategoryBookController
}

func (api *API) SetupRouter() {

	v1 := api.Echo.Group("/api/v1")

	user := v1.Group("/user")

	//user.POST("/user/login", api.UserController.Login)
	user.POST("/signup", api.UserController.Signup)
	user.POST("/signin", api.UserController.Login)
	user.POST("/update", api.UserController.Update, middleware.JWTMiddleware())
	user.POST("/update-role", api.UserController.UpdateRole, middleware.JWTMiddleware())

	categoryBook := user.Group("/category-book")
	categoryBook.POST("/save", api.CategoryBookController.SaveCategoryBook, middleware.JWTMiddleware())

	///pi.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
	url := echoSwagger.URL("http://localhost:1323/swagger/doc.json") //The url pointing to API definition
	api.Echo.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
}
