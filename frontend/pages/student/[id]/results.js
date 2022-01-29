import { getStudentWithResults } from '../../../utils/student-utils';
import React from 'react'
import { Button, Table } from 'semantic-ui-react'
import styles from '../../../styles/Home.module.css'
import Head from 'next/head'
import Router from 'next/router';


export async function getServerSideProps({ query }) {
    const student_id = query.id
    const student = await getStudentWithResults(student_id);

    return {
        props: {
            student_id,
            ...student
        }
    }
}

export default function ViewResults({ student_id, results }) {

    function backToStudentHome() {
        Router.push(`/student/${student_id}`)
    }

    const rows = results != null ? results.map(function (result) {
        return (
            <Table.Row key={result.mark_id}>
                <Table.Cell>{result.mark_id}</Table.Cell>
                <Table.Cell>{result.module.module_code}</Table.Cell>
                <Table.Cell>{result.module.module_name}</Table.Cell>
                <Table.Cell>{`${result.marks}%`}</Table.Cell>
            </Table.Row>
        )
    }) : 'There are no results'

    return (

        <div className={styles.container}>
            <Head>
                <title>View Results</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <h1 className={styles.title}>
                View Results
            </h1>

            <br />

            <Table celled>
                <Table.Header>
                    <Table.Row>
                        <Table.HeaderCell>ID</Table.HeaderCell>
                        <Table.HeaderCell>Module Code</Table.HeaderCell>
                        <Table.HeaderCell>Module Name</Table.HeaderCell>
                        <Table.HeaderCell>Marks</Table.HeaderCell>
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