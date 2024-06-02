<template>
  <div class="admin_user_info_menu">
    <a-dropdown class="admin_dropdown" @select="select">
      <div class="admin_user_info_menu_dropdown">
        <img :src="store.userInfo.avatar" alt="">
        <span class="admin_user_info_menu_dropdown_span">{{ store.userInfo.nick_name }}</span>
        <IconDown></IconDown>
      </div>
      <template #content>
        <template v-for="(item, index) in dopList" :key="index">
          <a-doption :value="item.name">{{ item.title }}</a-doption>
        </template>
      </template>
    </a-dropdown>
  </div>
</template>
<script lang="ts" setup>
import {IconDown} from "@arco-design/web-vue/es/icon";
import {useStore} from "@/stores";
import type {tabType} from "@/types";
import router from "@/router";

interface dopType extends tabType {
  type?: string
}

const store = useStore()

//管理员
let dopList: dopType[] = [
  {
    name: "home",
    title: "管理系统",
  },
  {
    name: "my_association",
    title: "我的协会",
  },
  {
    name: "my_activities",
    title: "我的活动",
  },
  {
    name: "logout",
    title: "注销退出",
  }
]

function select(value: any) {
  let val = value as string
  if (val === "logout") {
    store.logout()
    router.push({name: "index"})
    return
  }
  router.push({
    name: val
  })
}

</script>

<style lang="scss">

.admin_user_info_menu {
  img {
    width: 30px;
    height: 30px;
    border-radius: 50%;
  }

  .admin_user_info_menu_dropdown {
    display: flex;
    cursor: pointer;
    align-items: center;

    .admin_user_info_menu_dropdown_span {
      margin: 0 5px;
    }
  }
}

</style>
