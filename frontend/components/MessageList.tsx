import axios from "axios";
import { createRef, useState, useEffect } from "react";
import styles from '@/components/styles/MessageList.module.css'
import MessageView from '@/components/MessageView'

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
  useEffect(() => {
    console.log("TopicId: ", props.topicId)
    if (props.topicId == null) return
    document.scrollingElement?.scroll(0, 1)
    axios.get(`/v1/topics/${props.topicId}/messages`).then(resp => {
      console.log("Bingo: ", resp.data)
      setMessageList(new ResultList<any>(resp.data.messages))
    });
  }, [props.topicId])

  // we can also have messages added, removed or updated being sent as "properties"

  useEffect(() => {
    console.log("Received Topic Events: ", props.topicEvents)
    if (props.topicEvents && props.topicEvents.length > 0) {
      const currItems = messageList.items
      const newItems = [...currItems, ...props.addedMessages]
      setMessageList(new ResultList<any>(newItems))
    }
  }, [props.addedMessages])

  useEffect(() => {
    console.log("Messages Removed: ", props.removedMessages)
  }, [props.removedMessages])

  useEffect(() => {
    console.log("Messages Updated: ", props.updatedMessages)
  }, [props.updatedMessages])
  
  // const delegate = new ContentViewDelegate()

  return <div className={styles.container}>{
            messageList.items.map((message, index) => {
              return <MessageView
                          message={message}
                          key={message.id} />
            })
          }
          <div className={styles.bottomanchor}></div>
    </div>
}
