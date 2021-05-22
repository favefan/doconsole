<template>
  <page-header-wrapper>
    <a-card style="margin-top: 24px" :bordered="false" :title="containerCount + '个容器'">
      <div slot="extra">
        <a-radio-group v-model="status">
          <a-radio-button value="all">全部</a-radio-button>
          <a-radio-button value="processing">正在运行</a-radio-button>
          <a-radio-button value="waiting">未运行</a-radio-button>
        </a-radio-group>
        <a-input
          style="width: 280px; margin-left: 10px"
          type="text"
          allowClear
          placeholder="列表过滤"
          v-model="filterKey"
          @change="filterChange"
        >
          <a-icon type="filter" style="color: rgba(0, 0, 0, 0.45)" />
        </a-input>
        <a-button
          type="primary"
          icon="plus"
          style="margin-left: 10px"
          @click="toCreator()"
        >
          新建
        </a-button>
      </div>

      <a-list size="large">
        <!-- :pagination="{showSizeChanger: true, showQuickJumper: true, pageSize: 5, total: 50}" -->
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.Status + ' 端口:' + item.Ports[0]">
            <a-avatar
              slot="avatar"
              size="large"
              shape="square"
              :src="item.State === 'running' ? containerRunningIcon : containerStoppedIcon"
            />
            <a slot="title" style="font-size: 110%" @click="handleToContainerInformation(item.Id)">{{ item.Names[0] }}</a>
            <a slot="title" style="color: blue; font-size: 80%" @click="handleToImageInformation(item.ImageID)">&nbsp;&nbsp;{{ item.Image }}</a>
          </a-list-item-meta>
          <div slot="actions" class="action-button-list">
            <a-button
              type="circle"
              size="large"
              icon="code"
              title="CLI"
              :disabled="item.State === 'running' ? false : true"
            ></a-button>
            <a-button
              type="circle"
              size="large"
              :icon="item.State === 'running' ? 'poweroff' : 'caret-right'"
              @click="item.State === 'running' ? stop(item.Id) : start(item.Id)"
              title="启动或停止"
            ></a-button>
            <a-button
              type="circle"
              size="large"
              icon="reload"
              @click="restart(item.Id)"
              :disabled="item.State === 'running' ? false : true"
              title="重启"
            ></a-button>
            <a-popconfirm
              placement="topLeft"
              title="确定删除这个容器?"
              ok-text="是"
              cancel-text="我再想想"
              @confirm="confirm(item.Id)"
              @cancel="cancel"
            >
              <a-button type="circle" size="large" icon="delete" title="删除"></a-button>
            </a-popconfirm>
          </div>
          <div slot="actions">
            <a-dropdown placement="bottomRight">
              <a-menu class="more" slot="overlay">
                <a-menu-item><a>导出为镜像</a></a-menu-item>
                <a-menu-item><a>无</a></a-menu-item>
              </a-menu>
              <a-button type="circle" size="large" icon="more"></a-button>
            </a-dropdown>
          </div>
          <!-- <div class="list-content">
            <div class="list-content-item">
              <span>Owner</span>
              <p>1</p>
            </div>
            <div class="list-content-item">
              <span>开始时间</span>
              <p>2</p>
            </div>
            <div class="list-content-item">
              3<a-progress :percent="item.progress.value" :status="!item.progress.status ? null : item.progress.status" style="width: 180px" />
            </div>
          </div> -->
        </a-list-item>
      </a-list>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import ContainerStoppedIcon from '../../assets/icons/container-stopped-icon.png'
import ContainerRunningIcon from '../../assets/icons/container-running-icon.png'
// 演示如何使用 this.$dialog 封装 modal 组件

import {
  ContainerList,
  ContainerStart,
  ContainerStop,
  ContainerRestart,
  ContainerRemove
} from '@/api/containers'

export default {
  name: 'StandardList',
  components: {
  },
  data () {
    return {
      data: [],
      status: 'all',
      containerStoppedIcon: ContainerStoppedIcon,
      containerRunningIcon: ContainerRunningIcon,
      containerCount: 0,
      filterKey: ''
    }
  },
  // filters: {
  //   stateFilter (state) {
  //     // console.log(originVal)
  //     if (state === 'running') {
  //       return 'pause'
  //     } else {
  //       return 'caret-right'
  //     }
  //   }
  // },
  created () {
    this.freshList()
  },
  methods: {
    freshList () {
      ContainerList()
        .then((res) => {
          this.data = res.data
          this.containerCount = res.data.length
        })
        .catch((err) => {
          this.$message.error(`更新容器列表失败: ${err.message}`)
        })
    },
    filterChange () {
      // var key = this.filterKey.trim()
      // if (key.length === 0) {
      //   this.listData = this.networkListData
      // } else {
      //   this.listData = this.filterList
      // }
    },
    handleToContainerInformation (id) {
      this.$router.push({ path: `/container/information?id=${id}` })
    },
    handleToImageInformation (id) {
      this.$router.push({ path: `/image/information?id=${id}` })
    },
    start: function (id) {
      ContainerStart(id)
        .then((res) => {
          this.$message.success('容器启动成功')
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`容器启动失败: ${err.message}`)
        })
    },
    stop: function (id) {
      ContainerStop(id)
        .then((res) => {
          this.$message.success('容器停止成功')
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`容器停止失败: ${err.message}`)
        })
    },
    restart: function (id) {
      ContainerRestart(id)
        .then((res) => {
          this.$message.success('容器重启成功')
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`容器重启成功: ${err.message}`)
        })
    },
    confirm (id) {
      ContainerRemove(id)
        .then((res) => {
          this.$message.success('容器删除成功')
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`容器删除失败: ${err.message}`)
        })
    },
    cancel (e) {
      console.log(e)
    },
    toCreator () {
      this.$router.push({ path: `/container/creator` })
    }
  }
}
</script>

<style lang="less" scoped>
.ant-avatar-lg {
  width: 48px;
  height: 48px;
  line-height: 48px;
}

.list-content-item {
  color: rgba(0, 0, 0, 0.45);
  display: inline-block;
  vertical-align: middle;
  font-size: 14px;
  margin-left: 40px;
  span {
    line-height: 20px;
  }
  p {
    margin-top: 4px;
    margin-bottom: 0;
    line-height: 22px;
  }
}

.action-button-list * {
  margin-left: 10px;
}

.more {
  margin-top: 5px;
}

.operate {
  margin-bottom: 20px;
}
</style>
