import Head from 'next/head'
import Link from 'next/link'
import styles from '../../../styles/Home.module.css'
import { getTutor } from '../../../utils/tutor-utils';

export async function getServerSideProps({ query }) {
    const tutorID = query.id
    const tutor = await getTutor(tutorID);

    return {
        props: {
            ...tutor
        }
    }
}

export default function StudentHome({ tutor_id, last_name, first_name }) {
    let viewAllStudentsLink = `${tutor_id}/view`

    return (
        <div className={styles.container}>
            <Head>
                <title>Home</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    U are now signed in as <span className={styles.blueColor}> {last_name + " " + first_name}  </span>
                </h1>

                <br />

                <div className={styles.grid}>

                    <Link href={viewAllStudentsLink}>
                        <a className={styles.card}>
                            <h2>Search students &rarr;</h2>
                            <p>Search for all students and view their profile, modules, timetable, ratings, and comments.</p>
                        </a>
                    </Link>

                </div>
            </main>

        </div>
    )
}