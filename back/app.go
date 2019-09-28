package main

// import (
//     "net/http"
//     "log"
//     "github.com/gorilla/mux"
// )

// func YourHandler(w http.ResponseWriter, r *http.Request) {
//     w.Write([]byte("Gorilla!\n"))
// }

// func main() {
//     r := mux.NewRouter()
//     // Routes consist of a path and a handler function.
//     r.HandleFunc("/", YourHandler)

//     // Bind to a port and pass our router in
//     log.Fatal(http.ListenAndServe(":8080", r))
// }

// Test ajout Gonic

import "github.com/gin-gonic/gin"

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