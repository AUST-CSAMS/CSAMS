<template>
  <div v-if="store.isGeneral">
    <a-modal v-model:visible="submitVisible" :on-before-ok="submitassignment">
      <a-form ref="assignmentFormRef" :model="assignmentForm">
        <a-form-item :rules="[{required:true,message:'请填入作业'}]" :validate-trigger="['blur']"
        >
          <a-input v-model="assignmentForm.content" placeholder="作业内容"></a-input>
        </a-form-item>
      </a-form>

    </a-modal>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="activityListApi"
      no-add
      no-confirm
      no-delete
      no-edit
      search-placeholder="搜索活动">
      <template #activity_avatar="{record}:{record: activityRequest}">
        <a-avatar :imageUrl="record.avatar"></a-avatar>
      </template>
      <template #action_middle="">
        <a-button type="outline" @click="submit()">提交作业</a-button>
      </template>
    </admin_table>
  </div>
  <div v-else>
    <activity_create v-model:visible="visible" @ok="createOk"></activity_create>
    <a-modal v-model:visible="updateVisible" :on-before-ok="updateActivityOk" title="编辑活动">
      <a-form ref="formRef" :model="updateActivityForm">
        <a-form-item :rules="[{required:true,message:'请输入活动名'}]" :validate-trigger="['blur']"
                     field="name"
                     label="活动名"
        >
          <a-input v-model="updateActivityForm.title" placeholder="活动名"></a-input>
        </a-form-item>
        <a-form-item field="time" label="活动时间">
          <a-input v-model="updateActivityForm.create_at" placeholder="活动时间"></a-input>
        </a-form-item>
        <a-form-item field="time" label="活动地点">
          <a-input v-model="updateActivityForm.place" placeholder="活动地点"></a-input>
        </a-form-item>
        <a-form-item field="time" label="活动内容">
          <a-input v-model="updateActivityForm.content" placeholder="活动内容"></a-input>
        </a-form-item>
      </a-form>
    </a-modal>
    <a-modal v-model:visible="checkVisible" :on-before-ok="checkassignment">
      <div v-for="item in data.list">
        <div>{{ item.id }}</div>
        <div>{{ item.content }}</div>
      </div>
    </a-modal>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="activityListApi"
      add-label="生成活动"
      default-delete
      no-confirm
      search-placeholder="搜索活动"
      @add="visible=true"
      @edit="edit">
      <template #activity_avatar="{record}:{record: userInfoType}">
        <a-avatar :imageUrl="record.avatar"></a-avatar>
      </template>
      <template #action_middle="">
        <a-button type="outline" @click="check()">查看作业</a-button>
      </template>
    </admin_table>
  </div>
</template>

<script lang="ts" setup>
import Admin_table from "@/components/admin_table.vue";
import type {userInfoType} from "@/api/user_api";
import type {RecordType} from "@/components/admin_table.vue";
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import Activity_create from "@/components/activity_create.vue";
import {activityListApi, type activityRequest, activityUpdateApi, type activityUpdateRequest} from "@/api/activity_api";
import {useStore} from "@/stores";
import {type assignmentInfoType, assignmentListApi, assignmentSubmitApi} from "@/api/assignment_api";
import type {listDataType} from "@/api";

const store = useStore()

const data = reactive<listDataType<assignmentInfoType>>({
  list: [],
  count: 0
})

const assignmentForm = reactive<assignmentInfoType>({
  id: store.userInfo.user_id,
  content: ""
})
let columns = [
  {title: '活动名称', dataIndex: 'activity_name'},
  {title: '活动图片', slotName: 'activity_avatar'},
  {title: '联系人', dataIndex: 'name'},
  {title: '联系方式', dataIndex: 'tel'},
  {title: '开始时间', slotName: 'created_at'},
  {title: '结束时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]


const visible = ref(false)
const adminTable = ref(false)

function createOk() {
  adminTable.value.getList()
}

const updateVisible = ref(false)

const submitVisible = ref(false)
const checkVisible = ref(false)


const updateActivityForm = reactive<activityUpdateRequest>({
  id: 0,
  title: "",
  create_at: "",
  place: "",
  content: ""
})

interface activityInfoType {
  name: string
}

function edit(record: RecordType<activityInfoType>): void {
  updateActivityForm.title = record.name
  updateVisible.value = true
}

const formRef = ref()
const assignmentFormRef = ref()

async function updateActivityOk() {
  let val = await formRef.value.validate()
  if (val) return false

  let res = await activityUpdateApi(updateActivityForm)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  await adminTable.value.getList()
  return true
}

function submit() {
  submitVisible.value = true
}

async function submitassignment() {
  let val = await assignmentFormRef.value.validate()
  if (val) return
  let data: assignmentInfoType = {
    id: assignmentForm.id,
    content: assignmentForm.content
  }
  let res = await assignmentSubmitApi(data)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
}

function check() {
  checkVisible.value = true
}

async function checkassignment() {
  let res = await assignmentListApi()
  data.list = res.data.list
  data.count = res.data.count
}
</script>
