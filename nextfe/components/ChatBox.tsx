
import axios from "axios";
import React, { useState, useEffect, useCallback, useMemo } from "react";
import isHotkey from 'is-hotkey'
import styles from '@/components/styles/ChatBox.module.css'
import Auth from '@/core/Auth'

// Slate editor
import {
  Editor,
  BaseEditor,
  Transforms,
  createEditor,
  Descendant,
  Element as SlateElement,
} from 'slate'
import { ReactEditor, Editable, withReact, useSlate, Slate } from 'slate-react'
import { withHistory } from 'slate-history'
import { Button, Icon, Toolbar } from '@/components/slate/components'

import { Api } from '@/core/Api'
const api = new Api()

export default function Container(props: any) {
  /*
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
  }*/

  useEffect(() => {
    if (props.topicId == null) return
  }, [props.topicId])

  const initialValue = [
    {
      type: 'paragraph',
      children: [{ text: '' }],
    },
  ] as any[];
  const renderElement = useCallback((props: any) => <Element {...props} />, [])
  const renderLeaf = useCallback((props: any) => <Leaf {...props} />, [])
  const editor = useMemo(() => withHistory(withReact(createEditor())), [])

  return (
      <Slate editor={editor} initialValue={initialValue} >
        <div className={styles.header}>
          <Toolbar className={styles.header_toolbar}>
            <MarkButton format="bold" icon="format_bold" />
            <MarkButton format="italic" icon="format_italic" />
            <MarkButton format="underline" icon="format_underlined" />
            <MarkButton format="code" icon="code" />
            <BlockButton format="heading-one" icon="looks_one" />
            <BlockButton format="heading-two" icon="looks_two" />
            <BlockButton format="block-quote" icon="format_quote" />
            <BlockButton format="numbered-list" icon="format_list_numbered" />
            <BlockButton format="bulleted-list" icon="format_list_bulleted" />
            <BlockButton format="left" icon="format_align_left" />
            <BlockButton format="center" icon="format_align_center" />
            <BlockButton format="right" icon="format_align_right" />
            <BlockButton format="justify" icon="format_align_justify" />
          </Toolbar>
        </div>
        <div className={styles.inputarea}>
          <Editable         // Define a new handler which prints the key that was pressed.
              className = {styles.editable}
              renderElement={renderElement}
              renderLeaf={renderLeaf}
              placeholder="Enter some rich text…"
              spellCheck
              autoFocus
              onKeyDown={event => {
                for (const hotkey in HOTKEYS) {
                  if (isHotkey(hotkey, event as any)) {
                    event.preventDefault()
                    const mark = HOTKEYS[hotkey]
                    toggleMark(editor, mark)
                  }
                }
              }}
          />
        </div>
        <div className={styles.footer}>
          <Toolbar className={styles.footer_toolbar}>
            <SettingButton icon="send" />
          </Toolbar>
        </div>
        </Slate>
  )
}


const HOTKEYS = {
  'mod+b': 'bold',
  'mod+i': 'italic',
  'mod+u': 'underline',
  'mod+`': 'code',
}

const LIST_TYPES = ['numbered-list', 'bulleted-list']
const TEXT_ALIGN_TYPES = ['left', 'center', 'right', 'justify']


const Element = (props: any) => {
  let { attributes, children, element } = props
  const style = { textAlign: element.align }
  switch (element.type) {
    case 'block-quote':
      return (
        <blockquote style={style} {...attributes}>
          {children}
        </blockquote>
      )
    case 'bulleted-list':
      return (
        <ul style={style} {...attributes}>
          {children}
        </ul>
      )
    case 'heading-one':
      return (
        <h1 style={style} {...attributes}>
          {children}
        </h1>
      )
    case 'heading-two':
      return (
        <h2 style={style} {...attributes}>
          {children}
        </h2>
      )
    case 'list-item':
      return (
        <li style={style} {...attributes}>
          {children}
        </li>
      )
    case 'numbered-list':
      return (
        <ol style={style} {...attributes}>
          {children}
        </ol>
      )
    default:
      return (
        <p style={style} {...attributes}>
          {children}
        </p>
      )
  }
}

const Leaf = (props: any) => {
  let { attributes, children, leaf } = props
  if (leaf.bold) {
    children = <strong>{children}</strong>
  }

  if (leaf.code) {
    children = <code>{children}</code>
  }

  if (leaf.italic) {
    children = <em>{children}</em>
  }

  if (leaf.underline) {
    children = <u>{children}</u>
  }

  return <span {...attributes}>{children}</span>
}

const BlockButton = (props: any) => {
  let { format, icon } = props
  const editor = useSlate()
  return (
    <Button
      active={isBlockActive(
        editor,
        format,
        TEXT_ALIGN_TYPES.includes(format) ? 'align' : 'type'
      )}
      onMouseDown={(event: any) => {
        event.preventDefault()
        toggleBlock(editor, format)
      }}
    >
      <Icon>{icon}</Icon>
    </Button>
  )
}

const SettingButton = (props: any) => {
  const { icon, onClick } = props
  const editor = useSlate()
  return (
    <Button onMouseDown={onClick}
      // active={isMarkActive(editor, format)}
      >
      <Icon>{icon}</Icon>
    </Button>
  )
}

const MarkButton = (props: any) => {
  const { format, icon } = props
  const editor = useSlate()
  return (
    <Button
      active={isMarkActive(editor, format)}
      onMouseDown={(event: any) => {
        event.preventDefault()
        toggleMark(editor, format)
      }}
    >
      <Icon>{icon}</Icon>
    </Button>
  )
}

const toggleBlock = (editor: any, format: any) => {
  const isActive = isBlockActive(
    editor,
    format,
    TEXT_ALIGN_TYPES.includes(format) ? 'align' : 'type'
  )
  const isList = LIST_TYPES.includes(format)

  Transforms.unwrapNodes(editor, {
    match: n =>
      !Editor.isEditor(n) &&
      SlateElement.isElement(n) &&
      LIST_TYPES.includes(n.type) &&
      !TEXT_ALIGN_TYPES.includes(format),
    split: true,
  })
  let newProperties: Partial<SlateElement>
  if (TEXT_ALIGN_TYPES.includes(format)) {
    newProperties = {
      align: isActive ? undefined : format,
    }
  } else {
    newProperties = {
      type: isActive ? 'paragraph' : isList ? 'list-item' : format,
    }
  }
  Transforms.setNodes<SlateElement>(editor, newProperties)

  if (!isActive && isList) {
    const block = { type: format, children: [] }
    Transforms.wrapNodes(editor, block)
  }
}

const toggleMark = (editor: any, format: any) => {
  const isActive = isMarkActive(editor, format)

  if (isActive) {
    Editor.removeMark(editor, format)
  } else {
    Editor.addMark(editor, format, true)
  }
}

const isBlockActive = (editor: any, format: string, blockType = 'type') => {
  const { selection } = editor
  if (!selection) return false

  const [match] = Array.from(
    Editor.nodes(editor, {
      at: Editor.unhangRange(editor, selection),
      match: n =>
        !Editor.isEditor(n) &&
        SlateElement.isElement(n) &&
        n[blockType] === format,
    })
  )

  return !!match
}

const isMarkActive = (editor: any, format: string) => {
  const marks = Editor.marks(editor)
  return marks ? marks[format] === true : false
}

