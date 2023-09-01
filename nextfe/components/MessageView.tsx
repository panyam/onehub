
import styles from '@/components/styles/MessageView.module.css'
import './styles/MessageView.module.css'
import React, { useState, useEffect } from 'react'
import Link from 'next/link'
import Avatar from '@mui/material/Avatar';
import Moment from 'react-moment'
import moment from 'moment'
import 'moment-timezone'
import { Api } from '@/core/Api'
import { hashCode, pickRandColor } from '@/core/utils'
const api = new Api()

class ContentView {
  constructor(public message: any,
  public setMessage: (message: any) => void,
  public node: React.ReactNode) {
  }
}

export function UserInfo(props: any) {
  const { userid, createdAt } = props
  const [ format, setFormat ] = useState("LLLL")
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
    const currmom = moment(createdAt)
    if (currmom >= moment().startOf('day')) {
      // only show time and am/pm
      setFormat("h:mm:ss a")
    } else if (currmom >= moment().startOf('week')) {
      // show Sunday 28th time am/pm
      setFormat("ddd h:mm a")
    } else if (currmom >= moment().startOf('month')) {
      // Wed 12th HH:MM:SS AM/PM
      setFormat("ddd Do h:mm a")
    } else {
      setFormat("MMM Do YYYY, h:mm:ss a")
    }
  }, [props.userid, props.createdAt])

  const colorForName = (name: string) => {
    const hash = hashCode(name)
    return pickRandColor(100, hash%10)
/*
    const red = hash & 0xff
    const green = (hash >> 8) & 0xff
    const blue = (hash >> 16) & 0xff
    const color = `#${red.toString(16) }${green.toString(16) }${blue.toString(16)}`
    return color
*/
  }

  // All these times should be based on user local time
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
    <span className={styles.header_username} style={{color: colorForName(userName)}}>{userName}</span>
    <span className={styles.header_createdat}>
      <Moment date={createdAt} format={format} />
    </span>
  </>
}

export default function MessageView(props: {
  message: any,
}) {
  const { message } = props
  const [hovered, setHovered] = useState(false)
  const toggleHover = () => setHovered(!hovered)

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
    if (message.contentType == "text/plain" ||
        message.contentType == "chat/text") {
      return <>
        <div style={{wordWrap: "break-word"}}><p>{message.contentText}</p></div>
      </>
    }
    throw new Error("Invalid content type: " + message.contentType)
}
