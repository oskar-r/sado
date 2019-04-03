import * as parseDoc from './documentTypes'

export function connect (url, store) {
  setcookie('id', '72f91914-f3da-4f89-bc3a-b12fb9444cda', 1)
  const socket = new WebSocket(url)

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
    } catch (error) {
      console.error(error)
    }
  }
}

function setcookie (name, value, days) {
  var expires = ''

  if (days) {
    var date = new Date()
    date.setTime(date.getTime() + days * 24 * 60 * 60 * 1000) // ) removed
    expires = '; expires=' + date.toGMTString() // + added
  }
  document.cookie = name + '=' + value + expires + ';path=/;domain=localhost'
}