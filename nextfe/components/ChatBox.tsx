
import axios from "axios";
import React, { useState, useEffect, useCallback, useMemo } from "react";
import isHotkey from 'is-hotkey'
import styles from '@/components/styles/ChatBox.module.css'
import Auth from '@/core/Auth'
import PlainTextInput from '@/components/PlainTextInput'

// import { useQuill } from "react-quilljs";

import { Api } from '@/core/Api'
const api = new Api()

export default function Container(props: any) {
  useEffect(() => {
    if (props.topicId == null) return
  }, [props.topicId])

  const onNewMessage = async (msg: any): Promise<boolean> => {
    if (props.topicId == null) return false;
    const user = api.auth.ensureLoggedIn()
    if (user == null) {
      alert("You need to be logged in")
      return false
    }
    msg["topic_id"] = props.topicId 
    msg["user_id"] = user.id
    const resp = await api.createMessage(props.topicId, { "message": msg });
    if (props.onNewMessage != null) {
      props.onNewMessage(resp)
    }
    return true
  }

  return (
    <PlainTextInput onNewMessage={onNewMessage}/>
  )
}

