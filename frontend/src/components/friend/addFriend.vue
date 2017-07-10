<template>
  <div>
    <el-menu theme="dark" class="el-menu-demo" mode="horizontal">
        <el-menu-item index="1" v-on:click="toContacts">联系人</el-menu-item>
        <el-menu-item index="2" v-on:click="logout">退出登录</el-menu-item>
    </el-menu>
    <el-row>
        <el-col :span="12" :push="6">
            <div class="sub-title">查找好友</div>
            <el-input placeholder="请输入好友姓名" v-model="friend">
                <el-button slot="append" icon="search"></el-button>
            </el-input>
            <el-button type="primary" @click="addFriend">添加</el-button>
            <el-button type="primary" v-on:click="toContacts">返回联系人</el-button>
            <el-alert v-show="addFriAlert === 0" title="" @bind:type="addFriType">{{addFriMsg}}</el-alert>
        </el-col>
    </el-row>
    
  </div>
</template>
<script>
import auth from '../../service/auth'
import friend from '../../service/friend'

export default {
  data () {
    return {
      friend: '',
      addFriAlert: 1,
      addFriMsg: '',
      addFriType: 'success'
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
    addFriend: function () {
      var own = auth.getCurName()
      var addFriend = this.friend
      friend.addFriend(this, own, addFriend)
    },
    FriAlert: function (addFriAlert, addFriMsg, addFriType) {
      this.addFriAlert = addFriAlert
      this.addFriMsg = addFriMsg
      this.addFriType = addFriType
    }
  },
  created: function () {
    if (!auth.checkAuth()) {
      this.$router.push('/login')
    }
  }
}
</script>
