package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/uditkalra/swaggerRestApi/header"
	"github.com/uditkalra/swaggerRestApi/variables"
)

// function to update the csv with the new entries
func SaveToCSV() {
	file, err := os.Create("employees.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	//write the header row
	writer.Write(header.HeaderFields)

	for _, employee := range variables.Slice_of_employees {
		row := []string{
			strconv.Itoa(employee.ID),
			employee.Name,
			employee.Password,
			strconv.FormatBool(employee.IsAdmin),
			employee.Email,
			employee.PhoneNumber,
			employee.Department,
			employee.Role,
			employee.DateOfBirth,
		}
		writer.Write(row)
	}
}
