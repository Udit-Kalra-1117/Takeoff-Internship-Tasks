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

func addEmployee(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Add Employee =====")

	//declaring a variable called emp of type Employee structure to store the changes into it against the required employee
	emp := employee.Employee{}

	//asking the user for entering the name of the new employee record
	fmt.Print("Name: ")
	emp.Name, _ = reader.ReadString('\n')
	emp.Name = strings.TrimSpace(emp.Name)

	//asking the user for entering the id of the new employee record
	fmt.Print("ID: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	//converting the entered id parameter to integer
	id, erro := strconv.Atoi(idStr)
	//if there is error while converting the id parameter to integer then the application is stopped
	//if there is no error while converting the id parameter to integer the entered id parameter is stored against the new employee record
	if erro != nil {
		fmt.Println(erro)
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	emp.ID = id

	//asking the user for entering the password of the new employee record
	fmt.Print("Password: ")
	emp.Password, _ = reader.ReadString('\n')
	emp.Password = strings.TrimSpace(emp.Password)

	//asking the user for entering the email of the new employee record
	fmt.Print("Email: ")
	emp.Email, _ = reader.ReadString('\n')
	emp.Email = strings.TrimSpace(emp.Email)

	//asking the user for entering the phone-no of the new employee record
	fmt.Print("Phone No: ")
	phnoStr, _ := reader.ReadString('\n')
	phnoStr = strings.TrimSpace(phnoStr)
	//converting the entered phone-no parameter to integer
	phno, erro := strconv.Atoi(phnoStr)
	//if there is error while converting the phone-no parameter to integer then the application is stopped
	//if there is no error while converting the phone-no parameter to integer the entered id parameter is stored against the new employee record
	if erro != nil {
		fmt.Println(erro)
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	emp.PhoneNo = phno

	//asking the user for entering the department of the new employee record
	fmt.Print("Department: ")
	emp.Department, _ = reader.ReadString('\n')
	emp.Department = strings.TrimSpace(emp.Department)

	//asking the user for entering the role of the new employee record
	fmt.Print("Role: ")
	emp.Role, _ = reader.ReadString('\n')
	emp.Role = strings.TrimSpace(emp.Role)

	//asking the user for entering the date-of-birth of the new employee record
	fmt.Print("Date of Birth (YYYY-MM-DD: 2000-04-10): ")
	dobStr, _ := reader.ReadString('\n')
	dobStr = strings.TrimSpace(dobStr)
	//checking if the entered value for date of birth is according to the default format by using the time.Parse function
	dob, erro := time.Parse("2006-01-02", dobStr)
	//if there is error in the entered date of birth then the application is stopped
	//if there is no error then the entered date of birth value is stored in the date of birth parameter
	if erro != nil {
		fmt.Println("Invalid Date of Birth. Please try again.")
		fmt.Println("Exiting...")
		os.Exit(0)
	} else {
		emp.DateOfBirth = dob
	}

	//calling the AddEmployee from Employee --> employee.go to add the new employee record to the employee management system
	err := ems.AddEmployee(emp)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully.")
	}
}

// function to view/search the employee details based on id
func viewEmployeeDetails(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== View Employee Details =====")

	//asking the admin to enter the id to view the details of
	fmt.Print("Enter Employee ID: ")
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

	//displaying the employee details of the entered id
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

func updateEmployeeDetails(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Update Employee Details =====")

	//asking the admin to enter the id to update the details of
	fmt.Print("Enter Employee ID: ")
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
	inputPhNo, _ := reader.ReadString('\n')
	inputPhNo = strings.TrimSpace(inputPhNo)
	//converting the entered phone-no parameter to integer to check if entered value is integer
	inputPh, erro := strconv.Atoi(inputPhNo)
	//if there is error while converting the phone-no parameter to integer then the application is stopped
	//if there is no error while converting the phone-no parameter to integer the entered id parameter is stored against the new employee record
	if erro != nil {
		fmt.Println(erro)
		fmt.Println("Exiting...")
		os.Exit(0)
	}
	//checking if the user has entered any value for the specified parameter
	//if new value is entered then assigning the new value to be stored in the records
	//else storing the old unchanged value in the records
	if inputPhNo != "" {
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
	fmt.Print("Date of Birth (YYYY-MM-DD: 2000-04-10) (Leave blank and press Enter to keep existing value: ", emp.DateOfBirth.Format("2006-01-02"), "): ")
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

func deleteEmployee(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Delete Employee =====")

	//asking the admin to enter the id to delete the details of
	fmt.Print("Enter Employee ID: ")
	//take input of id from user until a new line occurs or enter is pressed on the keyboard
	idStr, _ := reader.ReadString('\n')
	//remove any newline character, white-spaces and tabs from entered id
	idStr = strings.TrimSpace(idStr)
	//convert the id to int to pass it to ViewEmployeeDetails function of Employee --> employee.go
	id, _ := strconv.Atoi(idStr)

	//calling the DeleteEmployee function from Employee --> employee.go to delete the employee details of the given id
	err := ems.DeleteEmployee(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee deleted successfully.")
	}
}

func searchEmployeeByName(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Search Employee by Name =====")

	//asking the admin to enter the name to search the employee details
	fmt.Print("Enter Employee Name: ")
	//take input of name from user until a new line occurs or enter is pressed on the keyboard
	name, _ := reader.ReadString('\n')
	//remove any newline character, white-spaces and tabs from entered name
	name = strings.TrimSpace(name)

	//calling the SearchEmployeeByName function from Employee --> employee.go to search the details of the employee as per the entered name
	emp, err := ems.SearchEmployeeByName(name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//displaying the details of the employee searched as per name of the employee
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

// function to display admin menu after the logging of an admin using their credentials
func AdminMenu(ems *employee.EmployeeManagementSystem) {
	// running an infinite for loop displaying the functionalities an admin can perform
	for {
		fmt.Println("\n===== Admin Menu =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View/Search Employee Details by ID")
		fmt.Println("3. Update Employee Details")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. List all Employees")
		fmt.Println("6. List Employees with Upcoming Birthdays")
		fmt.Println("7. Search Employee by Name")
		fmt.Println("8. Logout")

		//reading the user input by declaring a variable called choice
		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		//switch-case to perform the functions from the given list of available functions for admin
		switch choice {
		case 1:
			//functionality to add new employee details to the employee management system
			addEmployee(ems)
		case 2:
			//functionality to view/search employee details by entering the id of the desired employee
			viewEmployeeDetails(ems)
		case 3:
			//functionality to update employee details by entering the id of the desired employee
			updateEmployeeDetails(ems)
		case 4:
			//functionality to delete employee details by entering the id of the desired employee
			deleteEmployee(ems)
		case 5:
			//functionality to display all the employee details stored in the employee management system
			employee.ListAllEmployees(ems)
		case 6:
			//functionality to display all the employee details stored in the employee management system
			//who have upcoming birthday in the current ongoing month
			employee.ListUpcomingBirthdays(ems)
		case 7:
			//functionality to search employee details by entering the name of the desired employee
			searchEmployeeByName(ems)
		case 8:
			//functionality to logout the admin
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
