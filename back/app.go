package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/theosenoussaoui/project_go/back/controllers"
	"github.com/theosenoussaoui/project_go/back/db"
	"github.com/theosenoussaoui/project_go/back/middleware"
)

func main() {
	log.Println("Starting server...")

	db.Initialize()
	db.CreateSystemAdmin()

	r := gin.Default()
	route := r.Group("/")

	// Manage login (auth + generate JWT)
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// route.POST("/login", authMiddleware.LoginHandler)

	// Manage protected routes
	route.Use(authMiddleware.MiddlewareFunc())
	{
		users := route.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.CreateUser)
		}

		// votes := route.Group("/votes")
		// {
		// 	votes.GET("/", controllers.GetVotes)
		// 	votes.POST("/", controllers.CreateVote)
		// }
	}

	r.Run(":8080")
}
