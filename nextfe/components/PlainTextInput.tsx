
import axios from "axios";
import React, { useState, useEffect, useCallback, useMemo } from "react";
import isHotkey from 'is-hotkey'
import styles from './styles/PlainTextInput.module.css'

export default function Container(props: any) {
  const textAreaRef = React.createRef<HTMLTextAreaElement>()
  const onKeyUp = async (evt: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (isHotkey("mod+enter", evt) || isHotkey("ctrl+enter", evt) || isHotkey("cmd+enter", evt)) {
      const tarea = evt.target as HTMLTextAreaElement
      const contenttext = tarea.value.trim();
      if (contenttext.length > 0) {
        console.log(evt, evt.target)
        if (await props.onNewMessage({
            "content_text": contenttext,
            "content_type": "text/plain",
        })) {
          if (textAreaRef.current != null) {
            textAreaRef.current.value = ""
          }
        }
      }
    }
  }

  return (
    <div className={styles.container}>
      <div className={styles.header}></div>
      <div className={styles.inputarea}>
        <textarea ref={textAreaRef} className={styles.editable}
                  placeholder="Enter your message and press Ctrl-Enter or Cmd-Enter"
                  onKeyPress = {onKeyUp}>
        </textarea>
      </div>
    </div>
  )
}

