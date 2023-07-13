// package helloworld

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"reflect"
// 	"regexp"
// 	"time"

// 	"cloud.google.com/go/firestore"
// 	// "github.com/GoogleCloudPlatform/functions-framework-go/functions"
// 	"golang.org/x/crypto/bcrypt"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/status"
// )

// func init() {
// 	functions.HTTP("updateEMP", updateEMP)
// }

// // helloHTTP is an HTTP Cloud Function with a request parameter.
// func updateEMP(w http.ResponseWriter, r *http.Request) {
// 	// Parse the document ID from the URL path
// 	id := r.URL.Path[len("/employees/"):]

// 	// Initialize Firestore client
// 	ctx := context.Background()
// 	client, err := firestore.NewClient(ctx, "rest-api-391313")
// 	if err != nil {
// 		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
// 		return
// 	}
// 	defer client.Close()

// 	// Get the employee document by document ID
// 	docRef := client.Collection("employees").Doc(id)
// 	doc, err := docRef.Get(ctx)
// 	if err != nil {
// 		if status.Code(err) == codes.NotFound {
// 			// Employee not found, return an error message
// 			errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
// 			http.Error(w, errMsg, http.StatusNotFound)
// 			return
// 		}
// 		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
// 		return
// 	}

// 	// Check if the employee document exists
// 	if !doc.Exists() {
// 		errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
// 		http.Error(w, errMsg, http.StatusNotFound)
// 		return
// 	}

// 	// Parse the existing employee data into an Employee struct
// 	var employee Employee
// 	err = doc.DataTo(&employee)
// 	if err != nil {
// 		http.Error(w, "Failed to parse employee data", http.StatusInternalServerError)
// 		return
// 	}

// 	// Decode the request body to get the updated employee data
// 	var updatedEmployee Employee
// 	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
// 	if err != nil {
// 		http.Error(w, "Failed to decode request data", http.StatusBadRequest)
// 		return
// 	}

// 	// Validating and updating fields if provided

// 	// ID
// 	if updatedEmployee.ID != "" {
// 		http.Error(w, "Don't pass ID parameter in request body as it cannot be changed and ID parameter is auto-generated", http.StatusBadRequest)
// 		return
// 	}

// 	// Name
// 	if updatedEmployee.Name != "" {
// 		if !IsString(updatedEmployee.Name) {
// 			http.Error(w, "Invalid data type for Name field.\nPlease enter name as a string field", http.StatusBadRequest)
// 			return
// 		}
// 		employee.Name = updatedEmployee.Name

// 	}

// 	// Password
// 	if updatedEmployee.Password != "" {
// 		if !IsString(updatedEmployee.Password) {
// 			http.Error(w, "Invalid data type for Password field.\nPlease enter password as a string field", http.StatusBadRequest)
// 			return
// 		}
// 		// Hash the password before storing it
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedEmployee.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
// 			return
// 		}
// 		employee.Password = string(hashedPassword)
// 	}

// 	// Is Admin
// 	if updatedEmployee.IsAdmin != employee.IsAdmin {
// 		if reflect.TypeOf(updatedEmployee.IsAdmin).Kind() != reflect.Bool {
// 			http.Error(w, "Invalid data type for admin field.\nPlease enter is_admin field as a boolean value.\nEg: false", http.StatusBadRequest)
// 			return
// 		}
// 		employee.IsAdmin = updatedEmployee.IsAdmin
// 	}

// 	// Email
// 	if updatedEmployee.Email != "" {
// 		if !UpdatedValidEmail(updatedEmployee.Email) {
// 			http.Error(w, "Invalid email format.\nPlease enter email in string format.\nEg: \"abcd123@example.com\" ", http.StatusBadRequest)
// 			return
// 		}
// 		employee.Email = updatedEmployee.Email
// 	}

// 	// Phone Number
// 	if updatedEmployee.PhoneNumber != "" {
// 		if !UpdatedValidPhoneNumber(updatedEmployee.PhoneNumber) {
// 			http.Error(w, "Invalid phone number format.\nPlease enter phone_number in string format.\nEg: \"+911234567890\"", http.StatusBadRequest)
// 			return
// 		}
// 		employee.PhoneNumber = updatedEmployee.PhoneNumber
// 	}

// 	// Department
// 	if updatedEmployee.Department != "" {
// 		if !IsString(updatedEmployee.Department) {
// 			http.Error(w, "Invalid data type for Department field.\nPlease enter Department as a string field.", http.StatusBadRequest)
// 			return
// 		}
// 		employee.Department = updatedEmployee.Department
// 	}

// 	// Role
// 	if updatedEmployee.Role != "" {
// 		if !IsString(updatedEmployee.Role) {
// 			http.Error(w, "Invalid data type for Role field.\nPlease enter Role as a string field.", http.StatusBadRequest)
// 			return
// 		}
// 		employee.Role = updatedEmployee.Role
// 	}

// 	// Date of Birth
// 	if updatedEmployee.DateOfBirth != "" {
// 		if !UpdatedValidDateOfBirth(updatedEmployee.DateOfBirth) {
// 			http.Error(w, "Invalid date of birth format\nPlease enter date_of_birth in string format.\nEg: \"YYYY-MM-DD\"", http.StatusBadRequest)
// 			return
// 		}
// 		employee.DateOfBirth = updatedEmployee.DateOfBirth
// 	}

// 	// Update the employee document in Firestore
// 	_, err = docRef.Set(context.Background(), employee)
// 	if err != nil {
// 		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
// 		return
// 	}

// 	// Return success response
// 	response := map[string]string{"message": "Employee with ID " + id + " updated successfully"}
// 	responseJSON, err := json.Marshal(response)
// 	if err != nil {
// 		http.Error(w, "Failed to convert response data to JSON", http.StatusInternalServerError)
// 		return
// 	}

// 	// Set the response content type to application/json
// 	w.Header().Set("Content-Type", "application/json")

// 	// Write the JSON response
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(responseJSON)
// }

// // UpdatedValidEmail checks if an email address is in the correct format
// func UpdatedValidEmail(email string) bool {
// 	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
// 	return regex.MatchString(email)
// }

// // UpdatedValidPhoneNumber checks if a phone number starts with "+" and has a length of digits between 7 and 15
// func UpdatedValidPhoneNumber(phoneNumber string) bool {
// 	regex := regexp.MustCompile(`^\+[0-9]{7,15}$`)
// 	return regex.MatchString(phoneNumber)
// }

// // UpdatedValidDateOfBirth checks if a date of birth is in the format "YYYY-MM-DD"
// func UpdatedValidDateOfBirth(dateOfBirth string) bool {
// 	_, err := time.Parse("2006-01-02", dateOfBirth)
// 	return err == nil
// }

// func IsString(value interface{}) bool {
// 	_, ok := value.(string)
// 	return ok
// }

// type Employee struct {
// 	ID          string `json:"id" firestore:"ID"`
// 	Name        string `json:"name" parsing:"required" firestore:"Name"`
// 	Password    string `json:"password" parsing:"required" firestore:"Password"`
// 	IsAdmin     bool   `json:"is_admin" parsing:"required" firestore:"Is Admin"`
// 	Email       string `json:"email" parsing:"required" firestore:"Email"`
// 	PhoneNumber string `json:"phone_number" parsing:"required" firestore:"Phone Number"`
// 	Department  string `json:"department" parsing:"required" firestore:"Department"`
// 	Role        string `json:"role" parsing:"required" firestore:"Role"`
// 	DateOfBirth string `json:"date_of_birth" parsing:"required" firestore:"Date of Birth"`
// }
