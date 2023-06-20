package functions

import "regexp"

func IsValidPhoneNumber(phoneNumber string) bool {
	// Regular expression pattern for basic phone number validation
	// This pattern matches a sequence of which must start with + symbol with the country code of 2 digits
	// and this country code must start with digits fom 1 to 9
	// and after that the mobile number must be entered
	// which can range from 9 digit number to 14 digit number as per the country conditions and country code
	// so the total length of phone number is 11 to 14 digits
	phoneNumberRegex := regexp.MustCompile(`^\+[1-9]\d{11,14}$`)
	return phoneNumberRegex.MatchString(phoneNumber)
}
