package helloworld

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/api/iterator"
)

// declaring google cloud project id
var project_ID = "final-project-393405"

func init() {
	functions.HTTP("GetAllItemsByQuery", GetAllItemsByQuery)
}

// GetAllItemsByQuery is an HTTP Cloud Function with a request query parameter to get/display all the entries of grocery item
// which satisfy all the query conditions as provided by the user.
func GetAllItemsByQuery(w http.ResponseWriter, r *http.Request) {
	// Initializing the firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, project_ID)
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Retrieve query parameters from the request
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")
	price := r.URL.Query().Get("price")

	// create a based on the query parameter
	query := client.Collection("grocery_items_database").Query

	// for product name parameter
	if name != "" {
		query = query.Where("Product_Name", "==", name)
	}

	// for product category parameter
	if category != "" {
		query = query.Where("Category", "==", category)
	}

	// for product price parameter
	if price != "" {
		price_int, err := strconv.Atoi(price)
		if err != nil {
			http.Error(w, "Price parameter should be passed as numbers only", http.StatusBadRequest)
			return
		}
		query = query.Where("Price (Rs)", "==", price_int)
	}

	// Retrieve the items that match the query
	iter := query.Documents(ctx)
	var items []Grocery_item
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Error iterating through documents: %v", err)
			http.Error(w, "Failed to retrieve items", http.StatusInternalServerError)
			return
		}

		var item Grocery_item
		err = doc.DataTo(&item)
		if err != nil {
			log.Printf("Failed to parse item data: %v", err)
			http.Error(w, "Failed to parse items data", http.StatusInternalServerError)
			return
		}

		items = append(items, item)
	}

	// Checking if the query returned any items or is the results of get all cloud function is an empty list
	var responseJSON []byte
	if len(items) > 0 {
		// Convert items to JSON response
		responseJSON, err = json.Marshal(items)
		if err != nil {
			log.Printf("Failed to convert items to JSON: %v", err)
			http.Error(w, "Failed to convert items to JSON", http.StatusInternalServerError)
			return
		}
	} else {
		// Empty list case: write an empty JSON array
		responseJSON = []byte("[]")
	}

	// Set response headers and write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
