import express from 'express';
import cors from "cors"
import helmet from "helmet"
import studentRouter from './routers/student-router';
import moduleRouter from './routers/module-router';
import markRouter from './routers/mark-router';
import timetableRouter from './routers/timetable-router';
import ratingRouter from './routers/rating-router';
import commentRouter from './routers/comment-router';

const PORT = 4000;
const app = express();

app.use(helmet()); //safety
app.use(cors()); //safety
app.use(express.json()); //receive do respond with request

app.use('/api/v1/students', studentRouter)
app.use('/api/v1/modules', moduleRouter)
app.use('/api/v1/marks', markRouter)
app.use('/api/v1/timetable', timetableRouter)
app.use('/api/v1/ratings', ratingRouter)
app.use('/api/v1/comments', commentRouter)


app.listen(PORT, async () => {
    console.log(`Listening on port ${PORT}`);
});