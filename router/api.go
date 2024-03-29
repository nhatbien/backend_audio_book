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
	OrderController        controller.OrderController
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
			user.GET("/:id", api.UserController.SelectUserId)
		}
		//user.GET("/profile", api.UserController.GetProfile, middleware.JWTMiddleware())
		book := user.Group("/book")
		{
			book.POST("/save", api.BookController.SaveBook, middleware.JWTMiddleware())
			book.POST("/:id/update", api.BookController.UpdateBook, middleware.JWTMiddleware())
			book.GET("/:id", api.BookController.SelectBookById)
			book.DELETE("/:id", api.BookController.DeleteBook)
			book.GET("/all", api.BookController.SelectAllBook)
			book.GET("/search", api.BookController.SearchBookByName)
		}
		categoryBook := user.Group("/category-book")
		{
			categoryBook.POST("/save", api.CategoryBookController.SaveCategoryBook, middleware.JWTMiddleware())
			categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware())
			categoryBook.GET("/:id", api.CategoryBookController.GetCategoryBookById)
			categoryBook.DELETE("/:id", api.CategoryBookController.DeleteCategoryBook, middleware.JWTMiddleware())
			categoryBook.GET("/all", api.CategoryBookController.GetAllCategoryBook)
		}
		cart := user.Group("/cart")
		{
			cart.POST("/add", api.CartController.AddItemToCart, middleware.JWTMiddleware())
			cart.DELETE("/:id", api.CartController.DeleteCartItem, middleware.JWTMiddleware())
			cart.DELETE("/delete", api.CartController.DeleteCart, middleware.JWTMiddleware())
			cart.GET("/", api.CartController.SelectMyCart, middleware.JWTMiddleware())
			cart.GET("", api.CartController.SelectMyCart, middleware.JWTMiddleware())
			//categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware())
			cart.GET("/:id", api.CategoryBookController.GetCategoryBookById)
		}
		order := user.Group("/order")
		{
			order.POST("/save", api.OrderController.SaveOrder, middleware.JWTMiddleware())
			order.GET("/:id", api.OrderController.SelectOrderById)
			order.PUT(":id", api.OrderController.PutOrderStatus, middleware.JWTMiddleware())
			order.GET("/status/:status", api.OrderController.SelectOrderByStatus, middleware.JWTMiddleware())
			order.GET("/get/:status", api.OrderController.SelectAllBookOrderbyStatusAndUserId, middleware.JWTMiddleware())
			order.GET("/me/:status", api.OrderController.SelectAllOrderbyStatusAndUserId, middleware.JWTMiddleware())
			//categoryBook.POST("/:id/update", api.CategoryBookController.UpdateCategoryBook, middleware.JWTMiddleware())
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
		user.GET("/:id", api.UserController.SelectUserId).
			AddParamPath("id", "id", "string").
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
		categoryBook.DELETE("/:id", api.CategoryBookController.DeleteCategoryBook, middleware.JWTMiddleware()).
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
		book.GET("/search", api.BookController.SearchBookByName).
			AddParamQuery("name", "name", "nameBoook", false).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Book{}}, nil)
		book.DELETE("/:id", api.BookController.DeleteBook, middleware.JWTMiddleware()).
			AddParamPath("id", "id", "string").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Book{}}, nil)

	}
	cart := r.Group("Cart", "/api/v1/cart")
	{
		cart.POST("/add", api.CartController.AddItemToCart, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.CartItemSave{}, "body", "cart book update", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Cart{}}, nil)
		cart.DELETE("/:id", api.CartController.DeleteCartItem, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamPath("id", "id", "id").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Cart{}}, nil)
		cart.DELETE("/delete", api.CartController.DeleteCart, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Cart{}}, nil)
		cart.GET("/", api.CartController.SelectMyCart, middleware.JWTMiddleware()).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Cart{}}, nil)
	}

	order := r.Group("Order", "/api/v1/order")
	{
		order.POST("/save", api.OrderController.SaveOrder, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.OrderSave{}, "body", "order save ", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Order{}}, nil)
		order.PUT("/:id", api.OrderController.PutOrderStatus, middleware.JWTMiddleware()).
			SetSecurity("Authorization").
			AddParamBody(&request.OrderStatusChange{}, "body", "order put change status ", true).
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Order{}}, nil)
		order.GET("/:id", api.OrderController.SelectOrderById).
			AddParamPath("id", "id", "int ID").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &model.Order{}}, nil)
		order.GET("/status/:status", api.OrderController.SelectOrderByStatus, middleware.JWTMiddleware()).
			AddParamPath("status", "status", "int status").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &[]model.Order{}}, nil)
		order.GET("/get/:status", api.OrderController.SelectAllBookOrderbyStatusAndUserId, middleware.JWTMiddleware()).
			AddParamPath("status", "status", "int status").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &[]model.Book{}}, nil)
		order.GET("/me/:status", api.OrderController.SelectAllOrderbyStatusAndUserId, middleware.JWTMiddleware()).
			AddParamPath("status", "status", "int status").
			AddResponse(http.StatusOK, "success", &model.Response{Status: true, Message: "success", Data: &[]model.Book{}}, nil)
	}
}
