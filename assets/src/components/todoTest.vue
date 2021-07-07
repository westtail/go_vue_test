<template>
<div>
    <div>
      <label for="title">タイトル：</label>
      <input type="text" v-model="title" placeholder="タイトル">
    </div>
    <div>
      <label for="body">内容：</label>
      <textarea v-model="body" placeholder="内容"></textarea>
    </div>
    <div>
      <input type="submit" value="追加" @click="addList">
    </div>
    <ul v-for="(list,i) in lists" v-bind:key="i">
      <li>id: {{i}}, title: {{ list.title }}</li>
    </ul>
    <pre>
      {{$data}}
      <!--dataの中身を全て確認できる -->
    </pre>
    </div>
</template>

<script>
export default {
    name: 'Home',
    data: function() {
        return{
            title: null,
            body: null,
            lists: [
                {title: 'mytodo', body: 'contentcontent'}
            ]
        }
    },
    methods: {
    addList: function(){
      if(this.title === '' || this.body === '') return
      // dataオプションのlists配列にv-modelから取得したデータを追加する
      this.lists.push({title: this.title, body: this.body})
      // 追加し終わったら、フォームのvalueを空にする
      this.title = ''
      this.body = ''
    }
    },
}
</script>