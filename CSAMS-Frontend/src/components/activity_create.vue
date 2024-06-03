<template>
  <div>
    <a-modal :on-before-ok="createActivity" :visible="props.visible" title="创建活动"
             @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item :rules="[{required:true,message:'请输入活动名'}]" :validate-trigger="['blur']"
                     field="user_name"
                     label="活动名"
        >
          <a-input v-model="form.title" placeholder="活动名"></a-input>
        </a-form-item>
        <a-form-item field="time" label="活动时间">
          <a-input v-model="form.create_at" placeholder="活动时间"></a-input>
        </a-form-item>
        <a-form-item field="place" label="地点">
          <a-input v-model="form.place" placeholder="活动地点"></a-input>
        </a-form-item>
        <a-form-item field="content" label="内容">
          <a-input v-model="form.content" placeholder="活动内容"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {activityCreateApi, type activityRequest} from "@/api/activity_api";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  id: 0,
  name: "",
  create_at: "",
  place: "",
  content: ""
}

const form = reactive<activityRequest>({
  id: 0,
  title: "",
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
