import Vue from 'vue'
import App from './components/App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import router from './router'
import store from './store/index'
import Validator from './plugins/validator'

import BootstrapVue from 'bootstrap-vue'

Vue.use(BootstrapVue)

Vue.config.productionTip = false
Vue.use(VueAxios, axios)
Vue.use(Validator)
// Vue.use(VTooltip)
// Vue.use(Comment)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
