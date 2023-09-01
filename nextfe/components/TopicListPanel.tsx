
import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';
import ListItemText from '@mui/material/ListItemText';
import ListItemAvatar from '@mui/material/ListItemAvatar';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemIcon from '@mui/material/ListItemIcon';
import DeleteIcon from '@mui/icons-material/Delete';
import IconButton from '@mui/material/IconButton';
import Avatar from '@mui/material/Avatar';



import axios from "axios";
import React, { useState, useEffect } from "react";
import styles from '@/components/styles/TopicListPanel.module.css'
import TopicDetail from "./TopicDetail"
import { Api } from '@/core/Api'
const api = new Api()

class ResultList<T> {
  hasNext = false
  hasPrev = false
  constructor(public items: T[]) {
  }
}

export default function Container(props: any) {
  const [topicList, setTopicList] = useState(new ResultList<any>([]));
  const [selectedIndex, setSelectedIndex] = React.useState(1);
  const handleListItemClick = (
    event: React.MouseEvent<HTMLDivElement, MouseEvent>,
    index: number,
  ) => {
    if ((props.onTopicSelected || null) != null) {
      props.onTopicSelected(topicList.items[index])
    }
    setSelectedIndex(index);
  };



  useEffect(() => {
    api.getTopics().then(response => {
      const out = new ResultList<any>(response.topics)
      out.hasNext = response.nextPageKey.trim() != ""
      setTopicList(out)
    }).catch(error => {
      const out = new ResultList<any>([])
      out.hasNext = false
      setTopicList(out)
      console.log("Error: ", error)
    });
  }, [])

  const newButtonRef = React.createRef<HTMLButtonElement>()

  const onTopicDeleted = (topic: any) => {
    api.deleteTopic(topic.id).then(resp => {
      const newTopics = topicList.items.filter(t => t.id != topic.id)
      setTopicList(new ResultList(newTopics))
    })
  }

  const onSearchTopics = () => {}

  const onNewTopic = () => {
    const user = api.auth.ensureLoggedIn()
    if (user == null) {
      alert("You need to be logged in")
    }
    const result = prompt("Enter name of new topic")
    if (result != null && result.trim() != "") {
      if (newButtonRef.current != null) {
        newButtonRef.current.disabled = false
      }

      api.createTopic({
        "topic": {
        "name": result,
        "creator_id": user.id,
        },
      }).then(response => {
        const newTopics = [response.topic, ...topicList.items]
        const newResults = new ResultList<any>(newTopics )
        setTopicList(newResults)
      });
      if (newButtonRef.current != null) {
        newButtonRef.current.disabled = false
      }
    }
  }

  return (<>
  <div className={styles.header}>
    <h3>Topics</h3>
    <div className={styles.TopicsHeaderPanel}>
      <button onClick={onSearchTopics} className={styles.headerButton}>Search</button>
      <button onClick={onNewTopic} ref={newButtonRef} className={styles.headerButton}>New</button>
    </div>
  </div>
  <div className={styles.topiclist}>
    <List sx={{ width: '100%', maxWidth: 360}}>{
        topicList.items.map((topic, index) => {
          return <ListItem className = {styles.topicListItem} key={topic.id}>
            <ListItemButton
                selected={selectedIndex === 0}
                onClick={(event) => handleListItemClick(event, index)}>
                <ListItemText primary={topic.name} />
            </ListItemButton>
          </ListItem>
        })
    }</List>
  </div>
  <div className={styles.footer}>
  </div>
</>)
}
