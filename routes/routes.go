package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin_learn/handler"
	"github.com/gin_learn/middleware"
	"github.com/gin_learn/models/book"
	"github.com/gin_learn/models/usermodel"
	"gorm.io/gorm"
)

func InitiateRoute(app *gin.Engine, db *gorm.DB) {
	// book
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewHandler(bookService)

	// user
	userRepo := usermodel.NewUserRepository(db)
	userService := usermodel.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	v1 := app.Group("/v1")
	{
		// auth
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		v1.Use(middleware.IsAuthenticated())

		v1.GET("/user", userHandler.User)
		v1.POST("/logout", userHandler.Logout)
		// books
		v1.GET("/books", bookHandler.AllBooks)
		v1.GET("/books/:id", bookHandler.FindBookById)
		v1.POST("/books/:id", bookHandler.HapusBuku)
		v1.PUT("/books/:id", bookHandler.UpdateBuku)
		v1.POST("/books", bookHandler.CreateBookHandler)

		v1.POST("/upload", handler.Upload)
		// serve static files
		v1.Static("/upload", "uploads/")
	}
}
