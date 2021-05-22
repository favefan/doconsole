<template>
  <a-config-provider :locale="locale">
    <div id="app">
      <router-view v-if="isRouterAlive"/>
    </div>
  </a-config-provider>
</template>

<script>
import { domTitle, setDocumentTitle } from '@/utils/domUtil'
import { i18nRender } from '@/locales'

export default {
  provide () { // 在祖先组件中通过 provide 提供变量
    return {
      reload: this.reload //  声明一个变量
    }
  },
  data () {
    return {
      isRouterAlive: true // 控制 router-view 是否显示达到刷新效果
    }
  },
  computed: {
    locale () {
      // 只是为了切换语言时，更新标题
      const { title } = this.$route.meta
      title && (setDocumentTitle(`${i18nRender(title)} - ${domTitle}`))

      return this.$i18n.getLocaleMessage(this.$store.getters.lang).antLocale
    }
  },
  methods: {
    // provide中声明的变量
    // reload () {
    //   // 通过 this.isRouterAlive 控制 router-view 达到刷新效果
    //   this.isRouterAlive = false
    //   this.$nextTick(function () {
    //     this.isRouterAlive = true
    //   })
    // }
  }
}
</script>
