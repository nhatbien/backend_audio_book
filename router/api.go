package router

import (
	"backend/controller"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo           *echo.Echo
	UserController controller.UserController
}

func (api *API) SetupRouter() {

	v1 := api.Echo.Group("/api/v1")

	user := v1.Group("/user")

	//user.POST("/user/login", api.UserController.Login)
	user.POST("/signup", api.UserController.Signup)
	user.POST("/signin", api.UserController.Login)
	user.POST("/update", api.UserController.Update, middleware.JWTMiddleware())

}
