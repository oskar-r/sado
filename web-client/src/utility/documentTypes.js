export function byType (contentType) {
  var dc = 'document'
  switch (contentType) {
    case 'application/zip':
      dc = 'dataset'
      break
    case 'application/json':
      dc = 'dataset'
      break
    case 'text/plain':
      dc = 'dataset'
      break
    case 'text/csv':
      dc = 'dataset'
      break
    case 'application/x-gzip':
      dc = 'dataset'
      break
  }
  return dc
}
