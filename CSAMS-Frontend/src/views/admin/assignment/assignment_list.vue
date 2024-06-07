<template>
  <div>
    <a-modal v-model:visible="correctVisible" :on-before-ok="correctOK" title="批改作业">
      <a-form ref="assignmentcorrectFormRef" :model="assignmentcorrectForm">
        <a-form-item :rules="[{required:true}]" :validate-trigger="['blur']"
        >
          <a-input-number v-model="assignmentcorrectForm.id"></a-input-number>
        </a-form-item>
        <a-form-item :rules="[{required:true}]" :validate-trigger="['blur']"
        >
          <a-select v-model="assignmentcorrectForm.is_finish" :options="roleOptions" placeholder="是否完成"></a-select>
        </a-form-item>
      </a-form>

    </a-modal>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="assignmentListApi"
      default-delete
      no-add
      no-confirm
      no-search
      @edit="edit">
    </admin_table>
  </div>
</template>

<script lang="ts" setup>
import Admin_table from "@/components/admin_table.vue";
import {reactive, ref} from "vue";
import {
  type activityCorrectType, assignmentCorrectApi,
  assignmentListApi,
} from "@/api/assignment_api";
import {Message} from "@arco-design/web-vue";


const roleOptions = [
  {label: "完成", value: true},
  {label: "未完成", value: false},
]

const correctVisible = ref(false)

const assignmentcorrectForm = reactive<activityCorrectType>({
  id: 0,
  is_finish: false
})

const assignmentcorrectFormRef = ref()

let columns = [
  {title: 'ID', dataIndex: 'id'},
  {title: '活动ID', dataIndex: 'activity_id'},
  {title: '内容', dataIndex: 'content'},
  {title: '用户ID', dataIndex: 'user_id'},
  {title: '是否提交', dataIndex: 'is_submit'},
  {title: '是否完成', dataIndex: 'is_finish'},
  {title: '是否批改', dataIndex: 'is_correct'},
  {title: '操作', slotName: 'action'},
]


const adminTable = ref()

async function correctOK() {
  let val = await assignmentcorrectFormRef.value.validate()
  if (val) return
  let res = await assignmentCorrectApi(assignmentcorrectForm)
  console.log(res)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
}

function edit(): void {
  correctVisible.value = true
}


</script>
