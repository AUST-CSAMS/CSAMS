<template>
  <a-modal v-model:visible="visible" :footer="false" body-class="index_search_body" title="搜索" width="40%">
    <div class="head">
      <a-input-search ref="searchRef" v-model="params.key" placeholder="活动搜索" search-button
                      @search="search"
                      @keydown.enter="search"></a-input-search>
    </div>
    <template v-if="data.list.length">
      <div class="body">
        <div v-for="item in data.list" class="item" @click="goTo(item)">
          <div class="content">{{ item.title }}</div>
        </div>
      </div>
      <div class="footer">
        共搜索到结果{{ data.count }}条
      </div>
    </template>
  </a-modal>
  <div class="index_search">
    <IconSearch @click="showSearch"></IconSearch>
  </div>
</template>

<script lang="ts" setup>
import {IconSearch} from "@arco-design/web-vue/es/icon";
import {reactive, ref} from "vue";
import {nextTick} from "vue";
import type {listDataType, paramsType} from "@/api";
import {activitySearchApi, type activitySearchType} from "@/api/activity_api";

const visible = ref(false)

const params = reactive<paramsType>({})

const data = reactive<listDataType<activitySearchType>>({
  list: [],
  count: 0
})

const searchRef = ref()

async function search() {
  let res = await activitySearchApi(params)
  data.count = res.data.count
  data.list = res.data.list
}

function showSearch() {
  visible.value = true
  nextTick(() => {
    searchRef.value.focus()
  })
}

function goTo(item: activitySearchType) {
  window.open(`/article/${item.slug}`, "_blank")
}

</script>
<style lang="scss">
.index_search_body {
  padding: 0;

  .head {
    padding: 20px;
  }

  .body {
    max-height: 50vh;
    overflow-y: auto;
    color: var(--color-text-2);

    .item {
      padding: 10px 20px;
      cursor: pointer;

      &:hover {
        background-color: var(--color-fill-2);
      }

      .title {
        color: var(--color-text-1);
        font-size: 16px;
        margin-bottom: 5px;
      }
    }
  }

  .footer {
    text-align: center;
    padding: 10px 0;
    color: var(--color-text-2);
  }
}
</style>
