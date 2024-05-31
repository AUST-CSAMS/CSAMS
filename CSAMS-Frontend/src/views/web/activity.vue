<template>
  <div class="article_view">
    <index_nav></index_nav>
    <main>
      <div class="container">
        <div class="article_container">
          <div class="head">
            <div :id="data.name" class="title">{{ data.name }}</div>
            <div class="date">
              发布时间：{{ dateFormat(data.create_at) }}
            </div>
          </div>
          <article>
            <MdPreview v-model="data.content"></MdPreview>
          </article>
        </div>
        <aside>
          <association_info_preview></association_info_preview>
        </aside>
      </div>
    </main>
    <index_footer></index_footer>
  </div>
</template>
<script lang="ts" setup>
import {dateFormat} from "@/utils/date";
import {MdPreview} from "md-editor-v3";
import "md-editor-v3/lib/preview.css"
import Index_nav from "@/components/index_nav.vue";
import Association_info_preview from "@/components/association_info_preview.vue";
import Index_footer from "@/components/index_footer.vue";
import {onMounted, reactive, ref, watch} from "vue";
import {type activityCreateRequest, activityDetailApi} from "@/api/activity_api";
import {Message} from "@arco-design/web-vue";
import router from "@/router";
import {useRoute} from "vue-router";

const route = useRoute()


const id = ref<string>(route.params.id as string)


const data = reactive<activityCreateRequest>({
  name: "",
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
      width: var(--container_width);
      display: flex;
      justify-content: space-between;

      aside {
        width: 300px;

        .gvb_user_info_preview_component {
          margin-bottom: 20px;
        }

        .gvb_article_dict {
          .body {
            max-height: calc(100vh - 400px);
            overflow: auto;
          }
        }


        .article_actions {

          position: relative;

          &.isFixed {
            position: fixed;
            top: 80px;
            width: 300px;
          }
        }


        .gvb_article_action {
          margin-top: 20px;
          background-color: var(--bg1);
          border-radius: 20px;
          height: 50px;
          display: flex;
          align-items: center;

          > svg {
            flex: 1;
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 18px;
            color: var(--color-text-1);
            cursor: pointer;

            &:hover {
              color: var(--active);
            }

            &.active {
              color: var(--active);
            }
          }
        }

      }

      .article_container {
        width: calc(100% - 320px);

        .head {
          border-radius: 5px 5px 0 0;
          margin-bottom: 1px;
          background-color: var(--bg1);
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
            color: var(--color-text-2);
          }

          .tag {
            display: flex;
            align-items: center;
            justify-content: center;
            flex-wrap: wrap;

            > svg {
              margin-right: 10px;
              font-size: 16px;
              font-weight: 600;
            }

            .arco-tag {
              margin-right: 10px;

              &:last-child {
                margin-right: 0;
              }
            }
          }
        }

        .md-editor {
          background-color: var(--bg1);
        }

        .next_prev {
          background-color: var(--bg1);
          border-radius: 0 0 5px 5px;
          display: flex;
          justify-content: space-between;
          align-items: center;
          padding: 20px;
          margin-top: 1px;
          margin-bottom: 20px;

          a {
            color: var(--color-text-2);
          }
        }
      }

      .md-editor-catalog-link span:hover {
        color: var(--active);
      }

      .md-editor-catalog-active > span {
        color: var(--active);
      }
    }
  }
}
</style>
