package helloworld

import "strconv"

// Parse_bool converts a string to a boolean and returns false if the conversion fails
func Parse_bool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}