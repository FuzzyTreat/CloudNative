package main

import (
	"goweb01/data"
	"net/http"
	"strconv"

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

	// router.LoadHTMLGlob("templates/**") // where are the html templates located
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

// request, response
func handleGetAllEmployees(c *gin.Context) {
	emps := data.GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func handleGetOneEmployee(c *gin.Context) {
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employee := data.GetEmployee(numId)

	if employee == nil { // INTE HITTAT  /api/employee/
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func handleNewEmployees(c *gin.Context) {
	// TODO Add new
	// försöka få fram den JSON Employee som man skickat in
	var employee data.Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	data.CreateNewEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)
}
