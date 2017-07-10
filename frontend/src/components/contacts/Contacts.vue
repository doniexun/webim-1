<template>
    <div class="contacts">
        <el-row>
            <!-- left contact list -->
            <el-col :span="6" :push="3">
                <el-card class="c-left">
                    <div class="left-header">
                    {{curUsername}}
                    </div>
                    <div class="left-list">
                        <div class="left-contact" v-for="contact in contacts" v-on:click="startChat(contact)">
                            <span v-if="contact.msgNum>0">{{contact.msgNum}}</span>
                            <span>{{contact}}</span>
                        </div>
                    </div>
                </el-card>
            </el-col>
            <!-- end left contact list -->

            <!-- right chat box -->
            <el-col :span="12" :push="3">
                <el-card class="c-right">
                    <div class="right-chatbox-main-header">
                        {{chatTo}}
                    </div>
                    <div class="scroll-wrapper right-chatbox-main" id="right-chatbox-main">
                        <div v-for="msg in filterMsgBySender(curUsername, chatTo)" :key="msg.send_time" v-show="chatTo.length !== 0" class="chatGroupMsgList">
                            <p>
                                {{msg.sender}}-->{{msg.receiver}}:{{msg.msg}}
                            </p>
                        </div>
                    </div>
                    <el-input type="textarea" :rows="1" class="right-chat-texteditor" v-model="msgToSend">
                        {{msgToSend}}
                    </el-input>
                    <el-button type="primary" v-on:click="sendMsg" :disabled="msgToSend.length === 0 || chatTo.length === 0">
                        Send
                    </el-button>
                </el-card>
            </el-col>
            <!-- end right chat box -->

        </el-row>
    </div>
</template>
<script>
import auth from '../../service/auth'
import friend from '../../service/friend'
import message from '../../service/message'
import commondata from '../../service/commondata'

export default {
  data: function () {
    return {
      curUsername: '',
      contacts: [],
      msgToSend: '',
      chatTo: '',
      chatData: [],
      unreadMsgs: [],
      ws: null
    }
  },
  methods: {
    sendMsg: function () {
      // init ws if ws is null
      if (this.ws === null) {
        this.ws = message.createWS(this.curUsername)
      }
      this.ws.onerror = function () {
        alert('create ws error, please login')
        // clean sessionStorage login ingo
        auth.deleteLocalSession(this.curUsername)
        this.$router.push('/login')
      }
      // generate msg data
      var msgObj = {
        id: 0,
        sender: this.curUsername,
        receiver: this.chatTo,
        msg: this.msgToSend,
        send_time: message.getTimestamp(),
        state: 'msg_send'}
      // send data to server by websocket
      this.ws.send(JSON.stringify(msgObj))
      // sync locally
      this.chatData.push(msgObj)
      // clean msgToSend
      this.msgToSend = ''
    },
    startChat: function (contact) {
      this.chatTo = contact
    },
    filterMsgBySender: function (curUsername, chatTo) {
      // this methon used for this.chatData,
      // it filter chatData by sender
      return this.chatData.filter(function (msg) {
        return (msg.receiver === chatTo && msg.sender === curUsername) ||
            (msg.receiver === curUsername && msg.sender === chatTo)
      })
    },
    getUnreadMsgs: function () {
      var msgUnreadUrl = commondata.APIURL + '/api/v1/message/unread?receiver=' + this.curUsername
      var self = this
      this.$http.get(msgUnreadUrl).then(response => {
        var data = response.data
        if (data.status === 200) {
          self.unreadMsgs = data.data
          console.log('get all unread msgs: ', data.data)
        } else {
          alert('get unread msgs of error: ', data.data)
        }
      }).catch(err => {
        alert(err)
      })
    }
  },
  created: function () {
    if (!auth.checkAuth()) {
      this.$router.push('/login')
    }
    // var this2 = this
    this.curUsername = auth.getCurName()
    // list all friends
    friend.listFriends(this, this.curUsername)
    // get all unread msgs
    this.getUnreadMsgs()
    // set up websocket between client and server
    if (this.ws === null) {
      this.ws = message.createWS(this.curUsername)
    }
    // add message lister
    var self = this
    this.ws.addEventListener('message', function (e) {
      var msg = e.data
      console.log('receive new msg from ws: ', msg)
      self.chatData.push(JSON.parse(msg))
    })
    this.ws.addEventListener('error', function (e) {
      alert(e)
    })
    this.ws.addEventListener('close', function (e) {
      console.log('close ws of user: ', this.curUsername)
    })
  }
}
</script>
<style scope>
.c-left, .c-right {
    height: 600px;
}
.left-add, .left-header, .left-contact {
    height: 30px;
    border-bottom: 1px solid #d6d6d6;
}
.right-chatbox-main {
    height: 460px;
}
.scroll-wrapper {
    overflow-y: auto;
    padding: 0!important;
}
.right-chatbox-main-header {
    height: 30px;
    position: relative;
    border-bottom: 1px solid #d6d6d6;
}
.right-chat-texteditor {
    height: 40px;
    border-top: 1px solid #d6d6d6;
}
</style>
