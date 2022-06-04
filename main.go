package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin_learn/models/book"
	"github.com/gin_learn/models/usermodel"
	"github.com/gin_learn/routes"
	"github.com/gin_learn/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@/gin_books?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic("gak bisa konek database")
	}

	util.DB = db

	db.AutoMigrate(&book.Book{}, &usermodel.User{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	routes.InitiateRoute(r, db)

	// main
	// handler
	// service
	// repo
	// db
	// mysql

	r.Run()
}
