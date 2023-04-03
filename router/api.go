package router

import (
	"backend/controller"
	"backend/model"
	"backend/model/request"
	"net/http"

	// "backend/docs"
	"backend/middleware"

	"github.com/labstack/echo/v4"
	"github.com/pangpanglabs/echoswagger/v2"
)

type API struct {
	Echo                   *echo.Echo
	UserController         controller.UserController
	CategoryBookController controller.CategoryBookController
	BookController         controller.BookController
	CartController         controller.CartController
}

func (api *API) SetupRouter() {

	v1 := api.Echo.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			//user.POST("/user/login", api.UserController.Login)
			user.POST("/signup", api.UserController.Signup)
			user.POST("/signin", api.UserController.Login)
			user.POST("/update", api.UserController.Update, middleware.JWTMiddleware())
			user.POST("/update-role", api.UserController.UpdateRole, middleware.JWTMiddleware())
			user.GET("/profile", api.UserController.GetProfile, middleware.JWTMiddleware())
		}
		//user.GET("/profile", api.UserController.GetProfile, middleware.JWTMiddleware())
		book := user.Group("/book")
		{
			book.POST("/save", api.BookController.SaveBook, middleware.JWTMiddleware())
			book.POST("/:id/update", api.BookController.UpdateBook, middleware.JWTMiddleware())
			book.GET("/:id", api.BookController.SelectBookById)
			book.GET("/all", api.BookController.SelectAllBook)
		}
		categoryBook := user.Group("/category-book")
		{
			categoryBook.POST("/save", api.CategoryBookController.SaveCategoryBook, middleware.JWTMiddleware())
			categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware())
			categoryBook.GET("/:id", api.CategoryBookController.GetCategoryBookById)
			categoryBook.GET("/all", api.CategoryBookController.GetAllCategoryBook)
		}
		cart := user.Group("/cart")
		{
			cart.POST("/add", api.CartController.AddItemToCart, middleware.JWTMiddleware())
			cart.GET("/", api.CartController.SelectMyCart, middleware.JWTMiddleware())
			cart.GET("", api.CartController.SelectMyCart, middleware.JWTMiddleware())

			//categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware())
			cart.GET("/:id", api.CategoryBookController.GetCategoryBookById)
		}

	}

	//url := echoSwagger.URL("https://nhatbien.github.io/doc.json") //The url pointing to API definition
	//api.Echo.GET("/swagger/*", echoSwagger.EchoWrapHandler(url))
}

func (api *API) SetupSwagger() {
	/* respf := &&model.Response{
		Status:  false,
		Message: "failurea",
		Data:    nil,
	} */

	r := echoswagger.New(api.Echo, "/doc", &echoswagger.Info{
		Title:       "Audibook Example API",
		Description: "This is a sample server Audibook server.",
		Version:     "1.0.0",
	})

	r.AddSecurityAPIKey("Authorization", "Bearer", echoswagger.SecurityInHeader).
		SetRequestContentType("application/json", "application/x-www-form-urlencoded", "multipart/form-data").
		SetScheme("http")

	user := r.Group("User", "/api/v1/user")

	{
		user.POST("/signup", api.UserController.Signup).
			AddParamBody(&request.UserSignupRequest{}, "body", "user register", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: model.User{}}, nil)
		user.POST("/signin", api.UserController.Login).
			AddParamBody(&request.UserLoginRequest{}, "body", "user login", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: model.User{}}, nil)
		user.POST("/update", api.UserController.Update, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.UserUpdateRequest{}, "body", "user update profile", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: []byte("null")}, nil)
		user.POST("/update-role", api.UserController.UpdateRole, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.UserUpdateRoleRequest{}, "body", "user update role", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: []byte("null")}, nil)

		user.GET("/profile", api.UserController.GetProfile, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.User{}}, nil)
	}
	categoryBook := r.Group("CategoryBook", "/api/v1/category-book")
	{
		categoryBook.POST("/save", api.CategoryBookController.SaveCategoryBook, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.CategoryBookSave{}, "body", "category book save ", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.BookCategory{}}, nil)
		categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamPath("id", "id", "string").
			AddParamBody(&request.CategoryBookSave{}, "body", "category book update", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.BookCategory{}}, nil)

		categoryBook.GET("/all", api.CategoryBookController.GetAllCategoryBook).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &[]model.BookCategory{}}, nil)
		categoryBook.GET("/:id", api.CategoryBookController.GetCategoryBookById).
			AddParamPath("id", "id", "string").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.BookCategory{}}, nil)
	}
	book := r.Group("Book", "/api/v1/book")
	{
		book.POST("/save", api.BookController.SaveBook, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.BookSaveRequest{}, "body", " book save ", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Book{}}, nil)

		book.POST("/:id/update", api.BookController.UpdateBook, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamPath("id", "id", "string").
			AddParamBody(&request.BookUpdateRequest{}, "body", " book update", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Book{}}, nil)

		/* categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware()).
		SetSecurity("Authorization").
		AddParamPath("id", "id", "string").
		AddParamBody(&request.CategoryBookSave{}, "body", "category book update", true).
		AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.BookCategory{}}, nil)
		*/
		book.GET("/all", api.BookController.SelectAllBook).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &[]model.Book{}}, nil)
		book.GET("/:id", api.BookController.SelectBookById).
			AddParamPath("id", "id", "string").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Book{}}, nil)
	}
	cart := r.Group("Cart", "/api/v1/cart")
	{
		cart.POST("/add", api.CartController.AddItemToCart, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.CartItemSave{}, "body", "cart book update", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Cart{}}, nil)
		cart.GET("/", api.CartController.SelectMyCart).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Cart{}}, nil)
	}
}
