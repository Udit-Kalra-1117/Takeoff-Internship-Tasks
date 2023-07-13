package main

type Employee struct {
	ID          string `json:"id" firestore:"ID"`
	Name        string `json:"name" parsing:"required" firestore:"Name"`
	Password    string `json:"password" parsing:"required" firestore:"Password"`
	IsAdmin     bool   `json:"is_admin" parsing:"required" firestore:"Is Admin"`
	Email       string `json:"email" parsing:"required" firestore:"Email"`
	PhoneNumber string `json:"phone_number" parsing:"required" firestore:"Phone Number"`
	Department  string `json:"department" parsing:"required" firestore:"Department"`
	Role        string `json:"role" parsing:"required" firestore:"Role"`
	DateOfBirth string `json:"date_of_birth" parsing:"required" firestore:"Date of Birth"`
}
//parameters disallowed, error handling