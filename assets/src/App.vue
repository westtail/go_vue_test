<template>
  <div id="app">
    <Home name="ようこそ西尾亮太"/>
    test
    {{test}}
    test2
    {{test2}}
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
      test: null,
      test2: null
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/')
      .then(response => {
        console.log(response)
        this.test = response
      })
    this.$axios.get('http://localhost/api/test')
      .then(response => {
        console.log(response)
        this.test2 = response
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
