import axios from 'axios'

const API_URL = 'http://localhost:3001/api'

export default {
  send(message) {
    return axios.post(`${API_URL}/smtp/send`, message)
  },
}
