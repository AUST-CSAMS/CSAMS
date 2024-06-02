<template>
  <div class="login">
    <div class="login_mask">
      <login @ok="ok"></login>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Login from "@/views/login/login.vue";
import router from "@/router";


function ok() {
  let back = window.history.state.back
  if (back === null) {
    router.push({name: "index"})
    return
  }
  router.push(back)
}

</script>

<style lang="scss">
.login {
  background: url(/image/天空.jpg) 50%/cover no-repeat;
  width: 100%;
  height: 100vh;

  .login_mask {
    position: absolute;
    right: 0;
    width: 400px;
    height: 100%;
    background-color: rgba(212, 227, 236, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 40px;
  }
}

</style>
