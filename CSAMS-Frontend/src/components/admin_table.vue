<template>
  <div class="admin_table">
    <div class="admin_table_head">
      <div v-if="!props.noAdd" class="action_create">
        <a-button type="primary" @click="add">{{ addLabel }}</a-button>
      </div>
      <div v-if="!noActionGroup" class="action_group">
        <a-select v-model="actionValue" :options="actionOptions" allow-clear placeholder="操作"
                  style="width: 120px;"></a-select>
        <a-popconfirm v-if="!props.noConfirm" content="是否确认执行此操作?" @ok="actionMethod">
          <a-button v-if="actionValue !== undefined && actionValue !== ''" status="danger" type="primary">执行
          </a-button>
        </a-popconfirm>
        <a-button v-else v-if="actionValue !== undefined && actionValue !== ''" status="danger" type="primary"
                  @click="actionMethod">执行
        </a-button>

      </div>
      <div v-if="!noSearch" class="action_search">
        <a-input-search
          v-model="params.key"
          :placeholder="searchPlaceholder" @search="search"
          @keydown.enter="search"></a-input-search>
      </div>
      <slot name="action_other_search"></slot>
      <slot name="action_slot"></slot>
      <div class="action_flush">
        <a-button @click="flush">
          <IconRefresh></IconRefresh>
        </a-button>
      </div>
    </div>
    <a-spin :loading="isLoading" class="admin_table_data" tip="加载中">
      <div>
        <div class="admin_table_source">
          <a-table v-model:selectedKeys="selectedKeys" :columns="props.columns" :data="data.list"
                   :pagination="false"
                   :row-key=rowKey :row-selection="props.noCheck ? undefined : rowSelection">
            <template #columns>
              <template v-for="item in props.columns">
                <a-table-column v-if="item.render" :title="item.title as string">
                  <template #cell="data">
                    <component :is="item.render(data) as Component"></component>
                  </template>
                </a-table-column>
                <a-table-column v-else-if="!item.slotName" :data-index="item.dataIndex"
                                :title="item.title as string"></a-table-column>
                <a-table-column v-else :title="item.title as string">
                  <template v-if="item.slotName === 'action'" #cell="{record}">
                    <div class="gvd_cell_actions">
                      <slot :record="record" name="action_left"></slot>
                      <a-button v-if="!props.noEdit" type="primary" @click="edit(record)">编辑</a-button>
                      <slot :record="record" name="action_middle"></slot>
                      <a-popconfirm v-if="!props.noDelete" content="是否确认删除?" @ok="remove(record)">
                        <a-button status="danger" type="primary">删除</a-button>
                      </a-popconfirm>
                      <slot :record="record" name="action_right"></slot>
                    </div>
                  </template>
                  <template v-else-if="item.slotName === 'created_at'" #cell="{record}">
                    <span>{{ dateTimeFormat(record.created_at) }}</span>
                  </template>
                  <template v-else #cell="{record}">
                    <slot :name="item.slotName" :record="record"></slot>
                  </template>
                </a-table-column>
              </template>
            </template>
          </a-table>
        </div>
        <div v-if="!props.noPage" class="admin_table_page">
          <a-pagination v-model:current="params.page" :default-page-size="params.limit" :total="data.count"
                        show-jumper show-total @change="pageChange"/>
        </div>
      </div>
    </a-spin>
  </div>
</template>

<script lang="ts" setup>
import {dateTimeFormat} from "@/utils/date";
import {IconRefresh} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from "vue";
import type {Component} from "vue";
import {type baseResponse, defaultDeleteApi, type listDataType} from "@/api";
import type {paramsType} from "@/api";
import type {TableColumnData, TableRowSelection} from "@arco-design/web-vue";
import {Message} from "@arco-design/web-vue";

export interface actionOptionType {
  label: string
  value?: number
  callback?: (idList: (number | string)[]) => Promise<boolean>
}


interface Props {
  url: (params: paramsType) => Promise<baseResponse<listDataType<any>>> // 请求列表数据的api函数
  columns: TableColumnData[] // 字段
  limit?: number // 显示多少条，默认10条
  rowKey?: string // 选中的时候，按照什么选，默认是id
  addLabel?: string // 添加按钮的提示文字
  defaultDelete?: boolean // 是否启用默认删除
  noActionGroup?: boolean// 不启用操作组
  actionGroup?: actionOptionType[], // 操作组
  noCheck?: boolean // 不能选择
  noConfirm?: boolean // 关闭二次确认
  noAdd?: boolean // 没有添加
  noEdit?: boolean // 没有编辑
  noDelete?: boolean // 没有单独的删除
  searchPlaceholder?: string // 模糊匹配的提示词
  defaultParams?: paramsType & any //第一次查询的查询参数
  noPage?: boolean // 不要分页
  noSearch?: boolean //不要搜索
}

