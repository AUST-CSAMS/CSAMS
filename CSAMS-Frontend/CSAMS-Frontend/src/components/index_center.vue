<template>
  <div class="container">
    <div class="title">正在进行的活动</div>
    <div class="activity_source">
      <div v-for="item in data.list" class="item">
        <a :href="`/activities/${item.id}`" v-html="item.title"></a>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type {listDataType, paramsType} from "@/api";
import {activityListApi, type activityRequest,} from "@/api/activity_api";
import {reactive} from "vue";

const params = reactive<paramsType>({
  key: ""

})

const data = reactive<listDataType<activityRequest>>({
  list: [],
  count: 0
})

async function getData(p?: paramsType) {
  if (p) {
    Object.assign(params, p)
  }
  let res = await activityListApi(params)
  data.list = res.data.list
  data.count = res.data.count
}

getData()

</script>

<style lang="scss">
.container {
  width: 360px;
  border: rgb(11, 196, 221) solid 5px;
  border-radius: 20px;
  margin: 20px;
  justify-content: center;
  overflow-y: auto;


  .title {
    text-align: center;
    font-size: 18px;
    font-weight: bold;
    margin: 10px 0;
  }

  .activity_source {

    .item {
      display: flex;
      align-items: center;
      justify-content: center;
      font-size: 16px;

      .content {
        margin: 5px 0;
      }

      &:hover {
        cursor: pointer;
        background-color: rgba(212, 227, 236, 0.7);
      }
    }

  }


}
</style>
