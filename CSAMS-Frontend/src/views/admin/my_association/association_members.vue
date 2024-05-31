<template>
  <div>
    <user_create v-model:visible="visible" @ok="createOk"></user_create>
    <a-modal v-model:visible="updateVisible" :on-before-ok="updateUserOk" title="编辑成员">
      <a-form ref="formRef" :model="updateMemberForm">
        <a-form-item field="role" label="权限">
          <a-select v-model="updateMemberForm.role" :options="roleOptions" placeholder="选择角色"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="userListApi"
      add-label="创建成员"
      default-delete
      no-confirm
      search-placeholder="搜索成员"
      @add="visible=true"
      @edit="edit">
      <template #avatar="{record}:{record: userInfoType}">
        <a-avatar :imageUrl="record.avatar"></a-avatar>
      </template>
    </admin_table>
  </div>
</template>

<script lang="ts" setup>
import Admin_table from "@/components/admin_table.vue";
import {memberUpdateApi, type memberUpdateRequest, userListApi} from "@/api/user_api";
import type {userInfoType} from "@/api/user_api";
import type {RecordType} from "@/components/admin_table.vue";
import {reactive, ref} from "vue";
import User_create from "@/components/user_create.vue";
import {Message} from "@arco-design/web-vue";


const columns = [
  {title: '用户名', dataIndex: 'nick_name'},
  {title: '头像', slotName: 'avatar'},
  {title: '学号', dataIndex: 'number'},
  {title: '角色', dataIndex: 'role'},
  {title: '专业', dataIndex: 'major'},
  {title: '加入时间', slotName: 'created_at'},
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

const updateMemberForm = reactive<memberUpdateRequest>({
  role: 0,
})

function edit(record: RecordType<memberUpdateRequest>): void {
  updateMemberForm.role = record.role
  updateVisible.value = true
}

const formRef = ref()

async function updateUserOk() {
  let val = await formRef.value.validate()
  if (val) return false

  let res = await memberUpdateApi(updateMemberForm)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  await adminTable.value.getList()
  return true
}


</script>
