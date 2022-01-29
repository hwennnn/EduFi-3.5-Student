package models

type Student struct {
	StudentID   string `json:"student_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`

	Ratings   []Rating  `json:"ratings,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
	Modules   []Module  `json:"modules,omitempty"`
	Results   []Mark    `json:"results,omitempty"`
	Timetable []Lesson  `json:"timetable,omitempty"`
}

// Other data structures used in student fields
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

type Mark struct {
	MarkID        string  `json:"mark_id"`
	ModuleID      string  `json:"module_id"`
	StudentID     string  `json:"student_id"`
	Marks         float64 `json:"marks"`
	AdjustedMarks float64 `json:"adjusted_marks"`
	Module        Module  `json:"module"`
}

type Lesson struct {
	LessonID  string `json:"lesson_id"`
	ModuleID  string `json:"module_id"`
	LessonDay string `json:"lesson_day"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Module    Module `json:"module"`
}

type Rating struct {
	RatingID    string  `json:"rating_id"`
	CreatorID   string  `json:"creator_id"`
	CreatorType string  `json:"creator_type"`
	TargetID    string  `json:"target_id"`
	TargetType  string  `json:"target_type"`
	RatingScore float64 `json:"rating_score"`
	IsAnonymous bool    `json:"is_anonymous"`
	CreatedTime string  `json:"created_time"`
}

type Comment struct {
	CommentID   string `json:"comment_id"`
	CreatorID   string `json:"creator_id"`
	CreatorType string `json:"creator_type"`
	TargetID    string `json:"target_id"`
	TargetType  string `json:"target_type"`
	CommentData string `json:"comment_data"`
	IsAnonymous bool   `json:"is_anonymouse"`
	CreatedTime string `json:"created_time"`
}
