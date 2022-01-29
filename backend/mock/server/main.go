package main

import (
	"database/sql"
	"encoding/json"
	"os"

	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Tutor struct {
	TutorID      string `json:"tutor_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Descriptions string `json:"descriptions"`
}

type Module struct {
	ModuleID           string `json:"module_id"`
	ModuleCode         string `json:"module_code"`
	ModuleName         string `json:"module_name"`
	Synopsis           string `json:"synopsis"`
	LearningObjectives string `json:"learning_objectives"`
	Tutor              Tutor  `json:"tutor"`
}

type ModuleStudent struct {
	ModuleID  string `json:"module_id"`
	StudentID string `json:"student_id"`
}

type ModuleTutor struct {
	ModuleID string `json:"module_id"`
	TutorID  string `json:"tutor_id"`
}

type Mark struct {
	MarkID        string  `json:"mark_id"`
	StudentID     string  `json:"student_id"`
	Module        Module  `json:"module"`
	Marks         float64 `json:"marks"`
	AdjustedMarks float64 `json:"adjusted_marks"`
}

type Lesson struct {
	LessonID  string `json:"lesson_id"`
	ModuleID  string `json:"module_id"`
	LessonDay string `json:"lesson_day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

type Rating struct {
	RatingsID   string `json:"rating_id"`
	CreatorID   string `json:"creator_id"`
	CreatorType string `json:"creator_type"`
	TargetID    string `json:"target_id"`
	TargetType  string `json:"target_type"`
	RatingScore string `json:"rating_score"`
	IsAnonymous bool   `json:"is_anonymous"`
	CreatedTime int64  `json:"created_time"`
}

type Comment struct {
	CommentID   string `json:"comment_id"`
	CreatorID   string `json:"creator_id"`
	CreatorType string `json:"creator_type"`
	TargetID    string `json:"target_id"`
	TargetType  string `json:"target_type"`
	RatingScore string `json:"rating_score"`
	IsAnonymous bool   `json:"is_anonymouse"`
	CreatedTime int64  `json:"created_time"`
}

// this middleware will set the returned content type as application/json
// this helps reduce code redudancy, which originally has to be added in each response writer
func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(res, req)
	})
}

// This method is used to retrieve tutors from MySQL,
// and return the result in array of tutor json object
func getTutors(res http.ResponseWriter, req *http.Request) {
	database := openMockDB()

	var results []Tutor

	query := "SELECT * FROM Tutors"

	databaseResults, err := database.Query(query)

	if err != nil {
		panic(err.Error())
	}

	for databaseResults.Next() {
		// Map the tutor object to the record in the table
		var tutor Tutor
		err = databaseResults.Scan(&tutor.TutorID, &tutor.FirstName, &tutor.LastName, &tutor.Email, &tutor.Descriptions)
		if err != nil {
			panic(err.Error())
		}
		results = append(results, tutor)
	}

	// Returns all the tutors in JSON
	json.NewEncoder(res).Encode(results)

	closeMockDB(database)
}

// This method is used to retrieve tutors from MySQL,
// and return the result in array of tutor json object
func getModulesForStudent(res http.ResponseWriter, req *http.Request) {
	database := openMockDB()

	params := mux.Vars(req)
	studentID := params["studentid"]

	var results []Module

	queryModulesTaken := fmt.Sprintf("SELECT * FROM ModuleStudent WHERE StudentID='%s'", studentID)

	databaseResults, err := database.Query(queryModulesTaken)

	if err != nil {
		panic(err.Error())
	}

	for databaseResults.Next() {
		// Map the tutor object to the record in the table
		var moduleStudent ModuleStudent
		err = databaseResults.Scan(&moduleStudent.ModuleID, &moduleStudent.StudentID)
		if err != nil {
			panic(err.Error())
		}

		isExist, module := getModuleHelper(moduleStudent.ModuleID)

		if isExist {
			results = append(results, module)
		}
	}

	// Returns all the modules in JSON
	json.NewEncoder(res).Encode(results)

	closeMockDB(database)
}

// This helper method helps to query the module from the database,
// and return (boolean, module) tuple object
func getModuleHelper(moduleID string) (bool, Module) {
	database := openMockDB()

	query := fmt.Sprintf("SELECT * FROM Modules WHERE ModuleID='%s'", moduleID)
	databaseResults, err := database.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var isExist bool
	var module Module
	for databaseResults.Next() {
		err = databaseResults.Scan(&module.ModuleID, &module.ModuleCode, &module.ModuleName, &module.Synopsis, &module.LearningObjectives)
		if err != nil {
			panic(err.Error())
		}

		hasTutorForThisModule, moduleTutor := getTutorForSpecificModuleHelper(module.ModuleID)

		if hasTutorForThisModule {
			isTutorExist, tutor := getTutorHelper(moduleTutor.TutorID)

			if isTutorExist {
				module.Tutor = tutor
			}
		}

		isExist = true
	}

	closeMockDB(database)

	return isExist, module
}

// This helper method helps to query the tutor who teaches the specific module from the database,
// and return (boolean, ModuleTutor) tuple object
func getTutorForSpecificModuleHelper(moduleID string) (bool, ModuleTutor) {
	database := openMockDB()

	query := fmt.Sprintf("SELECT * FROM ModuleTutor WHERE ModuleID='%s'", moduleID)
	databaseResults, err := database.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var isExist bool
	var moduleTutor ModuleTutor
	for databaseResults.Next() {
		err = databaseResults.Scan(&moduleTutor.ModuleID, &moduleTutor.TutorID)

		if err != nil {
			panic(err.Error())
		}
		isExist = true
	}

	closeMockDB(database)

	return isExist, moduleTutor
}

// This helper method helps to query the tutor from the database,
// and return (boolean, tutor) tuple object
func getTutorHelper(tutorID string) (bool, Tutor) {
	database := openMockDB()

	query := fmt.Sprintf("SELECT * FROM Tutors WHERE TutorID='%s'", tutorID)
	databaseResults, err := database.Query(query)
	if err != nil {
		panic(err.Error())
	}

	var isExist bool
	var tutor Tutor
	for databaseResults.Next() {
		err = databaseResults.Scan(&tutor.TutorID, &tutor.FirstName, &tutor.LastName, &tutor.Email, &tutor.Descriptions)

		if err != nil {
			panic(err.Error())
		}
		isExist = true
	}

	closeMockDB(database)

	return isExist, tutor
}

func openMockDB() sql.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("APP_DB_USERNAME"), os.Getenv("APP_DB_PASSWORD"), os.Getenv("APP_DB_CONTAINER_NAME"), os.Getenv("APP_DB_PORT"), os.Getenv("APP_DB_NAME"))
	fmt.Println(connectionString)
	// Use mysql as driverName and a valid DSN as dataSourceName:
	database, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("EduFi Mock Database Connection Opened!")
	}

	return *database
}

func closeMockDB(database sql.DB) {
	database.Close()
	fmt.Println("EduFi Mock Database Connection Closed!")
}

func main() {
	router := mux.NewRouter()
	router.Use(middleware)

	router.HandleFunc("/api/v1/tutors/", getTutors).Methods("GET")
	router.HandleFunc("/api/v1/students/{studentid}/modules/", getModulesForStudent).Methods("GET")

	// enable cross-origin resource sharing (cors) for all requests
	handler := cors.AllowAll().Handler(router)

	fmt.Println("Mock database server -- Listening at port 9212")
	log.Fatal(http.ListenAndServe(":9212", handler))
}
