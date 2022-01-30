# EduFi API Documentation

|   Microservice   | Port |      Endpoint URL       |
| :--------------: | :--: | :---------------------: |
|     Frontend     | 9210 | http://10.31.11.12:9210 |
| Student REST API | 9211 | http://10.31.11.12:9211 |
|  Mock REST API   | 9211 | http://10.31.11.12:9212 |

## 1. Student Rest API Documentation

**Base URL: http://10.31.11.12:9211**

### Data Structures

|  Field Name   |  Type  |                                         Description                                         |
| :-----------: | :----: | :-----------------------------------------------------------------------------------------: |
|  student_id   | string |                           The unique ID which identifies student                            |
|     name      | string |                               The English name of the student                               |
| date_of_birth | string | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string |                                 The address of the student                                  |
| phone_number  | string |                               The phone number of the student                               |

---

### 1.1 **[GET]** api/v1/students/

It retrieves the students based on the request body parameters (if there is any). It supports filtering the students by putting the filtered condition in the request body parameters.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/students/
```

#### JSON Body Paremeters

|     Name      |  Type  |   Required   | Description                                                                                 |
| :-----------: | :----: | :----------: | ------------------------------------------------------------------------------------------- |
|     name      | string | Not required | The English name of the student                                                             |
| date_of_birth | string | Not required | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string | Not required | The address of the student                                                                  |
| phone_number  | string | Not required | The phone number of the student                                                             |

#### Response

The response will be a status code 200 and an array of student json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "student_id": "1",
    "name": "Wai Hou Man",
    "date_of_birth": "996076800000",
    "address": "BLK678B Jurong West, Singapore",
    "phone_number": "6511111111"
  },
  {
    "student_id": "2",
    "name": "Zachary Hong Rui Quan",
    "date_of_birth": "1007136000000",
    "address": "BLK123F Orchard Rd",
    "phone_number": "6512345678"
  },
  {
    "student_id": "3",
    "name": "Tee Yong Teng",
    "date_of_birth": "912441600000",
    "address": "BLK666A Punggol",
    "phone_number": "6533333333"
  }
]
```

---

### 1.2 **[GET]** api/v1/students/:studentid

It retrieves the student associated with the supplied studentid. A studentID must be supplied in the request query parameters.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/students/1
```

#### Response

The response will be a status code 200 and a student json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful. For example, it will return 404 if no record is found.

```json
{
  "student_id": "1",
  "name": "Wai Hou Man",
  "date_of_birth": "996076800000",
  "address": "BLK678B Jurong West, Singapore",
  "phone_number": "6511111111"
}
```

---

### 1.3 **[POST]** api/v1/students/:studentid

It creates a student in MySQL database by specific studentid. Information such as student_id, name, date_of_birth, address, and phone_number must be supplied in the request body during registration.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/students/4
```

#### JSON Body Paremeters

|     Name      |  Type  | Required | Description                                                                                 |
| :-----------: | :----: | :------: | ------------------------------------------------------------------------------------------- |
|  student_id   | string | Required | The unique ID which identifies student                                                      |
|     name      | string | Required | The English name of the student                                                             |
| date_of_birth | string | Required | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string | Required | The address of the student                                                                  |
| phone_number  | string | Required | The phone number of the student                                                             |

#### Request Body Example

```json
{
  "student_id": "4",
  "name": "hello world",
  "date_of_birth": "1007136000000",
  "address": "#12-123 BLK666 Woodlands, Singapore",
  "phone_number": "6544444444"
}
```

#### Response

Case 1: If the compulsory student information is not provided, it will return message which says the information is not correctly supplied<br>
Case 2: It will fail and return conflict status code if a student with same studentID is already found in the database<br>
Case 3: Otherwise, it will return success message with status created code

```text
201 - Student added: 4
409 - Duplicate student ID
422 - Please supply student information in JSON format
422 - The data in body and parameters do not match
```

---

### 1.4 **[PUT]** api/v1/students/:studentid

