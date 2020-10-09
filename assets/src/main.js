import Vue from 'vue'
import App from './App.vue'

// ここから
import axios from 'axios'
Vue.prototype.$axios = axios
// ここまでを追加

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
