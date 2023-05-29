package menu

import "fmt"

// function to display the required main menu after the application is started first time
func DisplayMainMenu() {
	fmt.Println("\n===== Employee Management System =====")
	fmt.Println("Login as an:-")
	fmt.Println("1. Admin")
	fmt.Println("2. Employee")
	fmt.Println("3. Exit")
}
