<template>
  <div class="header-info">
    <span>{{ title }}</span>
    <p>{{ value }}</p>
    <em v-if="bordered" />
  </div>
</template>

<script>
export default {
  name: 'OperationButtonGroup',
  props: {
    title: {
      type: String,
      default: ''
    },
    value: {
      type: String,
      default: ''
    },
    bordered: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style lang="less" scoped>
@import "~ant-design-vue/es/style/themes/default.less";

</style>
