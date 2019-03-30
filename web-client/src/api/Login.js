import Api from '@/api/Api'
import { setLocal, getLocal } from '@/api/LocalStore'

export default {
  logIn (username, password) {
    return new Promise((resolve, reject) => {
      Api().post('login', {
        username: username,
        password: password,
        role: 'user'
      }).then((response) => {
        Api().defaults.headers.common['Authorization'] = response.data.token
        setLocal({
          'logedIn': true,
          'username': response.data.username,
          'jwt': response.data.token,
          'activeRole': response.data.role
        })
        resolve({ identifier: response.data.token, username: response.data.username, activeRole: response.data.role })
      }).catch((error) => {
        localStorage.clear()
        setLocal({ 'logedIn': false })
        reject(error)
      })
    })
  },
  logOut () {
    localStorage.clear()
    setLocal({ 'logedIn': false })
  },
  logedIn () {
    const lo = getLocal('logedIn')
    if (!lo || lo === null) {
      return false
    }
    return true
  }
}
