import commondata from './commondata'
export default {
  addFriend (ctx, f1, f2) {
    var addFriendUrl = commondata.APIURL + '/api/v1/friend/add'
    var creds = this.formatFriend(f1, f2)
    ctx.$http.post(addFriendUrl, creds).then(response => {
      var data = response.data
      if (data.status === 200) {
        ctx.FriAlert(0, 'add friend success', 'success')
      } else if (data.status === 400) {
        ctx.FriAlert(0, data.data, 'error')
      } else {
        ctx.FriAlert(0, 'unknown error', 'error')
      }
    })
  },
  formatFriend (f1, f2) {
    var fmin, fmax
    if (f1 <= f2) {
      fmin = f1
      fmax = f2
    } else {
      fmin = f2
      fmax = f1
    }
    var creds = {
      fmin: fmin,
      fmax: fmax
    }
    return creds
  },
  listFriends (ctx, curUsername) {
    var listFriendUrl = commondata.APIURL + '/api/v1/friend/list?username=' + curUsername
    ctx.$http.get(listFriendUrl).then(response => {
      var data = response.data
      if (data.status === 200) {
        ctx.contacts = data.data
      } else {
        alert('get contacts error')
      }
    })
  },
  deleteFriend (ctx, curUsername, fname) {
    var deleteFriendUrl = commondata.APIURL + '/api/v1/friend/delete'
    var creds = this.formatFriend(curUsername, fname)
    ctx.$http.put(deleteFriendUrl, creds).then(response => {
      var data = response.data
      if (data.status === 200) {
        alert(data.data)
      } else {
        alert('dlete error , try again.')
      }
    })
  }
}
