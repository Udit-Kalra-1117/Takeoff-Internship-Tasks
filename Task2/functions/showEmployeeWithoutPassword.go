package functions

import "github.com/uditkalra/swaggerRestApi/structure"

// function to convert the entered employee details to the ShowEmployee structure
func ConvertToShowEmployee(employee structure.Employee) structure.ShowEmployee {
	return structure.ShowEmployee{
		ID:          employee.ID,
		Name:        employee.Name,
		IsAdmin:     employee.IsAdmin,
		Email:       employee.Email,
		PhoneNumber: employee.PhoneNumber,
		Department:  employee.Department,
		Role:        employee.Role,
		DateOfBirth: employee.DateOfBirth,
	}
}
