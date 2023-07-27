package helloworld

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

// declaring google cloud project id
var project_id = "final-project-393405"

func init() {
	functions.HTTP("BulkCreateItemsUsingCSV", BulkCreateItemsUsingCSV)
}

// BulkCreateItemsUsingCSV is an HTTP Cloud Function that accepts a CSV file containing grocery item details and creates the items in Firestore
func BulkCreateItemsUsingCSV(w http.ResponseWriter, r *http.Request) {
	// Parse the CSV file from the request
	csv_file, _, err := r.FormFile("csv")
	if err != nil {
		http.Error(w, "Failed to read the CSV file", http.StatusBadRequest)
		return
	}
	defer csv_file.Close()

	// Initializing Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, project_id)
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Get the total count of existing documents in the collection
	totalCount, err := Get_documents_count(ctx, client)
	if err != nil {
		http.Error(w, "Failed to get document count from Firestore", http.StatusInternalServerError)
		return
	}

	// Read the CSV file
	reader := csv.NewReader(csv_file)

	// Skip the first row (i.e. they are column headers)
	_, err = reader.Read()
	if err != nil {
		http.Error(w, "Failed to read CSV file", http.StatusInternalServerError)
		return
	}

	for {
		// Read each row from the CSV file
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			http.Error(w, "Failed to read CSV file", http.StatusInternalServerError)
			return
		}

		// Generate the new ID for the item
		newID := strconv.Itoa(totalCount + 1)

		// Create a new grocery item based on the row data
		item := Grocery_item{
			ID:             "GROCERY_ITEM_" + newID,
			Name:           row[0],
			Price:          Parse_int(row[1]),
			Category:       row[2],
			Weight:         row[3],
			Veg:            Parse_bool(row[4]),
			Brand:          row[5],
			Quantity:       Parse_int(row[6]),
			Pack_info:      row[7],
			Manufacturer:   row[8],
			Country_origin: row[9],
		}

		// Check if the values for availability, discount, and offers are provided in the CSV file
		if len(row) > 10 {
			item.Availability = Parse_bool(row[10])
		}
		if len(row) > 11 {
			item.Discount = Parse_bool(row[11])
		}
		if len(row) > 12 {
			item.Offers = Parse_bool(row[12])
		}

		// Create the grocery item in Firestore
		docRef := client.Collection("grocery_items_database").Doc(item.ID)
		_, err = docRef.Set(ctx, item)
		if err != nil {
			http.Error(w, "Failed to create grocery item in Firestore", http.StatusInternalServerError)
			return
		}

		// Increment the total count
		totalCount++
	}

	// return the success response along the unique id of the record created
	response := struct {
		Message string `json:"message"`
	}{
		Message: "New grocery item entries using bulk create operation successful.",
	}

	// convert response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to convert response to json", http.StatusInternalServerError)
		return
	}

	// setting response content type to application/json
	w.Header().Set("Content-type", "application/json")

	// write the JSON response
	w.WriteHeader(http.StatusCreated)
	w.Write(responseJSON)
}
