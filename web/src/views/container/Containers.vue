<template>
  <page-header-wrapper>
    <a-card style="margin-top: 24px" :bordered="false" :title="containerCount + '个容器'">
      <div slot="extra">
        <!-- <a-radio-group v-model="status">
          <a-radio-button value="all">全部</a-radio-button>
          <a-radio-button value="processing">正在运行</a-radio-button>
          <a-radio-button value="waiting">未运行</a-radio-button>
        </a-radio-group> -->
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
        <a-list-item :key="index" v-for="(item, index) in data">
          <a-list-item-meta :description="item.Status">
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
            <operation-button-group :data="item" :call="freshList"></operation-button-group>
          </div>
        </a-list-item>
      </a-list>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import ContainerStoppedIcon from '../../assets/icons/container-stopped-icon.png'
import ContainerRunningIcon from '../../assets/icons/container-running-icon.png'

import { ContainerList } from '@/api/containers'

import OperationButtonGroup from './components/OperationButtonGroup'

export default {
  name: 'Containers',
  components: {
    OperationButtonGroup
  },
  data () {
    return {
      data: [],
      cacheData: [],
      status: 'all',
      containerStoppedIcon: ContainerStoppedIcon,
      containerRunningIcon: ContainerRunningIcon,
      containerCount: 0,
      filterKey: ''
    }
  },
  created () {
    this.freshList()
  },
  computed: {
    filterList: function () {
      var key = this.filterKey
      var list = this.cacheData
      return list.filter(function (item) {
        return item.Names[0].toLowerCase().indexOf(key.toLowerCase()) !== -1
      })
    }
  },
  methods: {
    freshList () {
      ContainerList()
        .then((res) => {
          this.cacheData = res.data
          this.data = this.cacheData
          this.containerCount = res.data.length
        })
        .catch((err) => {
          this.$message.error(`更新容器列表失败: ${err.message}`)
        })
    },
    filterChange () {
      var key = this.filterKey.trim()
      if (key.length === 0) {
        this.data = this.cacheData
      } else {
        this.data = this.filterList
      }
    },
    handleToContainerInformation (id) {
      this.$router.push({ path: `/container/information?id=${id}` })
    },
    handleToImageInformation (id) {
      this.$router.push({ path: `/image/information?id=${id}` })
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
