import { timetableEndpointBaseURL } from '../config/baseURL';
import express from 'express';

const timetableRouter = express.Router();

// Redirect the requests to the student microservice
// HTTP 307 Temporary Redirect is used
// so that the method and the body of the original request are reused to perform the redirected request

timetableRouter.get("/:studentID", function (req, res) {
    let studentID = req.params.studentID
    res.redirect(307, `${timetableEndpointBaseURL}/${studentID}/`)
});

export default timetableRouter;