
import styles from '@/components/styles/MessageView.module.css'
import './styles/MessageView.module.css'
import React, { useState, useEffect } from 'react'
import Link from 'next/link'
import Avatar from '@mui/material/Avatar';
import Moment from 'react-moment'
import 'moment-timezone'
import { Api } from '@/core/Api'
const api = new Api()

class ContentView {
  constructor(public message: any,
  public setMessage: (message: any) => void,
  public node: React.ReactNode) {
  }
}

export function UserInfo(props: any) {
  const { userid, createdAt } = props
  const [ userName, setUserName ] = useState("NoName")
  const [ avatarInitials, setAvatarInitials ] = useState("AI")
  const [ avatarUrl, setAvatarUrl ] = useState(null)
  const [ avatarBG, setAvatarBG ] = useState("#ffffff")

  useEffect(() => {
    api.getUserInfo(userid).then((userinfo: any) => {
      const user = userinfo.user
      let avatar = (user.avatar || "").trim()
      if (avatar.startsWith("initials://")) {
        const initials = avatar.substring("initials://".length).trim()
        setAvatarInitials(initials)
      } else {
        setAvatarUrl(avatar)
      }
      setUserName(user.name)
    });
  }, [props.userid])
  return <>
    {
    avatarUrl == null ? 
      <Avatar className={styles.header_avatar}
              alt={`${userName} ${userid} - Image`}
      >{avatarInitials}</Avatar>
    : 
      <Avatar className={styles.header_avatar}
              src={avatarUrl}
              alt={`${userName} ${userid} - Image`} />
    }
    <span className={styles.header_username}>{userName}</span>
    <span className={styles.header_createdat}>
      <Moment unix date={createdAt} format="YYYY-MM-D hh:mm A" />
    </span>
  </>
}

export default function MessageView(props: {
  message: any,
}) {
  const { message } = props
  const [hovered, setHovered] = useState(false)
  const toggleHover = () => setHovered(!hovered)

  // Moment.globalTimezone = 'America/Los_Angeles'
  const contentView = createContentView(message)
  return (
    <div
      className={styles.container}
      onMouseEnter={toggleHover}
      onMouseLeave={toggleHover}
    >
      <div className={styles.userinfoarea}>
        <UserInfo userid={message.userId} createdAt={message.createdAt} />
      </div>
      <div className={styles.contentarea}>{contentView}</div>
    </div>
  )
}

export function createContentView(message: any): React.ReactNode {
    if (message.contentType == "chat/text") {
      return <>
        <div style={{wordWrap: "break-word"}}><p>{message.contentText}</p></div>
      </>
    }
    throw new Error("Invalid content type: " + message.contentType)
}
