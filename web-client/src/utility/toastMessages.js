function errorToast (msg) {
  return {
    message: '<div class="toast-title aria-label="Title" style=""> Error </div><div aria-live="polite" role="alertdialog" class="toast-message">' + msg + '</div>',
    style: {
      className: 'toast-error',
      icon: 'fa-exclamation-triangle'
    } }
}

function infoToast (msg) {
  return {
    message: '<div class="toast-title aria-label="Title" style=""> Uploaded </div><div aria-live="polite" role="alertdialog" class="toast-message" style="">' + msg + '</div>',
    style: {
      className: 'toast-info',
      icon: 'fa-file-upload'
    } }
}

export { errorToast, infoToast }
