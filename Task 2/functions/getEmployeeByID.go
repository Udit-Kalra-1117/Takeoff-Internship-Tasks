package functions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/uditkalra/swaggerRestApi/variables"
)

// GetEmployeeByID retrieves an employee by ID
// @Summary Get an employee by ID
// @Description Get an employee with the provided ID
// @Tags Employees
// @Param id path int true "Employee ID"
// @Produce json
// @Success 200 {object} structure.ShowEmployee
// @Failure 400 {object} views.ErrorResponse
// @Failure 404 {object} views.ErrorResponse
// @Router /api/v1/employees/{id} [get]
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	for _, employee := range variables.Slice_of_employees {
		if employee.ID == id {
			showEmployee := ConvertToShowEmployee(employee)

			// Return the showEmployee instead of employee
			json.NewEncoder(w).Encode(showEmployee)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}
