package functions

import (
	"time"

	"github.com/uditkalra/swaggerRestApi/constants"
)

func IsValidDateOfBirth(dateOfBirth string) bool {
	// Checking if the dateOfBirth is in the correct format
	_, err := time.Parse(constants.DateFormat, dateOfBirth)
	return err == nil
}
