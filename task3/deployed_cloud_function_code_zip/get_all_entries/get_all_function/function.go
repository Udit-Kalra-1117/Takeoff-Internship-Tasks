package helloworld

import (
	"context"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("getAllEMP", getAllEMP)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func getAllEMP(w http.ResponseWriter, r *http.Request) {
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
	// iter := client.Collection("employees123").Documents(ctx)
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

type Employee struct {
	ID          string `json:"id" firestore:"ID"`
	Name        string `json:"name" parsing:"required" firestore:"Name"`
	Password    string `json:"-" firestore:"Password"`
	IsAdmin     bool   `json:"is_admin" parsing:"required" firestore:"Is Admin"`
	Email       string `json:"email" parsing:"required" firestore:"Email"`
	PhoneNumber string `json:"phone_number" parsing:"required" firestore:"Phone Number"`
	Department  string `json:"department" parsing:"required" firestore:"Department"`
	Role        string `json:"role" parsing:"required" firestore:"Role"`
	DateOfBirth string `json:"date_of_birth" parsing:"required" firestore:"Date of Birth"`
}
