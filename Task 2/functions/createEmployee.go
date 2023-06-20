package functions

import (
	"encoding/json"
	"net/http"

	"github.com/uditkalra/swaggerRestApi/csv"
	"github.com/uditkalra/swaggerRestApi/structure"
	"github.com/uditkalra/swaggerRestApi/variables"
	"github.com/uditkalra/swaggerRestApi/views"
	"golang.org/x/crypto/bcrypt"
)

// CreateEmployee creates a new employee
// @Summary Create a new employee
// @Description Create a new employee with the provided details
// @Tags Employees
// @Accept json
// @Produce json
// @Param employee body structure.Employee true "Employee details"
// @Success 200 {object} structure.ShowEmployee
// @Failure 400 {object} views.ErrorResponse
// @Failure 500 {object} views.ErrorResponse
// @Router /api/v1/employees [post]
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var employee structure.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse {Error: "Invalid employee data"})
		return
	}

	// Validating the entered email address
	if !IsValidEmail(employee.Email) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse {Error: "Invalid email address"})
		return
	}

	// Validating the entered phone number
	if !IsValidPhoneNumber(employee.PhoneNumber) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse {Error: "Invalid phone number"})
		return
	}

	// Validating the entered date of birth
	if !IsValidDateOfBirth(employee.DateOfBirth) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse {Error: "Invalid date of birth"})
		return
	}

	// Hashing the newly entered password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(employee.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(views.ErrorResponse {Error: "Failed to hash the password"})
		return
	}

	employee.ID = GetNextAvailableID()
	employee.Password = string(hashedPassword)
	variables.Slice_of_employees = append(variables.Slice_of_employees, employee)
	csv.SaveToCSV()
	showEmployee := ConvertToShowEmployee(employee)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(showEmployee)
}
