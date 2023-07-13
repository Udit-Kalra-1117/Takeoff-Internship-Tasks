package main

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
)

func getAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "rest-api-391313")
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Retrieve all employees from the Firestore collection
	iter := client.Collection("employees").Documents(ctx)
	var employees []Employee

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}

		var employee Employee
		err = doc.DataTo(&employee)
		if err != nil {
			http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
			return
		}

		employee.Password = ""
		employees = append(employees, employee)
	}

	// Convert employees slice to JSON
	employeesJSON, err := json.Marshal(employees)
	if err != nil {
		http.Error(w, "Failed to convert employees to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(employeesJSON)
}
