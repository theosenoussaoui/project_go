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

	// Inserting admin in Docker PGSQL
	db.CreateSystemAdmin()

	r := gin.Default()
	route := r.Group("/")

	// Manage login : 
	// Authentificating + generating JWT

	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	route.POST("/login", authMiddleware.LoginHandler)

	// Manage protected routes
	route.Use(authMiddleware.MiddlewareFunc())
	{
		users := route.Group("/users")
		{
			users.GET("/", controllers.GetUsers)
			users.POST("/", controllers.CreateUser)
			users.PUT("/:uuid", controllers.UpdateUser)
			users.DELETE("/:uuid", controllers.DeleteUser)
		}

		votes := route.Group("/votes")
		{
			votes.GET("/", controllers.GetVotes)
			votes.POST("/", controllers.CreateVote)
			votes.PUT("/:uuid", controllers.UpdateVote)
			// votes.DELETE("/:uuid", controllers.DeleteVote)
		}
	}

	// Run on port 8080

	r.Run(":8080")
}
