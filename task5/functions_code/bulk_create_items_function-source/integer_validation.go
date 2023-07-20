package helloworld

import "strconv"

// Parse_int converts a string to an integer and returns 0 if the conversion fails
func Parse_int(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return i
}