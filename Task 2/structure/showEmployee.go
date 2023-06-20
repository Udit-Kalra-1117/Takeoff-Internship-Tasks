package structure

// structure to show employee details without showing the password
type ShowEmployee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	IsAdmin     bool   `json:"is_admin"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Department  string `json:"department"`
	Role        string `json:"role"`
	DateOfBirth string `json:"date_of_birth"`
}
