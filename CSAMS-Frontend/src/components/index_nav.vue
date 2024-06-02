<template>
  <div class="nav">
    <div class="nav-container">
      <div class="nav_left">
        <div class="logo">安徽理工大学协会管理系统</div>
        <div class="nav_item">首页</div>
        <index_search></index_search>
      </div>
      <div class="nav_right">
        <div v-if="!store.isLogin" class="login">
          <router-link :to="{name: 'login'}">登录</router-link>
        </div>
        <div v-else>
          <user_info_menu></user_info_menu>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {useStore} from "@/stores";
import User_info_menu from "@/components/user_info_menu.vue";
import Index_search from "@/components/index_search.vue";

const store = useStore();
</script>

<style lang="scss">
.nav {
  width: 100%;
  top: 0;
  display: flex;
  justify-content: center;
  height: 60px;
  background-color: rgba(212, 227, 236, 0.7);

  .nav-container {
    width: 1200px;
    display: flex;
    align-items: center;


    .nav_left {
      width: 70%;
      font-size: 16px;
      display: flex;
      align-items: center;

      .nav_item {
        margin-right: 20px;
      }

      .logo {
        width: 100px;
        margin-right: 30px;
      }
    }

    .nav_right {
      width: 30%;
      font-size: 16px;
      display: flex;
      align-items: center;
      justify-content: end;
    }
  }
}
</style>
