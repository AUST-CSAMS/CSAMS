<template>
  <div class="main">
    <div class="article_list">
      <div v-for="item in data.list" class="item">
        <div>协会信息</div>
        <div>协会名: {{ item.association_name }}</div>
        <div>宣言: {{ item.introduction }}</div>
        <div>教师: {{ item.teacher_name }}</div>
        <div>会长: {{ item.president }}</div>
        <div>成立时间: {{ dateTimeFormat(item.create_at) }}</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {
  type associationInfoType, associationListApi,
} from "@/api/association_api";
import type {listDataType, paramsType} from "@/api";
import {dateTimeFormat} from "@/utils/date";

const params = reactive<paramsType>({
  key: ""
})

const data = reactive<listDataType<associationInfoType>>({
  list: [],
  count: 0
})


async function getData(p?: paramsType) {
  if (p) {
    Object.assign(params, p)
  }
  let res = await associationListApi(params)
  data.list = res.data.list
  data.count = res.data.count
}

getData()


</script>

<style lang="scss">
.main {
  min-height: 100vh;

  .article_list {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    grid-column-gap: 10px;
    grid-row-gap: 5px;

    .item {
      margin: 10px 0;
      padding: 10px;
      width: 100%;
      border-radius: 15px;
      border: rgb(11, 196, 221) solid 5px;
      font-size: 14px;
    }
  }
}

</style>
