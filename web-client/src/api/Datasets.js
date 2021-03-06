
import Api from '@/api/Api'

// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  get () {
    return new Promise((resolve, reject) => {
      Api().get('datasets/list').then((response) => {
        resolve(response.data)
      }).catch((error) => {
        reject(error)
      })
    })
  },
  preview (file) {
    return new Promise((resolve, reject) => {
      console.log(file)
      Api().get('datasets/preview?file=' + btoa(file)).then((response) => {
        resolve(response.data)
      }).catch((error) => {
        reject(error)
      })
    })
  }
}
