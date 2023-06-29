package function

import (
	"time"

	"github.com/uditkalra/emsGcpApi/constants"
)

// isValidDateOfBirth checks if a date of birth is in the format "YYYY-MM-DD"
func IsValidDateOfBirth(dateOfBirth string) bool {
	_, err := time.Parse(constants.DateFormat, dateOfBirth)
	return err == nil
}
