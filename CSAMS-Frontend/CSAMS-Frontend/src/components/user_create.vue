<template>
  <div>
    <a-modal :on-before-ok="createMember" :visible="props.visible" title="创建成员"
             @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item :rules="[{required:true,message:'请输入用户id'}]" :validate-trigger="['blur']"
                     field="user_name"
                     label="用户id"
        >
          <a-input v-model="form.id" placeholder="id"></a-input>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="form.role" :options="roleOptions" placeholder="选择角色"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {memberCreateApi, type memberCreateRequest} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  id: "",
  role: 1
}

const form = reactive<memberCreateRequest>({
  role: 1,
  id: ""
})

const formRef = ref()

const roleOptions = [
  {label: "普通用户", value: 1},
  {label: "管理员", value: 2},
  {label: "教师", value: 3},
]


async function createMember() {
  let val = await formRef.value.validate()
  if (val) return false

  let res = await memberCreateApi(form)
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
