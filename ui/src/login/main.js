import Vue from 'vue'
import App from './App.vue'
import request from '../request'

import 'bootstrap/dist/css/bootstrap.min.css'

Vue.config.productionTip = false

Vue.prototype.auth = request.auth

new Vue({
  render: h => h(App)
}).$mount('#app')
