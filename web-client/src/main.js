import Vue from 'vue'
import App from './components/App.vue'
import axios from 'axios'
import VueAxios from 'vue-axios'
import router from './router'
import store from './store/index'
import Validator from './plugins/validator'
import Toasted from 'vue-toasted'

import BootstrapVue from 'bootstrap-vue'
import { library } from '@fortawesome/fontawesome-svg-core'
import { faUserSecret } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faUserSecret)

Vue.component('font-awesome-icon', FontAwesomeIcon)

Vue.use(BootstrapVue)
Vue.use(Toasted, {
  theme: 'toasted-primary',
  position: 'bottom-right',
  iconPack: 'fontawesome',
  duration: 5000
})

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
