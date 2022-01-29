CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost';

CREATE DATABASE IF NOT EXISTS Edufi_MockDB;

USE Edufi_MockDB;

-- Create Tutors Table
CREATE TABLE IF NOT EXISTS Tutors
(
    TutorID VARCHAR(36) NOT NULL,
    FirstName VARCHAR(30) NOT NULL,
    LastName VARCHAR(30) NOT NULL,
    Email VARCHAR(100) NOT NULL,
    Descriptions VARCHAR(100) NOT NUll,
    CONSTRAINT PK_Tutor PRIMARY KEY (TutorID)
);

-- Insert Tutor Data
INSERT IGNORE INTO Tutors (TutorID, FirstName, LastName, Email, Descriptions) VALUES ('1', 'Wen Qiang', 'Wesley Teo', 'wesleytwq@gmail.com', 'This is the description for Tutor --1');
INSERT IGNORE INTO Tutors (TutorID, FirstName, LastName, Email, Descriptions) VALUES ('2', 'Kheng Hian', 'Low', 'lowkh@gmail.com', 'This is the description for Tutor --2');


-- Create Modules Table
CREATE TABLE IF NOT EXISTS Modules
(
    ModuleID VARCHAR(36) NOT NULL,
    ModuleCode VARCHAR(30) NOT NULL,
    ModuleName VARCHAR(30) NOT NULL,
    Synopsis VARCHAR(100) NOT NULL,
    LearningObjectives VARCHAR(100) NOT NUll,
    CONSTRAINT PK_Module PRIMARY KEY (ModuleID)
);

-- Insert Module Data
INSERT IGNORE INTO Modules (ModuleID, ModuleCode, ModuleName, Synopsis, LearningObjectives) VALUES ('1', 'CM', 'Computing Math', 'Learn MATH', 'UNION INTERSEC');
INSERT IGNORE INTO Modules (ModuleID, ModuleCode, ModuleName, Synopsis, LearningObjectives) VALUES ('2', 'PRG1', 'Programming 1', 'Learn introductory programming', 'PYTHON PROGRAMMING');


-- Create ModuleStudent Table (To indicate the relationship between module and student)
CREATE TABLE IF NOT EXISTS ModuleStudent
(
    ModuleID VARCHAR(36) NOT NULL,
    StudentID VARCHAR(36) NOT NULL,
    CONSTRAINT PK_ModuleStudent PRIMARY KEY (ModuleID, StudentID)
);

-- Insert ModuleStudent Data
INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("1", "1"); -- Student with ID 1 takes Module with ID 1
INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("2", "1"); -- Student with ID 1 takes Module with ID 2

INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("1", "2"); -- Student with ID 2 takes Module with ID 1
INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("2", "3"); -- Student with ID 3 takes Module with ID 2

-- Create ModuleTutor Table (To indicate the relationship between module and tutor)
CREATE TABLE IF NOT EXISTS ModuleTutor
(
    ModuleID VARCHAR(36) NOT NULL,
    TutorID VARCHAR(36) NOT NULL,
    CONSTRAINT PK_ModuleTutor PRIMARY KEY (ModuleID, TutorID)
);

-- Insert ModuleTutor Data
INSERT IGNORE INTO ModuleTutor (ModuleID, TutorID) VALUES ("1", "1"); -- Tutor with ID 1 takes Module with ID 1
INSERT IGNORE INTO ModuleTutor (ModuleID, TutorID) VALUES ("2", "2"); -- Tutor with ID 2 takes Module with ID 2


-- Create Marks Table
CREATE TABLE IF NOT EXISTS Marks
(
    MarkID VARCHAR(36) NOT NULL,
    ModuleID VARCHAR(36) NOT NULL,
    StudentID VARCHAR(36) NOT NULL,
    Marks FLOAT NOT NULL,
    AdjustedMarks FLOAT NULL,
    CONSTRAINT PK_Marks PRIMARY KEY (MarkID)
);

-- Insert Marks Data
INSERT IGNORE INTO Marks (MarkID, ModuleID, StudentID, Marks, AdjustedMarks) VALUES ("1", "1", "1", 81.5, 85.0); -- Student with ID 1 scores 81.5 marks in Module with ID 1 (Adjusted to 85.0 marks later)
INSERT IGNORE INTO Marks (MarkID, ModuleID, StudentID, Marks) VALUES ("2", "2", "1", 95.0); -- Student with ID 1 scores 95.0 marks in Module with ID 2

INSERT IGNORE INTO Marks (MarkID, ModuleID, StudentID, Marks, AdjustedMarks) VALUES ("3", "1", "2", 75.5, 80.0); -- Student with ID 2 scores 75.5 marks in Module with ID 1 (Adjusted to 80.0 marks later)
INSERT IGNORE INTO Marks (MarkID, ModuleID, StudentID, Marks) VALUES ("4", "2", "3", 85.0); -- Student with ID 3 scores 85.0 marks in Module with ID 1

