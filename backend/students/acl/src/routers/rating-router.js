import { ratingEndpointBaseURL } from '../config/baseURL';
import express from 'express';

const ratingRouter = express.Router();

// Redirect the requests to the rating microservice
// HTTP 307 Temporary Redirect is used
// so that the method and the body of the original request are reused to perform the redirected request

ratingRouter.get("/:studentID", function (req, res) {
    let studentID = req.params.studentID

    /* 
    As of 5 February 2022, the developer of Rating REST API didnt upload any documentation for the usage,
    although I find his code in Github (not completely done yet).

    Hence, this request will be redirected to the mock API.
    */

    res.redirect(307, `${ratingEndpointBaseURL}/${studentID}/`)
});

export default ratingRouter;