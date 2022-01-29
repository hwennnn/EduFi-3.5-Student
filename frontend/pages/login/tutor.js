import React, { useState } from 'react';
import { Input, Button } from 'semantic-ui-react'
import Head from 'next/head'
import Router from 'next/router';
import styles from '../../styles/Home.module.css'
import { isTutorExist } from '../../utils/tutor-utils';


export default function LoginTutor() {
    const [tutorID, setTutorID] = useState(''); // '' is the initial state value

    async function loginAsTutor() {
        if (tutorID != '') {
            let exist = await isTutorExist(tutorID);
            if (exist) {
                Router.push(`/tutor/${tutorID}`)
            } else {
                alert("Incorrect tutor credential information to login!")
            }
        }
    }

    return (
        <div className={styles.container}>
            <Head>
                <title>Login</title>
                <meta name="description" content="Generated by create next app" />
                <link rel="icon" href="/favicon.ico" />
            </Head>

            <main className={styles.main}>
                <h1 className={styles.title}>
                    Login as a tutor
                </h1>

                <p className={styles.description}>
                    You will be prompted to enter your <span className={styles.blueColor}>tutor id</span> in order to login.
                </p>

                <Input value={tutorID} onChange={e => setTutorID(e.target.value)} focus placeholder='Enter your tutor id' />
                <br />
                <Button onClick={loginAsTutor} primary>Login</Button>
            </main>

        </div>
    )
}
