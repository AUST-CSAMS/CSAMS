<template>
  <div v-if="store.isTeacher">
    <user_create v-model:visible="visible" @ok="createOk"></user_create>
    <a-modal v-model:visible="updateVisible" :on-before-ok="updateUserOk" title="编辑成员">
      <a-form ref="formRef" :model="updateMemberForm">
        <a-form-item field="posts" label="职位">
          <a-select v-model="updateMemberForm.posts" :options="roleOptions" placeholder="职位"></a-select>
        </a-form-item>
      </a-form>
    </a-modal>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="associationMemberListApi"
      add-label="创建成员"
      no-confirm
      no-delete
      no-edit
      no-search
      @add="visible=true">
      <template #action_middle="{record}:{record: associationMemberType}">
        <a-button type="primary" @click="editmember(record.user_id)">编辑</a-button>
      </template>
      <template #action_right="{record}:{record: associationMemberType}">
        <a-button type="secondary" @click="deletemember(record.user_id)">删除</a-button>
      </template>

    </admin_table>
  </div>
  <div v-else>
    <admin_table
      ref="adminTable"
      :columns="columns"
      :url="associationMemberListApi"
      no-add
      no-confirm
      no-delete
      no-edit
      no-search
      @add="visible=true">
    </admin_table>
  </div>
</template>

<script lang="ts" setup>
import Admin_table from "@/components/admin_table.vue";
import {memberUpdateApi, type memberUpdateRequest} from "@/api/user_api";
import type {RecordType} from "@/components/admin_table.vue";
import {reactive, ref} from "vue";
import User_create from "@/components/user_create.vue";
import {Message} from "@arco-design/web-vue";
import {useStore} from "@/stores";
import {
  type associationDeleteType,
  associationMemberDeleteApi,
  associationMemberListApi,
  type associationMemberType
} from "@/api/association_api";

const store = useStore()

const roleOptions = [
  {label: "会长", value: "会长"},
  {label: "成员", value: "成员"},
]

let columns = [
  {title: 'id', dataIndex: 'user_id'},
  {title: '职位', dataIndex: 'posts'},
  {title: '加入时间', slotName: 'created_at'},
  {title: '操作', slotName: 'action'},
]

if (store.isGeneral) {
  columns = [
    {title: 'id', dataIndex: 'user_id'},
    {title: '职位', dataIndex: 'posts'},
    {title: '加入时间', slotName: 'created_at'},
  ]
}


const visible = ref(false)
const adminTable = ref()

function createOk() {
  adminTable.value.getList()
}


const updateVisible = ref(false)

const updateMemberForm = reactive<memberUpdateRequest>({
  id: 0,
  posts: "",
})

function editmember(id: number): void {
  updateMemberForm.id = id
  updateVisible.value = true
}

const formRef = ref()

async function updateUserOk() {
  let res = await memberUpdateApi(updateMemberForm)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  await adminTable.value.getList()
  return true
}

async function deletemember(id: number) {
  let deleteform: associationDeleteType = {
    id: [id]
  }
  console.log(deleteform)
  let res = await associationMemberDeleteApi(deleteform)
  if (res.code) {
    Message.error(res.msg)
    return
  }
  Message.success(res.msg)
  await adminTable.value.getList()
  return true
}

</script>
