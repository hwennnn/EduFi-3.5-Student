CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'localhost';

/* CREATE database Edufi_Student; */

USE Edufi_Student;

CREATE TABLE Students
(
    StudentID VARCHAR(36) NOT NULL,
    Name VARCHAR(30) NOT NULL,
    DateOfBirth VARCHAR(30) NOT NULL,
    Address VARCHAR(100) NOT NULL,
    PhoneNumber VARCHAR(30) NOT NUll,
    CONSTRAINT PK_Student PRIMARY KEY (StudentID)
);

-- Insert Student Data
INSERT INTO Students (StudentID, Name, DateOfBirth, Address, PhoneNumber) VALUES ('1', 'Wai Hou Man', '996076800000', 'BLK678B Jurong West, Singapore', '6511111111');
INSERT INTO Students (StudentID, Name, DateOfBirth, Address, PhoneNumber) VALUES ('2' ,'Zachary Hong Rui Quan', '1007136000000', 'BLK123F Orchard Rd', '6512345678');
INSERT INTO Students (StudentID, Name, DateOfBirth, Address, PhoneNumber) VALUES ('3', 'Tee Yong Teng', '912441600000', 'BLK666A Punggol', '6533333333');
