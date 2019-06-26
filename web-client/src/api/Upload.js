
import Api from '@/api/Api'

// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  post (file) {
    return new Promise((resolve, reject) => {
      Api().post('datasets/upload?name=' + btoa(file.name) + '&content-type=' + btoa(file.type),
        file).then((response) => {
        console.log(response)
        resolve(true)
      }).catch((error) => {
        reject(error)
      })
    })
  }
}
