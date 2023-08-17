
import axios from "axios";
import React, { useState, useEffect } from "react";
import styles from '@/components/styles/ChatBox.module.css'
import Auth from '@/core/Auth'
import { Api } from '@/core/Api'
const api = new Api()

export default function Container(props: any) {
  const textAreaRef = React.createRef<HTMLTextAreaElement>()
  const onKeyUp = (evt: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (props.topicId != null && evt.code == "Enter" && evt.ctrlKey) {
      const tarea = evt.target as HTMLTextAreaElement
      const contenttext = tarea.value.trim();
      if (contenttext.length > 0) {
        const user = new Auth().ensureLoggedIn()
        if (user == null) {
          alert("You need to be logged in")
        } else {
          console.log(evt, evt.target)
          api.createMessage(props.topicId, {
            "message": {
              "topic_id": props.topicId,
              "content_text": contenttext,
              "content_type": "chat/text",
              "user_id": user.id,
            },
          }).then(resp => {
            if (props.onNewMessage != null) {
              props.onNewMessage(resp)
            }
            if (textAreaRef.current != null) {
              textAreaRef.current.value = ""
            }
          });
        }
      }
    }
  }

  useEffect(() => {
    if (props.topicId == null) return
  }, [props.topicId])

  return (<div className={styles.container}>
    <div className={styles.header}>header</div>
    <div className={styles.inputarea}>
      <textarea ref={textAreaRef} className={styles.textarea}
                placeholder="Enter your message and press Ctrl-Enter"
                onKeyPress = {onKeyUp}>
      </textarea>
    </div>
  </div>)
}
