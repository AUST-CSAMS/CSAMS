<template>
  <div class="admin_menu">
    <a-menu
      v-model:open-keys="openKeys"
      v-model:selected-keys="selectedKeys"
      @menu-item-click="clickMenu"
    >
      <template v-for="item in menuList" :key="item.name">
        <a-menu-item v-if="item.child?.length === 0" :key="item.name">
          {{ item.title }}
        </a-menu-item>
        <a-sub-menu v-if="item.child?.length!==0 " :key="item.name">
          <template #title>{{ item.title }}</template>
          <a-menu-item v-for="sub in item.child" :key="sub.name">
            {{ sub.title }}
          </a-menu-item>
        </a-sub-menu>
      </template>
    </a-menu>
  </div>
</template>

<script lang="ts" setup>
import {ref, watch} from "vue";
import {useRoute, useRouter} from "vue-router";
import {useStore} from "@/stores";

const store = useStore()
const route = useRoute()
const router = useRouter()


interface MenuType {
  title: string
  name?: string
  child?: MenuType[]
}

//管理员能看的
let menuList: MenuType[] = [
  {title: "首页", name: "home", child: []},
  {
    title: "个人中心", name: "user_center", child: [
      {title: "我的信息", name: "user_info"},
    ]
  },
  {
    title: "我的协会", name: "my_association", child: [
      {title: "协会信息", name: "association_info"},
      {title: "协会成员", name: "association_members"},
    ]
  },
  {
    title: "我的活动", name: "my_activities", child: [
      {title: "参与活动", name: "get_activities"},
    ]
  },
  {
    title: "活动作业", name: "assignments", child: [
      {title: "作业列表", name: "assignment_list"},
    ]
  },
]

//普通用户能看的
if (store.isGeneral) {
  menuList = [
    {title: "首页", name: "home", child: []},

    {
      title: "个人中心", name: "user_center", child: [
        {title: "我的信息", name: "user_info"},
      ]
    },
    {
      title: "我的协会", name: "my_association", child: [
        {title: "协会信息", name: "association_info"},
        {title: "协会成员", name: "association_members"},
      ]
    },
    {
      title: "我的活动", name: "my_activities", child: [
        {title: "参与活动", name: "get_activities"},
      ]
    },
  ]
}
//展开的子菜单 key 数组 string[]
const openKeys = ref([route.matched[0].name])
//选中的菜单项 key 数组 string[]
const selectedKeys = ref([route.name])

//menu-item-click 点击菜单项时触发
function clickMenu(name: string) {
  router.push({
    name: name,
  })
}

watch(() => route.name, () => {
  selectedKeys.value = [route.name]
  openKeys.value = [route.matched[1].name]
})

</script>


<style lang="scss">
</style>
