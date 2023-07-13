package employee

import (
	"fmt"
	"time"
)

// declaring the employee structure with the desired parameters along with their types
type Employee struct {
	Name        string
	ID          int
	Password    string
	Email       string
	PhoneNo     int
	Department  string
	Role        string
	DateOfBirth time.Time
}

// declaring the employee-user structure with the desired parameters along with their types
type EmployeeUser struct {
	ID       int
	Password string
}

// declaring the employee-management-system structure to store the slices of employees
type EmployeeManagementSystem struct {
	employees []Employee
}

// function to initialize an empty slice of type employee structure
func (ems *EmployeeManagementSystem) Initialize() {
	ems.employees = make([]Employee, 0)
}

// function to add a new employee record
func (ems *EmployeeManagementSystem) AddEmployee(emp Employee) error {
	//checking if the employee with the same ID already exists
	for _, e := range ems.employees {
		if e.ID == emp.ID {
			//if employee exists then return an error
			return fmt.Errorf("Employee with ID %d already exists", emp.ID)
		}
	}
	//if employee with the entered id does not exist then the record with the entered details is added to the slice of employees
	ems.employees = append(ems.employees, emp)
	//return no errors
	return nil
}

// function to delete the employee details with the desired id
func (ems *EmployeeManagementSystem) DeleteEmployee(id int) error {
	//checking if the employee with the desired id exists
	for i, emp := range ems.employees {
		if emp.ID == id {
			//if employee with the desired id exists then the employee is deleted from the slice of employees
			ems.employees = append(ems.employees[:i], ems.employees[i+1:]...)
			//return no errors
			return nil
		}
	}
	//if employee with the desired id does not exist then an error is returned
	return fmt.Errorf("Employee with ID %d not found", id)
}

// function to view/search employee details as per the entered id
func (ems *EmployeeManagementSystem) ViewEmployeeDetails(id int) (Employee, error) {
	//checking if the employee with the desired id exists
	for _, emp := range ems.employees {
		//if employee is found, return the details of the employee and no errors
		if emp.ID == id {
			return emp, nil
		}
	}
	//if employee with the desired id does not exist then an empty slice of employee and an error is returned
	return Employee{}, fmt.Errorf("Employee with ID: %d not found", id)
}

// function to update the employee details as per the entered id
func (ems *EmployeeManagementSystem) UpdateEmployeeDetails(id int, newEmp Employee) error {
	//checking if the employee with the desired id exists
	for i, emp := range ems.employees {
		//if employee is found, update the employee details and return no errors
		if emp.ID == id {
			ems.employees[i] = newEmp
			return nil
		}
	}
	//if employee with the desired id does not exist then an error is returned
	return fmt.Errorf("Employee with ID %d not found", id)
}

// function to display all the employee records stored in the slice of employees
func ListAllEmployees(ems *EmployeeManagementSystem) {
	fmt.Println("\n===== Employee List =====")
	//loop through the employees slice and print the details of all employees
	for _, emp := range ems.employees {
		fmt.Println("--------------------------------------")
		fmt.Printf("\nName: %s, \nID: %d, \nEmail: %s, \nPhone No: %d, \nDepartment: %s, \nRole: %s, \nDate of Birth: %s\n",
			emp.Name, emp.ID, emp.Email, emp.PhoneNo, emp.Department, emp.Role, emp.DateOfBirth.Format("2006-01-02"))
		fmt.Println("--------------------------------------")
	}
	fmt.Println("===========================================")
}

// function to display the details of the employees who have their birthday in the current ongoing month
func ListUpcomingBirthdays(ems *EmployeeManagementSystem) {
	fmt.Println("\n===== Employee(s) with Upcoming Birthdays =====")
	//loop through the employees slice and print the details of all employees which have upcoming birthday in the current month
	for _, emp := range ems.employees {
		//checking if any employee has birthday in the current month and also
		//checking if the birthday is today or in any of the remaining days of the current ongoing month
		if emp.DateOfBirth.Month() == time.Now().Month() && emp.DateOfBirth.Day() >= time.Now().Day() {
			fmt.Println("--------------------------------------")
			fmt.Printf("\nName: %s, \nID: %d, \nEmail: %s, \nPhone No: %d, \nDepartment: %s, \nRole: %s, \nDate of Birth: %s\n",
				emp.Name, emp.ID, emp.Email, emp.PhoneNo, emp.Department, emp.Role, emp.DateOfBirth.Format("2006-01-02"))
			fmt.Println("--------------------------------------")
		}
	}
	fmt.Println("===========================================")
}

// function to search and display the employee details as per the entered name
func (ems *EmployeeManagementSystem) SearchEmployeeByName(name string) (Employee, error) {
	//searching for the employee details with the entered name
	for _, emp := range ems.employees {
		//if employee is found, return the details of the employee and no errors
		if emp.Name == name {
			return emp, nil
		}
	}
	//if employee is not found, return an empty slice and an error
	return Employee{}, fmt.Errorf("Employee with Name: %s not found", name)
}

// function to login an employee with their credentials
func (ems *EmployeeManagementSystem) EmployeeLogin(name, pass string) (EmployeeUser, error) {
	//declaring an empty employee-user structure
	employee := EmployeeUser{}
	//searching for employee with the entered name and password
	for _, emp := range ems.employees {
		//if employee is found, return the employee-user structure with the id and password of the employee and no errors
		if emp.Name == name && emp.Password == pass {
			employee.ID = emp.ID
			employee.Password = emp.Password
			return employee, nil
		}
	}
	//if employee is not found, return an empty employee-user structure and an error
	return EmployeeUser{}, fmt.Errorf("\nInvalid Credentials")
}
