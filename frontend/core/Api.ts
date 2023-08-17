import axios from "axios";

export class Api {
  async getUserInfo(userid: string): Promise<any> {
    const resp = await axios.get(`/v1/users/${userid}`)
    return resp.data
  }

  async getTopicInfo(topicid: string): Promise<any> {
    const resp = await axios.get(`/v1/topics/${topicid}`)
    return resp.data
  }
}
