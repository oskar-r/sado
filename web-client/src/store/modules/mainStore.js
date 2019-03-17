import BaseData from '../../api/BaseData'
import Upload from '../../api/Upload'

const state = {
  username: localStorage.username,
  identifier: localStorage.identifier,
  error: {
    code: 0,
    message: ''
  },
  baseData: {
    numberOfUsers: 0,
    latestUsers: [],
    activeUsers: 0,
    activeUsersWeekly: 0
  },
  appRoutes: [],
  roles: [],
  activeRole: ''
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
  getNumberOfUsers () {
    return state.baseData.numberOfUsers
  },
  getLatestUsers () {
    return state.baseData.latestUsers
  },
  getActiveUsers () {
    return state.baseData.activeUsers
  },
  getActiveUsersWeekly () {
    return state.baseData.activeUsersWeekly
  },
  appRoutes () {
    return state.appRoutes
  },
  roles () {
    return state.roles
  },
  getActiveRole () {
    return state.activeRole
  }
}
// mutatations
const mutations = {
  setLogoOutState (state, newState) {
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
  setActiveRoles (state, role) {
    state.activeRole = role
  },
  setBaseData (state, value) {
    state.baseData = value
  }
}
// actions
const actions = {
  logOut ({ commit }) {
    commit('setLogoOutState', false)
  },
  getBaseData ({ commit }) {
    BaseData.get().then((resp) => {
      commit('setBaseData', {
        numberOfUsers: resp.number_of_users,
        latestUsers: resp.latest_registrations,
        activeUsers: resp.active_users,
        activeUsersWeekly: resp.active_users_weekly
      })
    }).catch((error) => {
      console.error(error.response)
      var e = {
        message: 'Can´t get response from server - Is internet working properly?',
        code: ''
      }
      if (error.response !== undefined && error.response.status !== undefined) {
        switch (error.response.status) {
          case 403:
            e.message = error.response.data
            e.code = error.response.status
            commit('setLogoOutState', false)
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
  getAppConfig ({ commit }) {
    BaseData.getAppConfig().then((resp) => {
      if (resp.UI_routers !== undefined) {
        resp.UI_routers.sort((a, b) => a.order - b.order)
        commit('setAppRoutes', resp.UI_routers)
      } else {
        commit('setAppRoutes', [])
      }
      commit('setRoles', resp.roles)
      commit('setActiveRoles', resp.active_role)
    }).catch((error) => {
      console.error(error.response)
      var e = {
        message: 'Can´t get response from server - Is internet working properly?',
        code: ''
      }
      if (error.response !== undefined && error.response.status !== undefined) {
        switch (error.response.status) {
          case 403:
            e.message = error.response.data
            e.code = error.response.status
            commit('setLogoOutState', false)
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
