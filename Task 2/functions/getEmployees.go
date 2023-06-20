package functions

import (
	"encoding/json"
	"net/http"

	"github.com/uditkalra/swaggerRestApi/structure"
	"github.com/uditkalra/swaggerRestApi/variables"
)

// GetEmployees retrieves all employees
// @Summary Get all employees
// @Description Get a list of all employees
// @Tags Employees
// @Produce json
// @Success 200 {array} structure.ShowEmployee
// @Router /api/v1/employees [get]
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	showEmployees := make([]structure.ShowEmployee, len(variables.Slice_of_employees))
	for i, employee := range variables.Slice_of_employees {
		showEmployees[i] = ConvertToShowEmployee(employee)
	}

	// Return the showEmployees instead of employees
	json.NewEncoder(w).Encode(showEmployees)
}
