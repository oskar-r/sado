import axios from 'axios'
require('dotenv').config()

export default () => {
  return axios.create({
    baseURL: process.env.VUE_APP_BASE_API,
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
      'Authorization': 'Bearer ' + JSON.parse(localStorage.getItem('jwt'))
    }
  })
}
