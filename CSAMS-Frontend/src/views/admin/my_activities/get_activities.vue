<template>
  <div v-if="store.isTeacher">
    <activity_create v-model:visible="visible" @ok="createOk"></activity_create>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="activityListTeacherApi"
      add-label="生成活动"
      default-delete
      no-confirm
      no-edit
      no-search
      @add="visible=true">
      <template #image="{record}:{record: activityRequest}">
        <a-avatar :imageUrl="record.image"></a-avatar>
      </template>
      <template #action_middle="{record}:{record: activityRequest}">
        <a-button @click="quit(record.id)">活动截止</a-button>
      </template>

    </admin_table>
  </div>
  <div v-else>
    <a-modal v-model:visible="submitVisible" :on-before-ok="submitassignment">
      <a-form ref="assignmentFormRef" :model="assignmentsubmitForm">
        <a-form-item :rules="[{required:true,message:'请填入作业'}]" :validate-trigger="['blur']"
                     field="id"
        >
          <a-input-number v-model="assignmentsubmitForm.id" placeholder="活动id"></a-input-number>
        </a-form-item>

        <a-form-item :rules="[{required:true,message:'请填入作业'}]" :validate-trigger="['blur']"
                     field="content"
        >
          <a-input v-model="assignmentsubmitForm.content" placeholder="作业内容"></a-input>
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
      no-search>
      <template #activity_avatar="{record}:{record: activityRequest}">
        <a-avatar :imageUrl="record.image"></a-avatar>
      </template>
      <template #image="{record}:{record: activityRequest}">
        <a-avatar :imageUrl="record.image"></a-avatar>
      </template>
      <template #action_middle="">
        <a-button type="outline" @click="submit()">提交作业</a-button>
      </template>
    </admin_table>
  </div>
</template>

<script lang="ts" setup>
import Admin_table from "@/components/admin_table.vue";
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import Activity_create from "@/components/activity_create.vue";
import {
  activityEndApi,
  type activityJoinType,
  activityListApi,
  activityListTeacherApi,
  type activityRequest,
} from "@/api/activity_api";
import {useStore} from "@/stores";
import {
  assignmentSubmitApi, type assignmentSubmitFormType,
  type assignmentSubmitType,
} from "@/api/assignment_api";

const store = useStore()


const assignmentsubmitForm = reactive<assignmentSubmitType>({
  id: 0,
  content: ""
})
let columns = [
  {title: '活动ID', dataIndex: 'id'},
  {title: '活动名称', dataIndex: 'activity_name'},
  {title: '活动图片', slotName: 'image'},
  {title: '联系人', dataIndex: 'responsible_person'},
  {title: '联系方式', dataIndex: 'tel'},
  {title: '开始时间', slotName: 'created_at'},
  {title: '结束时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]


const visible = ref(false)
const adminTable = ref()

function createOk() {
  adminTable.value.getList()
}


const submitVisible = ref(false)


const assignmentFormRef = ref()


function submit() {
  submitVisible.value = true
}

async function quit(id: number) {
  let quitform: activityJoinType = {
    id: id
  }
  let res = await activityEndApi(quitform)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
}

async function submitassignment() {
  let val = await assignmentFormRef.value.validate()
  if (val) return

  let res = await assignmentSubmitApi(assignmentsubmitForm)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
}
</script>
