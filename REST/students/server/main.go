package main

import (
	"database/sql"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	models "students/models"
	utils "students/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// global database handler object
var db *sql.DB

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
	var results []models.Student

	params := req.URL.Query()

	// Customise the field query from request query parameters
	formmatedFieldQuery := utils.FormattedStudentQueryField(params)
	query := fmt.Sprintf("SELECT * FROM Students %s", formmatedFieldQuery)

	databaseResults, err := db.Query(query)

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
		results = append(results, student)
	}

	// Returns all the students in JSON
	json.NewEncoder(res).Encode(results)
}

// This method is used to retrieve a student from MySQL by specific studentID,
// and return the result in json otherwise return 404 code
func getStudent(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	studentID := params["studentid"]

	isStudentExist, student := getStudentHelper(studentID)

	if isStudentExist {
		json.NewEncoder(res).Encode(student)
	} else {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("404 - No student found"))
	}
}

// This helper method helps to query the student from the database,
// and return (boolean, student) tuple object
func getStudentHelper(driverID string) (bool, models.Student) {
	query := fmt.Sprintf("SELECT * FROM Students WHERE StudentID='%s'", driverID)
	databaseResults, err := db.Query(query)
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
		isExist = true
	}

	return isExist, student
}

// This method is used to create a student in MySQL by specific studentID,
// Case 1: If the compulsory student information is not provided, it will return message which says the information is not correctly supplied
// Case 2: It will fail and return conflict status code if a student with same studentID is already found in the database
// Case 3: Otherwise, it will return success message with status created code
func postStudent(res http.ResponseWriter, req *http.Request) {
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
			return
		}

		if studentID != newStudent.StudentID {
			res.WriteHeader(http.StatusUnprocessableEntity)
			res.Write([]byte("422 - The data in body and parameters do not match"))
			return
		}

		// check if student exists; add only if student does not exist
		isStudentExist, _ := getStudentHelper(studentID)

		if !isStudentExist {
			query := fmt.Sprintf("INSERT INTO Drivers VALUES ('%s', '%s', '%s', '%s', '%s')", newStudent.StudentID, newStudent.Name, newStudent.DateOfBirth, newStudent.Address, newStudent.PhoneNumber)

			_, err := db.Query(query)

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
}

// This method is used for either creating or updating the student depends whether the studentID exists
// Case 1: If studentID exists, update the student using the information retrieved from request body
// Case 2: If studentID does not exist, create the student using the information retrieved from request body
func putStudent(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	studentID := params["studentid"]

	var newStudent models.Student
	reqBody, err := ioutil.ReadAll(req.Body)

	if err == nil {
		json.Unmarshal(reqBody, &newStudent)

		if studentID != newStudent.StudentID {
			res.WriteHeader(http.StatusUnprocessableEntity)
			res.Write([]byte("422 - The data in body and parameters do not match"))
			return
		}

		// check if student exists; add only if student does not exist, else update
		isStudentExist, _ := getStudentHelper(studentID)

		if !isStudentExist {
			if !utils.IsStudentJsonCompleted(newStudent) {
				res.WriteHeader(http.StatusUnprocessableEntity)
				res.Write([]byte("422 - Please supply student information in JSON format"))
				return
			}

			query := fmt.Sprintf("INSERT INTO Drivers VALUES ('%s', '%s', '%s', '%s', '%s')", newStudent.StudentID, newStudent.Name, newStudent.DateOfBirth, newStudent.Address, newStudent.PhoneNumber)

			_, err := db.Query(query)

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
				return
			}

			query := fmt.Sprintf("UPDATE Drivers SET %s WHERE StudentID='%s'", formattedUpdateFieldQuery, newStudent.StudentID)

			_, err := db.Query(query)

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
}

func main() {

	// Use mysql as driverName and a valid DSN as dataSourceName:
	db, _ = sql.Open("mysql", "user:password@tcp(db:3306)/Edufi_Student")

	fmt.Println("Database opened")

	router := mux.NewRouter()
	router.Use(middleware)
	router.HandleFunc("/api/v1/students", getStudents).Methods("GET")
	router.HandleFunc("/api/v1/students/{studentID}", getStudent).Methods("GET")
	router.HandleFunc("/api/v1/students/{studentID}", postStudent).Methods("POST")
	router.HandleFunc("/api/v1/students/{studentID}", putStudent).Methods("PUT")

	// enable cross-origin resource sharing (cors) for all requests
	handler := cors.AllowAll().Handler(router)

	fmt.Println("Student database server -- Listening at port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
