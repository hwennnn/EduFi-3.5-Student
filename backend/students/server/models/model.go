package models

type Student struct {
	StudentID   string `json:"student_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
}
