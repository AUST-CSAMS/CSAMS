<template>
  <div class="article_view">
    <index_nav></index_nav>
    <main>
      <div class="container">
        <div class="article_container">
          <div class="head">
            <div class="title">{{ data.activity_name }}</div>
            <div class="date">
              开始时间：{{ dateFormat(data.startTime) }}
            </div>
            <div class="date">
              结束时间：{{ dateFormat(data.endTime) }}
            </div>
            <div class="date=">
              限制：{{ data.limit }}
            </div>
          </div>
          <article>
            <MdPreview v-model="data.introduction"></MdPreview>
          </article>
        </div>
      </div>
      <Button type="primary" @click="activityjoin">参加</Button>
    </main>
    <index_footer></index_footer>
  </div>
</template>
<script lang="ts" setup>
import {dateFormat} from "@/utils/date";
import {MdPreview} from "md-editor-v3";
import "md-editor-v3/lib/preview.css"
import Index_nav from "@/components/index_nav.vue";
import Index_footer from "@/components/index_footer.vue";
import {onMounted, reactive, ref, watch} from "vue";
import {
  activityInfoApi,
  activityJoinApi,
  type activityJoinType,
  type activityRequest
} from "@/api/activity_api";
import {Button, Message} from "@arco-design/web-vue";
import router from "@/router";
import {useRoute} from "vue-router";
import {useStore} from "@/stores";

const route = useRoute()

const store = useStore()

const id = ref<number>(parseInt(route.params.id as string))


const data = reactive<activityRequest>({
  id: 0,
  activity_name: "",
  startTime: "",
  endTime: "",
  location: "",
  introduction: "",
  image: "",
  responsible_person: "",
  tel: 0,
  score: 0,
  limit: []
})

async function getData() {
  let res = await activityInfoApi(id.value)
  console.log(id.value)
  console.log(res.code)
  console.log(dateFormat(data.startTime))
  console.log(data.startTime)
  if (res.code) {
    Message.warning("活动不存在")
    return
  }
  Object.assign(data, res.data)
}

async function activityjoin() {
  let form: activityJoinType = {
    id: data.id
  }
  let res = await activityJoinApi(form)
  if (res.code) {
    Message.error(res.msg)
    return false
  }
  Message.success(res.msg)
}

watch(() => route.params, () => {
  if (route.params.id) {
    id.value = Number(route.params.id as string)
    getData()
  }
}, {immediate: true, deep: true})

onMounted(() => {
  let hash = route.hash
  if (hash !== "") {
    setTimeout(() => {
      let dom = document.querySelector(hash)
      if (dom) {
        let top = dom.getBoundingClientRect().top
        document.documentElement.scrollTo({
          top: top - 60,
          behavior: "smooth"
        })
      }
    }, 150)


  }
})

</script>

<style lang="scss">
.article_view {
  main {
    width: 100%;
    display: flex;
    justify-content: center;
    background-color: var(--bg);
    padding-top: 20px;
    padding-bottom: 20px;
    min-height: 83vh;

    .container {
      width: 1200px;
      display: flex;
      justify-content: center;

      .article_container {

        .head {
          border-radius: 5px 5px 0 0;
          margin-bottom: 1px;
          padding: 20px;
          display: flex;
          flex-direction: column;
          align-items: center;

          .title {
            font-size: 20px;
            font-weight: 600;
            margin-bottom: 10px;
          }

          .date {
            margin-bottom: 10px;
          }
        }
      }
    }
  }
}
</style>
