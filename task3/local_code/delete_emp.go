package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
)

func deleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the document ID from the URL path
	id := r.URL.Path[len("/employees/"):]

	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "rest-api-391313")
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Delete the employee document by document ID
	docRef := client.Collection("employees").Doc(id)
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		if docSnapshot != nil && !docSnapshot.Exists() {
			// Employee not found, return an error message
			errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
			http.Error(w, errMsg, http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	// Delete the employee document
	_, err = docRef.Delete(ctx)
	if err != nil {
		http.Error(w, "Failed tp delete employee", http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]string{"message": "Employee with ID " + id + " deleted successfully"}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to convert response data to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
