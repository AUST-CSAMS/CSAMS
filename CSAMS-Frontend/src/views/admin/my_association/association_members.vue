<template>
  <div v-if="store.isTeacher">
    <user_create v-model:visible="visible" @ok="createOk"></user_create>
    <a-modal v-model:visible="updateVisible" :on-before-ok="updateUserOk" title="编辑成员">
      <a-form ref="formRef" :model="updateMemberForm">
        <a-form-item field="id" label="id">
          <a-input-number v-model="updateMemberForm.id" placeholder="id"></a-input-number>
        </a-form-item>
        <a-form-item field="posts" label="职位">
          <a-input v-model="updateMemberForm.posts" placeholder="职位"></a-input>
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
      no-search
      @add="visible=true"
      @edit="edit">
      <template #action_middle="{record}:{record: associationMemberType}">
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
import type {activityJoinType, activityRequest} from "@/api/activity_api";

const store = useStore()


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

function edit(record: RecordType<memberUpdateRequest>): void {
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
