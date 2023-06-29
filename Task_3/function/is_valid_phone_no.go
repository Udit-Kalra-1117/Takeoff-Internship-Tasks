package function

import "regexp"

// isValidPhoneNumber checks if a phone number starts with "+" and has a length of digits between 7 and 15
func IsValidPhoneNumber(phoneNumber string) bool {
	regex := regexp.MustCompile(`^\+[0-9]{7,15}$`)
	return regex.MatchString(phoneNumber)
}
