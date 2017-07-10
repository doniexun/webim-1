<template>
  <div class="menu">
    <el-menu theme="dark" mode="horizontal">
        <el-menu-item index="1" v-on:click="toAddFriend">添加好友</el-menu-item>
        <el-menu-item index="2" v-on:click="toDeleteFriend">删除好友</el-menu-item>
        <el-menu-item index="3" v-on:click="logout">退出登录</el-menu-item>
    </el-menu>
    <contacts></contacts>
  </div>
</template>
<script type="text/javascript">
import contacts from '@/components/contacts/Contacts'
import addFriend from '@/components/friend/addFriend'
import auth from '../service/auth'

export default {
  data () {
    return {}
  },
  methods: {
    toAddFriend: function () {
      this.$router.push('/addFriend')
    },
    toDeleteFriend: function () {
      this.$router.push('/deleteFriend')
    },
    logout: function () {
      var creds = {
        username: auth.getCurName()
      }
      auth.logout(this, creds)
    }
  },
  components: {
    contacts,
    addFriend
  },
  created: function () {
    if (!auth.checkAuth()) {
      this.$router.push('/login')
    }
  }
}
</script>
