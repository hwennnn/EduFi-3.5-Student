import { getStudentsWithAllInformation } from '../../../utils/student-utils';
import React, { useEffect, useState } from 'react'
import { Button, Table, Input } from 'semantic-ui-react'
import styles from '../../../styles/Home.module.css'
import Head from 'next/head'
import Router from 'next/router';
import { formatDateStringFromMs, formatDateTimeStringFromMs } from '../../../utils/date-utils';

export async function getServerSideProps({ query }) {
    const student_id = query.id
    const students = (await getStudentsWithAllInformation()).filter((student) => student.student_id != student_id);

    return {
        props: {
            student_id,
            students
        }
    }
}

export default function SearchStudents({ student_id, students }) {

    function formatBasicInfo(student) {
        let result = []

        result.push(`Name: ${student.name}`)
        result.push(`Address: ${student.address}`)
        result.push(`Phone Number: ${student.phone_number}`)
        result.push(`Date Of Birth: ${formatDateStringFromMs(parseInt(student.date_of_birth))}`)

        return result.join("\n\n")
    }

    function formatModuleInfo(modules) {
        let result = []
        let index = 1

        for (const module of modules) {
            result.push(`${index}. ${module.module_name} (${module.module_code}) taught by Tutor ${module.tutor.last_name + " " + module.tutor.first_name}`)
            ++index
        }

        return result.join("\n\n")
    }

    function formatTimetableInfo(timetables) {
        let result = []
        let index = 1

        for (const timetable of timetables) {
            result.push(`${index}. ${timetable.lesson_day} ${timetable.start_time}-${timetable.end_time}: ${timetable.module.module_name} (${timetable.module.module_code})`)
            ++index
        }

        return result.join("\n\n")
    }

    function formatRatingsDetails(ratings) {
        // sort the created time in descending order
        ratings.sort((second, first) => parseInt(first.created_time) - parseInt(second.created_time))

        let result = []
        result.push(`Average Ratings: ${formatRatings(ratings)} (${ratings.length})`)
        let index = 1

        for (const rating of ratings) {
            let senderName = rating.is_anonymous ? 'Anonymous User' : `${rating.creator_type} #${rating.creator_id}`

            result.push(`${index}. ${senderName} gave ${rating.rating_score} ratings at ${formatDateTimeStringFromMs(parseInt(rating.created_time))}.`)
            ++index
        }

        return result.join("\n\n")
    }

    function formatRatings(ratings) {
        let ratingCount = 0;
        let totalRatings = 0;

        for (const rating of ratings) {
            ratingCount += 1;
            totalRatings += rating.rating_score;
        }

        return totalRatings / ratingCount;
    }

    function formatCommentsDetails(comments) {
        if (comments.length == 0) return "No comments available."

        // sort the created time in descending order
        comments.sort((second, first) => parseInt(first.created_time) - parseInt(second.created_time))

        let result = []

        let index = 1

        for (const comment of comments) {
            let senderName = comment.is_anonymous ? 'Anonymous User' : `${comment.creator_type} #${comment.creator_id}`

            result.push(`${index}. ${senderName} left a comment "${comment.comment_data}" at ${formatDateTimeStringFromMs(parseInt(comment.created_time))}.`)
            ++index
        }

        return result.join("\n\n")
    }

    const [keyword, setKeyword] = useState("");
    const [rows, setRows] = useState([])

    useEffect(() => {
        let filteredStudents = students;

        if (keyword != "") {
            filteredStudents = students.filter((student) => student.name.toLowerCase().includes(keyword.toLowerCase()));
        }

        if (filteredStudents.length == 0) {
            setRows('There are no students based on the filter')
        } else {
            setRows(filteredStudents.map(function (student) {
                return (
                    <Table.Row key={student.student_id}>
                        <Table.Cell verticalAlign='top'>{student.student_id}</Table.Cell>
                        <Table.Cell verticalAlign='top' style={{ whiteSpace: "pre-line" }}>{formatBasicInfo(student)}</Table.Cell>
                        <Table.Cell verticalAlign='top' style={{ whiteSpace: "pre-line" }}>{formatModuleInfo(student.modules)}</Table.Cell>
                        <Table.Cell verticalAlign='top' style={{ whiteSpace: "pre-line" }}>{formatTimetableInfo(student.timetable)}</Table.Cell>
                        <Table.Cell verticalAlign='top' style={{ whiteSpace: "pre-line" }}>{formatRatingsDetails(student.ratings)}</Table.Cell>
                        <Table.Cell verticalAlign='top' style={{ whiteSpace: "pre-line" }}>{formatCommentsDetails(student.comments)}</Table.Cell>
                    </Table.Row>
                )
            }));
        }

    }, [keyword]);

    function backToStudentHome() {
        Router.push(`/student/${student_id}`)
    }

    return (

        <div className={styles.container}>
            <Head>
                <title>Search Other Students</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <h1 className={styles.title}>
                Search Other Students
            </h1>

            <br />

            <Input value={keyword} onChange={e => setKeyword(e.target.value)} maxLength="30" fluid placeholder='Please enter the name of the student' />

            <br />


            <Table celled>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>StudentID</Table.HeaderCell>
                        <Table.HeaderCell>Basic Info</Table.HeaderCell>
                        <Table.HeaderCell>Modules Taken</Table.HeaderCell>
                        <Table.HeaderCell>Timetable</Table.HeaderCell>
                        <Table.HeaderCell>Ratings</Table.HeaderCell>
                        <Table.HeaderCell>Comments</Table.HeaderCell>
                    </Table.Row>
                </Table.Header>

                <Table.Body>
                    {rows}
                </Table.Body>

            </Table>

            <br />
            <Button primary onClick={backToStudentHome} type='submit'>Back To Home</Button>
        </div>
    )
}