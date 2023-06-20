package functions

import "regexp"

func IsValidEmail(email string) bool {
	// Regular expression pattern for basic email validation
	// This pattern checks for some alpha-numeric till a @ symbol is encountered
	// and again for some alpha-numeric characters till a . is encountered
	// and 2 or more than 2 alphabets after the dot should be present
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
