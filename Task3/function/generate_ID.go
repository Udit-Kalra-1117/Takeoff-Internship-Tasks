package function

import "time"

// generateID generates a unique ID for the employee
func GenerateID() string {
	// Generate a unique ID using a suitable algorithm
	// This is just a simple example; you may want to use a more robust approach
	return "employee-" + time.Now().Format("20060102150405")
}
