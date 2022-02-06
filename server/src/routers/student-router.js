import { studentEndpointBaseURL } from '../config/baseURL';
import express from 'express';
import axios from 'axios';
const url = require('url');

const studentRouter = express.Router();

// Redirect the requests to the student microservice

studentRouter.get("/", async function (req, res) {
    try {
        const result = await axios.get(url.format({
            pathname: `${studentEndpointBaseURL}`,
            query: req.query,
        }));

        res.status(result.status).json(result.data);
    } catch (exception) {
        res.status(exception.response.status).json(exception.response.statusText);
    }
});

studentRouter.get("/:studentID", async function (req, res) {
    try {
        let studentID = req.params.studentID;
        const result = await axios.get(url.format({
            pathname: `${studentEndpointBaseURL}/${studentID}`,
            query: req.query,
        }));

        res.status(result.status).json(result.data);
    } catch (exception) {
        res.status(exception.response.status).json(exception.response.statusText);
    }
});

studentRouter.post("/:studentID", async function (req, res) {
    try {
        let studentID = req.params.studentID;
        let body = req.body;
        const result = await axios.post(`${studentEndpointBaseURL}/${studentID}`, body);

        res.status(result.status).json(result.data);
    } catch (exception) {
        res.status(exception.response.status).json(exception.response.statusText);
    }
});

studentRouter.put("/:studentID", async function (req, res) {
    try {
        let studentID = req.params.studentID;
        let body = req.body;
        const result = await axios.put(`${studentEndpointBaseURL}/${studentID}`, body);

        res.status(result.status).json(result.data);
    } catch (exception) {
        res.status(exception.response.status).json(exception.response.statusText);
    }
});

export default studentRouter;