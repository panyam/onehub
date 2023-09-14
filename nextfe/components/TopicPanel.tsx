
import axios from "axios";
import React, { useState, useEffect } from "react";
import styles from '@/components/styles/TopicPanel.module.css'
import Auth from '@/core/Auth'
import ChatBox from '@/components/ChatBox'
import MessageList from '@/components/MessageList'
import { Api } from '@/core/Api'
const api = new Api()

class ResultList<T> {
  hasNext = false
  hasPrev = false
  constructor(public items: T[]) {
  }
}

export default function Container(props: any) {
  const [ topicName, setTopicName ] = useState("")
  const [ topicEvents, setTopicEvents ] = useState([] as any[])
  useEffect(() => {
    if (props.topicId) {
      api.getTopicInfo(props.topicId).then((topicinfo: any) => {
        console.log("Loaded Topic Info: ", topicinfo)
        setTopicName(topicinfo.topic.name)
      });
    }
  }, [props.topicId])

  const onNewMessage = (msg: any) => {
    console.log("NewMsg: ", msg)
    setTopicEvents([{
      "type": "new_message",
      "value": msg,
    }])
  }

  return (<>
  <div className={styles.header}><h2 className={styles.title_header}>{topicName}</h2></div>
  <div className={styles.footer}></div>
  <div className={styles.integrations}></div>
  <div className={styles.messagelist}>
    <MessageList topicId = {props.topicId} topicEvents={topicEvents} />
  </div>
  <div className={styles.chatbox}>
    <ChatBox topicId = {props.topicId} onNewMessage = {onNewMessage}/>
  </div>
</>)
}
