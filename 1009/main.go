package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"systementor.se/apidemo1009/data"
)

type PageView struct {
	Title  string
	Rubrik string
}

func handleAbout(c *gin.Context) {

	c.HTML(http.StatusOK, "about.html", &PageView{Title: "testint testing", Rubrik: data.GetEmployee(1).Namn})
}

func handleStart(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte("<html><body>Hello</body></html>"))
}

var config Config

func Readconfig(config *Config) {

}

func main() {
	Readconfig(&config)
	//	readConfig(&config)

	data.Init(config.Database.File,
		config.Database.Server,
		config.Database.Database,
		config.Database.Username,
		config.Database.Password,
		config.Database.Port)

	r := gin.Default()
	r.LoadHTMLGlob("templates/**")
	// GET all
	// GET one (by id)
	// POST Create new
	// PUT, DELETE

	r.GET("/", handleStart)
	r.GET("/about", handleAbout)
	r.GET("/api/employee", data.HandleGetAllEmployees)
	r.GET("/api/employee/:id", data.HandleGetOneEmployee)
	r.POST("/api/employee", data.HandleNewEmployees) // SKA JU Employee skickas med som JSON
	// PUT = replace (ALLA properties)
	// PATCH = update a few properties
	r.PUT("/api/employee/:id", data.HandleUpdateEmployee)    // SKA JU Employee skickas med som JSON
	r.DELETE("/api/employee/:id", data.HandleDeleteEmployee) // SKA JU Employee skickas med som JSON

	// r.GET("/omoss", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 		"nu":      "rast",
	// 	})
	// })
	r.Run(":45000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
