package handlers

import (
	"employee-api/database"
	"employee-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var emp models.Employee

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, _ := database.DB.Exec(
		"INSERT INTO employees(name,email,position,salary) VALUES(?,?,?,?)",
		emp.Name, emp.Email, emp.Position, emp.Salary,
	)

	id, _ := result.LastInsertId()
	emp.ID = int(id)

	c.JSON(http.StatusCreated, emp)
}

func GetEmployees(c *gin.Context) {
	rows, _ := database.DB.Query(
		"SELECT id,name,email,position,salary FROM employees",
	)
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee
		rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary)
		employees = append(employees, emp)
	}

	c.JSON(http.StatusOK, employees)
}

func GetEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp models.Employee

	err := database.DB.QueryRow(
		"SELECT id,name,email,position,salary FROM employees WHERE id=?",
		id,
	).Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, emp)
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var emp models.Employee

	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Exec(
		"UPDATE employees SET name=?, email=?, position=?, salary=? WHERE id=?",
		emp.Name, emp.Email, emp.Position, emp.Salary, id,
	)

	empID, _ := strconv.Atoi(id)
	emp.ID = empID

	c.JSON(http.StatusOK, emp)
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	database.DB.Exec("DELETE FROM employees WHERE id=?", id)

	c.JSON(http.StatusOK, gin.H{"message": "Employee deleted"})
}
