package functions

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/uditkalra/swaggerRestApi/csv"
	"github.com/uditkalra/swaggerRestApi/variables"
	"github.com/uditkalra/swaggerRestApi/views"
)

// DeleteEmployee deletes an employee by ID
// @Summary Delete an employee by ID
// @Description Delete an employee with the provided ID
// @Tags Employees
// @Param id path int true "Employee ID"
// @Success 200 {object} views.SuccessResponse
// @Failure 400 {object} views.ErrorResponse
// @Failure 404 {object} views.ErrorResponse
// @Router /api/v1/employees/{id} [delete]
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Invalid employee ID"})
		return
	}

	for i, employee := range variables.Slice_of_employees {
		if employee.ID == id {
			variables.Slice_of_employees = append(variables.Slice_of_employees[:i], variables.Slice_of_employees[i+1:]...)
			csv.SaveToCSV()
			json.NewEncoder(w).Encode(views.SuccessResponse{Message: "Employee with id " + strconv.Itoa(id) + " has been deleted successfully"})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(views.ErrorResponse{Error: "Employee not found"})
}
