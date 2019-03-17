
import Api from '@/api/Api'

// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  post (file) {
    return new Promise((resolve, reject) => {
      /*
      var oReq = new XMLHttpRequest()
      oReq.open('POST', 'http://localhost:8091/user/upload?name=' + file.name, true)
      // oReq.setRequestHeader('Content-Length',toString(file.size))
      oReq.onprogress = (e) => {
        console.log(e)
      }
      oReq.onerror = (error) => {
        reject(error)
      }
      oReq.onloadend = (resp) => {
        resolve(resp)
      }
      // let reader = new FileReader();
      oReq.send(file)
      */
      Api().post('user/upload?name=' + file.name,
        file).then((response) => {
        console.log(response)
        resolve(true)
      }).catch((error) => {
        reject(error)
      })
    })
  }
}
