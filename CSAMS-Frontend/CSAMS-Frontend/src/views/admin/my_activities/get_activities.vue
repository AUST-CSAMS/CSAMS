<template>
  <div>
    <activity_create v-model:visible="visible" @ok="createOk"></activity_create>
    <a-modal v-model:visible="updateVisible" :on-before-ok="updateActivityOk" title="编辑活动">
      <a-form ref="formRef" :model="updateActivityForm">
        <a-form-item :rules="[{required:true,message:'请输入活动名'}]" :validate-trigger="['blur']"
                     field="name"
                     label="活动名"
        >
          <a-input v-model="updateActivityForm.title" placeholder="活动名"></a-input>
        </a-form-item>
        <a-form-item field="role" label="权限">
          <a-select v-model="updateActivityForm.place" placeholder="选择角色"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="userListApi"
      add-label="生成活动"
      default-delete
      no-confirm
      search-placeholder="搜索活动"
      @add="visible=true"
      @edit="edit">
      <template #activity_avatar="{record}:{record: userInfoType}">
        <a-avatar :imageUrl="record.avatar"></a-avatar>
      </template>
    </admin_table>
  </div>
</template>

<script lang="ts" setup>
import Admin_table from "@/components/admin_table.vue";
import {userListApi} from "@/api/user_api";
import type {userInfoType} from "@/api/user_api";
import type {RecordType} from "@/components/admin_table.vue";
import {reactive, ref} from "vue";
import {Message} from "@arco-design/web-vue";
import Activity_create from "@/components/activity_create.vue";
import {activityUpdateApi, type activityUpdateRequest} from "@/api/activity_api";


const columns = [
  {title: '活动名称', dataIndex: 'activity_name'},
  {title: '活动图片', slotName: 'activity_avatar'},
  {title: '联系人', dataIndex: 'name'},
  {title: '联系方式', dataIndex: 'tel'},
  {title: '开始时间', slotName: 'created_at'},
  {title: '结束时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]

const roleOptions = [
  {label: "普通用户", value: 2},
  {label: "管理员", value: 1},
  {label: "教职工", value: 3},
]


const visible = ref(false)
const adminTable = ref()

function createOk() {
  adminTable.value.getList()
}


const updateVisible = ref(false)

const updateActivityForm = reactive<activityUpdateRequest>({
  title: "",
  create_at: "",
  place: "",
  content: ""
})

interface activityInfoType {
  name: string
}

function edit(record: RecordType<activityInfoType>): void {
  updateActivityForm.name = record.name
  updateVisible.value = true
}

const formRef = ref()

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


</script>
