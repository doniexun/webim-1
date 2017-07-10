<template>
  <div class="register">
    <el-input placeholder="username" type="text" v-model="username" :icon="inputi">
      <template slot="prepend"> 昵称：</template>
    </el-input>
    <el-input placeholder="password" type="password" v-model="password" :icon="inputi">
      <template slot="prepend"> 密码：</template>
    </el-input>
    <el-button @click="userregister()">注册</el-button>
    <el-button @click="toLogin()">登录</el-button>
    <p>{{msg}}</p>
  </div>
</template>
<script type="text/javascript">
import auth from '../../service/auth'

export default {
  props: {
    prop: {
      type: Object,
      default () {
        return {}
      }
    }
  },
  components: {

  },
  data: function () {
    return {
      username: '',
      password: '',
      inputi: '',
      msg: ''
    }
  },
  computed: {

  },
  methods: {
    userregister: function () {
      var creds = {
        username: this.username,
        password: this.password
      }
      auth.register(this, creds)
    },
    toLogin: function () {
      this.msg = ''
      this.$emit('cancel')
    }
  },
  created: function () {}
}
</script>
<style scoped>
$red: #FF4949;
.register {
  div {
    margin-bottom: .8em;
  }
  .el-select .el-input {
    width: 30%;
  }
}
</style>
