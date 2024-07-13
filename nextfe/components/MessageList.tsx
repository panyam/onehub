import axios from "axios";
import React, { createRef, useState, useEffect } from "react";
import styles from '@/components/styles/MessageList.module.css'
import MessageView from '@/components/MessageView'

import { Api } from '@/core/Api'
const api = new Api()

class ResultList<T> {
  hasNext = false
  hasPrev = false
  constructor(public items: T[]) {
  }
}

/**
 * Should this just be a dummy table view or also have smarts on how to load things?
 *
 * If it is a dummy table view then the holder of this view needs to pass info and listen
 * to events on when to load things etc (ie gone to top so load more and pass a new list)
 *
 * If not this has to have delegates to load entries etc
 */
export default function Container(props: any) {
  const [messageList, setMessageList] = useState(new ResultList<any>([]));
  const [msglistElem, setMsgListElem] = useState<Element | null>(null);
  const [msgscrollerElem, setMsgScrollerElem] = useState<Element | null>(null);
  const [userMap, setUserMap] = useState<Map<string, any>>(new Map<string, any>());
  useEffect(() => {
    console.log("TopicId: ", props.topicId)
    if (props.topicId == null) return
    api.getMessages(props.topicId).then(resp => {
      console.log("Bingo: ", resp)
      // get all suers in this list of messages
      const userids = new Set<string>()
      for (const msg of resp.messages) {
        userids.add(msg.creatorId)
      }
      const messages = resp.messages
      const newUserMap = new Map(userMap)
      api.getUserInfos(Array.from(userids.values())).then(resp => {
        for (const uid in resp.users) { newUserMap.set(uid, resp.users[uid]) }
        setUserMap(newUserMap)
        setMessageList(new ResultList<any>(messages))
        setTimeout(() => { scrollTo(-1) }, 0)
      });
    });
  }, [props.topicId, msglistElem, msgscrollerElem])

  // we can also have messages added, removed or updated being sent as "properties"

  useEffect(() => {
    console.log("Received Topic Events: ", props.topicEvents)
    const currItems = messageList.items
    if (props.topicEvents && props.topicEvents.length > 0) {
      const newItems = [...currItems]
      for (const tevent of props.topicEvents) {
        if (tevent.type == "new_message") {
          newItems.push(tevent.value)
        }
      }
      setMessageList(new ResultList<any>(newItems))
    }
  }, [props.topicEvents])

  const scrollTo= (y: number) => {
    if (msglistElem && msgscrollerElem) {
      if (y >= 0) {
        msglistElem.scrollTo({top: 0, behavior: "auto"})
      } else {
        const totalHeight = msgscrollerElem.getBoundingClientRect().height
        msglistElem.scrollTo({top: totalHeight - (y + 1), behavior: "auto"})
      }
    }
  }

  /*
  useEffect(() => {
    console.log("Messages Removed: ", props.removedMessages)
  }, [props.removedMessages])

  useEffect(() => {
    console.log("Messages Updated: ", props.updatedMessages)
  }, [props.updatedMessages])
  */
  
  // const delegate = new ContentViewDelegate()

  return  (
    <div ref={setMsgListElem} className={styles.container}>
      <div ref={setMsgScrollerElem} className={styles.msgscroller}>{
              messageList.items.map((message, index) => {
                return <MessageView
                          user = {userMap.get(message.creatorId)}
                          message={message}
                          key={message.id} />
              })
            }
            <div className={styles.bottomanchor}></div>
      </div>
    </div>
    )
}
