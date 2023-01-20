package main

import (
	"backend/controller"
	"backend/db"
	"backend/helper"
	"backend/repository/repository_impl"
	"backend/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "pjpjmn01",
		Dbname:   "backend_grab",
	}
	sql.Connect()
	e := echo.New()
	e.Use(middleware.CORS())

	structValidator := helper.NewStructValidaten()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	userController := controller.UserController{
		UserRepo: repository_impl.NewUserRepo(sql)}

	api := router.API{
		Echo:           e,
		UserController: userController,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3001"))
	defer sql.Close()

}
