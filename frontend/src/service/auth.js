import commondata from './commondata'
export default {
  user: {
    username: '',
    authenticated: false
  },
  validateUser (creds) {
    if (creds.username === '') {
      alert('please input username')
      return false
    }
    if (creds.password === '') {
      alert('please input password')
      return false
    }

    return true
  },
  login (ctx, creds) {
    if (!this.validateUser(creds)) {
      return
    }
    var self = this
    var loginUrl = commondata.APIURL + '/api/v1/user/login'
    console.log('start login...')
    ctx.$http.post(loginUrl, creds)
    .then(function (response) {
      var data = response.data
      if (data.status === 200) {
        sessionStorage.setItem('username', data.data)
        self.user.authenticated = true
        self.user.username = data.data
        ctx.$router.push('/home')
      } else if (data.status === 400) {
        alert(data.data)
        return
      } else {
        alert('unknown error, please again')
        return
      }
    }).catch(function (error) {
      alert(error)
    })
  },
  logout (ctx, creds) {
    var logoutUrl = commondata.APIURL + '/api/v1/user/logout'
    ctx.$http.post(logoutUrl, creds).then(response => {
      var data = response.data
      if (data.status === 200) {
        sessionStorage.removeItem('username')
        this.user.authenticated = false
        this.user.username = ''
        ctx.$router.push('/login')
      }
    })
  },
  checkAuth () {
    var username = sessionStorage.getItem('username')
    return (username !== null) // username is object
  },
  getCurName () {
    this.checkAuth()
    return sessionStorage.getItem('username')
  },
  register (ctx, creds) {
    if (!this.validateUser(creds)) {
      return
    }
    var registerUrl = commondata.APIURL + '/api/v1/user/register'
    ctx.$http.post(registerUrl, creds).then(response => {
      var data = response.data
      if (data.status === 200) {
        ctx.msg = data.data
        ctx.username = ''
        ctx.password = ''
        return
      } else if (data.status === 400) {
        ctx.msg = data.data
        return
      } else {
        ctx.msg = 'unknown error'
        return
      }
    })
  },
  deleteLocalSession (username) {
    sessionStorage.removeItem(username)
  }
}