It is either used for creating or updating the student depends whether the studentID exists. It allows updating fields for name, date_of_birth, address, and phone_number by putting the fields in the request body.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/students/4
```

#### JSON Body Paremeters (for update)

|     Name      |  Type  |   Required   | Description                                                                                 |
| :-----------: | :----: | :----------: | ------------------------------------------------------------------------------------------- |
|  student_id   | string |   Required   | The unique ID which identifies student                                                      |
|     name      | string | Not required | The English name of the student                                                             |
| date_of_birth | string | Not required | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string | Not required | The address of the student                                                                  |
| phone_number  | string | Not required | The phone number of the student                                                             |

#### Request Body Example

```json
{
  "student_id": "4",
  "phone_number": "6522222222"
}
```

#### Response

Case 1: If studentID exists, update the student using the information retrieved from request body<br>
Case 2: If studentID does not exist, create the student using the information retrieved from request body

```text
201 - Student added: 4
202 - Student updated: 4
422 - Please supply student information in JSON format
422 - The data in body and parameters do not match
```

## 2. Mock REST API Documentation

As mentioned earlier in the main README, as of 29 January 2022, I was having difficulties to consumer other microservices, hence I decided to create a [mock database](../backend/mock/database) and [mock server](../backend/mock/server) to satisfy the necessary requirements and for demo purposes.

**Base URL: http://10.31.11.12:9212**

### 2.1 **[GET]** api/v1/tutors/

It retrieves the tutors from the mock MySQL database.

#### Endpoint URL

```bash
http://10.31.11.12:9212/api/v1/tutors/
```

#### Response

The response will be a status code 200 and an array of tutor json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "tutor_id": "1",
    "first_name": "Wen Qiang",
    "last_name": "Wesley Teo",
    "email": "wesleytwq@gmail.com",
    "descriptions": "This is the description for Tutor #1"
  },
  {
    "tutor_id": "2",
    "first_name": "Kheng Hian",
    "last_name": "Low",
    "email": "lowkh@gmail.com",
    "descriptions": "This is the description for Tutor #2"
  }
]
```

---

### 2.2 **[GET]** api/v1/tutors/:tutorid

It retrieves the tutor associated with the supplied tutorid. A tutorID must be supplied in the request query parameters.

#### Endpoint URL

```bash
http://10.31.11.12:9212/api/v1/tutors/1/
```

#### Response

The response will be a status code 200 and a student json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful. For example, it will return 404 if no record is found.

```json
{
  "tutor_id": "1",
  "first_name": "Wen Qiang",
  "last_name": "Wesley Teo",
  "email": "wesleytwq@gmail.com",
  "descriptions": "This is the description for Tutor #1"
}
```

---

### 2.3 **[GET]** api/v1/modules/:studentid

It retrieves the modules taken by the student with the given studentID from the mock MySQL database.

#### Endpoint URL

```bash
http://10.31.11.12:9212/api/v1/modules/1/
```

#### Response

The response will be a status code 200 and an array of module json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "module_id": "1",
    "module_code": "CM",
    "module_name": "Computing Math",
    "synopsis": "Learn MATH",
    "learning_objectives": "UNION INTERSEC",
    "tutor": {
      "tutor_id": "1",
      "first_name": "Wen Qiang",
      "last_name": "Wesley Teo",
      "email": "wesleytwq@gmail.com",
      "descriptions": "This is the description for Tutor #1"
    }
  },
  {
    "module_id": "2",
    "module_code": "PRG1",
    "module_name": "Programming 1",
    "synopsis": "Learn introductory programming",
    "learning_objectives": "PYTHON PROGRAMMING",
    "tutor": {
      "tutor_id": "2",
      "first_name": "Kheng Hian",
      "last_name": "Low",
      "email": "lowkh@gmail.com",
      "descriptions": "This is the description for Tutor #2"
    }
  }
]
```

---

### 2.4 **[GET]** api/v1/marks/:studentid

It retrieves the results of the student with the given studentID from the mock MySQL database.

#### Endpoint URL

```bash
http://10.31.11.12:9212/api/v1/marks/1/
```

#### Response

The response will be a status code 200 and an array of marks json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "mark_id": "1",
    "module_id": "1",
    "student_id": "1",
    "marks": 81.5,
    "adjusted_marks": 85,
    "module": {
      "module_id": "1",
      "module_code": "CM",
      "module_name": "Computing Math",
      "synopsis": "Learn MATH",
      "learning_objectives": "UNION INTERSEC",
      "tutor": {
        "tutor_id": "1",
        "first_name": "Wen Qiang",
        "last_name": "Wesley Teo",
        "email": "wesleytwq@gmail.com",
        "descriptions": "This is the description for Tutor #1"
      }
    }
  },
  {
    "mark_id": "2",
    "module_id": "2",
    "student_id": "1",
    "marks": 95,
    "adjusted_marks": 0,
    "module": {
      "module_id": "2",
      "module_code": "PRG1",
      "module_name": "Programming 1",
      "synopsis": "Learn introductory programming",
      "learning_objectives": "PYTHON PROGRAMMING",
      "tutor": {
        "tutor_id": "2",
        "first_name": "Kheng Hian",
        "last_name": "Low",
        "email": "lowkh@gmail.com",
        "descriptions": "This is the description for Tutor #2"
      }
    }
  }
]
```

