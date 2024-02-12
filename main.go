package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello from golang in Vercel")
	})

	app.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Printf("Listening on port %s", port)

	app.Run(":" + port)
}
