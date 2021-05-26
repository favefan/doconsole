<template>
  <div v-if="data" class="action-button-list">
    <a-button
      shape="circle"
      type="primary"
      size="large"
      :icon="data.State === 'running' ? 'poweroff' : 'caret-right'"
      @click="data.State === 'running' ? stop(data.Id) : start(data.Id)"
      title="启动或停止"
    ></a-button>
    <a-button
      shape="circle"
      size="large"
      icon="reload"
      @click="restart(data.Id)"
      :disabled="data.State === 'running' ? false : true"
      title="重启"
    ></a-button>
    <a-popconfirm
      placement="topLeft"
      title="确定删除这个容器?"
      ok-text="是"
      cancel-text="我再想想"
      @confirm="confirm(data.Id)"
      @cancel="cancel"
    >
      <a-button shape="circle" type="danger" size="large" icon="delete" title="删除"></a-button>
    </a-popconfirm>
  </div>
</template>

<script>
import { ContainerStart, ContainerStop, ContainerRestart, ContainerRemove } from '@/api/containers'

export default {
  name: 'OperationButtonGroup',
  props: {
    data: {
      type: Object,
      required: true,
      default: () => null
    },
    call: {
      type: Function,
      default: null,
      required: true
    },
    outpage: {
      type: Boolean,
      default: false
    }
  },
  methods: {
    start: function (id) {
      ContainerStart(id)
        .then((res) => {
          this.$message.success('容器启动成功')
          this.call()
        })
        .catch((err) => {
          this.$message.error(`容器启动失败: ${err.message}`)
        })
    },
    stop: function (id) {
      ContainerStop(id)
        .then((res) => {
          this.$message.success('容器停止成功')
          this.call()
        })
        .catch((err) => {
          this.$message.error(`容器停止失败: ${err.message}`)
        })
    },
    restart: function (id) {
      ContainerRestart(id)
        .then((res) => {
          this.$message.success('容器重启成功')
          this.call()
        })
        .catch((err) => {
          this.$message.error(`容器重启成功: ${err.message}`)
        })
    },
    confirm (id) {
      var _this = this
      ContainerRemove(id)
        .then((res) => {
          this.$message.success('容器删除成功')
          if (_this.outpage) {
            _this.$router.push({ path: `/container/containers` })
          } else {
            _this.call()
          }
        })
        .catch((err) => {
          this.$message.error(`容器删除失败: ${err.message}`)
        })
    },
    cancel (e) {
      console.log(e)
    }
  }
}
</script>

<style lang="less" scoped>
.action-button-list * {
  margin-left: 10px;
}
</style>
