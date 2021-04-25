<template>
  <global-footer class="footer custom-render">
    <template v-slot:links>
      <!-- <a href="https://www.github.com/vueComponent/pro-layout" target="_blank">Pro Layout</a> -->
      <a href="https://www.github.com/vueComponent/ant-design-vue-pro" target="_blank">Ant Design Pro - Github</a>
      <!-- <a href="https://www.github.com/sendya/" target="_blank">@Sendya</a> -->
    </template>
    <template v-slot:copyright>
      <!-- <a href="https://github.com/vueComponent" target="_blank">vueComponent</a> -->
      <a href="#" target="_blank">doconsole - modified by favefan</a>
    </template>
  </global-footer>
</template>

<script>
import { GlobalFooter } from '@ant-design-vue/pro-layout'

export default {
  name: 'ProGlobalFooter',
  components: {
    GlobalFooter
  }
}
</script>
