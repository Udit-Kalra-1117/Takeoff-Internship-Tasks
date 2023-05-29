package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	employee "github.com/uditkalra/ems/Employee"
)

// function to view the logged in employee details
func viewMyDetails(employee *employee.EmployeeUser, ems *employee.EmployeeManagementSystem) {
	//checking if the employee with the given id exists or not by calling the ViewEmployeeDetails from the Employee --> employee.go
	emp, err := ems.ViewEmployeeDetails(employee.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//displaying the employee details
	fmt.Println("--------------------------------------")
	fmt.Println("Name:", emp.Name)
	fmt.Println("ID:", emp.ID)
	// fmt.Println("Password:", emp.Password)
	fmt.Println("Email:", emp.Email)
	fmt.Println("Phone No:", emp.PhoneNo)
	fmt.Println("Department:", emp.Department)
	fmt.Println("Role:", emp.Role)
	fmt.Println("Date of Birth:", emp.DateOfBirth.Format("2006-01-02"))
	fmt.Println("--------------------------------------")
}

// function to update the logged in employee details
func updateMyDetails(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Update Employee Details =====")

	//for security purposes asking the employee its id once again
	fmt.Print("Re-Enter your unique and non-changeable Employee ID: ")
	//take input of id from user until a new line occurs or enter is pressed on the keyboard
	idStr, _ := reader.ReadString('\n')
	//remove any newline character, white-spaces and tabs from entered id
	idStr = strings.TrimSpace(idStr)
	//convert the id to int to pass it to ViewEmployeeDetails function of Employee --> employee.go
	id, _ := strconv.Atoi(idStr)

	//checking if the employee with the given id exists or not by calling the ViewEmployeeDetails from the Employee --> employee.go
	emp, err := ems.ViewEmployeeDetails(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//declaring a variable called newEmp of type Employee structure to store the changes into it against the required employee
	newEmp := employee.Employee{}

	//as id is unique and cannot be changed once assigned so assigning the old id to id in the newEmp variable
	newEmp.ID = emp.ID
	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Name (Leave blank and press Enter to keep existing value: ", emp.Name, "): ")
	inputN, _ := reader.ReadString('\n')
	inputN = strings.TrimSpace(inputN)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputN != "" {
		newEmp.Name = inputN
	} else {
		newEmp.Name = emp.Name
	}

	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Password (Leave blank and press Enter to keep existing value): ")
	inputPa, _ := reader.ReadString('\n')
	inputPa = strings.TrimSpace(inputPa)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputPa != "" {
		newEmp.Password = inputPa
	} else {
		newEmp.Password = emp.Password
	}

	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Email (Leave blank and press Enter to keep existing value: ", emp.Email, "): ")
	inputE, _ := reader.ReadString('\n')
	inputE = strings.TrimSpace(inputE)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputE != "" {
		newEmp.Email = inputE
	} else {
		newEmp.Email = emp.Email
	}

	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Phone No (Leave blank and press Enter to keep existing value: ", emp.PhoneNo, "): ")
	inputPh, _ := reader.ReadString('\n')
	inputPh = strings.TrimSpace(inputPh)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputPh != "" {
		newEmp.PhoneNo = inputPh
	} else {
		newEmp.PhoneNo = emp.PhoneNo
	}

	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Department (Leave blank and press Enter to keep existing value: ", emp.Department, "): ")
	inputD, _ := reader.ReadString('\n')
	inputD = strings.TrimSpace(inputD)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputD != "" {
		newEmp.Department = inputD
	} else {
		newEmp.Department = emp.Department
	}

	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Role (Leave blank and press Enter to keep existing value: ", emp.Role, "): ")
	inputR, _ := reader.ReadString('\n')
	inputR = strings.TrimSpace(inputR)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputR != "" {
		newEmp.Role = inputR
	} else {
		newEmp.Role = emp.Role
	}

	//asking the user to enter new value along with displaying the current stored value
	//and telling user to don't input anything and just press enter to keep the old entry i.e. to not modify the specified parameter
	fmt.Print("Date of Birth (YYYY-MM-DD) (Leave blank and press Enter to keep existing value: ", emp.DateOfBirth.Format("2006-01-02"), "): ")
	dobStr, _ := reader.ReadString('\n')
	dobStr = strings.TrimSpace(dobStr)
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if dobStr != "" {
		//checking if the entered value for date of birth is according to the default format by using the time.Parse function
		dob, erro := time.Parse("2006-01-02", dobStr)
		//if there is error in the entered date of birth then the application is stopped
		if erro != nil {
			fmt.Println("Invalid Date of Birth. Please try again")
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		// if there is no error then the new entered value is stored in the date of birth parameter
		newEmp.DateOfBirth = dob
	} else {
		newEmp.DateOfBirth = emp.DateOfBirth
	}

	//updating the employee record in the employee management system slice
	err = ems.UpdateEmployeeDetails(id, newEmp)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee details updated successfully.")
	}
}

// function to display employee menu after the logging of an employee using their credentials
func EmployeeMenu(employee *employee.EmployeeUser, ems *employee.EmployeeManagementSystem) {
	//running an infinite for loop displaying the functionalities an employee can perform
	for {
		fmt.Println("\n===== Employee Menu =====")
		fmt.Println("1. View My Details")
		fmt.Println("2. Update My Details")
		fmt.Println("3. Logout")

		//reading the user input by declaring a variable called choice
		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		//switch-case to perform the functions from the given list of available functions for employees
		switch choice {
		case 1:
			//displaying the logged in employee's details
			viewMyDetails(employee, ems)
		case 2:
			//updating the logged in employee's details
			updateMyDetails(ems)
		case 3:
			//logging out the logged in employee
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
