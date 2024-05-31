<template>
  <div class="container">
    <div class="title">已经结束的活动</div>
    <div class="activity_source">
      <!--      <div v-for="item in data.list" class="item" @click="goTo(item)">-->
      <!--        <div class="content">{{ item.title }}</div>-->
      <!--      </div>-->
      <div class="item">
        <div class="content">爱心活动</div>
      </div>
      <div class="item">
        <div class="content">爱心活动</div>
      </div>
      <div class="item">
        <div class="content">爱心活动</div>
      </div>
      <div class="item">
        <div class="content">爱心活动</div>
      </div>
      <div class="item">
        <div class="content">爱心活动</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import type {listDataType, paramsType} from "@/api";
import {activitySearchApi, type activitySearchType} from "@/api/activity_api";
import {reactive} from "vue";

const params = reactive<paramsType>({})

const data = reactive<listDataType<activitySearchType>>({
  list: [],
  count: 0
})

async function search() {
  let res = await activitySearchApi(params)
  data.count = res.data.count
  data.list = res.data.list
}

search()

function goTo(item: activitySearchType) {
  window.open(`/article/${item.slug}`, "_blank")
}
</script>

<style lang="scss">
.container {
  width: 360px;
  border: rgb(11, 196, 221) solid 5px;
  border-radius: 20px;
  margin: 20px;
  justify-content: center;

  .title {
    text-align: center;
    font-size: 18px;
    font-weight: bold;
    margin: 10px 0;
  }

  .activity_source {
    overflow-y: auto;

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
