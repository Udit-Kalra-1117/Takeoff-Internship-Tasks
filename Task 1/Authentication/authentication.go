package authentication

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	employee "github.com/uditkalra/ems/Employee"
	functions "github.com/uditkalra/ems/Functions"
	"golang.org/x/crypto/ssh/terminal"
)

// LoginAsAdmin function with passed parameters as EmployeeManagementSystem
func LoginAsAdmin(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Admin Login =====")
	fmt.Print("Username: ")
	//take input of username from user until a new line occurs or enter is pressed on the keyboard
	username, _ := reader.ReadString('\n')
	//remove any newline character, white-spaces and tabs from entered username
	username = strings.TrimSpace(username)

	// fmt.Print("Password: ")
	// //take input of password from user until a new line occurs or enter is pressed on the keyboard
	// password, _ := reader.ReadString('\n')
	// //remove any newline character, white-spaces and tabs from entered password
	// password = strings.TrimSpace(password)

	fmt.Print("Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	//check if username and password are valid
	if username == "admin" && password == "admin123" {
		fmt.Println("Login successful as an Admin!!")
		functions.AdminMenu(ems)
	} else {
		fmt.Println("Invalid credentials. Please try again.")
	}
}

// LoginAsEmployee function with passed parameters as EmployeeManagementSystem
func LoginAsEmployee(ems *employee.EmployeeManagementSystem) {
	//declaring a reusable-reader to read user inputs
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n===== Employee Login =====")
	fmt.Print("Employee Name: ")
	//take input of username from user until a new line occurs or enter is pressed on the keyboard
	name, _ := reader.ReadString('\n')
	//remove any newline character, white-spaces and tabs from entered username
	name = strings.TrimSpace(name)

	// fmt.Print("Password: ")
	// //take input of password from user until a new line occurs or enter is pressed on the keyboard
	// password, _ := reader.ReadString('\n')
	// //remove any newline character, whites-paces and tabs from entered password
	// password = strings.TrimSpace(password)

	fmt.Print("Password: ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)

	//checking if username and password are valid by receiving the results and error parameters from the EmployeeLogin function
	//called from Employee --> employee.go after passing the user entered username and password parameters to the EmployeeLogin function
	res, err := ems.EmployeeLogin(name, password)
	//printing the error parameter if it is not null
	if err != nil {
		fmt.Println(err)
		return
	}
	//if there is no error then the EmployeeMenu is displayed by calling it from Functions --> efunctions.go
	//after logging in the employee with the user entered parameters
	fmt.Println("\nLogin Successful!!")
	fmt.Println("Your Employee ID:", res.ID)
	functions.EmployeeMenu(&res, ems)
}
