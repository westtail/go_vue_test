<template>
  <div id="app">
    <Home name="ようこそ西尾亮太"/>
    <input v-model="message" placeholder="edit me">
    <div class="btn-container"><button type="submit" class="btn btn-primary">送信</button></div>
    <p>Message is: {{ message }}</p>
    {{ param_test}}
  </div>
</template>

<script>
import Home from './components/Home.vue'

export default {
  name: 'app',
  components: {
    Home
  },
  data: function() {
    return{
      message: null,
      param_test: null
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/').then(response => {
        console.log(response)
        this.message = response.data
    })
    this.$axios.get('http://localhost/api/hoge?key=test').then(response => {
        console.log(response)
        this.param_test = response.data
    })
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
