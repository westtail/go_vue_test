<template>
  <div id="app">
    <h1>TodoApp</h1>
    <div>
      <label for="input">入力：</label>
      <input type="text" v-model="input" placeholder="入力値">
    </div>
    <div>
      <input type="submit" value="送信" @click="getInput">
    </div>
    <button @click="getDB"></button>
    {{ message}}
    {{ param}}
    <TodoTest></TodoTest>
  </div>
</template>

<script>
import TodoTest from './components/todoTest.vue'

export default {
  name: 'app',
  components: {
    TodoTest
  },
  data: function() {
    return{
      message: null,
      param: null,
      input: null,
    }
  },
  // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
  created () {
    this.$axios.get('http://localhost/api/connect').then(response => {
        console.log(response)
        this.message = response.data
    })
  },
  methods: {
    getInput:function(){
      this.$axios.get('http://localhost/api/params',{
        params:{
          key: this.input
        }
      }).then(response => {
        console.log(response)
        this.param = response.data
      })
    },
    getDB:function(){
      this.$axios.get('http://localhost/api/tasks').then(response => {
        console.log(response.data)
        //this.param = response.data
      })
    }
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
