package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uditkalra/ems_api/docs"
	"golang.org/x/crypto/bcrypt"
)

// dateformat to compare the entered date of birth
const dateFormat = "2006-01-02"

// column names in the excel file
var headerFields = []string{
	"ID",
	"Name",
	"Password",
	"IsAdmin",
	"Email",
	"PhoneNumber",
	"Department",
	"Role",
	"DateOfBirth",
}

// employee struct
type Employee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"is_admin"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Department  string `json:"department"`
	Role        string `json:"role"`
	DateOfBirth string `json:"date_of_birth"`
}

// slice of employee records
var employees []Employee

// error message to be shown in case any error occurs
type ErrorResponse struct {
	Error string `json:"error"`
}

// success message to be shown in case of success of the operation desired
type SuccessResponse struct {
	Message string `json:"message"`
}

// @title Implementing and Documenting Employee Management System API in Go using Swagger
// @version 1
// @Description This is the implementation and documentation of the Employee Management System API in Go using Swagger

// @contact.name Udit Kalra
// @contact.url https://github.com/Udit-Kalra-1117
// @contact.email kalra.udit15@gmail.com

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	fmt.Println("\nWelcome to Employee Management System API Implementation and Documentation in GO using Swagger!\n")
	// Check if CSV file exists
	if _, err := os.Stat("employees.csv"); err == nil {
		// If file exists, load the data from the CSV file
		loadFromCSV()
	} else {
		// If file doesn't exist, create an empty file
		file, err := os.Create("employees.csv")
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")

	employees_database := v1.Group("/employees")
	{
		employees_database.GET("/", getEmployees)
		employees_database.GET("/:id", getEmployeeByID)
		employees_database.POST("/", createEmployee)
		employees_database.PUT("/:id", updateEmployee)
		employees_database.DELETE("/:id", deleteEmployee)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// GetEmployees godoc
// @Summary Get all employees
// @Description Retrieves a list of all employees
// @Tags Employees
// @Produce json
// @Success 200 {array} Employee
// @Router /employees [get]
func getEmployees(c *gin.Context) {
	c.JSON(http.StatusOK, employees)
}

// GetEmployeeByID godoc
// @Summary Get an employee by ID
// @Description Retrieves a specific employee by ID
// @Tags Employees
// @Produce json
// @Param id path int true "Employee ID"
// @Success 200 {object} Employee
// @Failure 404 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /employees/{id} [get]
func getEmployeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Employee not found"})
		return
	}

	for _, employee := range employees {
		if employee.ID == id {
			c.JSON(http.StatusOK, employee)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "Employee not found"})
}

// CreateEmployee godoc
// @Summary Create an employee
// @Description Creates a new employee record
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body Employee true "Employee object"
// @Success 201 {object} Employee
// @Router /employees [post]
func createEmployee(c *gin.Context) {
	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid employee data"})
		return
	}

	// Validating the entered email address
	if !isValidEmail(employee.Email) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid email address"})
		return
	}

	// Validating the entered phone number
	if !isValidPhoneNumber(employee.PhoneNumber) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid phone number"})
		return
	}

	// Validating the entered date of birth
	if !isValidDateOfBirth(employee.DateOfBirth) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid date of birth"})
		return
	}

	// Hashing the entered password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to hash the password"})
		return
	}

	// Generating the next available ID
	employee.ID = getNextAvailableID()
	employee.Password = string(hashedPassword)
	employees = append(employees, employee)
	saveToCSV()
	c.JSON(http.StatusCreated, employee)
}

// append to appropriate ID if the excel has entries already
func getNextAvailableID() int {
	highestID := 0
	for _, employee := range employees {
		if employee.ID > highestID {
			highestID = employee.ID
		}
	}
	return highestID + 1
}

