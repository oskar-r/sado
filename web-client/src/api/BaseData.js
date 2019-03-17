
import Api from '@/api/Api'
import { setLocal } from '@/api/LocalStore'
// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  changeRoleHeader (role) {
    setLocal({ 'activeRole': role })
  },
  get () {
    return new Promise((resolve, reject) => {
      Api().get('admin/basestats').then((response) => {
        const latest = response.data.latest_registrations.map(element => {
          if (element.school === undefined) {
            element.school = ''
          }
          return {
            id: element.id,
            school: element.school.split(',')[0] !== undefined ? element.school.split(',')[0] : 'N/A',
            municiplaity: element.school.split(',')[1] !== undefined ? element.school.split(',')[1] : 'N/A',
            email: element.email !== undefined ? element.email : 'N/A',
            registration: element.registration_date
          }
        })
        resolve({
          latest_registrations: latest,
          number_of_users: response.data.number_of_users,
          active_users: response.data.active_users,
          active_users_weekly: response.data.active_users_weekly
        })
      }).catch((error) => {
        reject(error)
      })
    })
  },
  getAppConfig () {
    return new Promise((resolve, reject) => {
      Api().get('admin/config').then((response) => {
        resolve(response.data)
      }).catch((error) => {
        reject(error)
      })
    })
  }
}
