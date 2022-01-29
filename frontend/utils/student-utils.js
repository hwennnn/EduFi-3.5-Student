import axios from 'axios';
import { clientRequestBaseUrl, serverRequestBaseUrl } from './globals';

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

export async function getStudentWithModules(studentID) {
    const response = await axios.get(`${serverRequestBaseUrl}/students/${studentID}?modules=true`);

    return response.data;
}

export async function getStudentWithResults(studentID) {
    const response = await axios.get(`${serverRequestBaseUrl}/students/${studentID}?marks=true`);

    return response.data;
}

export async function getStudentWithTimetables(studentID) {
    const response = await axios.get(`${serverRequestBaseUrl}/students/${studentID}?timetable=true`);

    return response.data;
}

export async function getStudentsWithRatings() {
    const response = await axios.get(`${serverRequestBaseUrl}/students?ratings=true`);

    return response.data;
}

export async function getStudentsWithAllInformation() {
    const response = await axios.get(`${serverRequestBaseUrl}/students?modules=true&marks=true&timetable=true&ratings=true&comments=true`);

    return response.data;
}