package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"math/rand"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("CreateITEM", CreateITEM)
}

// declaring the google cloud project id
var project_id = "final-project-393405"

// CreateITEM is an HTTP Cloud Function with a request parameter to create new entry of grocery item.
func CreateITEM(w http.ResponseWriter, r *http.Request) {

	var new_item Grocery_item
	err := r.ParseMultipartForm(20 << 20) // setting the maximum limit of image that can be uploaded to be 20 MB
	if err != nil {
		http.Error(w, "Failed to parse sent request body please re-check all the parameters and their types and resend the request", http.StatusBadRequest)
		return
	}

	// checking if id parameter is not passed in the request
	passed_id := r.FormValue("Id")
	if passed_id != "" {
		http.Error(w, "ID parameter is not required to be passed in the request as it is auto-generated and cannot be changed", http.StatusBadRequest)
		return
	}

	// checking if name parameter is passed and is passed as a string value only
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name parameter is a required field and its value should be passed as a string in the request", http.StatusBadRequest)
		return
	}
	if !IsString(name) {
		http.Error(w, "The value of name parameter should be string consisting of only alphabets", http.StatusBadRequest)
		return
	}
	new_item.Name = NormalizeName(name)

	// checking if price parameter is passed and is passed as an integer value only
	new_item.Price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "Price parameter is a required parameter and its value should be passed as an integer value in the request", http.StatusBadRequest)
		return
	}

	// checking if category parameter is passed and is passed as a string value only
	category := r.FormValue("category")
	if category == "" {
		http.Error(w, "Category parameter is a required field and its value should be passed as a string in the request", http.StatusBadRequest)
		return
	}
	if !IsString(category) {
		http.Error(w, "The value of category parameter should be string consisting of only alphabets", http.StatusBadRequest)
		return
	}
	new_item.Category = NormalizeName(category)

	// checking if weight parameter is passed
	weight := r.FormValue("weight")
	if weight == "" {
		http.Error(w, "Weight parameter is a required field and its value should be passed as a string in the request\nThis value can be of alphanumeric type like: 500ML or 1kg", http.StatusBadRequest)
		return
	}
	new_item.Weight = weight

	// checking if the vegetarian parameter is passed and is passed as a boolean value only
	new_item.Veg, err = strconv.ParseBool(r.FormValue("veg"))
	if err != nil {
		http.Error(w, "Vegetarian parameter is a required field and its value should be passed as a boolean value in the request", http.StatusBadRequest)
		return
	}

	// checking if the brand parameter is passed and is passed as a string value only
	brand := r.FormValue("brand")
	if brand == "" {
		http.Error(w, "Brand parameter is a required field and its value should be passed as a string in the request", http.StatusBadRequest)
		return
	}
	if !IsString(brand) {
		http.Error(w, "The value of brand parameter should be string consisting of only alphabets", http.StatusBadRequest)
		return
	}
	new_item.Brand = NormalizeName(brand)

	// checking if quantity parameter is passed and is passed as an integer value only
	new_item.Quantity, err = strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Error(w, "Quantity parameter is a required parameter and its value should be passed as an integer value in the request", http.StatusBadRequest)
		return
	}

	// checking if the pack info parameter is passed and is passed as a string value only
	pack_info := r.FormValue("pack_info")
	if pack_info == "" {
		http.Error(w, "Pack info parameter is a required field and its value should be passed as a string in the request", http.StatusBadRequest)
		return
	}
	if !IsString(pack_info) {
		http.Error(w, "The value of pack info parameter should be string consisting of only alphabets", http.StatusBadRequest)
		return
	}
	new_item.Pack_info = NormalizeName(pack_info)

	// checking if the manufacturer parameter is passed and is passed as a string value only
	manufacturer := r.FormValue("manufacturer")
	if manufacturer == "" {
		http.Error(w, "Manufacturer parameter is a required field and its value should be passed as a string in the request", http.StatusBadRequest)
		return
	}
	if !IsString(manufacturer) {
		http.Error(w, "The value of manufacturer parameter should be string consisting of only alphabets", http.StatusBadRequest)
		return
	}
	new_item.Manufacturer = NormalizeName(manufacturer)

	// checking if the country origin parameter is passed and is passed as a string value only
	origin := r.FormValue("country_origin")
	if origin == "" {
		http.Error(w, "Country of origin parameter is a required field and its value should be passed as a string in the request", http.StatusBadRequest)
		return
	}
	if !IsString(origin) {
		http.Error(w, "The value of country of origin parameter should be string consisting of only alphabets", http.StatusBadRequest)
		return
	}
	new_item.Country_origin = NormalizeName(origin)

	// checking if the availability parameter is passed and if it is passed it is passed as a boolean value only
	available := r.FormValue("availability")
	if available != "" {
		availability, err := strconv.ParseBool(available)
		if err != nil {
			http.Error(w, "Availability parameter value should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		new_item.Availability = availability
	}

	// checking if the discount parameter is passed and if it is passed it is passed as a boolean value only
	discount := r.FormValue("discount")
	if discount != "" {
		discount_available, err := strconv.ParseBool(discount)
		if err != nil {
			http.Error(w, "Discount parameter value should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		new_item.Discount = discount_available
	}

	// checking if the offers parameter is passed and if it is passed it is passed as a boolean value only
	offers := r.FormValue("offers")
	if offers != "" {
		offers_available, err := strconv.ParseBool(offers)
		if err != nil {
			http.Error(w, "Offers parameter value should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		new_item.Offers = offers_available
	}

	// reading the image file uploaded in the request
	file, _, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			http.Error(w, "Image file not found in the request.\nImage parameter is a required field.\nPlease provide an image file in the image parameter while passing a request", http.StatusBadRequest)
			return
		} else {
			http.Error(w, "Failed to read image file", http.StatusBadRequest)
			return
		}
	}
	defer file.Close()
	new_item.Image = &file

	// initializing firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, project_id)
	if err != nil {
		http.Error(w, "Failed to create a firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	//initializing cloud storage client
	ctx_storage := context.Background()
	client_storage, err := storage.NewClient(ctx_storage)
	if err != nil {
		http.Error(w, "Failed to create a cloud storage bucket", http.StatusInternalServerError)
		return
	}
	defer client_storage.Close()

	// a query to avoid duplicate entries by checking through the records in the database
	// to find a match with the product name, price and weight of the request made
	query := client.Collection("grocery_items_database").
		Where("Product_Name", "==", new_item.Name).
		Where("Price (Rs)", "==", new_item.Price).
		Where("Weight", "==", new_item.Weight).
		Limit(1)

	existing_item, err := query.Documents(ctx).GetAll()
	if err != nil {
		http.Error(w, "Failed to check if grocery items  already exists from the database", http.StatusInternalServerError)
		return
	}

	// if the grocery item record already exists then throw an error
	if len(existing_item) > 0 {
		http.Error(w, "Grocery item with the same Product Name, Weight and Price already exists.\nYou can update the quantity or any parameter by using the appropriate link for updating records", http.StatusBadRequest)
		return
	}

	// calculating the number of records present in the database to generate unique id for records
	existing_item, err = client.Collection("grocery_items_database").Documents(ctx).GetAll()
	if err != nil {
		http.Error(w, "Failed to get Grocery Item", http.StatusInternalServerError)
		return
	}
	// logic to generate id for each new record in the database
	new_id := "GROCERY_ITEM_" + strconv.Itoa(len(existing_item)+1)
	new_item.ID = new_id

	// generate a new document in the firestore with grocery item data and auto generated unique record id
	docRef := client.Collection("grocery_items_database").Doc(new_id)
	_, err = docRef.Set(ctx, new_item)
	if err != nil {
		http.Error(w, "Failed to create a new entry in the database with the provided details of the grocery item", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// assigning the name of the image with the unique item record id
	imageFileName := fmt.Sprintf("%s.jpg", docRef.ID)

	// defining a globally unique bucket name
	bucket_name := "udit-grocery-items-images-final-project-bucket"

	// creating universally accessible i.e. publicly accessible bucket
	bucket, err := createPublicBucket(ctx_storage, client_storage, bucket_name)
	if err != nil {
		http.Error(w, "Failed to create cloud storage bucket in the given project", http.StatusInternalServerError)
		return
	}

	// creating an object with the image file name variable assigned few steps above
	obj := bucket.Object(imageFileName)
	wc := obj.NewWriter(ctx_storage)
	if _, err := io.Copy(wc, file); err != nil {
		http.Error(w, "Failed to upload image to cloud storage", http.StatusInternalServerError)
		return
	}
	if err := wc.Close(); err != nil {
		http.Error(w, "Failed to close cloud storage writer", http.StatusInternalServerError)
		return
	}

	// generating an image url for the uploaded image to assign it against the appropriate firestore grocery item document
	imageURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s?%d", bucket_name, imageFileName, random_int())

	// updating url of the image stored in the publicly accessible cloud storage bucket against the appropriate firestore document
	_, err = docRef.Update(ctx, []firestore.Update{
		{
			Path:  "Image",
			Value: imageURL,
		},
	})
	if err != nil {
		http.Error(w, "Failed to update grocery item with the image URL", http.StatusInternalServerError)
		return
	}

	// return the success response along the unique id of the record created
	response := struct {
		Message string `json:"message"`
	}{
		Message: "New grocery item entry created with the item id as: " + new_item.ID,
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

// function to create a publicly bucket which stores all the images for the grocery item records
func createPublicBucket(ctx context.Context, client *storage.Client, bucketName string) (*storage.BucketHandle, error) {

	bucket := client.Bucket(bucketName)

	// checking if the bucket already exists
	_, err := bucket.Attrs(ctx)
	if err == nil {
		// bucket already exists
		return bucket, nil
	}

	// if not create a new bucket
	if err := bucket.Create(ctx, project_id, &storage.BucketAttrs{
		Location: "us",
	}); err != nil {
		return nil, err
	}

	return bucket, nil
}

// // randomInt generates a random number to be used as a query parameter in image URLs
func random_int() int {
	return rand.Intn(9999999999-1000000000+1) + 1000000000
}
