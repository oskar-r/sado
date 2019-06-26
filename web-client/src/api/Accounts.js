import Api from '@/api/Api'

// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  create (accountdetails) {
    return new Promise((resolve, reject) => {
      Api().post('admin/create-account', {
        username: accountdetails.username,
        password: accountdetails.password,
        my_bucket: accountdetails.bucket
      }).then((response) => {
        resolve(response.data)
      }).catch((error) => {
        reject(error)
      })
    })
  }
}
