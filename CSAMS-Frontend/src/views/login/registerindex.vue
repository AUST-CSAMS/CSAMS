<template>
  <div class="register">
    <div class="register_mask">
      <register @ok="ok"></register>
    </div>
  </div>
</template>

<script lang="ts" setup>
import router from "@/router";
import Register from "@/views/login/register.vue";


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
.register {
  background: url(/image/天空.jpg) 50%/cover no-repeat;
  width: 100%;
  height: 100vh;

  .register_mask {
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
