<template>
  <div>
    <a-modal :on-before-ok="createActivity" :visible="props.visible" title="创建活动"
             @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item :rules="[{required:true,message:'请输入活动名'}]" :validate-trigger="['blur']"
                     field="user_name"
                     label="活动名"
        >
          <a-input v-model="form.name" placeholder="活动名"></a-input>
        </a-form-item>
        <a-form-item field="time" label="活动时间">
          <a-select v-model="form.content" placeholder="选择角色"></a-select>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="form.content" placeholder="选择角色"></a-select>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="form.content" placeholder="选择角色"></a-select>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="form.content" placeholder="选择角色"></a-select>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="form.content" placeholder="选择角色"></a-select>
        </a-form-item>

      </a-form>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {activityCreateApi, type activityCreateRequest} from "@/api/activity_api";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  name: "",
  create_at: "",
  place: "",
  content: ""
}

const form = reactive<activityCreateRequest>({
  name: "",
  create_at: "",
  place: "",
  content: ""
})

const formRef = ref()


async function createActivity() {
  let val = await formRef.value.validate()
  if (val) return false

  let res = await activityCreateApi(form)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  emits("update:visible", false)
  emits("ok")
  Object.assign(form, defaultForm)
}


</script>
<style lang="scss" scoped>

</style>
