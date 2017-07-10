import commondata from './commondata'
export default {
  createWS (username) {
    if (window.WebSocket === undefined) {
      alert('your browser does not support websocket!')
      return
    }
    var userWSURL = commondata.WSURL + '/api/v1/message/ws/' + username
    return new WebSocket(userWSURL)
  },
  getTimestamp () {
    return Math.floor(new Date().getTime() / 1000)
  },
  getUnreadMsgs (ctx, receiver) {
    var msgUnreadUrl = commondata.APIURL + '/api/v1/message/unread?receiver=' + receiver
    ctx.$http.get(msgUnreadUrl).then(response => {
      var data = response.data
      console.log('request unread msgs: ', data)
      if (data.status === 200) {
        ctx.unreadMsgs = data.data
        ctx.test = 'set done'
      } else {
        alert('get unread msgs of ', receiver, ' error: ', data.data)
      }
    }).catch(err => {
      alert(err)
      return
    })
  }
}
