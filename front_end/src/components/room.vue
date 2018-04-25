<template>
  <div class="container">
  <div class="columns col-10 col-mx-auto">
    <!-- <div class="column col-2 col-xs-2" id="user_list">
      <div class="panel" style="height:600px; overflow:auto;">
        <div class="panel-header">
          <div class="panel-title">在线用户(32)</div>
        </div>
        <div class="pabel-body" style="overflow:auto;">
          <p>张三</p>
          <p>张三</p>
          <p>张三李四</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三asdfasdf</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三</p>
          <p>张三</p>
        </div>
      </div>
    </div> -->
    <div class="column col-12 col-xs-12" id="content">
      <div class="panel" style="height:600px;">
        <div class="panel-header">
          <div class="panel-title">匿名聊天群</div>
        </div>
        <div class="panel-body" style="overflow:auto;">
          <div class="columns" v-for="c in msg" :key="c._id">
            <div class="user_name">
              <span class="label label-secondary" v-if="Cid != c.cid">{{c.username}}</span>
              <span class="label label-success" v-else>{{c.username}}</span>
            </div>
            <div class="text-gray">
              <span class="label">{{c.create_time}}</span>
            </div>
            <br />
            <div class="float-right">{{c.content}}</div>
          </div>
        </div>
        <div class="panel-footer">
          <div class="input-group">
            <input v-model="userName" type="text" class="form-input col-2" placeholder="name">
            <input v-model="Content" type="text" class="form-input" name="content" placeholder="Hello">
            <button class="btn btn-primary input-group-btn" v-on:click="submit">Send</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: 'room',
  data () {
    return {
      msg: [],
      userName: '',
      Content: '',
      Cid: 0
    }
  },
  mounted () {
    this.Cid = this.$md5(Date.now().toString())
    console.log(this.Cid)
    this.get_content_list()
  },
  methods: {
    scrollToEnd: function () {
      var container = this.$el.querySelector('.panel-body')
      console.log(container.scrollHeight)
      container.scrollTop = container.scrollHeight
    },
    get_content_list: function () {
      this.$ajax({
        method: 'get',
        url: '/content'
      }).then(response => {
        var data = response.data
        this.msg = data
        this.scrollToEnd()
      })
    },
    submit: function () {
      if (!this.userName.trim() || !this.Content.trim()) {
        alert('姓名和内容不能为空！')
      }
      var ts = Date.now()
      var time = new Date(ts)
      var timeStr = time.getFullYear() + '-' + (time.getMonth() + 1) + '-' + time.getDate() + ' ' + time.getHours() + ':' + time.getMinutes() + ':' + time.getSeconds()
      this.$ajax({
        method: 'post',
        url: '/submit',
        data: {'username': this.userName, 'content': this.Content, 'create_time': timeStr, 'cid': this.Cid},
        headers: {'Content-Type': 'application/json'}
      }).then(response => {
        if (response.data.result) {
          this.get_content_list()
          this.Content = ''
        }
      })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
