package functions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/uditkalra/swaggerRestApi/csv"
	"github.com/uditkalra/swaggerRestApi/structure"
	"github.com/uditkalra/swaggerRestApi/variables"
	"github.com/uditkalra/swaggerRestApi/views"
	"golang.org/x/crypto/bcrypt"
)

// UpdateEmployee updates an existing employee
// @Summary Update an existing employee
// @Description Update an existing employee with the provided ID and details
// @Tags Employees
// @Accept json
// @Produce json
// @Param id path int true "Employee ID"
// @Param employee body structure.Employee true "Updated employee details"
// @Success 200 {object} structure.ShowEmployee
// @Failure 400 {object} views.ErrorResponse
// @Failure 404 {object} views.ErrorResponse
// @Failure 500 {object} views.ErrorResponse
// @Router /employees/{id} [put]
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Invalid employee ID"})
		return
	}

	var updatedEmployee structure.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Invalid employee data"})
		return
	}

	// Validating the entered email address
	if updatedEmployee.Email != "" && !IsValidEmail(updatedEmployee.Email) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Invalid email address"})
		return
	}

	// Validating the entered phone number
	if updatedEmployee.PhoneNumber != "" && !IsValidPhoneNumber(updatedEmployee.PhoneNumber) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Invalid phone number"})
		return
	}

	// Validating the entered date of birth
	if updatedEmployee.DateOfBirth != "" && !IsValidDateOfBirth(updatedEmployee.DateOfBirth) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Invalid date of birth"})
		return
	}

	for i, employee := range variables.Slice_of_employees {
		if employee.ID == id {
			// Updating the specific parameter if it's provided in the request
			if updatedEmployee.Name != "" {
				employee.Name = updatedEmployee.Name
			}
			if updatedEmployee.Role != "" {
				employee.Role = updatedEmployee.Role
			}
			if !updatedEmployee.IsAdmin {
				employee.IsAdmin = updatedEmployee.IsAdmin
			}
			if updatedEmployee.IsAdmin {
				employee.IsAdmin = updatedEmployee.IsAdmin
			}
			if updatedEmployee.Email != "" {
				employee.Email = updatedEmployee.Email
			}
			if updatedEmployee.PhoneNumber != "" {
				employee.PhoneNumber = updatedEmployee.PhoneNumber
			}
			if updatedEmployee.DateOfBirth != "" {
				employee.DateOfBirth = updatedEmployee.DateOfBirth
			}
			if updatedEmployee.Password != "" {
				// Hashing the newly entered password
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updatedEmployee.Password), bcrypt.DefaultCost)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Failed to hash the password"})
					return
				}
				employee.Password = string(hashedPassword)
			}
			updatedEmployee.ID = id
			variables.Slice_of_employees[i] = employee
			csv.SaveToCSV()
			json.NewEncoder(w).Encode(views.SuccessResponse{Message: "Employee with id " + strconv.Itoa(id) + " has been updated successfully"})
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}
