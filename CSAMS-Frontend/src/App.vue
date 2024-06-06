<template>
  <RouterView/>
</template>


<script lang="ts" setup>
import {RouterView} from 'vue-router'
import {useStore} from "@/stores";

const store = useStore()
store.loadToken()
</script>

