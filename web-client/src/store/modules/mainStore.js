import BaseData from '../../api/BaseData'
import Upload from '../../api/Upload'
import Login from '../../api/Login'
import Datasets from '../../api/Datasets'

const state = {
  logedIn: Login.logedIn(),
  username: localStorage.username,
  identifier: localStorage.identifier,
  error: {
    code: 0,
    message: ''
  },
  datasets: [],
  documents: [],
  appRoutes: [],
  roles: [],
  activeRole: '',
  socket: {
    isConnected: false,
    message: '',
    reconnectError: false
  }
}
// getters
const getters = {
  isLogedIn () {
    return state.logedIn
  },
  getUsername () {
    return state.username
  },
  getIdentifier () {
    return state.identifier
  },
  errorMessage () {
    return state.error
  },
  getBaseData () {
    return state.baseData
  },
  appRoutes () {
    return state.appRoutes
  },
  roles () {
    return state.roles
  },
  getActiveRole () {
    return state.activeRole
  },
  getDatasets () {
    return state.datasets
  },
  getNamedDatasets (name) {
    var i = -1
    state.datasets.forEach((item, index) => {
      if (item.name === name) {
        i = index
      }
    })
    if (i > -1) {
      return state.datasets[i]
    }
    return null
  }
}
// mutatations
const mutations = {
  setLogedIn (state, newState) {
    state.logedIn = newState
  },
  setAppRoutes (state, routes) {
    state.appRoutes = routes
  },
  setErrorMessage (state, error) {
    if (error !== undefined) {
      state.error = error
    } else {
      state.error = {
        code: 0,
        message: ''
      }
    }
  },
  resetErrorMessage (state) {
    state.error = {
      code: 0,
      message: ''
    }
  },
  setIdentifier (state, value) {
    state.identifier = value
  },
  setUsername (state, value) {
    state.username = value
  },
  setRoles (state, roles) {
    state.roles = roles
  },
  setActiveRole (state, role) {
    state.activeRole = role
  },
  setDatasets (state, sets) {
    state.datasets = sets
  },
  setDataset (state, set) {
    state.datasets.push(set)
  },
  setPreviewDetails (state, preview, filename) {
    state.datasets.forEach((item, index) => {
      if (item.name === filename) {
        // state.datasets[index].preview = preview
      }
    })
  }
}
// actions
const actions = {
  logIn ({ commit, dispatch }, creds) {
    return new Promise((resolve, reject) => {
      Login.logIn(creds.username, creds.password).then((resp) => {
        console.log(resp)
        commit('setLogedIn', true)
        commit('setIdentifier', resp.identifier)
        commit('setUsername', resp.username)
        commit('setActiveRole', resp.role)
        dispatch('getAppConfig')
        dispatch('getMyDatasets')
        resolve(resp)
      }).catch((error) => {
        commit('setErrorMessage', { message: 'Ett fel intäffade vid inloggning', code: 403 })
        console.log(error)
        commit('setLogedIn', false)
        commit('setIdentifier', '')
        commit('setUsername', '')
        reject(error)
      })
    }).catch((error) => {
      console.log(error)
    })
  },
  updateDatasets ({ commit }, dataset) {
    commit('setDataset', dataset)
  },
  logOut ({ commit }) {
    commit('setLogedIn', false)
  },
  getAppConfig ({ commit }) {
    BaseData.getAppConfig().then((resp) => {
      console.log(resp)
      if (resp.routes !== undefined) {
        resp.routes.sort((a, b) => a.order - b.order)
        commit('setAppRoutes', resp.routes)
      } else {
        commit('setAppRoutes', [])
      }
      commit('setRoles', resp.roles)
      commit('setActiveRole', resp.active_role)
    }).catch((error) => {
      console.error(error.response)
      var e = {
        message: 'Can´t get response from server - Is internet working properly?',
        code: ''
      }
      if (error.response !== undefined && error.response.status !== undefined) {
        switch (error.response.status) {
          case 401:
            e.message = error.response.data
            e.code = error.response.status
            commit('setLogedIn', false)
            break
          case 403:
            e.message = error.response.data
            e.code = error.response.status
            this._vm.toasted.show('Token have expired, login again')
            commit('setLogedIn', false)
            break
          case 500:
            e.message = 'Server error'
            e.code = error.response.status
            commit('setErrorMessage', e)
            break
        }
      }
    })
  },
  getMyDatasets ({ commit }) {
    Datasets.get().then((resp) => {
      commit('setDatasets', resp)
    }).catch((error) => {
      console.log(error)
    })
  },
  previewFile ({ commit }, filename) {
    return new Promise((resolve, reject) => {
      Datasets.preview(filename).then((resp) => {
        commit('setPreviewDetails', resp, filename)
        resolve(resp)
      }).catch((error) => {
        console.error(error)
        reject(error)
      })
    })
  },
  showErrorMessage ({ commit }, error) {
    commit('setErrorMessage', error)
  },
  resetErrorMessage ({ commit }) {
    commit('resetErrorMessage')
  },
  upload ({ commit, dispatch }, file) { // Change header and reload config from the server !!Remeber to clean up vuex
    Upload.post(file).then((resp) => {
      console.log(resp)
    }).catch((err) => {
      console.error(err)
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
  // mutations
}
