import axios from "axios";

import Auth from './Auth'

function toTopic(topic: any): any {
  const base = topic.base
  delete(topic["base"])
  return {...topic, ...base}
}

function toMessage(message: any): any {
  const contbase = message.contentBase
  const msgbase = message.base
  delete(message["base"])
  delete(message["contentBase"])
  return {...message, ...contbase, ...msgbase}
}

export class Api {
  auth = new Auth()

  basicAuthParamsFor(username: string) {
    return {
      username: username,
      password: username + "123",
    }
  }

  get basicAuthParams(): any {
    const username = (this.auth.loggedInUser || {}).id || ""
    return this.basicAuthParamsFor(username)
  }

  getApiPath(path: string): string {
    if (path.startsWith("/")) {
      return `/api/v1{path}`
    } else {
      return `/api/v1/${path}`
    }
  }

  async createUser(userid: string, fullname: string): Promise<any> {
    const resp = await axios.post(this.getApiPath(`users`), {
      "user": {
        "id": userid,
        "name": fullname,
      }
    }, {auth: this.basicAuthParamsFor(userid)})
    return resp.data
  }

  async updateUser(userid: string, fullname: string): Promise<any> {
    const resp = await axios.patch(this.getApiPath(`users/${userid}`), {
      "user": {
        "name": fullname,
      },
      "update_mask": "name"
    }, {auth: this.basicAuthParamsFor(userid)})
    return resp.data
  }

  async getUserInfos(userids: string[]): Promise<any> {
    const userIds = userids.map(uid => "ids=" + uid)
    const path = this.getApiPath(`users:batchGet?${userIds.join('&')}`)
    const resp = await axios.get(path, {auth: this.basicAuthParamsFor("admin")})
    return resp.data
  }

  async getUserInfo(userid: string): Promise<any> {
    const resp = await axios.get(this.getApiPath(`users/${userid}`), {auth: this.basicAuthParamsFor(userid)})
    return resp.data
  }

  async getTopicInfo(topicid: string): Promise<any> {
    const resp = await axios.get(this.getApiPath(`topics/${topicid}`), {auth: this.basicAuthParams})
    resp.data.topic = toTopic(resp.data.topic)
    return resp.data
  }

  async getTopics(): Promise<any> {
    const resp = await axios.get(this.getApiPath(`topics`), {auth: this.basicAuthParams})
    resp.data.topics.forEach((topic: any, index: number) => {
      resp.data.topics[index] = toTopic(topic)
    })
    return resp.data
  }

  async deleteTopic(topicid: string): Promise<any> {
    const resp = await axios.delete(this.getApiPath(`topics/${topicid}`), {auth: this.basicAuthParams})
    return resp.data
  }

  async createTopic(topic: any): Promise<any> {
    const resp = await axios.post(this.getApiPath("topics"), topic, {auth: this.basicAuthParams})
    return resp.data
  }

  async getMessages(topicId: string): Promise<any> {
    const resp = await axios.get(this.getApiPath(`topics/${topicId}/messages`), {auth: this.basicAuthParams})
    resp.data.messages.forEach((message: any, index: number) => {
      resp.data.messages[index] = toMessage(message)
    })
    return resp.data
  }

  async createMessage(topicId: string, message: any): Promise<any> {
    const path = this.getApiPath(`topics/${topicId}/messages`)
    const resp = await axios.post(path, {"messages": [message]}, {auth: this.basicAuthParams})
    return toMessage(resp.data.messages[0])
  }
}
