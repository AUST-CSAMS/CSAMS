<template>
  <Form ref="formRef" :label-col-props="{span: 0}" :model="form" :wrapper-col-props="{span: 24}"
        class="login_form">
    <div class="title">用户登录</div>
    <FormItem :rules="[{required:true,message:'请输入用户名'}]" :validate-trigger="['blur']"
              field="user_name"
              label="用户名"
    >
      <Input v-model="form.user_name" placeholder="用户名">
        <template #prefix>
          <icon-user/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入密码'}]" :validate-trigger="['blur']"
              field="password"
              label="密码"
    >
      <Input v-model="form.password" placeholder="密码" type="password">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <Button type="primary" @click="login">登录</Button>
    <Button type="secondary">
      <router-link to="enroll">注册</router-link>
    </Button>
  </Form>
</template>

<script lang="ts" setup>
import {IconLock, IconUser} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from "vue";
import {Form, FormItem, Input, Button, Message} from "@arco-design/web-vue";
import {loginApi, type loginType} from "@/api/user_api";
import {useStore} from "@/stores";

const emits = defineEmits(["ok"])
const store = useStore()

const form = reactive<loginType>({
  user_name: "",
  password: "",
})

const formRef = ref()

function formReset() {
  formRef.value.resetFields(Object.keys(form))
  formRef.value.clearValidate(Object.keys(form))
}

defineExpose({
  formReset
})

async function login() {
  let val = await formRef.value.validate()
  if (val) {
    return
  }
  let res = await loginApi(form)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  await store.setToken(res.data)
  emits("ok")
}


</script>

<style lang="scss">
.login_form {
  .title {
    font-size: 24px;
    font-weight: 700;
    text-align: center;
    margin-bottom: 10px;
  }

  button {
    margin-top: 10px;
  }
}
</style>
