import Head from 'next/head'
import Link from 'next/link'
import styles from '../../../styles/Home.module.css'
import { getStaticPathForStudents, getStudent } from '../../../utils/student-utils';

export async function getStaticProps({ params }) {
    const studentID = params.id
    const student = await getStudent(studentID, false);

    return {
        props: {
            ...student
        }
    }
}

export async function getStaticPaths() {
    const paths = await getStaticPathForStudents();

    return {
        paths,
        fallback: false
    }
}

export default function StudentHome({ student_id, name }) {
    let viewParticularsLink = `${student_id}/view`
    let updateParticulasLink = `${student_id}/edit`;

    return (
        <div className={styles.container}>
            <Head>
                <title>Home</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    U are now signed in as <span className={styles.blueColor}> {name}  </span>
                </h1>

                <br />

                <div className={styles.grid}>
                    <Link href={viewParticularsLink}>
                        <a className={styles.card}>
                            <h2>View Particulars &rarr;</h2>
                            <p>View the student's identification number, name, date of birth, address and phone number.</p>
                        </a>
                    </Link>

                    <Link href={updateParticulasLink}>
                        <a className={styles.card}>
                            <h2>Update Particulars &rarr;</h2>
                            <p>Edit the student's particulars including name, date of birth, address and phone number.</p>
                        </a>
                    </Link>

                </div>
            </main>

        </div>
    )
}