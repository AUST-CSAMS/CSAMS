<template>
  <div class="admin_bread_crumbs">
    <a-breadcrumb>
      <!-- route.matched 包含当前活动路由及其父路由的信息 -->
      <a-breadcrumb-item v-for="item in route.matched">{{ item.meta.title }}</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
</template>

<script lang="ts" setup>
import {useRoute} from "vue-router";

const route = useRoute()
</script>
<style lang="scss">
</style>
