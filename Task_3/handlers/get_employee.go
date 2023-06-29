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
	functions.HTTP("GetEmployee", GetEmployeeHandler)
}

// getEmployeeHandler retrieves an employee by ID
func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the employee ID from the request URL
	id := r.URL.Path[len("/employee/"):]

	// Get the employee data from Firestore
	doc, err := variables.Client.Collection("employees").Doc(id).Get(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Deserialize the document data into an Employee object
	var employee structure.Employee
	doc.DataTo(&employee)

	// Return the employee as the response
	w.Header().Set("Content-Type", "application/json")
	showEmployee := function.EmployeeOutput(employee)
	json.NewEncoder(w).Encode(showEmployee)
}
