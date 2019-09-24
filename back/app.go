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

	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST users",
		})
	})

	r.PUT("/users/:uuid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT users",
		})
	})

	r.DELETE("/users/:uuid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE users",
		})
	})

	r.POST("/votes", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST votes",
		})
	})

	r.GET("/votes/:uuid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET votes",
		})
	})

	r.PUT("/votes/:uuid", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT votes",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}