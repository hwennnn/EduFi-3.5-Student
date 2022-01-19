import axios from 'axios';
import { serverRequestBaseUrl, requestConfig, clientRequestBaseUrl } from './globals';

export async function getStaticPathForStudents() {
    const response = await axios.get(`${serverRequestBaseUrl}/students/`, requestConfig);

    return response.data.map((student) => {
        return {
            params: {
                id: student.student_id
            }
        }
    })
}

export async function isStudentExist(studentID, sentFromClientBrowser = true) {
    try {
        const response = await axios.get(`${sentFromClientBrowser ? clientRequestBaseUrl : serverRequestBaseUrl}/students/${studentID}`, requestConfig)
        console.log(response.status)
    } catch (err) {
        return false
    }

    return true
}

export async function getStudent(studentID) {
    const response = await axios.get(`${serverRequestBaseUrl}/students/${studentID}`, requestConfig)

    return response.data
}
