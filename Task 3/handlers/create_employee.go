package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/uditkalra/emsGcpApi/function"
	"github.com/uditkalra/emsGcpApi/structure"
	"github.com/uditkalra/emsGcpApi/variables"
)

func init() {
	functions.HTTP("CreateEmployee", CreateEmployeeHandler)
}

// createEmployeeHandler handles the creation of a new employee
func CreateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into an Employee object
	var employee structure.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the fields
	if employee.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if employee.Email == "" || !function.IsValidEmail(employee.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}
	if employee.PhoneNumber == "" || !function.IsValidPhoneNumber(employee.PhoneNumber) {
		http.Error(w, "Invalid phone number format", http.StatusBadRequest)
		return
	}
	if employee.DateOfBirth == "" || !function.IsValidDateOfBirth(employee.DateOfBirth) {
		http.Error(w, "Invalid date of birth format", http.StatusBadRequest)
		return
	}
	if employee.Password == "" {
		http.Error(w, "Password is required", http.StatusBadRequest)
		return
	}

	// Hash the password
	hashedPassword, err := function.HashPassword(employee.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	employee.Password = hashedPassword

	// Set the creation time
	employee.CreationTime = time.Now()

	// Generate a new ID for the employee
	employee.ID = function.GenerateID()

	// Save the employee data to Firestore
	_, err = variables.Client.Collection("employees").Doc(employee.ID).Set(context.Background(), employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created employee as the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	showEmployee := function.EmployeeOutput(employee)
	json.NewEncoder(w).Encode(showEmployee)
}
