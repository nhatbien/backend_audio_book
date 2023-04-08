package main

import (
	"backend/controller"
	"backend/db"
	_ "backend/docs"

	"backend/helper"
	"backend/repository/repo_impl"
	"backend/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// @title AudioBook Example API
// @version 1.0.1
// @description This is a sample server AudioBook server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Nginx
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3001
// @BasePath /
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

	bookController := controller.BookController{
		BookRepo: repo_impl.NewBookRepo(sql)}

	cartController := controller.CartController{
		CartRepo: repo_impl.NewCartRepo(sql)}
	orderController := controller.OrderController{
		OrderRepo: repo_impl.NewOrderRepo(sql)}

	api := router.API{
		Echo:                   e,
		UserController:         userController,
		CategoryBookController: categoryBookController,
		BookController:         bookController,
		CartController:         cartController,
		OrderController:        orderController,
	}

	api.SetupRouter()
	api.SetupSwagger()

	e.Logger.Fatal(e.Start(":3001"))
	//	defer sql.Close()

}
