package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

//don't change module example.com/gcf in cloud function's go.mod file add other lines than module line from ur local go.mod file to the cloud's go.mod file

func main() {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "rest-api-391313")
	if err != nil {
		fmt.Printf("Failed to create a firestore client: %v", err)
	}
	defer client.Close()

	fmt.Println("Main function of EMS is executing...")
	// router := mux.NewRouter()                                              // server initialization
	// router.HandleFunc("/employees", getAllEmployeesHandler).Methods("GET") // handling request
	// router.HandleFunc("/employees/{id}", deleteEmployeeHandler).Methods("DELETE")
	// router.HandleFunc("/employees", addEmployeeHandler).Methods("POST")
	// router.HandleFunc("/employees/{id}", updateEmployeeHandler).Methods("PUT")
	// router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// log.Println("Server started on port 8080")
	// log.Fatal(http.ListenAndServe(":8080", router)) // server started
}
