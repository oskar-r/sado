export function byType (contentType) {
  var dc = 'file'
  switch (contentType) {
    case 'application/zip':
      dc = 'file-archive'
      break
    case 'application/json':
      dc = 'file-code'
      break
    case 'text/plain':
      dc = 'file-alt'
      break
    case 'text/csv':
      dc = 'file-csv'
      break
    case 'application/x-gzip':
      dc = 'file-archive'
      break
    case 'text/javascript':
      dc = 'file-code'
      break
  }
  return dc
}
