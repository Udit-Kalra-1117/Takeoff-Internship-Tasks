package structure

import "time"

// Employee represents the employee structure
type Employee struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	IsAdmin      bool      `json:"is_admin"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Department   string    `json:"department"`
	Role         string    `json:"role"`
	DateOfBirth  string    `json:"date_of_birth"`
	CreationTime time.Time `json:"creation_time"`
}
