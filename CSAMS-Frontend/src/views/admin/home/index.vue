<template>
  <association_create v-model:visible="visible"></association_create>
  <Button v-if="store.isTeacher" @click="associationcreate">创建协会</Button>
</template>

<script lang="ts" setup>
import Association_create from "@/components/association_create.vue";
import {ref} from "vue";
import {useStore} from "@/stores";

const store = useStore()

let visible = ref(false)

function associationcreate() {
  visible.value = true;
}

</script>

<style lang="scss">

</style>
