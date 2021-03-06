
import Api from '@/api/Api'
import { setLocal } from '@/api/LocalStore'
// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  changeRoleHeader (role) {
    setLocal({ 'activeRole': role })
  },
  getAppConfig () {
    return new Promise((resolve, reject) => {
      Api().get('user/config').then((response) => {
        console.log(response)
        resolve(response.data)
      }).catch((error) => {
        reject(error)
      })
    })
  }
}
