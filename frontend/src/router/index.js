import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/components/login/Login'
import Home from '@/components/Home'
import addFriend from '@/components/friend/addFriend'
import deleteFriend from '@/components/friend/deleteFriend'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/home',
      name: 'Home',
      component: Home
    }, {
      path: '/login',
      name: 'Login',
      component: Login
    }, {
      path: '/addFriend',
      name: 'addFriend',
      component: addFriend
    }, {
      path: '/deleteFriend',
      name: 'deleteFriend',
      component: deleteFriend
    }
  ]
})
