package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/uditkalra/swaggerRestApi/header"
	"github.com/uditkalra/swaggerRestApi/structure"
	"github.com/uditkalra/swaggerRestApi/variables"
)

// function to load previous records from the existing csv if any
func LoadFromCSV() {
	file, err := os.Open("employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// initializing a reader and making it reading all the fields as per the length of the headerFields struct defined above
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = len(header.HeaderFields)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i, row := range rows {
		if i == 0 {
			// Skipping header row
			continue
		}
		id, _ := strconv.Atoi(row[0])
		isAdmin, _ := strconv.ParseBool(row[3])

		// mapping the appropriate elements to the appropriate columns in the excel file
		employee := structure.Employee{
			ID:          id,
			Name:        row[1],
			Password:    row[2],
			IsAdmin:     isAdmin,
			Email:       row[4],
			PhoneNumber: row[5],
			Department:  row[6],
			Role:        row[7],
			DateOfBirth: row[8],
		}
		variables.Slice_of_employees = append(variables.Slice_of_employees, employee)
	}
}
