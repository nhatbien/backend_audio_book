package main

import (
	"backend/controller"
	"backend/db"
	"backend/helper"
	"backend/repository/repo_impl"
	"backend/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     "3306",
		User:     "audio_book",
		Password: "Nhutkhung123@",
		Dbname:   "audio_book",
	}
	sql.Connect()
	e := echo.New()
	e.Use(middleware.CORS())

	structValidator := helper.NewStructValidaten()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	userController := controller.UserController{
		UserRepo: repo_impl.NewUserRepo(sql)}
	categoryBookController := controller.CategoryBookController{
		CategoryBookRepo: repo_impl.NewCategoryBookRepo(sql)}

	api := router.API{
		Echo:                   e,
		UserController:         userController,
		CategoryBookController: categoryBookController,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3001"))
	//	defer sql.Close()

}
