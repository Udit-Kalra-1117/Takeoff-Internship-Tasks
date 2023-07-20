package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// declaring the google project id
var project_id = "task5-393405"

func init() {
	functions.HTTP("GetAnItemByID", GetAnItemByID)
}

// GetAnITEM is an HTTP Cloud Function to get details of a grocery item record from the firestore client by getting the ID of the grocery item record
func GetAnItemByID(w http.ResponseWriter, r *http.Request) {
	// Initializing the firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, project_id)
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Parse the document ID from the URL path
	id := r.URL.Path[len("/grocery_items_database/"):]

	// Get the grocery item document by document ID
	docRef := client.Collection("grocery_items_database").Doc(id)
	doc, err := docRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// Grocery item record not found, return an error message
			errMsg := fmt.Sprintf("Grocery item record with ID %s does not exist", id)
			http.Error(w, errMsg, http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve grocery item record from database", http.StatusInternalServerError)
		return
	}

	// Check if the grocery item record document exists
	if !doc.Exists() {
		errMsg := fmt.Sprintf("Grocery item record with ID %s does not exist", id)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}

	// Parse the grocery item record data into an Grocery item struct
	var item Grocery_item
	err = doc.DataTo(&item)
	if err != nil {
		http.Error(w, "Failed to parse grocery item record data from the firestore", http.StatusInternalServerError)
		return
	}

	// Convert the employee struct to JSON
	itemJSON, err := json.Marshal(item)
	if err != nil {
		http.Error(w, "Failed to convert response data to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(itemJSON)
}
