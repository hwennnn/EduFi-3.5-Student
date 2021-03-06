import express from 'express';
import cors from "cors"
import helmet from "helmet"
import studentRouter from './routers/student-router';
import tutorRouter from './routers/tutor-router';

const PORT = 5000;
const app = express();

app.use(helmet()); //safety
app.use(cors()); //safety
app.use(express.json()); //receive do respond with request

app.use('/api/v1/students', studentRouter)
app.use('/api/v1/tutors', tutorRouter)


app.listen(PORT, async () => {
    console.log(`Listening on port ${PORT}`);
});