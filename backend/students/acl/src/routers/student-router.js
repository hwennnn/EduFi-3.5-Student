import { studentEndpointBaseURL } from '../config/baseURL';
import express from 'express';
const url = require('url');

const studentRouter = express.Router();

// Redirect the requests to the student microservice
// HTTP 307 Temporary Redirect is used
// so that the method and the body of the original request are reused to perform the redirected request

studentRouter.get("/", function (req, res) {
    res.redirect(url.format({
        pathname: `${studentEndpointBaseURL}`,
        query: req.query,
    }))
});

studentRouter.get("/:studentID", function (req, res) {
    let studentID = req.params.studentID
    res.redirect(307, `${studentEndpointBaseURL}/${studentID}`)
});

studentRouter.post("/:studentID", function (req, res) {
    let studentID = req.params.studentID
    res.redirect(307, `${studentEndpointBaseURL}/${studentID}`)
});

studentRouter.put("/:studentID", function (req, res) {
    let studentID = req.params.studentID
    res.redirect(307, `${studentEndpointBaseURL}/${studentID}`)
});

export default studentRouter;