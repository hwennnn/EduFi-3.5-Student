import { mockServerEndpointBaseURL } from '../config/baseURL';
import express from 'express';
import axios from 'axios';

const tutorRouter = express.Router();

// Redirect the requests to the tutor microservice

/* 
As of 5 February 2022, the developer of Tutor REST API didnt upload any documentation for the usage,
although I find his code in Github.

Hence, this request will be redirected to the mock API.
*/

tutorRouter.get("/", async function (req, res) {
    try {
        const result = await axios.get(`${mockServerEndpointBaseURL}/tutors/`);

        res.status(result.status).json(result.data);
    } catch (exception) {
        res.status(exception.response.status).json(exception.response.statusText);
    }
});

tutorRouter.get("/:tutorID", async function (req, res) {
    try {
        let tutorID = req.params.tutorID;
        console.log(tutorID)
        const result = await axios.get(`${mockServerEndpointBaseURL}/tutors/${tutorID}/`);

        res.status(result.status).json(result.data);
    } catch (exception) {
        res.status(exception.response.status).json(exception.response.statusText);
    }
});

export default tutorRouter;