---

### 2.5 **[GET]** api/v1/timetable/:studentid

It retrieves the timetable taken by the student with the given studentID from the mock MySQL database.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/timetable/1/
```

#### Response

The response will be a status code 200 and an array of timetable json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "lesson_id": "1",
    "module_id": "1",
    "lesson_day": "Monday",
    "start_time": "0900",
    "end_time": "1100",
    "module": {
      "module_id": "1",
      "module_code": "CM",
      "module_name": "Computing Math",
      "synopsis": "Learn MATH",
      "learning_objectives": "UNION INTERSEC",
      "tutor": {
        "tutor_id": "1",
        "first_name": "Wen Qiang",
        "last_name": "Wesley Teo",
        "email": "wesleytwq@gmail.com",
        "descriptions": "This is the description for Tutor #1"
      }
    }
  },
  {
    "lesson_id": "2",
    "module_id": "2",
    "lesson_day": "Tuesday",
    "start_time": "1400",
    "end_time": "1600",
    "module": {
      "module_id": "2",
      "module_code": "PRG1",
      "module_name": "Programming 1",
      "synopsis": "Learn introductory programming",
      "learning_objectives": "PYTHON PROGRAMMING",
      "tutor": {
        "tutor_id": "2",
        "first_name": "Kheng Hian",
        "last_name": "Low",
        "email": "lowkh@gmail.com",
        "descriptions": "This is the description for Tutor #2"
      }
    }
  }
]
```

---

### 2.6 **[GET]** api/v1/ratings/:studentid

It retrieves the ratings received by the student with the given studentID from the mock MySQL database.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/ratings/1/
```

#### Response

The response will be a status code 200 and an array of ratings json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "rating_id": "1",
    "creator_id": "2",
    "creator_type": "Student",
    "target_id": "1",
    "target_type": "Student",
    "rating_score": 4,
    "is_anonymous": false,
    "created_time": "1643440926037"
  },
  {
    "rating_id": "4",
    "creator_id": "3",
    "creator_type": "Student",
    "target_id": "1",
    "target_type": "Student",
    "rating_score": 3,
    "is_anonymous": true,
    "created_time": "1643446626037"
  }
]
```

---

### 2.7 **[GET]** api/v1/comments/:studentid

It retrieves the comments received by the student with the given studentID from the mock MySQL database.

#### Endpoint URL

```bash
http://10.31.11.12:9211/api/v1/comments/1/
```

#### Response

The response will be a status code 200 and an array of comments json object if successful, otherwise it would be an error code with a corresponding status message if unsuccessful.

```json
[
  {
    "comment_id": "1",
    "creator_id": "2",
    "creator_type": "Student",
    "target_id": "1",
    "target_type": "Student",
    "comment_data": "Best teammates ever",
    "is_anonymouse": false,
    "created_time": "1643440926037"
  },
  {
    "comment_id": "4",
    "creator_id": "3",
    "creator_type": "Student",
    "target_id": "1",
    "target_type": "Student",
    "comment_data": "Programming god",
    "is_anonymouse": true,
    "created_time": "1643446626037"
  }
]
```

---
