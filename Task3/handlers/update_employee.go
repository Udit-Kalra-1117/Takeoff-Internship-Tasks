package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/uditkalra/emsGcpApi/function"
	"github.com/uditkalra/emsGcpApi/structure"
	"github.com/uditkalra/emsGcpApi/variables"
)

func init() {
	functions.HTTP("UpdateEmployee", UpdateEmployeeHandler)
}

// updateEmployeeHandler updates an employee by ID
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the employee ID from the request URL
	id := r.URL.Path[len("/employee/"):]

	// Get the existing employee data from Firestore
	doc, err := variables.Client.Collection("employees").Doc(id).Get(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Deserialize the document data into an Employee object
	var existingEmployee structure.Employee
	doc.DataTo(&existingEmployee)

	// Parse the request body into an Employee object
	var updatedEmployee structure.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update only the provided fields
	if updatedEmployee.Name != "" {
		existingEmployee.Name = updatedEmployee.Name
	}
	if updatedEmployee.Password != "" {
		password, err := function.HashPassword(updatedEmployee.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		existingEmployee.Password = password
	}
	if updatedEmployee.IsAdmin != existingEmployee.IsAdmin {
		existingEmployee.IsAdmin = updatedEmployee.IsAdmin
	}
	if updatedEmployee.Email != "" {
		// Validate the email format
		if function.IsValidEmail(updatedEmployee.Email) {
			existingEmployee.Email = updatedEmployee.Email
		} else {
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}
	}
	if updatedEmployee.PhoneNumber != "" {
		// Validate the phone number format
		if function.IsValidPhoneNumber(updatedEmployee.PhoneNumber) {
			existingEmployee.PhoneNumber = updatedEmployee.PhoneNumber
		} else {
			http.Error(w, "Invalid phone number format", http.StatusBadRequest)
			return
		}
	}
	if updatedEmployee.Department != "" {
		existingEmployee.Department = updatedEmployee.Department
	}
	if updatedEmployee.Role != "" {
		existingEmployee.Role = updatedEmployee.Role
	}
	if updatedEmployee.DateOfBirth != "" {
		// Validate the date of birth format
		if function.IsValidDateOfBirth(updatedEmployee.DateOfBirth) {
			existingEmployee.DateOfBirth = updatedEmployee.DateOfBirth
		} else {
			http.Error(w, "Invalid date of birth format", http.StatusBadRequest)
			return
		}
	}

	// Update the employee data in Firestore
	_, err = variables.Client.Collection("employees").Doc(id).Set(context.Background(), existingEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated employee as the response
	w.Header().Set("Content-Type", "application/json")
	showEmployee := function.EmployeeOutput(existingEmployee)
	json.NewEncoder(w).Encode(showEmployee)
}
