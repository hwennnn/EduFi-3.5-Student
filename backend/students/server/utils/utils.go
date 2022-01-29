package utils

import (
	"encoding/json"
	"strings"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	models "students/models"
)

const serverEndpointBaseURL = "http://acl:4000/api/v1"

var moduleEndpointBaseURL = fmt.Sprintf("%s/module", serverEndpointBaseURL)
var marksEndpointBaseURL = fmt.Sprintf("%s/module", serverEndpointBaseURL)
var timetableEndpointBaseURL = fmt.Sprintf("%s/timetable", serverEndpointBaseURL)
var ratingsEndpointBaseURL = fmt.Sprintf("%s/ratings", serverEndpointBaseURL)
var commentsEndpointBaseURL = fmt.Sprintf("%s/comments", serverEndpointBaseURL)

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

// This method will send a request to module microservice
// in order to retrieve module information for specific studentID
func FetchModules(studentID string) []models.Module {
	var result []models.Module

	url := fmt.Sprintf("%s/%s/", moduleEndpointBaseURL, studentID)

	// Create a new request using http
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result
}

// This method will send a request to marks microservice
// in order to retrieve mark information for specific studentID
func FetchMarks(studentID string) []models.Mark {
	var result []models.Mark

	url := fmt.Sprintf("%s/%s/", marksEndpointBaseURL, studentID)

	// Create a new request using http
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result
}

// This method will send a request to timetable microservice
// in order to retrieve timetable information for specific studentID
func FetchTimetable(studentID string) []models.Lesson {
	var result []models.Lesson

	url := fmt.Sprintf("%s/%s/", timetableEndpointBaseURL, studentID)

	// Create a new request using http
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result
}

// This method will send a request to ratings microservice
// in order to retrieve ratings information for specific studentID
func FetchRatings(studentID string) []models.Rating {
	var result []models.Rating

	url := fmt.Sprintf("%s/%s/", ratingsEndpointBaseURL, studentID)

	// Create a new request using http
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result
}

// This method will send a request to comments microservice
// in order to retrieve comments information for specific studentID
func FetchComments(studentID string) []models.Comment {
	var result []models.Comment

	url := fmt.Sprintf("%s/%s/", commentsEndpointBaseURL, studentID)

	// Create a new request using http
	req, _ := http.NewRequest(http.MethodGet, url, nil)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &result)

	return result
}
