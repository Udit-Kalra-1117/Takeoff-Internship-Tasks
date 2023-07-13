package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "rest-api-391313")
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Parse the document ID from the URL path
	id := r.URL.Path[len("/employees/"):]

	// Get the employee document by document ID
	docRef := client.Collection("employees").Doc(id)
	doc, err := docRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// Employee not found, return an error message
			errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
			http.Error(w, errMsg, http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	// Check if the employee document exists
	if !doc.Exists() {
		errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}

	// Parse the employee data into an Employee struct
	var employee Employee
	err = doc.DataTo(&employee)
	if err != nil {
		http.Error(w, "Failed to parse employee data", http.StatusInternalServerError)
		return
	}

	// Exclude the password field from the JSON response
	employee.Password = ""

	// Convert the employee struct to JSON
	employeeJSON, err := json.Marshal(employee)
	if err != nil {
		http.Error(w, "Failed to convert response data to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(employeeJSON)
}
