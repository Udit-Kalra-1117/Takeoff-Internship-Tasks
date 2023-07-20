package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// declaring google cloud project id
var project_id = "task5-393405"

func init() {
	functions.HTTP("UpdateITEM", UpdateITEM)
}

// UpdateITEM is an HTTP Cloud Function with a request parameter to update an entry of a grocery item in the firestore.
func UpdateITEM(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Failed to retrieve grocery item record", http.StatusInternalServerError)
		return
	}

	// Check if the grocery item document exists
	if !doc.Exists() {
		errMsg := fmt.Sprintf("Grocery item record with ID %s does not exist", id)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}

	// Parse the existing grocery item record data into an Grocery_item struct
	var item_record Grocery_item
	err = doc.DataTo(&item_record)
	if err != nil {
		http.Error(w, "Failed to parse already existing grocery item record's data", http.StatusInternalServerError)
		return
	}

	err = r.ParseMultipartForm(20 << 20) // setting the maximum limit of image that can be uploaded to be 20 MB
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

	// checking if name parameter is passed and if it is passed it is passed as a string value only
	name := r.FormValue("name")
	if name != "" {
		if !IsString(name) {
			http.Error(w, "The value of name parameter should be a string consisting of only alphabets", http.StatusBadRequest)
			return
		}
		item_record.Name = NormalizeName(name)
	}

	// checking if price parameter is passed and if it is passed it is passed as an integer value only
	price := r.FormValue("price")
	if price != "" {
		price_int, err := strconv.Atoi(price)
		if err != nil {
			http.Error(w, "Value of price parameter should be passed as an integer value in the request", http.StatusBadRequest)
			return
		}
		item_record.Price = price_int
	}

	// checking if category parameter is passed and if it is passed it is passed as a string value only
	category := r.FormValue("category")
	if category != "" {
		if !IsString(category) {
			http.Error(w, "The value of category parameter should be a string consisting of only alphabets", http.StatusBadRequest)
			return
		}
		item_record.Category = NormalizeName(category)
	}

	// checking if weight parameter is passed and if it is passed it is passed as a string value only
	weight := r.FormValue("weight")
	if weight != "" {
		item_record.Weight = weight
	}

	// checking if veg parameter is passed and if it is passed it is passed as a boolean value only
	veg := r.FormValue("veg")
	if veg != "" {
		veg_bool, err := strconv.ParseBool(veg)
		if err != nil {
			http.Error(w, "Value of vegetarian parameter should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		item_record.Veg = veg_bool
	}

	// checking if brand parameter is passed and if it is passed it is passed as a string value only
	brand := r.FormValue("brand")
	if brand != "" {
		if !IsString(brand) {
			http.Error(w, "The value of brand parameter should be a string consisting of only alphabets", http.StatusBadRequest)
			return
		}
		item_record.Brand = NormalizeName(brand)
	}

	// checking if quantity parameter is passed and if it is passed it is passed as an integer value only
	quantity := r.FormValue("quantity")
	if quantity != "" {
		quantity_int, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Value of quantity parameter should be passed as an integer value in the request", http.StatusBadRequest)
			return
		}
		item_record.Quantity = quantity_int
	}

	// checking if packet info parameter is passed and if it is passed it is passed as a string value only
	packaging := r.FormValue("pack_info")
	if packaging != "" {
		if !IsString(packaging) {
			http.Error(w, "The value of pack info parameter should be a string consisting of only alphabets", http.StatusBadRequest)
			return
		}
		item_record.Pack_info = NormalizeName(packaging)
	}

	// checking if manufacturer parameter is passed and if it is passed it is passed as a string value only
	manufacturer := r.FormValue("manufacturer")
	if manufacturer != "" {
		if !IsString(manufacturer) {
			http.Error(w, "The value of manufacturer parameter should be a string consisting of only alphabets", http.StatusBadRequest)
			return
		}
		item_record.Manufacturer = NormalizeName(manufacturer)
	}

	// checking if country of origin parameter is passed and if it is passed it is passed as a string value only
	origin_country := r.FormValue("country_origin")
	if origin_country != "" {
		if !IsString(origin_country) {
			http.Error(w, "The value of country of origin parameter should be a string consisting of only alphabets", http.StatusBadRequest)
			return
		}
		item_record.Country_origin = NormalizeName(origin_country)
	}

	// checking if availability parameter is passed and if it is passed it is passed as a boolean value only
	availability := r.FormValue("availability")
	if availability != "" {
		availability_bool, err := strconv.ParseBool(availability)
		if err != nil {
			http.Error(w, "Value of availability parameter should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		item_record.Availability = availability_bool
	}

	// checking if discount parameter is passed and if it is passed it is passed as a boolean value only
	discount := r.FormValue("discount")
	if discount != "" {
		discount_bool, err := strconv.ParseBool(discount)
		if err != nil {
			http.Error(w, "Value of discount parameter should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		item_record.Discount = discount_bool
	}

	// checking if offers parameter is passed and if it is passed it is passed as a boolean value only
	offers := r.FormValue("offers")
	if offers != "" {
		offers_bool, err := strconv.ParseBool(offers)
		if err != nil {
			http.Error(w, "Value of offers parameter should be passed as a boolean value in the request", http.StatusBadRequest)
			return
		}
		item_record.Offers = offers_bool
	}

	buc_name := "udit-grocery-items-images-task5-bucket"

	// reading the image file if uploaded in the request
	file, _, err := r.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			// image file not found in the request so skip updating the image
			url := fmt.Sprintf("https://storage.googleapis.com/%s/%s.jpg?%d", buc_name, docRef.ID, random_int())
			item_record.Image_url = url
		} else {
			http.Error(w, "Failed to read image file", http.StatusBadRequest)
			return
		}
	} else {
		defer file.Close()
		item_record.Image = &file

		//initializing cloud storage client
		ctx_storage := context.Background()
		client_storage, err := storage.NewClient(ctx_storage)
		if err != nil {
			http.Error(w, "Failed to create a cloud storage bucket", http.StatusInternalServerError)
			return
		}
		defer client_storage.Close()

		// assigning the name of the image with the unique item record id
		imageFileName := fmt.Sprintf("%s.jpg", docRef.ID)

		// defining a globally unique bucket name
		bucket_name := "udit-grocery-items-images-task5-bucket"

		// creating universally accessible i.e. publicly accessible bucket
		bucket, err := createPublicBucket(ctx_storage, client_storage, bucket_name)
		if err != nil {
			http.Error(w, "Failed to create cloud storage bucket in the given project", http.StatusInternalServerError)
			return
		}

		// creating object with the image file name variable assigned few steps above
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

		item_record.Image_url = imageURL
	}

	// updating the grocery item record document in the firestore database
	_, err = docRef.Set(ctx, item_record)
	if err != nil {
		http.Error(w, "Failed to update the grocery item document in the firestore database", http.StatusInternalServerError)
		return
	}

	// return the success response along the unique id of the record created
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Grocery item entry record updated in the firestore database with the item id: " + item_record.ID,
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
