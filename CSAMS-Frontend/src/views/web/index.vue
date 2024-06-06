<template>
  <div class="index_view">
    <index_nav></index_nav>
    <main>
      <div class="index_web_container">
        <index_left></index_left>
        <index_center></index_center>
        <index_right></index_right>
      </div>
    </main>
    <index_footer></index_footer>
  </div>

</template>

<script lang="ts" setup>

import Index_nav from "@/components/index_nav.vue";
import Index_footer from "@/components/index_footer.vue";
import Index_left from "@/components/index_left.vue";
import Index_right from "@/components/index_right.vue";
import Index_center from "@/components/index_center.vue";
</script>

<style lang="scss">
.index_view {
  background: url(/image/天空.jpg) 50%/cover no-repeat fixed;
  min-height: 100vh;

  main {
    width: 100%;
    justify-content: center;
    display: flex;
    height: 622px;

    .index_web_container {
      width: 1200px;
      display: flex;
    }
  }
}
</style>
