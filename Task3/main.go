package main

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/uditkalra/emsGcpApi/handlers"
)

// var Client *firestore.Client

func main() {
	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "rest-api-391313")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Define API routes
	http.HandleFunc("/employee", handlers.CreateEmployeeHandler)
	http.HandleFunc("/employee", handlers.GetEmployeesHandler)
	http.HandleFunc("/employee/{id}", handlers.GetEmployeeHandler)
	http.HandleFunc("/employee/{id}", handlers.UpdateEmployeeHandler)
	http.HandleFunc("/employee/{id}", handlers.DeleteEmployeeHandler)

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
