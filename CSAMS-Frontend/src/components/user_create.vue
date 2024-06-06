<template>
  <div>
    <a-modal :on-before-ok="createMember" :visible="props.visible" title="创建成员"
             @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item :rules="[{required:true,message:'请输入用户id'}]" :validate-trigger="['blur']"
                     field="id"
                     label="用户id"
        >
          <a-input v-model="form.id" placeholder="id"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {memberCreateApi, type memberCreateFormRequest, type memberCreateRequest} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  id: "",
}

const form = reactive<memberCreateFormRequest>({
  id: ""
})

const formRef = ref()


async function createMember() {
  let val = await formRef.value.validate()
  if (val) return false

  let createform: memberCreateRequest = {
    id: parseInt(form.id),
  }
  console.log(form.id)
  let res = await memberCreateApi(createform)
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
