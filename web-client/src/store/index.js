import Vue from 'vue'
import Vuex from 'vuex'
import mainStore from './modules/mainStore'

Vue.use(Vuex)

const debug = process.env.NODE_ENV !== 'production'

export default new Vuex.Store({
  modules: {
    mainStore
  },
  strict: debug
  // plugins: debug ? [createLogger()] : []
})
