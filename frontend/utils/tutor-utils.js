import axios from 'axios';
import { serverRequestBaseUrl, clientRequestBaseUrl } from './globals';


export async function isTutorExist(tutorID, sentFromClientBrowser = true) {
    try {
        const response = await axios.get(`${sentFromClientBrowser ? clientRequestBaseUrl : serverRequestBaseUrl}/tutors/${tutorID}`)
        console.log(response.status)
    } catch (err) {
        return false
    }

    return true
}

export async function getTutor(tutorID) {
    const response = await axios.get(`${serverRequestBaseUrl}/tutors/${tutorID}/`)

    return response.data
}