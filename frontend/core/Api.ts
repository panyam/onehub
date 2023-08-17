import axios from "axios";

export class Api {
  getApiPath(path: string): string {
    if (path.startsWith("/")) {
      return `/api/v1{path}`
    } else {
      return `/api/v1/${path}`
    }
  }

  async getUserInfo(userid: string): Promise<any> {
    const resp = await axios.get(this.getApiPath(`users/${userid}`))
    return resp.data
  }

  async getTopicInfo(topicid: string): Promise<any> {
    const resp = await axios.get(this.getApiPath(`topics/${topicid}`))
    return resp.data
  }

  async getTopics(): Promise<any> {
    const resp = await axios.get(this.getApiPath(`topics`))
    return resp.data
  }

  async deleteTopic(topicid: string): Promise<any> {
    const resp = await axios.delete(this.getApiPath(`topics/${topicid}`))
    return resp.data
  }

  async createTopic(topic: any): Promise<any> {
    const resp = await axios.post("/v1/topics", topic)
    return resp.data
  }

  async getMessages(topicId: string): Promise<any> {
    const resp = await axios.get(this.getApiPath(`topics/${topicId}/messages`))
    return resp.data
  }

  async createMessage(topicId: string, message: any): Promise<any> {
    const resp = await axios.post(this.getApiPath(`topics/${topicId}/messages`), message)
    return resp.data
  }
}
