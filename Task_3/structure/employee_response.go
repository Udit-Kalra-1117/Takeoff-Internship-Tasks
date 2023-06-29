package structure

import "time"

// Create a custom struct for response serialization (excluding the HashedPassword field)
type EmployeeResponse struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	IsAdmin      bool      `json:"is_admin"`
	Email        string    `json:"email"`
	PhoneNumber  string    `json:"phone_number"`
	Department   string    `json:"department"`
	Role         string    `json:"role"`
	DateOfBirth  string    `json:"date_of_birth"`
	CreationTime time.Time `json:"creation_time"`
}
