<template>
  <div :class="wrpCls">
    <!-- <avatar-dropdown :menu="showMenu" :current-user="currentUser" :class="prefixCls" /> -->
    <!-- <select-lang :class="prefixCls" /> -->
    <host-selection></host-selection>
  </div>
</template>

<script>
// import AvatarDropdown from './AvatarDropdown'
import HostSelection from './HostSelection'
// import SelectLang from '@/components/SelectLang'

export default {
  name: 'RightContent',
  components: {
    // AvatarDropdown,
    // SelectLang
    HostSelection
  },
  props: {
    prefixCls: {
      type: String,
      default: 'ant-pro-global-header-index-action'
    },
    isMobile: {
      type: Boolean,
      default: () => false
    },
    topMenu: {
      type: Boolean,
      required: true
    },
    theme: {
      type: String,
      required: true
    }
  },
  data () {
    return {
      // showMenu: true,
      // currentUser: {}
      // HostSelectionList: []
    }
  },
  computed: {
    wrpCls () {
      return {
        'ant-pro-global-header-index-right': true,
        [`ant-pro-global-header-index-${(this.isMobile || !this.topMenu) ? 'light' : this.theme}`]: true
      }
    }
  },
  mounted () {
    // setTimeout(() => {
    //   this.currentUser = {
    //     name: 'Admin'
    //   }
    // }, 1500)
  }
}
</script>
