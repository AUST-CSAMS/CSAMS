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
      <div>协会信息</div>
      <a-form ref="formRef" :label-col-props="{span: 5}" :model="form"
              :wrapper-col-props="{span:19}">
        <a-form-item label="协会名">
          <a-input v-model="form.name" @change="associationInfoUpdate"></a-input>
        </a-form-item>
        <a-form-item label="标识">
          <a-avatar :image-url="form.avatar" @click="showCropper"></a-avatar>
        </a-form-item>
        <a-form-item label="宣言">
          <a-textarea v-model="form.introduction" :auto-size="{minRows: 3, maxRows: 3}" placeholder="简介"
                      @change="associationInfoUpdate"></a-textarea>
        </a-form-item>
        <a-form-item label="成立时间">
          <span>{{ form.create_at }}</span>
        </a-form-item>
        <a-form-item label="会长">
          <a-input v-model="form.president" @change="associationInfoUpdate"></a-input>
        </a-form-item>
        <a-form-item label="负责教师">
          <a-input v-model="form.teacher" @change="associationInfoUpdate"></a-input>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import Admin_cropper from "@/components/admin_cropper.vue";
import {
  associationInfoApi,
  type associationInfoType, associationInfoUpdateApi, type associationInfoUpdateType
} from "@/api/association_api";

const formRef = ref()


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


const form = reactive<associationInfoType>({
  name: "",
  avatar: "",
  create_at: "",
  teacher: "",
  president: "",
  introduction: ""
})

function onConfirm(val: string) {
  form.avatar = val
  associationInfoUpdate()
}

async function getData() {
  let res = await associationInfoApi()
  Object.assign(form, res.data)
}

async function associationInfoUpdate() {
  let val = await formRef.value.validate()
  if (val) return

  let data: associationInfoUpdateType = {
    name: form.name,
    introduction: form.introduction,
    teacher: form.teacher,
    president: form.president,
    avatar: form.avatar,
  }
  let res = await associationInfoUpdateApi(data)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
}

getData()


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