// UpdateEmployee godoc
// @Summary Update an employee by ID
// @Description Updates a specific employee by ID
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body Employee true "Employee object"
// @Success 200 {object} Employee
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /employees/{id} [put]
func updateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid employee ID"})
		return
	}

	var updatedEmployee Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid employee data"})
		return
	}

	// Validating the entered email address
	if updatedEmployee.Email != "" && !isValidEmail(updatedEmployee.Email) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid email address"})
		return
	}

	// Validating the entered phone number
	if updatedEmployee.PhoneNumber != "" && !isValidPhoneNumber(updatedEmployee.PhoneNumber) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid phone number"})
		return
	}

	// Validating the entered date of birth
	if updatedEmployee.DateOfBirth != "" && !isValidDateOfBirth(updatedEmployee.DateOfBirth) {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid date of birth"})
		return
	}

	for i, employee := range employees {
		if employee.ID == id {
			// Updating the specific parameter if it's provided in the request
			if updatedEmployee.Name != "" {
				employee.Name = updatedEmployee.Name
			}
			if updatedEmployee.Role != "" {
				employee.Role = updatedEmployee.Role
			}
			if updatedEmployee.IsAdmin {
				employee.IsAdmin = updatedEmployee.IsAdmin
			}
			if updatedEmployee.Email != "" {
				employee.Email = updatedEmployee.Email
			}
			if updatedEmployee.PhoneNumber != "" {
				employee.PhoneNumber = updatedEmployee.PhoneNumber
			}
			if updatedEmployee.DateOfBirth != "" {
				employee.DateOfBirth = updatedEmployee.DateOfBirth
			}
			if updatedEmployee.Password != "" {
				// Hashing the newly entered password
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedEmployee.Password), bcrypt.DefaultCost)
				if err != nil {
					c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to has the password"})
					return
				}
				employee.Password = string(hashedPassword)
			}
			updatedEmployee.ID = id
			employees[i] = employee
			saveToCSV()
			c.JSON(http.StatusOK, employee)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "Employee not found"})
}

// DeleteEmployee godoc
// @Summary Delete an employee by ID
// @Description Deletes a specific employee by ID
// @Tags Employees
// @Param id path int true "Employee ID"
// @Success 202 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /employees/{id} [delete]
func deleteEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid employee ID"})
		return
	}

	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			saveToCSV()
			c.JSON(http.StatusAccepted, SuccessResponse{Message: "Employee record deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "Employee not found"})
}

func isValidEmail(email string) bool {
	// Regular expression pattern for basic email validation
	// This pattern checks for some alpha-numeric till a @ symbol is encountered
	// and again for some alpha-numeric characters till a . is encountered
	// and 2 or more than 2 alphabets after the dot should be present
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Regular expression pattern for basic phone number validation
	// This pattern matches a sequence of 10 digits
	phoneNumberRegex := regexp.MustCompile(`^\d{10}$`)
	return phoneNumberRegex.MatchString(phoneNumber)
}

func isValidDateOfBirth(dateOfBirth string) bool {
	// Check if the dateOfBirth is in the correct format
	_, err := time.Parse(dateFormat, dateOfBirth)
	if err != nil {
		return false
	}
	return true
}

// function to load previous records from the existing csv if any
func loadFromCSV() error {
	file, err := os.Open("employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initializing a reader and making it reading all the fields as per the length of the headerFields struct defined above
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = len(headerFields)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// empty slice of type Employee
	employees = []Employee{}

	for i, record := range records {
		if i == 0 {
			//skip the header row
			continue
		}
		id, _ := strconv.Atoi(record[0])
		isAdmin, _ := strconv.ParseBool(record[3])

		// mapping the appropriate elements to the appropriate columns in the excel file
		employee := Employee{
			ID:          id,
			Name:        record[1],
			Password:    record[2],
			IsAdmin:     isAdmin,
			Email:       record[4],
			PhoneNumber: record[5],
			Department:  record[6],
			Role:        record[7],
			DateOfBirth: record[8],
		}

		employees = append(employees, employee)
	}
	return nil
}

// function to update the csv with the new entries
func saveToCSV() error {
	file, err := os.Create("employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//write the header row
	err = writer.Write(headerFields)
	if err != nil {
		return err
	}

	for _, employee := range employees {
		record := []string{
			strconv.Itoa(employee.ID),
			employee.Name,
			employee.Password,
			strconv.FormatBool(employee.IsAdmin),
			employee.Email,
			employee.PhoneNumber,
			employee.Department,
			employee.Role,
			employee.DateOfBirth,
		}
		if err := writer.Write(record); err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
