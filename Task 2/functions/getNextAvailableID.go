package functions

import "github.com/uditkalra/swaggerRestApi/variables"

// append to appropriate ID if the excel has entries already or generate the appropriate next employee id
func GetNextAvailableID() int {
	highestID := 0
	for _, employee := range  variables.Slice_of_employees {
		if employee.ID > highestID {
			highestID = employee.ID
		}
	}
	return highestID + 1
}
