import Vue from 'vue'
import Router from 'vue-router'
import Home from './views/Home.vue'
import Query from './views/Query.vue'
import Upload from './views/Upload.vue'
import Datasets from './views/Datasets.vue'
import Accounts from './views/Accounts.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/query',
      name: 'query',
      component: Query
    },
    {
      path: '/upload',
      name: 'upload',
      component: Upload
    },
    {
      path: '/datasets',
      name: 'datasets',
      component: Datasets
    },
    {
      path: '/accounts',
      name: 'accounts',
      component: Accounts
    }
  ]
})
