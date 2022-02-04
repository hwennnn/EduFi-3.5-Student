import { markEndpointBaseURL } from '../config/baseURL';
import express from 'express';

const markRouter = express.Router();

// Redirect the requests to the mark microservice
// HTTP 307 Temporary Redirect is used
// so that the method and the body of the original request are reused to perform the redirected request

markRouter.get("/:studentID", function (req, res) {
    let studentID = req.params.studentID

    /* 
    As of 29 January 2022, I was having difficulty to consume http://10.31.11.12:9201/api/v1/marks/1
    to retrieve marks by a specific student. It returns plain text for the result.

    Hence, this request will be redirected to the mock API.
    */

    res.redirect(307, `${markEndpointBaseURL}/${studentID}/`)
});

export default markRouter;