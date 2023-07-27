package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// declaring google cloud project id
var project_id = "final-project-393405"

func init() {
	functions.HTTP("DeleteItemByID", DeleteItemByID)
}

// DeleteItemByID is an HTTP Cloud Function with a request parameter to delete an existing entry of grocery item from the firestore
func DeleteItemByID(w http.ResponseWriter, r *http.Request) {
	// Parse the document ID from the URL path
	id := r.URL.Path[len("/grocery_items_database/"):]

	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, project_id)
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Delete the grocery item document by document ID by checking if it exists or not
	docRef := client.Collection("grocery_items_database").Doc(id)
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		if docSnapshot != nil && !docSnapshot.Exists() {
			// Grocery item not found, return an error message
			errMsg := fmt.Sprintf("Grocery item with ID %s does not exist", id)
			http.Error(w, errMsg, http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to delete grocery item record", http.StatusInternalServerError)
		return
	}

	// Delete the grocery item document if it exists
	_, err = docRef.Delete(ctx)
	if err != nil {
		http.Error(w, "Failed to delete grocery item record", http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]string{"message": "Grocery item record with ID " + id + " deleted successfully"}
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
