import axios from "axios"

const API_URL = "http://localhost:3001/api"

export default {
  kontrolEt(domain) {
    return axios.post(`${API_URL}/dnscheck`, { domain })
  }
}