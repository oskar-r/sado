export function setLocal (kv) {
  if (process.browser !== undefined && process.browser) {
    try {
      for (var key in kv) {
        if (typeof key === 'string') {
          localStorage.setItem(key, JSON.stringify(kv[key]))
        } else {
          console.error('Invalid' + key)
        }
      }
    } catch (err) {
      console.error(err)
    }
  }
}

export function getLocal (key) {
  if (process.browser !== undefined && process.browser) {
    try {
      if (typeof key === 'string') {
        try {
          var value = JSON.parse(localStorage.getItem(key))
          return value
        } catch (err) {
          if (err.name === 'SyntaxError') {
            return localStorage.getItem(key)
          }
          console.error('Error:' + err)
        }
      }
    } catch (err) {
      console.error(err)
    }
  }
  return null
}
