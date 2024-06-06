<template>
  <div>
    <a-modal :on-before-ok="createAssociation" :visible="props.visible" title="创建协会"
             @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item field="id" label="协会id">
          <a-input v-model="form.id" placeholder="协会id"></a-input>
        </a-form-item>
        <a-form-item :rules="[{required:true,message:'请输入协会名'}]" :validate-trigger="['blur']"
                     field="association_name"
                     label="协会名"
        >
          <a-input v-model="form.association_name" placeholder="协会名"></a-input>
        </a-form-item>
        <a-form-item field="introduction" label="宣言">
          <a-input v-model="form.introduction" placeholder="宣言"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {associationCreateApi, type associationCreateFormType, type associationCreateType} from "@/api/association_api";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  id: "",
  association_name: "",
  introduction: "",
}

const form = reactive<associationCreateFormType>({
  id: "",
  association_name: "",
  introduction: ""
})

const formRef = ref()


async function createAssociation() {
  let val = await formRef.value.validate()
  if (val) return false
  let createform: associationCreateType = {
    id: parseInt(form.id),
    association_name: form.association_name,
    introduction: form.introduction
  }
  let res = await associationCreateApi(createform)
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
