<template>
  <div>
    <el-menu theme="dark" class="el-menu-demo" mode="horizontal">
        <el-menu-item index="1" v-on:click="toContacts">联系人</el-menu-item>
        <el-menu-item index="2" v-on:click="logout">退出登录</el-menu-item>
    </el-menu>
    <el-row>
        <el-col :span="12" :push="6">
            <div class="sub-title">删除好友</div>
            <el-alert v-for="c in contacts" :key="c" title="" type="info" @close="deleteFriend(c)">{{c}}</el-alert>
        </el-col>
    </el-row>   
  </div>
</template>
<script>
import auth from '../../service/auth'
import friend from '../../service/friend'

export default {
  data: function () {
    return {
      contacts: []
    }
  },
  methods: {
    toContacts: function () {
      this.$router.push('/home')
    },
    logout: function () {
      var creds = {
        username: auth.getCurName()
      }
      auth.logout(this, creds)
    },
    deleteFriend: function (c) {
      var curUsername = auth.getCurName()
      friend.deleteFriend(this, curUsername, c)
    }
  },
  created: function () {
    if (!auth.checkAuth()) {
      this.$router.push('/login')
    }
    // get friend lists
    var curUsername = auth.getCurName()
    friend.listFriends(this, curUsername)
  }
}
</script>
