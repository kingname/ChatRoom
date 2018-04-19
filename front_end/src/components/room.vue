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
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
          <div class="columns">
            [<div class="user_name">张三</div>][<div class="text-gray">2018-09-08 18:00:00</div>]
            :
            <div class="float-left">测试asdfjkasdfaskldfdfjasdfjasdfjkasfkaskdsjf</div>
          </div>
        </div>
        <div class="panel-footer">
          <form action="/submit" method="POST">
          <div class="input-group">
            <input type="text" class="form-input col-2" name="username" placeholder="name">
            <input type="text" class="form-input" name="content" placeholder="Hello">
            <button class="btn btn-primary input-group-btn">Send</button>
          </div>
          </form>
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
      msg: 'Welcome to Your Vue.js App'
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
