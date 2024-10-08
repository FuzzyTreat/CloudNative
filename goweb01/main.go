package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", start)
	router.GET("/about", about)
	router.Run(":45000")
}

func start(c *gin.Context) {
	c.String(200, "<h1>Hello world!</h1>")
}

func about(c *gin.Context) {
	c.String(200, "<b>About this!</b>")
}
