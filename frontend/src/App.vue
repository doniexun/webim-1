<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
import auth from './service/auth'

export default {
  name: 'app',
  created: function () {
    if (auth.checkAuth()) {
      this.$router.push('/home')
    } else {
      this.$router.push('/login')
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
