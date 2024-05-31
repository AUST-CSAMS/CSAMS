<template>
  <div class="admin">
    <aside>
      <div class="logo">安徽理工大学协会管理系统
      </div>
      <admin_menu></admin_menu>
    </aside>
    <main>
      <div class="head">
        <bread_crumbs></bread_crumbs>
        <div class="function_area">
          <IconHome class="action_icon" @click="goIndex"></IconHome>
          <user_info_menu></user_info_menu>
        </div>
      </div>
      <div class="container">
        <router-view v-slot="{Component}">
          <transition mode="out-in" name="fade">
            <component :is="Component"></component>
          </transition>
        </router-view>
      </div>
    </main>

  </div>
</template>
<script lang="ts" setup>
import {IconHome} from '@arco-design/web-vue/es/icon';
import {useRouter, useRoute} from "vue-router";
import Admin_menu from "@/components/admin_menu.vue";
import Bread_crumbs from "@/components/bread_crumbs.vue";
import User_info_menu from "@/components/user_info_menu.vue";
import {useStore} from "@/stores";

const store = useStore()
store.loadAssociationInfo()

const router = useRouter()

function goIndex() {
  router.push({
    name: "index"
  })
}

</script>
<style lang="scss">
.admin {
  display: flex;
  height: 100vh;
  overflow-y: hidden;
  background-color: #e9eef1;;

  aside {
    width: 240px;
    border-right: 1px solid;
    height: 100vh;

    .logo {
      height: 90px;
      font-size: 18px;
      display: flex;
      align-items: center;
      border-bottom: 1px solid;
      margin: 10px;
    }
  }


  main {
    width: calc(100% - 240px);
    transition: all .3s;

    .head {
      width: 100%;
      height: 60px;
      border-bottom: 1px solid;
      display: flex;
      justify-content: space-between;
      padding: 0 20px;
      align-items: center;

      .function_area {
        display: flex;
        align-items: center;

        .action_icon {
          margin-right: 10px;
        }
      }
    }


    .container {
      height: calc(100vh - 90px);
      padding: 20px;
      overflow-y: auto;
      overflow-x: hidden;
    }

  }
}

.fade-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

.fade-enter-active {
  transform: translateX(-30px);
  opacity: 0;
}

.fade-enter-to {
  transform: translateX(0px);
  opacity: 1;
}

.fade-leave-active, .fade-enter-active {
  transition: all 0.3s ease-out;
}

</style>
