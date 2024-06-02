<template>
  <div class="enroll">
    <div class="enroll_mask">
      <Form ref="formRef" :label-col-props="{span: 0}" :model="form" :wrapper-col-props="{span: 24}"
            class="login_form">
        <div class="title">用户注册</div>
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
        <Button type="primary" @click="enroll" @ok="ok">注册</Button>
      </Form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {IconLock, IconUser} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from "vue";
import {Form, FormItem, Input, Button, Message} from "@arco-design/web-vue";
import {enrollApi, type loginType} from "@/api/user_api";
import router from "@/router";

const emits = defineEmits(["ok"])

const form = reactive<loginType>({
  user_name: "",
  password: "",
})

const formRef = ref()

async function enroll() {
  let val = await formRef.value.validate()
  if (val) {
    return
  }
  let res = await enrollApi(form)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  emits("ok")
}

function ok() {
  let back = window.history.state.back
  if (back === null) {
    router.push({name: "index"})
    return
  }
  router.push(back)
}


</script>

<style lang="scss">
.enroll {
  background: url(/image/天空.jpg) 50%/cover no-repeat;
  width: 100%;
  height: 100vh;

  .enroll_mask {
    position: absolute;
    right: 0;
    width: 400px;
    height: 100%;
    background-color: rgba(212, 227, 236, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    padding: 0 40px;

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

  }
}
</style>
