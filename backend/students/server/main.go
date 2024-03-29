package main

import (
	"database/sql"
	"encoding/json"
	"os"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	models "students/models"
	utils "students/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// this middleware will set the returned content type as application/json
// this helps reduce code redudancy, which originally has to be added in each response writer
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

// This method is used to retrieve students from MySQL,
// and return the result in array of student json object
func getStudents(res http.ResponseWriter, req *http.Request) {
	studentDB := openStudentsDB()

	var results []models.Student

	params := req.URL.Query()

	query := "SELECT * FROM Students"

	databaseResults, err := studentDB.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for databaseResults.Next() {
		// Map the student object to the record in the table
		var student models.Student
		err = databaseResults.Scan(&student.StudentID, &student.Name, &student.DateOfBirth, &student.Address, &student.PhoneNumber)
		if err != nil {
			panic(err.Error())
		}

		// Check if the params request to retrieve related module information (?modules=true)
		if len(params["modules"]) > 0 && params["modules"][0] == "true" {
			// Fetch module information by sending http request to module microservice
			student.Modules = utils.FetchModules(student.StudentID)
		}

		// Check if the params request to retrieve related marks information (?marks=true)
		if len(params["marks"]) > 0 && params["marks"][0] == "true" {
			// Fetch marks information by sending http request to marks microservice
			student.Results = utils.FetchMarks(student.StudentID)
		}

		// Check if the params request to retrieve related timetable information (?timetable=true)
		if len(params["timetable"]) > 0 && params["timetable"][0] == "true" {
			// Fetch timetable information by sending http request to timetable microservice
			student.Timetable = utils.FetchTimetable(student.StudentID)
		}

		// Check if the params request to retrieve related ratings information (?ratings=true)
		if len(params["ratings"]) > 0 && params["ratings"][0] == "true" {
			// Fetch ratings information by sending http request to ratings microservice
			student.Ratings = utils.FetchRatings(student.StudentID)
		}

		// Check if the params request to retrieve related comments information (?comments=true)
		if len(params["comments"]) > 0 && params["comments"][0] == "true" {
			// Fetch comments information by sending http request to comments microservice
			student.Comments = utils.FetchComments(student.StudentID)
		}

		results = append(results, student)
	}

	// Returns all the students in JSON
	json.NewEncoder(res).Encode(results)

	closeStudentsDB(studentDB)
}

// This method is used to retrieve a student from MySQL by specific studentID,
// and return the result in json otherwise return 404 code
func getStudent(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	studentID := params["studentid"]

	query := req.URL.Query()

	isStudentExist, student := getStudentHelper(studentID, query)

	if isStudentExist {
		json.NewEncoder(res).Encode(student)
	} else {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("404 - No student found"))
	}
}

// This helper method helps to query the student from the database,
// and return (boolean, student) tuple object
func getStudentHelper(studentID string, params url.Values) (bool, models.Student) {
	studentDB := openStudentsDB()

	query := fmt.Sprintf("SELECT * FROM Students WHERE StudentID='%s'", studentID)
	databaseResults, err := studentDB.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var isExist bool
	var student models.Student
	for databaseResults.Next() {
		err = databaseResults.Scan(&student.StudentID, &student.Name, &student.DateOfBirth, &student.Address, &student.PhoneNumber)
		if err != nil {
			panic(err.Error())
		}

		// Check if the params request to retrieve related module information (?modules=true)
		if len(params["modules"]) > 0 && params["modules"][0] == "true" {
			// Fetch module information by sending http request to module microservice
			student.Modules = utils.FetchModules(student.StudentID)
		}

		// Check if the params request to retrieve related marks information (?marks=true)
		if len(params["marks"]) > 0 && params["marks"][0] == "true" {
			// Fetch marks information by sending http request to marks microservice
			student.Results = utils.FetchMarks(student.StudentID)
		}

		// Check if the params request to retrieve related timetable information (?timetable=true)
		if len(params["timetable"]) > 0 && params["timetable"][0] == "true" {
			// Fetch timetable information by sending http request to timetable microservice
			student.Timetable = utils.FetchTimetable(student.StudentID)
		}

		// Check if the params request to retrieve related ratings information (?ratings=true)
		if len(params["ratings"]) > 0 && params["ratings"][0] == "true" {
			// Fetch ratings information by sending http request to ratings microservice
			student.Ratings = utils.FetchRatings(student.StudentID)
		}

		// Check if the params request to retrieve related comments information (?comments=true)
		if len(params["comments"]) > 0 && params["comments"][0] == "true" {
			// Fetch comments information by sending http request to comments microservice
			student.Comments = utils.FetchComments(student.StudentID)
		}

		isExist = true
	}

	closeStudentsDB(studentDB)

	return isExist, student
}

