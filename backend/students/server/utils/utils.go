package utils

import (
	"fmt"
	"net/url"
	"strings"

	models "students/models"
)

// This method is to return boolean value whether the given student information is completed
func IsStudentJsonCompleted(student models.Student) bool {
	StudentID := strings.TrimSpace(student.StudentID)
	name := strings.TrimSpace(student.Name)
	dateOfBirth := strings.TrimSpace(student.DateOfBirth)
	address := strings.TrimSpace(student.Address)
	phoneNumber := strings.TrimSpace(student.PhoneNumber)

	return StudentID != "" && name != "" && dateOfBirth != "" && address != "" && phoneNumber != ""
}

// This method is to convert the field query from request query parameters,
// to the sql syntax code
func FormmatedUpdateStudentQueryField(newStudent models.Student) string {
	var fields []string

	if newStudent.Name != "" {
		fields = append(fields, fmt.Sprintf("Name='%s'", newStudent.Name))
	}

	if newStudent.DateOfBirth != "" {
		fields = append(fields, fmt.Sprintf("DateOfBirth='%s'", newStudent.DateOfBirth))
	}

	if newStudent.Address != "" {
		fields = append(fields, fmt.Sprintf("Address='%s'", newStudent.Address))
	}

	if newStudent.PhoneNumber != "" {
		fields = append(fields, fmt.Sprintf("PhoneNumber='%s'", newStudent.PhoneNumber))
	}

	return strings.Join(fields, ", ")
}

// This method is to convert the field query from request query parameters,
// to the sql syntax code
func FormattedStudentQueryField(availableStatus url.Values) string {
	var results string

	// if len(availableStatus) > 0 && availableStatus[0] != "" {
	// 	results += fmt.Sprintf("AvailableStatus = '%s'", availableStatus[0])
	// }

	if results == "" {
		return ""
	}

	return "WHERE " + results
}
