<template>
  <div class="user_info_view">
    <div class="left">
      <admin_cropper ref="clipperRef"
                     :allow-type-list="clipperData.allowTypeList"
                     :fixed-number="clipperData.fixedNumber"
                     :fixed-number-aider="clipperData.fixedNumberAider"
                     :limit-size="clipperData.limitSize"
                     :preview-width="clipperData.previewWidth"
                     :type="clipperData.type"
                     @confirm="onConfirm">

      </admin_cropper>
      <div>用户信息</div>
      <a-form ref="formRef" :label-col-props="{span: 5}" :model="form"
              :wrapper-col-props="{span:19}">
        <a-form-item label="姓名">
          <span>{{ form.name }}</span>
        </a-form-item>
        <a-form-item label="头像">
          <a-avatar :image-url="form.avatar" @click="showCropper"></a-avatar>
        </a-form-item>
        <a-form-item label="学工号">
          <span>{{ form.id }}</span>
        </a-form-item>
        <a-form-item label="身份">
          <span>{{ form.role }}</span>
        </a-form-item>
        <a-form-item label="专业">
          <span>{{ form.major }}</span>
        </a-form-item>
        <a-form-item label="积分">
          <span>{{ form.score }}</span>
        </a-form-item>
        <a-form-item label="诚信度">
          <span>{{ form.integrity }}</span>
        </a-form-item>
      </a-form>
      <div>操作</div>
      <div class="action_group">
        <a-button type="primary" @click="updatePasswordVisible=true">修改密码</a-button>
      </div>
      <admin_update_password v-model:visible="updatePasswordVisible"></admin_update_password>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import type {userInfoType} from "@/api/user_api";
import {userInfoApi, userInfoUpdateApi} from "@/api/user_api";
import type {userInfoUpdateType} from "@/api/user_api";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";
import Admin_cropper from "@/components/admin_cropper.vue";
import Admin_update_password from "@/components/admin_update_password.vue";

const formRef = ref()
const store = useStore()


interface IClipper {
  type: string // 上传类型
  allowTypeList: string[] // 接收允许上传的图片类型
  limitSize: number // 限制大小
  fixedNumber: number[] // 截图框的宽高比例
  fixedNumberAider?: number[] // 侧边栏收起截图框的宽高比例
  previewWidth: number // 预览宽度
  previewWidthAider?: number // 侧边栏收起预览宽度
}

const clipperData = ref<IClipper>({
  type: '',
  allowTypeList: [],
  limitSize: 1,
  fixedNumber: [],
  previewWidth: 0
})

const clipperRef = ref()

function showCropper() {
  clipperData.value = {
    type: 'browserLogo', // 该参数可根据实际要求修改类型
    allowTypeList: ['png', 'jpg', 'jpeg',], // 允许上传的图片格式
    limitSize: 1, // 限制的大小
    fixedNumber: [1, 1],  // 截图比例，可根据实际情况进行修改
    previewWidth: 100 // 预览宽度
  }
  // 打开裁剪组件
  clipperRef.value.uploadFile("")
}


const form = reactive<userInfoType>({
  id: 0,
  password: "",
  name: "",
  age: 0,
  gender: "",
  major: 0,
  avatar: "",
  tel: 0,
  role: "",
  integrity: 100,
  score: 0,
  associationID: 0
})

async function getData() {
  let res = await userInfoApi()
  Object.assign(form, res.data)
}

getData()

async function userInfoUpdate() {
  let data: userInfoUpdateType = {
    avatar: form.avatar,
  }
  let res = await userInfoUpdateApi(data)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  // 同步到store里面
  store.setUserInfo("avatar", form.avatar)
}


function onConfirm(val: string) {
  form.avatar = val
  userInfoUpdate()
}

const updatePasswordVisible = ref(false)

</script>

<style lang="scss">
.user_info_view {
  display: flex;
  background-color: var(--bg1);
  padding: 20px;
  border-radius: 30px;
  height: calc(100vh - 130px);

  .left {
    width: 40%;

    .arco-form {
      margin-top: 10px;
    }

    .action_group {
      margin-top: 10px;

      > button {
        margin-left: 50px;
      }
    }
  }

}
</style>
