import React, { useState, useEffect, useContext } from "react";

import { createRoot } from "react-dom/client";
import * as FlexLayout from "flexlayout-react";

import Head from 'next/head'
import Image from 'next/image'
import { Inter } from 'next/font/google'
import styles from '@/styles/Home.module.css'
import TopicListPanel from '@/components/TopicListPanel'
import DialogModal from '@/components/DialogModal'
import TopicPanel from '@/components/TopicPanel'
import * as Names from '@/core/names'
import { Api } from '@/core/Api'
const api = new Api()

const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  const signinDialogRef = React.createRef<HTMLDialogElement>()
  const usernameInputRef = React.createRef<HTMLInputElement>()
  const fullnameInputRef = React.createRef<HTMLInputElement>()
  const usernameInputErrorRef = React.createRef<HTMLSpanElement>()
  const fullnameInputErrorRef = React.createRef<HTMLSpanElement>()
  const [fullname, setFullName] = React.useState("")
  const [currTopicId, setCurrTopicId] = React.useState(null)
  const [ isLoggedIn, setIsLoggedIn ] = React.useState(api.auth.isLoggedIn)
  const [isDialogOpened, setIsDialogOpened] = useState(false);
  const [signinButtonLabel, setSigninButtonLabel] = useState("Signin")

  const onTopicSelected = (topic: any) => {
    console.log("Selected: ", topic)
    setCurrTopicId(topic.id)
  }

  const onSigninButtonClicked = async () => {
    if (isLoggedIn) {
      api.auth.logout()
      setIsLoggedIn(false)
      setFullName("")
    } else {
      while (true) {
        let userid = prompt("Enter a unique username for yourself.  This is test only and can be an email, or a phone number or a userid consisting of numbers and letters")
        if (userid == null) {
          setFullName("")
          return
        }

        if (userid.trim().length == 0) continue;

        const randname = Names.RandomName()
        try {
          let userinfo = await api.getUserInfo(userid.trim())
          console.log("UserInfo: ", userinfo)
          if (userinfo.user.name.trim() == "") {
            let fullname = prompt("Enter your full name", randname)
            if (fullname == null || fullname.trim() == "") fullname = randname
            const resp = await api.updateUser(userid, fullname)
            userinfo = resp.user
          }
          api.auth.setLoggedInUser(userid, userinfo.user)
          setFullName(userinfo.user.name)
        } catch (e: any) {
          if (e.response && e.response.status == 404) {
            // create a user
            let fullname = prompt("Enter your full name", randname)
            if (fullname == null || fullname.trim() == "") fullname = randname
            const resp = await api.createUser(userid, fullname)
            const userinfo = resp.user
            api.auth.setLoggedInUser(userid, userinfo.user)
            setFullName(fullname)
          }
          console.log("Error: ",  e)
          return
        }
        return
      }
    }
  }

  const onClose = (label: string): boolean => {
    let result = true
    if (!fullnameInputRef.current) return false
    if (!usernameInputRef.current) return false
    usernameInputErrorRef.current!.innerHTML = ""
    fullnameInputErrorRef.current!.innerHTML = ""
    const userid = usernameInputRef.current.value.trim() || ""
    const fullname = fullnameInputRef.current.value.trim() || ""
    if (userid == "") {
      result = false
      usernameInputErrorRef.current!.innerHTML = "Invalid username"
    }
    if (fullname == "") {
      result = false
      fullnameInputErrorRef.current!.innerHTML = "Invalid fullname"
    }
    if (result) {
      api.auth.setLoggedInUser(userid, {"name": fullname})
      setFullName(fullname)
      setIsLoggedIn(true)
      setIsDialogOpened(false)
    } else {
    }
    return result
  }

  useEffect(() => {
    if (isLoggedIn) {
      setFullName(api.auth.loggedInUser.name)
      setSigninButtonLabel("Signout")
    }
  }, [])

  const flexFactory = (node: any) => {
    var component = node.getComponent();
    if (component === "button") {
      return <button>{node.getName()}</button>;
    }
  }

  const flexModel = {model: FlexLayout.Model.fromJson({
    global: {},
    borders: [],
    layout: {
        type: "row",
        weight: 100,
        children: [
            {
                type: "tabset",
                weight: 50,
                children: [
                    {
                        type: "tab",
                        name: "One",
                        component: "button",
                    }
                ]
            },
            {
                type: "tabset",
                weight: 50,
                children: [
                    {
                        type: "tab",
                        name: "Two",
                        component: "button",
                    }
                ]
            }
        ]
    }
  })};

  return (
  <>
    <Head>
    <title>OneHub Client Demo</title>
    <meta name="description" content="Generated by create next app" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <link rel="icon" href="/favicon.ico" />
    </Head>
    {/*
    <DialogModal
      title="Signin to Onehub"
      isOpened={isDialogOpened}
      buttons={["Signin"]}
      onClose={onClose}
    >
      <hr/>
      <form className={styles.loginForm}>
        <label className={styles.loginFormLabel}>Full name</label>
        <input ref={fullnameInputRef}
               className={styles.loginFormInput}
               placeholder="Enter your display name" />
        <span className={styles.loginFormErrorSpan}
              ref={fullnameInputErrorRef}></span>
        <label className={styles.loginFormLabel}>Username</label>
        <input ref={usernameInputRef}
               className={styles.loginFormInput}
               placeholder="Enter your userid/username" />
        <span className={styles.loginFormErrorSpan}
              ref={usernameInputErrorRef}></span>
      </form>
    </DialogModal>
    */}
    <main className={styles.main}>
      <div className={styles.header}>
        <h2>OneHub Playground</h2>
        <div className={styles.headerRightToolbar}>
            <span className={styles.headerLoggedInUser}>{fullname}</span>
            <button onClick={onSigninButtonClicked}>{signinButtonLabel}</button>
        </div>
      </div>
      {/*
      <div className={styles.mainpanel}>
	      <FlexLayout.Layout model={flexModel} factory={flexFactory}/>
      </div>
      */}
      <div className={styles.left}>
        <TopicListPanel onTopicSelected={onTopicSelected}/>
      </div>
      <div className={styles.centerpanel}>
        <TopicPanel topicId = {currTopicId} />
      </div>
      <div className={styles.footer}></div>
    </main>
  </>
  )
}
