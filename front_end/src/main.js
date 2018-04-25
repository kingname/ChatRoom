// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'
import md5 from 'js-md5'

require('./assets/css/spectre.css')

Vue.config.productionTip = false
Vue.prototype.$ajax = axios
Vue.prototype.$md5 = md5

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
