package helloworld

import (
  "encoding/json"
  "strconv"
  "net/http"
  "context"
  "time"
  "regexp"

  "cloud.google.com/go/firestore"
  "golang.org/x/crypto/bcrypt"
  "github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
  functions.HTTP("CreateEMP", createEMP)
}

// isValidEmail checks if an email address is in the correct format
func IsValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

// isValidPhoneNumber checks if a phone number starts with "+" and has a length of digits between 7 and 15
func IsValidPhoneNumber(phoneNumber string) bool {
	regex := regexp.MustCompile(`^\+[0-9]{7,15}$`)
	return regex.MatchString(phoneNumber)
}

// isValidDateOfBirth checks if a date of birth is in the format "YYYY-MM-DD"
func IsValidDateOfBirth(dateOfBirth string) bool {
	_, err := time.Parse("2006-01-02", dateOfBirth)
	return err == nil
}

// helloHTTP is an HTTP Cloud Function with a request parameter.
func createEMP(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON data from the request body into an Employee struct
	var employee Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Failed to decode request data.\nPlease enter the value of all the parameters of Employee structure as a string except is_admin.\nThe value of is_admin parameter is to be entered as a boolean.", http.StatusBadRequest)
		return
	}

	// checking if ID is passed in request body
	if employee.ID != "" {
		http.Error(w, "Don't pass ID parameter in request body as cannot change ID parameter and ID parameter is auto-generated", http.StatusBadRequest)
		return
	}

	// Validate phone number format
	if employee.PhoneNumber != "" && !IsValidPhoneNumber(employee.PhoneNumber) {
		http.Error(w, "Invalid phone number format\nPhone number should be entered as: \"+91123456789\"", http.StatusBadRequest)
		return
	}

	// Validate date of birth format
	if employee.DateOfBirth != "" && !IsValidDateOfBirth(employee.DateOfBirth) {
		http.Error(w, "Invalid date of birth format\nDate of Birth should be entered as: \"2002-10-25\"", http.StatusBadRequest)
		return
	}

	// Validate email format
	if employee.Email != "" && !IsValidEmail(employee.Email) {
		http.Error(w, "Invalid email format\nEmail should be entered as: \"abcd@gmail.com\"", http.StatusBadRequest)
		return
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	employee.Password = string(hashedPassword)

	// Initialize Firestore client
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, "rest-api-391313")
	if err != nil {
		http.Error(w, "Failed to initialize Firestore client", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Check if an employee with the same email or phone number already exists
	query := client.Collection("employees").
		Where("Email", "==", employee.Email).
		Where("Phone Number", "==", employee.PhoneNumber).
		Limit(1)

	existingEmployees, err := query.Documents(ctx).GetAll()
	if err != nil {
		http.Error(w, "Failed to check if employee exists", http.StatusInternalServerError)
		return
	}

	// If an employee with the same email or phone number already exists, return an error
	if len(existingEmployees) > 0 {
		http.Error(w, "Employee with the same Email and Phone Number already exists", http.StatusBadRequest)
		return
	}

	// Get the number of existing employees to generate the new employee ID
	existingEmployees, err = client.Collection("employees").Documents(ctx).GetAll()
	if err != nil {
		http.Error(w, "Failed to get employees", http.StatusInternalServerError)
		return
	}

	// Generate the new employee ID
	newEmployeeID := "EMP" + strconv.Itoa(len(existingEmployees)+1)

	// Set the employee ID and document ID to the generated ID
	employee.ID = newEmployeeID

	// Create a new document in the Firestore collection with the employee data and generated ID
	docRef := client.Collection("employees").Doc(newEmployeeID)
	_, err = docRef.Set(ctx, employee)
	if err != nil {
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	// Return the generated employee ID in the response
	response := struct {
		EmployeeID string `json:"employee_id"`
	}{
		EmployeeID: newEmployeeID,
	}

	// Convert the response to JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to convert response to JSON", http.StatusInternalServerError)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
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