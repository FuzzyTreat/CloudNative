package main

import (
	"goweb01/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageView struct {
	CurrentUser string
	PageTitle   string
	Title       string
	Text        string
}

func main() {
	data.Init()

	router := gin.Default()

	router.LoadHTMLGlob("templates/**") // where are the html templates located
	router.GET("/", start)
	router.GET("/about", about)

	router.GET("/api/employee", handleGetAllEmployees)
	router.GET("/api/employee/:id", handleGetOneEmployee)
	router.POST("/api/employee", handleNewEmployees) // SKA JU Employee skickas med som JSON

	router.Run(":45000")
}

func start(c *gin.Context) {
	c.String(200, "<h1>Hello world!</h1>")
	//c.HTML(200, "home.html", &PageView{PageTitle: "Some page", Title: "Välkommen!", Text: "Lite text på sidan"})
}

func about(c *gin.Context) {
	c.String(200, "<b>About this!</b>")
}

func mySelf(c *gin.Context) {
	c.JSON(http.StatusOK, &MySelf{Name: "MySelf", Age: 13, Birthday: "1950-05-05", City: "Nowhere"})
}
