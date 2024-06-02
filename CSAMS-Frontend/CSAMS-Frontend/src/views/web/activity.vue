<template>
  <div class="article_view">
    <index_nav></index_nav>
    <main>
      <div class="container">
        <div class="article_container">
          <div class="head">
            <div :id="data.title" class="title">{{ data.title }}</div>
            <div class="date">
              发布时间：{{ dateFormat(data.create_at) }}
            </div>
          </div>
          <article>
            <MdPreview v-model="data.content"></MdPreview>
          </article>
        </div>
      </div>
      <Button type="primary" @click="">参加</Button>
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
import {activityDetailApi, type activityRequest} from "@/api/activity_api";
import {Button, Message} from "@arco-design/web-vue";
import router from "@/router";
import {useRoute} from "vue-router";

const route = useRoute()


const id = ref<string>(route.params.id as string)


const data = reactive<activityRequest>({
  id: 0,
  title: "",
  create_at: "",
  place: "",
  content: ""
})

async function getData() {
  let res = await activityDetailApi(id.value)
  if (res.code) {
    Message.warning("文章不存在")
    await router.push({
      name: "article_notfound"
    })
    return
  }
  Object.assign(data, res.data)
}

watch(() => route.params, () => {
  if (route.params.id) {
    id.value = route.params.id as string
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

    .container {
      width: 1200px;
      display: flex;
      justify-content: space-between;

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
