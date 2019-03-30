
import Api from '@/api/Api'
import { getLocal, setLocal } from '@/api/LocalStore'
// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  changeRoleHeader (role) {
    setLocal({ 'activeRole': role })
  },
  getAppConfig () {
    return new Promise((resolve, reject) => {
      var role = getLocal('activeRole')
      if (role !== '' && role !== null) {
        Api().get(role + '/config').then((response) => {
          resolve(response.data)
        }).catch((error) => {
          reject(error)
        })
      }
    })
  }
}