-- Create Lesson Table
CREATE TABLE IF NOT EXISTS Lessons(
    LessonID VARCHAR(36) NOT NULL,
    ModuleID VARCHAR(36) NOT NULL,
    LessonDay VARCHAR(10) NOT NULL,
    StartTime VARCHAR(4) NOT NULL,
    EndTime VARCHAR(4) NOT NULL,
    CONSTRAINT PK_Lesson PRIMARY KEY (LessonID, ModuleID)
);

-- Insert Lesson Data
INSERT IGNORE INTO Lessons (LessonID, ModuleID, LessonDay, StartTime, EndTime) VALUES("1", "1", "Monday", "0900", "1100");
INSERT IGNORE INTO Lessons (LessonID, ModuleID, LessonDay, StartTime, EndTime) VALUES("2", "2", "Tuesday", "1400", "1600");

-- Create LessonStudent Table
CREATE TABLE IF NOT EXISTS LessonStudent(
    LessonID VARCHAR(36) NOT NULL,
    StudentID VARCHAR(36) NOT NULL,

    CONSTRAINT PK_Lesson PRIMARY KEY (LessonID, StudentID)
);

-- Insert LessonStudent Data
INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("1", "1");
INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("1", "2");

INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("2", "1");
INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("2", "3");

-- Create Ratings Table
CREATE TABLE IF NOT EXISTS Ratings(
	RatingID VARCHAR(36) NOT NULL,
	CreatorID VARCHAR(36) NOT NULL,
	CreatorType VARCHAR(255) NOT NULL,
	TargetID VARCHAR(36) NOT NULL,
    TargetType VARCHAR(255) NOT NULL,
	RatingScore FLOAT NOT NULL,
    IsAnonymous tinyint(1) NOT NULL,
    CreatedTime int(64) NOT NULL,

    CONSTRAINT PK_Rating PRIMARY KEY (RatingID)
);

-- Insert Rating Data
INSERT IGNORE INTO Ratings (RatingID, CreatorID, CreatorType, TargetID, TargetType, RatingScore, IsAnonymous, CreatedTime) 
VALUES("1", "2", "Student", "1", "Student", 4.0, 0, 1643440926037);

INSERT IGNORE INTO Ratings (RatingID, CreatorID, CreatorType, TargetID, TargetType, RatingScore, IsAnonymous, CreatedTime) 
VALUES("2", "2", "Student", "3", "Student", 5.0, 0, 1643440925433);

INSERT IGNORE INTO Ratings (RatingID, CreatorID, CreatorType, TargetID, TargetType, RatingScore, IsAnonymous, CreatedTime) 
VALUES("3", "1", "Student", "2", "Student", 4.5, 0, 1643550926037);

INSERT IGNORE INTO Ratings (RatingID, CreatorID, CreatorType, TargetID, TargetType, RatingScore, IsAnonymous, CreatedTime) 
VALUES("4", "3", "Student", "1", "Student", 3.0, 1, 1643446626037);


-- Create Comments Table
CREATE TABLE IF NOT EXISTS Comments(
	CommentID VARCHAR(36) NOT NULL,
	CreatorID VARCHAR(36) NOT NULL,
	CreatorType VARCHAR(255) NOT NULL,
	TargetID VARCHAR(36) NOT NULL,
    TargetType VARCHAR(255) NOT NULL,
	CommentData VARCHAR(500) NOT NULL,
    IsAnonymous tinyint(1) NOT NULL,
    CreatedTime int(64) NOT NULL,

    CONSTRAINT PK_Comment PRIMARY KEY (CommentID)
);

-- Insert Comment Data
INSERT IGNORE INTO Comments (CommentID, CreatorID, CreatorType, TargetID, TargetType, CommentData, IsAnonymous, CreatedTime) 
VALUES("1", "2", "Student", "1", "Student", "Best teammates ever", 0, 1643440926037);

INSERT IGNORE INTO Comments (CommentID, CreatorID, CreatorType, TargetID, TargetType, CommentData, IsAnonymous, CreatedTime) 
VALUES("2", "2", "Student", "3", "Student", "Hello how are you, im under the water woooo", 0, 1643440925433);

INSERT IGNORE INTO Comments (CommentID, CreatorID, CreatorType, TargetID, TargetType, CommentData, IsAnonymous, CreatedTime) 
VALUES("3", "1", "Student", "2", "Student", "Supp bro", 0, 1643550926037);

INSERT IGNORE INTO Comments (CommentID, CreatorID, CreatorType, TargetID, TargetType, CommentData, IsAnonymous, CreatedTime) 
VALUES("4", "3", "Student", "1", "Student", "Programming god", 1, 1643446626037);