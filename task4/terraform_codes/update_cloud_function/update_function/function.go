package helloworld

import (
  "encoding/json"
  "fmt"
  "time"
  "regexp"
  "net/http"
  "context"

  "cloud.google.com/go/firestore"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
  "github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
  functions.HTTP("updateEMP", updateEMP)
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func updateEMP(w http.ResponseWriter, r *http.Request) {
  // Parse the document ID from the URL path
	id := r.URL.Path[len("/employee_database/"):]

	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "terraform-task-392713")
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Get the employee document by document ID
	docRef := client.Collection("employee_database").Doc(id)
	doc, err := docRef.Get(ctx)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			// Employee not found, return an error message
			errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
			http.Error(w, errMsg, http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	// Check if the employee document exists
	if !doc.Exists() {
		errMsg := fmt.Sprintf("Employee with ID %s does not exist", id)
		http.Error(w, errMsg, http.StatusNotFound)
		return
	}

	// Parse the existing employee data into an Employee struct
	var employee Employee
	err = doc.DataTo(&employee)
	if err != nil {
		http.Error(w, "Failed to parse already existing employee data", http.StatusInternalServerError)
		return
	}

	// Decode the request body to get the updated employee data
	var updatedEmployee Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, "Failed to decode request data.\nPlease enter the value of all the parameters of Employee structure as a string except is_admin.\nThe value of is_admin parameter is to be entered as a boolean.", http.StatusBadRequest)
		return
	}

	// Validating and updating fields if provided

	// ID
	if updatedEmployee.ID != "" {
		http.Error(w, "Don't pass ID parameter in request body as cannot change ID parameter and ID parameter is auto-generated", http.StatusBadRequest)
		return
	}

	// Name
	if updatedEmployee.Name != "" {
		employee.Name = updatedEmployee.Name
	}

	// Password
	if updatedEmployee.Password != "" {
		// Hash the password before storing it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedEmployee.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		employee.Password = string(hashedPassword)
	}

	// Is Admin
	if updatedEmployee.IsAdmin != employee.IsAdmin {
		employee.IsAdmin = updatedEmployee.IsAdmin
	}

	// Email
	if updatedEmployee.Email != "" {
		if !UpdatedValidEmail(updatedEmployee.Email) {
			http.Error(w, "Invalid email format", http.StatusBadRequest)
			return
		}
		employee.Email = updatedEmployee.Email
	}

	// Phone Number
	if updatedEmployee.PhoneNumber != "" {
		if !UpdatedValidPhoneNumber(updatedEmployee.PhoneNumber) {
			http.Error(w, "Invalid phone number format", http.StatusBadRequest)
			return
		}
		employee.PhoneNumber = updatedEmployee.PhoneNumber
	}

	// Department
	if updatedEmployee.Department != "" {
		employee.Department = updatedEmployee.Department
	}

	// Role
	if updatedEmployee.Role != "" {
		employee.Role = updatedEmployee.Role
	}

	// Date of Birth
	if updatedEmployee.DateOfBirth != "" {
		if !UpdatedValidDateOfBirth(updatedEmployee.DateOfBirth) {
			http.Error(w, "Invalid date of birth format", http.StatusBadRequest)
			return
		}
		employee.DateOfBirth = updatedEmployee.DateOfBirth
	}

	// Update the employee document in Firestore
	_, err = docRef.Set(context.Background(), employee)
	if err != nil {
		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}

	// Return success response
	response := map[string]string{"message": "Employee with ID " + id + " updated successfully"}
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

// UpdatedValidEmail checks if an email address is in the correct format
func UpdatedValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// UpdatedValidPhoneNumber checks if a phone number starts with "+" and has a length of digits between 7 and 15
func UpdatedValidPhoneNumber(phoneNumber string) bool {
	regex := regexp.MustCompile(`^\+[0-9]{7,15}$`)
	return regex.MatchString(phoneNumber)
}

// UpdatedValidDateOfBirth checks if a date of birth is in the format "YYYY-MM-DD"
func UpdatedValidDateOfBirth(dateOfBirth string) bool {
	_, err := time.Parse("2006-01-02", dateOfBirth)
	return err == nil
}

type Employee struct {
	ID          string `json:"id" firestore:"ID"`
	Name        string `json:"name" parsing:"required" firestore:"Name"`
	Password    string `json:"password" parsing:"required" firestore:"Password"`
	IsAdmin     bool   `json:"is_admin" parsing:"required" firestore:"Is Admin"`
	Email       string `json:"email" parsing:"required" firestore:"Email"`
	PhoneNumber string `json:"phone_number" parsing:"required" firestore:"Phone Number"`
	Department  string `json:"department" parsing:"required" firestore:"Department"`
	Role        string `json:"role" parsing:"required" firestore:"Role"`
	DateOfBirth string `json:"date_of_birth" parsing:"required" firestore:"Date of Birth"`
}