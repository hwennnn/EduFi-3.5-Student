import { studentEndpointBaseURL } from '../config/baseURL';
import express from 'express';
import axios from 'axios';
const url = require('url');

const studentRouter = express.Router();

// Redirect the requests to the driver microservice

studentRouter.get("/", async function (req, res) {
    const result = await axios.get(url.format({
        pathname: `${studentEndpointBaseURL}`,
        query: req.query,
    }));

    res.status(result.status).json(result.data);
});

studentRouter.get("/:studentID", async function (req, res) {
    let studentID = req.params.studentID;
    const result = await axios.get(`${studentEndpointBaseURL}/${studentID}`);

    res.status(result.status).json(result.data);
});

studentRouter.post("/:studentID", async function (req, res) {
    let studentID = req.params.studentID;
    let body = req.body;
    const result = await axios.post(`${studentEndpointBaseURL}/${studentID}`, body);

    res.status(result.status).json(result.data);
});

studentRouter.put("/:studentID", async function (req, res) {
    let studentID = req.params.studentID;
    let body = req.body;
    const result = await axios.put(`${studentEndpointBaseURL}/${studentID}`, body);

    res.status(result.status).json(result.data);
});

export default studentRouter;