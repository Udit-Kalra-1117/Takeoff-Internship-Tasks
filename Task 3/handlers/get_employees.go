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
	functions.HTTP("GetEmployees", GetEmployeesHandler)
}

// getEmployeesHandler retrieves all employees from the database
func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Query all employees from Firestore
	docs, err := variables.Client.Collection("employees").Documents(context.Background()).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Iterate over the documents and deserialize them into Employee objects
	// var employees []Employee
	for _, doc := range docs {
		var employee structure.Employee
		doc.DataTo(&employee)
		w.Header().Set("Content-Type", "application/json")
		showEmployee := function.EmployeeOutput(employee)
		json.NewEncoder(w).Encode(showEmployee)
		// employees = append(employees, employee)
	}

	// // Return the employees as the response
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(employees)
}
