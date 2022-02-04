import { moduleEndpointBaseURL } from '../config/baseURL';
import express from 'express';

const moduleRouter = express.Router();

// Redirect the requests to the module microservice
// HTTP 307 Temporary Redirect is used
// so that the method and the body of the original request are reused to perform the redirected request

moduleRouter.get("/:studentID", function (req, res) {
    let studentID = req.params.studentID

    /* 
    As of 29 January 2022, no one had implemented an endpoint which retrieves the modules taken by a student.
    Hence, this request will be redirected to the mock API.
    */

    res.redirect(307, `${moduleEndpointBaseURL}/${studentID}/`)
});

export default moduleRouter;