const props = defineProps<Props>()

//默认值
const {
  limit = 10,
  rowKey = "id",
  addLabel = "添加",
  searchPlaceholder = "搜索"
} = props


export type RecordType<T> = T

const emits = defineEmits<{
  (e: "add"): void // 添加事件
  (e: "edit", record: RecordType<any>): void// 编辑事件
  (e: "remove", idList: (number | string)[]): void// 删除事件
}>()

const selectedKeys = ref<number[] | string[]>([]);
const rowSelection = reactive<TableRowSelection>({
  type: 'checkbox',
  showCheckedAll: true,
  onlyCurrent: false,
});

const actionOptions = ref<actionOptionType[]>([
  {label: "批量删除", value: 0},
])

//初始化操作组
function initActionGroup() {
  if (!props.actionGroup) return
  for (let i = 0; i < props.actionGroup.length; i++) {
    actionOptions.value.push({
      label: props.actionGroup[i].label,
      value: i + 1,
      callback: props.actionGroup[i].callback
    })
  }
}

initActionGroup()
const actionValue = ref<number | undefined | "">(undefined)

function actionMethod() {
  if (actionValue.value === "") {
    return
  }
  if (actionValue.value === 0) {
    if (selectedKeys.value.length === 0) {
      Message.warning("请选择要删除的记录")
      return
    }
    removeIdData(selectedKeys.value)
    return
  }
  const action = actionOptions.value[actionValue.value as number]
  if (!action.callback) {
    return
  }
  action.callback(selectedKeys.value).then(res => {
    if (res) {
      selectedKeys.value = []
      getList()
      return
    }
  })
}

function add() {
  emits("add")
}

function edit(record: RecordType<any>) {
  emits("edit", record)
}

// 从列表页的api里面匹配路径
const urlRegex = /\.get\("(.*?)",/


// 删除单个
async function remove(record: RecordType<any>) {
  let id = record[rowKey]

  await removeIdData([id])
}

// 批量删除
async function removeIdData(idList: (number | string)[]) {
  if (props.defaultDelete) {
    let regexResult = urlRegex.exec(props.url.toString())
    if (regexResult === null || regexResult.length !== 2) {
      return
    }
    let res = await defaultDeleteApi(regexResult[1], idList)
    if (res.code) {
      Message.error(res.msg)
      return
    }
    Message.success(res.msg)
    selectedKeys.value = []
    getList()
    return;
  }
  emits("remove", idList)
}

const data = reactive<listDataType<any>>({
  list: [],
  count: 0,
})

const params = reactive<paramsType>({
  page: 1,
  limit: limit,
  key: ""
})

const isLoading = ref(false)

//查询
async function getList(p?: paramsType & any) {
  if (p) {
    Object.assign(params, p)
  }
  isLoading.value = true
  let res = await props.url(params)
  console.log(res)
  isLoading.value = false
  if (res.code) {
    Message.error(res.msg)
    return
  }
  data.list = res.data.list
  data.count = res.data.count
  console.log(1)
  console.log(data)
}

function pageChange() {
  getList()
}

// 搜索
function search() {
  params.page = 1
  getList()
}

function flush() {
  Message.success("刷新成功")
  getList()
}

getList(props.defaultParams)

function clearData() {
  data.list = []
  data.count = 0
}


defineExpose({
  getList,
  clearData,
})


</script>

<style lang="scss">
.admin_table {
  background-color: var(--bg1);
  border-radius: 30px;

  .admin_table_head {
    padding: 20px 20px 10px 20px;
    border-bottom: 1px solid var(--bg);
    display: flex;
    align-items: center;
    position: relative;

    > div {
      margin-right: 10px;
    }

    .action_group {
      display: flex;

      button {
        margin-left: 10px;
      }
    }

    .action_filter {
      display: flex;

      > .arco-select {
        margin-right: 10px;
      }
    }

    .action_flush {
      position: absolute;
      right: 20px;
      margin-right: 0;

      button {
        padding: 0 10px;
      }
    }
  }

  .admin_table_data {
    width: 100%;
    padding: 10px 20px 20px 20px;

    .admin_table_source {
      .gvd_cell_actions {
        > button {
          margin-right: 10px;

          &:last-child {
            margin-right: 0;
          }
        }
      }
    }

    .admin_table_page {
      display: flex;
      justify-content: center;
      margin-top: 20px;
    }
  }

}
</style>
