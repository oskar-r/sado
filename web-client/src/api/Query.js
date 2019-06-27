import Api from '@/api/Api'

// import { setLocal } from '@/api/LocalStore'
// const apiPrefix = process.env.BASE_API

export default {
  async run (query) {
    const regex = /.*FROM\s(?<dataset>.*\.(csv|tar\.gz))\s.*/gm
    const m = regex.exec(query)
    if (m.groups.dataset === undefined) {
      const e = new Error('no valid dataset name (must be csv or tar.gz)')
      return e
    }
    try {
      const result = await Api().post('datasets/query', {
        query: query,
        dataset: m.groups.dataset,
        record_delimiter: '\n',
        field_delimiter: ',',
        output: 'json'
      })
      return result
    } catch (e) {
      throw e
    }
  }
}
