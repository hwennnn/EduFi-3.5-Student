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
INSERT IGNORE INTO Tutors (TutorID, FirstName, LastName, Email, Descriptions) VALUES ('1', 'Wen Qiang', 'Wesley Teo', 'wesleytwq@gmail.com', 'This is the description for Tutor #1');
INSERT IGNORE INTO Tutors (TutorID, FirstName, LastName, Email, Descriptions) VALUES ('2', 'Kheng Hian', 'Low', 'lowkh@gmail.com', 'This is the description for Tutor #2');


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
    StudentID VARCHAR(36) NOT NULL
    CONSTRAINT PK_ModuleStudent PRIMARY KEY (ModuleID, StudentID)
);

-- Insert Module Data
INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("1", "1") # Student with ID 1 takes Module with ID 1
INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("2", "1") # Student with ID 1 takes Module with ID 2

INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("1", "2") # Student with ID 2 takes Module with ID 1

INSERT IGNORE INTO ModuleStudent (ModuleID, StudentID) VALUES ("2", "3") # Student with ID 3 takes Module with ID 2

-- Create Marks Table
CREATE TABLE IF NOT EXISTS Marks
(
    ModuleID VARCHAR(36) NOT NULL,
    StudentID VARCHAR(36) NOT NULL,
    Marks FLOAT NOT NULL,
    AdjustedMarks FLOAT NULL,
    CONSTRAINT PK_Marks PRIMARY KEY (ModuleID, StudentID)
);

-- Insert Marks Data
INSERT IGNORE INTO Marks (ModuleID, StudentID, Marks, AdjustedMarks) VALUES ("1", "1", 81.5, 85.0) # Student with ID 1 scores 81.5 marks in Module with ID 1 (Adjusted to 85.0 marks later)
INSERT IGNORE INTO Marks (ModuleID, StudentID, Marks) VALUES ("2", "1", 95.0) # Student with ID 1 scores 95.0 marks in Module with ID 2

INSERT IGNORE INTO Marks (ModuleID, StudentID, Marks, AdjustedMarks) VALUES ("1", "2", 75.5, 80.0) # Student with ID 2 scores 75.5 marks in Module with ID 1 (Adjusted to 80.0 marks later)
INSERT IGNORE INTO Marks (ModuleID, StudentID, Marks) VALUES ("2", "3", 85.0) # Student with ID 3 scores 85.0 marks in Module with ID 1

-- Create Class Table
CREATE TABLE IF NOT EXISTS Class(
    ClassID VARCHAR(36) NOT NULL,
    LessonID VARCHAR(36) NOT NULL,

    CONSTRAINT PK_Class PRIMARY KEY (ClassID, LessonID)
)

-- Insert Class Data
INSERT IGNORE INTO Class (ClassID, ModuleID) VALUES("1", "1")
INSERT IGNORE INTO Class (ClassID, ModuleID) VALUES("1", "2")


-- Create Lesson Table
CREATE TABLE IF NOT EXISTS Lesson(
    LessonID VARCHAR(36) NOT NULL,
    ClassID VARCHAR(36) NOT NULL,
    LessonDay VARCHAR(10) NOT NULL,
    StartTime VARCHAR(4) NOT NULL,
    EndTime VARCHAR(4) NOT NULL,
    CONSTRAINT PK_Lesson PRIMARY KEY (LessonID, ClassID)
);

-- Insert Lesson Data
INSERT IGNORE INTO Lesson (LessonID, ClassID, LessonDay, StartTime, EndTime) VALUES("1", "1", "Monday", "0900", "1100")
INSERT IGNORE INTO Lesson (LessonID, ClassID, LessonDay, StartTime, EndTime) VALUES("2", "2", "Tuesday", "1400", "1600")

-- Create LessonStudent Table
CREATE TABLE IF NOT EXISTS LessonStudent(
    LessonID VARCHAR(36) NOT NULL,
    StudentID VARCHAR(36) NOT NULL,

    CONSTRAINT PK_Lesson PRIMARY KEY (LessonID, StudentID)
);

-- Insert LessonStudent Data
INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("1", "1")
INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("1", "2")

INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("2", "1")
INSERT IGNORE INTO LessonStudent (LessonID, StudentID) VALUES("2", "3")



