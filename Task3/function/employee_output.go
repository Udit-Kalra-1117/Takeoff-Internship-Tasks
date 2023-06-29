package function

import "github.com/uditkalra/emsGcpApi/structure"

// Create the response object
func EmployeeOutput(employee structure.Employee) structure.EmployeeResponse {
	return structure.EmployeeResponse{
		ID:           employee.ID,
		Name:         employee.Name,
		IsAdmin:      employee.IsAdmin,
		Email:        employee.Email,
		PhoneNumber:  employee.PhoneNumber,
		Department:   employee.Department,
		Role:         employee.Role,
		DateOfBirth:  employee.DateOfBirth,
		CreationTime: employee.CreationTime,
	}
}
