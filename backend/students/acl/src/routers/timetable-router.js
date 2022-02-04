import { timetableEndpointBaseURL } from '../config/baseURL';
import express from 'express';

const timetableRouter = express.Router();

// Redirect the requests to the timetable microservice
// HTTP 307 Temporary Redirect is used
// so that the method and the body of the original request are reused to perform the redirected request

timetableRouter.get("/:studentID", function (req, res) {
    let studentID = req.params.studentID

    /* 
    As of 29 January 2022, the developer of Timetable REST API had not updated his endpoint url
    to the exact port used in the server, hence I was not able to consume the microservice.
    Besides, I tried to run his code locally but the API returned html object instead of array of json object.

    Hence, this request will be redirected to the mock API.
    */

    res.redirect(307, `${timetableEndpointBaseURL}/${studentID}/`)
});

export default timetableRouter;