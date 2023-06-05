package main

import (
	"fmt"
	"os"

	authentication "github.com/uditkalra/ems/Authentication"
	employee "github.com/uditkalra/ems/Employee"
	menu "github.com/uditkalra/ems/Menu"
)

// main function declaration
func main() {
	//declaring a new employee management system named ems with the functionalities and properties of
	//EmployeeManagementSystem defined in Employee --> employees.go
	ems := employee.EmployeeManagementSystem{}
	//calling the initialize function on the declared ems variable
	ems.Initialize()

	//running an infinite loop to display the options to login as the desired entity of the application
	for {
		//calling the DisplayManiMenu function to display the main menu from Menu --> menu.go
		menu.DisplayMainMenu()

		//reading the user input by declaring a variable called choice
		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		//switch-case to login as an Admin or an Employee
		switch choice {
		case 1:
			//calling the LoginAsAdmin function to login as an Admin from Authentication --> authentication.go
			authentication.LoginAsAdmin(&ems)
		case 2:
			//calling the LoginAsEmployee function to login as an Employee from Authentication --> authentication.go
			authentication.LoginAsEmployee(&ems)
		case 3:
			//case to stop the application
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			//if the choice is not 1, 2 or 3 then print an error message
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

//Admin
// Username: admin
// Password: admin123
//sensitive tags - json
