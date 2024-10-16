package data

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POJO - POCO
// POGO

type Employee struct {
	Id   int    `json:"id"`
	Age  int    `json:"age"`
	City string `json:"city"`
	Namn string `json:"name"`
}

// MEDLEMSFUNKTION
// emp.CalculateSalary()
func (emp Employee) CalculateSalary() int {
	if emp.Namn == "Stefan" {
		return 1000
	}
	return 10
}

// request, response
func HandleGetAllEmployees(c *gin.Context) {
	emps := GetAllEmployees()
	c.IndentedJSON(http.StatusOK, emps)
}

func HandleGetOneEmployee(c *gin.Context) {
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employee := GetEmployee(numId)

	if employee == nil { // INTE HITTAT  /api/employee/
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Finns inte"})
	} else {
		c.IndentedJSON(http.StatusOK, employee)
	}
}

func HandleUpdateEmployee(c *gin.Context) { // PUT Bodyn JSON {"Age":16, "City":"Nacka"}"
	// /api/employee/21
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employeeFromDB := GetEmployee(numId) // DEN SOM REDAN FINNS I DATABASEN !!!
	if employeeFromDB == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		var employeeJson Employee

		if err := c.BindJSON(&employeeJson); err != nil {
			return
		}
		employeeJson.Id = numId
		UpdateEmployee(employeeJson)
		//data. Save(&employeeJson)
		c.IndentedJSON(http.StatusOK, employeeJson)
	}
}

func HandleDeleteEmployee(c *gin.Context) { // PUT Bodyn JSON {"Age":16, "City":"Nacka"}"
	// /api/employee/21
	id := c.Param("id") // "a"
	numId, _ := strconv.Atoi(id)
	employeeFromDB := GetEmployee(numId) // DEN SOM REDAN FINNS I DATABASEN !!!
	if employeeFromDB == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		DeleteEmployee(employeeFromDB)
		//data. Save(&employeeJson)
		c.IndentedJSON(http.StatusNoContent, gin.H{"message": "Deleted ok"})
	}
}

func HandleNewEmployees(c *gin.Context) {
	// TODO Add new
	// försöka få fram den JSON Employee som man skickat in
	var employee Employee
	if err := c.BindJSON(&employee); err != nil {
		return
	}
	CreateNewEmployee(employee)
	c.IndentedJSON(http.StatusCreated, employee)
}
