<template>
  <div class="user_info_view">
    <div class="left">
      <div>协会信息</div>
      <a-form ref="formRef" :label-col-props="{span: 5}" :model="form"
              :wrapper-col-props="{span:19}">
        <a-form-item label="协会名">
          <a-input v-model="form.association_name" @change="associationInfoUpdate"></a-input>
        </a-form-item>
        <a-form-item label="宣言">
          <a-textarea v-model="form.introduction" :auto-size="{minRows: 3, maxRows: 3}" placeholder="简介"
                      @change="associationInfoUpdate"></a-textarea>
        </a-form-item>
        <a-form-item label="成立时间">
          <span>{{ dateTimeFormat(form.create_at) }}</span>
        </a-form-item>
        <a-form-item label="会长">
          <a-input v-model="form.president" @change="associationInfoUpdate"></a-input>
        </a-form-item>
        <a-form-item label="负责教师">
          <a-input v-model="form.teacher_name" @change="associationInfoUpdate"></a-input>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import {
  associationInfoApi,
  type associationInfoType, associationInfoUpdateApi, type associationInfoUpdateType
} from "@/api/association_api";
import {dateTimeFormat} from "@/utils/date";

const formRef = ref()

const form = reactive<associationInfoType>({
  id: 0,
  association_name: "",
  create_at: "",
  teacher_name: "",
  president: "",
  introduction: ""
})


async function getData() {
  let res = await associationInfoApi()
  console.log(res)
  Object.assign(form, res.data)
  console.log(form)

}

async function associationInfoUpdate() {
  let val = await formRef.value.validate()
  if (val) return

  let data: associationInfoUpdateType = {
    name: form.association_name,
    introduction: form.introduction,
    teacher: form.teacher_name,
    president: form.president,
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
