package main

import "src/routes"
import "github.com/gin-gonic/gin"

func main() {
	router := routes.NewRoutes()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}