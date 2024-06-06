<template>
  <div>
    <a-modal :on-before-ok="createActivity" :visible="props.visible" title="创建活动"
             @cancel="emits('update:visible', false)">
      <a-form ref="formRef" :model="form">
        <a-form-item field="id" label="活动id">
          <a-input v-model="form.id" placeholder="活动id"></a-input>
        </a-form-item>
        <a-form-item field="activity_name" label="活动名">
          <a-input v-model="form.activity_name" placeholder="活动名"></a-input>
        </a-form-item>
        <a-form-item field="startTime" label="开始时间">
          <a-input v-model="form.startTime" placeholder="开始时间"></a-input>
        </a-form-item>
        <a-form-item field="endTime" label="结束时间">
          <a-input v-model="form.endTime" placeholder="结束时间"></a-input>
        </a-form-item>
        <a-form-item field="location" label="活动地点">
          <a-input v-model="form.location" placeholder="活动地点"></a-input>
        </a-form-item>
        <a-form-item field="introduction" label="活动介绍">
          <a-input v-model="form.introduction" placeholder="活动介绍"></a-input>
        </a-form-item>
        <a-form-item label="活动图片">
          <a-upload
            v-model:file-list="fileList"
            :headers="{token: store.userInfo.token}"
            action="/api/image"
            image-preview
            list-type="picture-card"
            name="image"
            @success="imageUploadSuccessEvent"
          ></a-upload>
        </a-form-item>
        <a-form-item field="score" label="活动分数">
          <a-input v-model="form.score" placeholder="活动分数"></a-input>
        </a-form-item>
        <a-form-item field="limit" label="活动限制">
          <a-input v-model="form.limit" placeholder="活动限制"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
  </div>
</template>
<script lang="ts" setup>
import {reactive, ref} from "vue";
import {type FileItem, Message} from "@arco-design/web-vue";
import {
  activityCreateApi,
  type activityCreateFormRequest, type activityCreateRequest,
} from "@/api/activity_api";
import {useStore} from "@/stores";

const props = defineProps({
  visible: {
    type: Boolean,
  }
})

const store = useStore()


const fileList = ref<FileItem[]>([])


const emits = defineEmits(["update:visible", "ok"])

const defaultForm = {
  id: 0,
  activity_name: "",
  startTime: "",
  endTime: "",
  location: "",
  introduction: "",
  image: "",
  responsible_person: "",
  tel: 0,
  score: 0,
  limit: []
}

const form = reactive<activityCreateFormRequest>({
  id: "",
  activity_name: "",
  startTime: "",
  endTime: "",
  location: "",
  introduction: "",
  image: "",
  score: "",
  limit: ""
})

const formRef = ref()


async function createActivity() {
  let val = form.limit.split(" ")
  let createform: activityCreateRequest = {
    id: parseInt(form.id),
    activity_name: form.activity_name,
    startTime: form.startTime,
    endTime: form.endTime,
    location: form.location,
    introduction: form.introduction,
    image: form.image,
    score: parseInt(form.score),
    limit: val
  }
  console.log(createform)
  let res = await activityCreateApi(createform)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  emits("update:visible", false)
  emits("ok")
  Object.assign(form, defaultForm)
}

function imageUploadSuccessEvent(fileItem: FileItem) {
  if (fileItem.response.msg != '成功') {
    Message.error(fileItem.response.msg)
    return
  } else {
    Message.success(fileItem.response.msg)
    form.image = fileItem.response.data
  }
}

</script>
<style lang="scss" scoped>

</style>
