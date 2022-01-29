import axios from 'axios';
import { serverRequestBaseUrl, clientRequestBaseUrl, mockServerBaseUrl } from './globals';

export async function getStaticPathForStudents() {
    const response = await axios.get(`${serverRequestBaseUrl}/students`);

    return response.data.map((student) => {
        return {
            params: {
                id: student.student_id
            }
        }
    })
}

export async function getStudents() {
    const response = await axios.get(`${serverRequestBaseUrl}/students/`);

    return response.data;
}

export async function isStudentExist(studentID, sentFromClientBrowser = true) {
    try {
        const response = await axios.get(`${sentFromClientBrowser ? clientRequestBaseUrl : serverRequestBaseUrl}/students/${studentID}`)
        console.log(response.status)
    } catch (err) {
        return false
    }

    return true
}

export async function getStudent(studentID) {
    const response = await axios.get(`${serverRequestBaseUrl}/students/${studentID}`)

    return response.data
}

export async function getModules(studentID) {
    const response = await axios.get(`${mockServerBaseUrl}/students/${studentID}/modules/`);

    return response.data;
}

export async function getResults(studentID) {
    const response = await axios.get(`${mockServerBaseUrl}/students/${studentID}/results/`);

    return response.data;
}

export async function getTimetables(studentID) {
    const response = await axios.get(`${mockServerBaseUrl}/students/${studentID}/timetable/`);

    return response.data;
}