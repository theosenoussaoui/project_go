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

<<<<<<< HEAD
	route.POST("/login", authMiddleware.LoginHandler)

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
=======
func main() {
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/users", UserController.CreateUser)
	r.PUT("/users/:uuid", UserController.UpdateUser)
	r.DELETE("/users/:uuid", UserController.DeleteUser)

	r.POST("/votes", VoteController.CreateVote)
	r.GET("/votes/:uuid", VoteController.GetVote)
	r.PUT("/votes/:uuid", VoteController.UpdateVote)

	r.Run() // listen and serve on 0.0.0.0:8080
}
>>>>>>> 57b3d25dce5e1fad5897c253004a09d67b5d9692
