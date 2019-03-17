const Validator = {
  install (Vue, options) {
    Vue.prototype.$validateData = (data) => {
      let users = []
      try {
        users = data.filter((elem, row) => {
          if (elem[0] !== undefined && elem[0] !== 'FirstName' && elem[0] !== 'First name' && elem[0] !== 'Förnamn') {
            return elem
          }
        }).map((elem, row) => {
          let user = {}
          if (elem[0] !== undefined) {
            user.fname = elem[0]
          } else {
            throw new Error('Wrong first name on row ' + (parseInt(row) + 1))
          }
          if (elem[1] !== undefined) {
            user.lname = elem[1]
          } else {
            throw new Error('Wrong last name on row ' + (parseInt(row) + 1))
          }
          if (elem[2] !== undefined && emailValid(elem[2])) {
            user.email = elem[2]
          } else {
            throw new Error('Wrong email on row ' + (parseInt(row) + 1))
          }
          if (elem[3] !== undefined && schoolValid(elem[3])) {
            user.school = elem[3]
          } else {
            throw new Error('Wrong school ' + (parseInt(row) + 1))
          }
          if (elem[4] !== undefined && schoolTypeValid(elem[4])) {
            user.schoolType = elem[4]
          } else {
            throw new Error('Wrong school type ' + (parseInt(row) + 1))
          }
          if (elem[5] !== undefined && validDate(elem[5])) {
            user.startDate = parseDate(elem[5])
          } else {
            throw new Error('Wrong start date format ' + (parseInt(row) + 1))
          }
          if (elem[6] !== undefined && validDate(elem[6])) {
            user.endDate = parseDate(elem[6])
          } else {
            throw new Error('Wrong end date format ' + (parseInt(row) + 1))
          }
          if (elem[7] !== undefined && permissionValid(elem[7])) {
            user.permission = elem[7]
          } else {
            throw new Error('Wrong type code ' + (parseInt(row) + 1))
          }
          return user
        })
      } catch (error) {
        console.log(error)
        throw error
      }
      return users
    }
  }
}
export default Validator
/* Used in both upload and modify */

function parseDate (eDate) {
  return eDate.toLocaleDateString('sv-SE')
}
function validDate (date) {
  var re = /\d{4}-[0-1]{1}[0-9]{1}-[0-3]{1}[0-9]{1}/
  return re.test(parseDate(date))
}

function permissionValid (data) {
  return true
}

function emailValid (data) {
  var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
  return re.test(data)
}

function schoolValid (data) {
  return true
}

function schoolTypeValid (data) {
  var re = /(^gym|kom|gr1|gr2|gr3|sfi|sar|spe)$/
  return re.test(data)
}