// This method is used to create a student in MySQL by specific studentID,
// Case 1: If the compulsory student information is not provided, it will return message which says the information is not correctly supplied
// Case 2: It will fail and return conflict status code if a student with same studentID is already found in the database
// Case 3: Otherwise, it will return success message with status created code
func postStudent(res http.ResponseWriter, req *http.Request) {
	studentDB := openStudentsDB()

	params := mux.Vars(req)
	studentID := params["studentid"]

	// read the body string sent to the service
	var newStudent models.Student
	reqBody, err := ioutil.ReadAll(req.Body)

	if err == nil {
		// convert JSON to object
		json.Unmarshal(reqBody, &newStudent)

		if !utils.IsStudentJsonCompleted(newStudent) {
			res.WriteHeader(http.StatusUnprocessableEntity)
			res.Write([]byte("422 - Please supply student information in JSON format"))
			closeStudentsDB(studentDB)
			return
		}

		if studentID != newStudent.StudentID {
			res.WriteHeader(http.StatusUnprocessableEntity)
			res.Write([]byte("422 - The data in body and parameters do not match"))
			closeStudentsDB(studentDB)
			return
		}

		// check if student exists; add only if student does not exist
		isStudentExist, _ := getStudentHelper(studentID, map[string][]string{})

		if !isStudentExist {
			query := fmt.Sprintf("INSERT INTO Students VALUES ('%s', '%s', '%s', '%s', '%s')", newStudent.StudentID, newStudent.Name, newStudent.DateOfBirth, newStudent.Address, newStudent.PhoneNumber)

			_, err := studentDB.Query(query)

			if err != nil {
				panic(err.Error())
			}

			res.WriteHeader(http.StatusCreated)
			res.Write([]byte("201 - Student added: " + studentID))
		} else {
			res.WriteHeader(http.StatusConflict)
			res.Write([]byte("409 - Duplicate student ID"))
		}
	} else {
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte("422 - Please supply student information in JSON format"))
	}

	closeStudentsDB(studentDB)
}

// This method is used for either creating or updating the student depends whether the studentID exists
// Case 1: If studentID exists, update the student using the information retrieved from request body
// Case 2: If studentID does not exist, create the student using the information retrieved from request body
func putStudent(res http.ResponseWriter, req *http.Request) {
	studentDB := openStudentsDB()

	params := mux.Vars(req)
	studentID := params["studentid"]

	var newStudent models.Student
	reqBody, err := ioutil.ReadAll(req.Body)

	if err == nil {
		json.Unmarshal(reqBody, &newStudent)

		if studentID != newStudent.StudentID {
			res.WriteHeader(http.StatusUnprocessableEntity)
			res.Write([]byte("422 - The data in body and parameters do not match"))
			closeStudentsDB(studentDB)
			return
		}

		// check if student exists; add only if student does not exist, else update
		isStudentExist, _ := getStudentHelper(studentID, map[string][]string{})

		if !isStudentExist {
			if !utils.IsStudentJsonCompleted(newStudent) {
				res.WriteHeader(http.StatusUnprocessableEntity)
				res.Write([]byte("422 - Please supply student information in JSON format"))
				closeStudentsDB(studentDB)
				return
			}

			query := fmt.Sprintf("INSERT INTO Students VALUES ('%s', '%s', '%s', '%s', '%s')", newStudent.StudentID, newStudent.Name, newStudent.DateOfBirth, newStudent.Address, newStudent.PhoneNumber)

			_, err := studentDB.Query(query)

			if err != nil {
				panic(err.Error())
			}

			res.WriteHeader(http.StatusCreated)
			res.Write([]byte("201 - Student added: " + studentID))
		} else {
			formattedUpdateFieldQuery := utils.FormmatedUpdateStudentQueryField(newStudent)

			// means there is no valid field can be updated
			if formattedUpdateFieldQuery == "" {
				res.WriteHeader(http.StatusUnprocessableEntity)
				res.Write([]byte("422 - Please supply student information in JSON format"))
				closeStudentsDB(studentDB)
				return
			}

			query := fmt.Sprintf("UPDATE Students SET %s WHERE StudentID='%s'", formattedUpdateFieldQuery, newStudent.StudentID)

			_, err := studentDB.Query(query)

			if err != nil {
				panic(err.Error())
			}

			res.WriteHeader(http.StatusAccepted)
			res.Write([]byte("202 - Student updated: " + studentID))
		}

	} else {
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte("422 - Please supply student information in JSON format"))
	}

	closeStudentsDB(studentDB)
}

func openStudentsDB() sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_CONTAINER_NAME"), os.Getenv("APP_DB_PORT"), os.Getenv("APP_DB_NAME"))

	// Use mysql as driverName and a valid DSN as dataSourceName:
	studentDB, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("EduFi Student Database Connection Opened!")

	}

	return *studentDB
}

func closeStudentsDB(studentDB sql.DB) {
	studentDB.Close()
	fmt.Println("EduFi Student Database Connection Closed!")
}

func main() {

	router := mux.NewRouter()
	router.Use(middleware)
	router.HandleFunc("/api/v1/students", getStudents).Methods("GET")
	router.HandleFunc("/api/v1/students/{studentid}", getStudent).Methods("GET")
	router.HandleFunc("/api/v1/students/{studentid}", postStudent).Methods("POST")
	router.HandleFunc("/api/v1/students/{studentid}", putStudent).Methods("PUT")

	// enable cross-origin resource sharing (cors) for all requests
	handler := cors.AllowAll().Handler(router)

	fmt.Println("Student database server -- Listening at port 9211")
	log.Fatal(http.ListenAndServe(":9211", handler))
}
