# EduFi Student API Documentation

**Base URL: http://localhost:80/server/**

# Data Structures

|  Field Name   |  Type  |                                         Description                                         |
| :-----------: | :----: | :-----------------------------------------------------------------------------------------: |
|  student_id   | string |                           The unique ID which identifies student                            |
|     name      | string |                               The English name of the student                               |
| date_of_birth | string | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string |                                 The address of the student                                  |
| phone_number  | string |                               The phone number of the student                               |

---

## 1.1 **[GET]** api/v1/students/

It retrieves the students based on the request body parameters (if there is any). It supports filtering the students by putting the filtered condition in the request body parameters.

### Endpoint URL

```bash
http://localhost:80/server/api/v1/students/
```

### JSON Body Paremeters

|     Name      |  Type  |   Required   | Description                                                                                 |
| :-----------: | :----: | :----------: | ------------------------------------------------------------------------------------------- |
|     name      | string | Not required | The English name of the student                                                             |
| date_of_birth | string | Not required | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string | Not required | The address of the student                                                                  |
| phone_number  | string | Not required | The phone number of the student                                                             |

### Response

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

## 1.2 **[GET]** api/v1/students/:studentid

It retrieves the student associated with the supplied studentid. A studentID must be supplied in the request query parameters.

### Endpoint URL

```bash
http://localhost:80/server/api/v1/students/1
```

### Response

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

## 1.3 **[POST]** api/v1/students/:studentid

It creates a student in MySQL database by specific studentid. Information such as student_id, name, date_of_birth, address, and phone_number must be supplied in the request body during registration.

### Endpoint URL

```bash
http://localhost:80/server/api/v1/students/4
```

### JSON Body Paremeters

|     Name      |  Type  | Required | Description                                                                                 |
| :-----------: | :----: | :------: | ------------------------------------------------------------------------------------------- |
|  student_id   | string | Required | The unique ID which identifies student                                                      |
|     name      | string | Required | The English name of the student                                                             |
| date_of_birth | string | Required | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string | Required | The address of the student                                                                  |
| phone_number  | string | Required | The phone number of the student                                                             |

### Request Body Example

```json
{
  "student_id": "4",
  "name": "hello world",
  "date_of_birth": "1007136000000",
  "address": "#12-123 BLK666 Woodlands, Singapore",
  "phone_number": "6544444444"
}
```

### Response

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

## 1.4 **[PUT]** api/v1/students/:studentid

It is either used for creating or updating the student depends whether the studentID exists. It allows updating fields for name, date_of_birth, address, and phone_number by putting the fields in the request body.

### Endpoint URL

```bash
http://localhost:80/server/api/v1/students/4
```

### JSON Body Paremeters (for update)

|     Name      |  Type  |   Required   | Description                                                                                 |
| :-----------: | :----: | :----------: | ------------------------------------------------------------------------------------------- |
|  student_id   | string |   Required   | The unique ID which identifies student                                                      |
|     name      | string | Not required | The English name of the student                                                             |
| date_of_birth | string | Not required | The date of birth of the student (The datetime is stored in milliseconds since epoch time.) |
|    address    | string | Not required | The address of the student                                                                  |
| phone_number  | string | Not required | The phone number of the student                                                             |

### Request Body Example

```json
{
  "student_id": "4",
  "phone_number": "6522222222"
}
```

### Response

Case 1: If studentID exists, update the student using the information retrieved from request body<br>
Case 2: If studentID does not exist, create the student using the information retrieved from request body

```text
201 - Student added: 4
202 - Student updated: 4
422 - Please supply student information in JSON format
422 - The data in body and parameters do not match
```
