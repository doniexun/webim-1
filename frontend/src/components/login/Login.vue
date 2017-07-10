<template>
  <div class="login">
    <div class="login-wrapper">
        <div class="login-form">
          <loginInsert v-show="loginStatus === 0" @register="loginStatus = 1"></loginInsert>
          <register v-show="loginStatus === 1" @cancel="loginStatus = 0" @commit="loginStatus = 1"></register>
        </div>
      </div>
  </div>
</template>
<script type="text/javascript">
import register from '@/components/register/Register'
import loginInsert from '@/components/loginInsert/loginInsert'
export default {
  data: function () {
    return {
      loginStatus: 0
    }
  },
  components: {
    register,
    loginInsert
  }
}
</script>
<style scoped>
.login {
  display: block;
  height: 200px;
  justify-content: center;
  align-items: center;
  position: absolute;
  top: 0;
  left: 0;
  bottom: 0;
  width: 100%;
}
</style>
