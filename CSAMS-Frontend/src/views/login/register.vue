<template>

  <Form ref="formRef" :label-col-props="{span: 0}" :model="form" :wrapper-col-props="{span: 24}"
        class="register_form">
    <div class="title">用户注册</div>
    <FormItem :rules="[{required:true,message:'请输入id'}]" :validate-trigger="['blur']"
              field="id"
              label="id"
    >
      <Input v-model="form.id" placeholder="id">
        <template #prefix>
          <icon-user/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入密码'}]" :validate-trigger="['blur']"
              field="password"
              label="密码"
    >
      <Input v-model="form.password" placeholder="密码">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入姓名'}]" :validate-trigger="['blur']"
              field="name"
              label="姓名"
    >
      <Input v-model="form.name" placeholder="姓名">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入年龄'}]" :validate-trigger="['blur']"
              field="age"
              label="年龄"
    >
      <Input v-model="form.age" placeholder="年龄">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入性别'}]" :validate-trigger="['blur']"
              field="gender"
              label="性别"
    >
      <Input v-model="form.gender" placeholder="性别">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入身份'}]" :validate-trigger="['blur']"
              field="role"
              label="身份"
    >
      <Input v-model="form.role" placeholder="身份">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入专业'}]" :validate-trigger="['blur']"
              field="password"
              label="专业"
    >
      <Input v-model="form.major" placeholder="专业">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <FormItem :rules="[{required:true,message:'请输入电话'}]" :validate-trigger="['blur']"
              field="tel"
              label="电话"
    >
      <Input v-model="form.tel" placeholder="电话">
        <template #prefix>
          <icon-lock/>
        </template>
      </Input>
    </FormItem>
    <Button type="primary" @click="register">注册</Button>
  </Form>

</template>

<script lang="ts" setup>
import {IconLock, IconUser} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from "vue";
import {Form, FormItem, Input, Button, Message} from "@arco-design/web-vue";
import {registerApi, type registerFormType, type registerType} from "@/api/user_api";

const emits = defineEmits(["ok"])

const form = reactive<registerType>({
  id: "",
  password: "",
  name: "",
  age: "",
  gender: "",
  role: "",
  major: "",
  tel: ""
})

const formRef = ref()

async function register() {
  let val = await formRef.value.validate()
  if (val) {
    return
  }
  let registerform: registerFormType = {
    id: parseInt(form.id),
    password: form.password,
    name: form.name,
    age: parseInt(form.age),
    gender: form.gender,
    role: form.role,
    major: form.major,
    tel: parseInt(form.tel)
  }
  let res = await registerApi(registerform)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  emits("ok")
}

</script>

<style lang="scss">
.register_form {
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
