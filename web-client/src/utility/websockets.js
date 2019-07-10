import * as parseDoc from './documentTypes'
import { infoToast } from './toastMessages'
require('dotenv').config()

export function connect (url, store, toast, vm) {
  console.log('Connnecting websocket with :' + store.getters['mainStore/getUserID'])
  setcookie('id', store.getters['mainStore/getUserID'], 1)
  const socket = new WebSocket(url)
  vm.$ws = socket
  socket.onopen = function () {
    console.log('Connected to ' + url)
  }
  socket.onmessage = function (event) {
    var data = JSON.parse(event.data)
    data.last_modified = data.time
    data.category = parseDoc.byType(event.data.content_type)
    try {
      console.log(store)
      store.dispatch('mainStore/updateDatasets', data)
      const it = infoToast(data.name)
      toast.info(it.message, it.style)
    } catch (error) {
      console.error(error)
    }
  }
  socket.onclose = function (event) {
    console.log(event)
    console.log('SOCKET CLOSED')
  }
}

function setcookie (name, value, days) {
  var expires = ''

  if (days) {
    var date = new Date()
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000) // ) removed
    expires = '; expires=' + date.toGMTString() // + added
  }
  document.cookie = name + '=' + value + expires + ';path=/;domain=' + process.env.VUE_APP_DOMAIN
}
