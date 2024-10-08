package main

import "github.com/gin-gonic/gin"

type PageView struct {
	CurrentUser string
	PageTitle   string
	Title       string
	Text        string
}

func main() {
	router := gin.Default()

	// router.LoadHTMLGlob("templates/**") // where are the html templates located
	router.GET("/", start)
	router.GET("/about", about)
	router.Run(":45000")
}

func start(c *gin.Context) {
	c.String(200, "<h1>Hello world!</h1>")
	//c.HTML(200, "home.html", &PageView{PageTitle: "Some page", Title: "Välkommen!", Text: "Lite text på sidan"})
}

func about(c *gin.Context) {
	c.String(200, "<b>About this!</b>")
}
