<template>
  <div class="login-insert">
    <el-input placeholder="username" v-model="username">
      <template slot="prepend"> 昵称：</template>
    </el-input>
    <el-input placeholder="password" v-model="password" type="password" :icon="inputi">
      <template slot="prepend"> 密码：</template>
    </el-input>
    <el-button @click="login">登录</el-button>
    <el-button @click="register">没有帐号？注册</el-button>
  </div>
</template>
<script type="text/javascript">
import auth from '../../service/auth'

export default {
  data: function () {
    return {
      username: '',
      password: '',
      inputi: ''
    }
  },
  methods: {
    login: function () {
      var creds = {
        username: this.username,
        password: this.password
      }
      auth.login(this, creds)
    },
    register: function () {
      this.$emit('register')
    }
  }
}
</script>
<style scoped>
.login-insert {
  div {
    margin-bottom: .8em;
  }
  .el-select .el-input {
    width: 30%;
  }
}
</style>
