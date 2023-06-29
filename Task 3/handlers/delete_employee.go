package handlers

import (
	"context"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/uditkalra/emsGcpApi/variables"
)

func init() {
	functions.HTTP("DeleteEmployee", DeleteEmployeeHandler)
}

// deleteEmployeeHandler deletes an employee by ID
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the employee ID from the request URL
	id := r.URL.Path[len("/employee/"):]

	// Delete the employee from Firestore
	_, err := variables.Client.Collection("employees").Doc(id).Delete(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Employee deleted successfully"))
}
