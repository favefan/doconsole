<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
      :title="imagesCount + ' 个镜像 - 总大小: ' + imagesTotalSize + ' MB'"
    >
      <div slot="extra">
        <a-tooltip title="从仓库中搜索">
          <a-button
            shape="circle"
            size="large"
            icon="search"
            @click="showDrawer" />
        </a-tooltip>
        <a-tooltip title="拉取镜像">
          <a-button
            type="primary"
            shape="circle"
            size="large"
            icon="cloud-download"
            style="margin-left: 10px"
            @click="showModal('DockerHub', '')"
          />
        </a-tooltip>

      </div>

      <div>
        <a-popconfirm
          placement="top"
          title="确定删除选中的镜像?"
          ok-text="是"
          cancel-text="我再想想"
          @confirm="batchRemove"
        >
          <a-button
            type="danger"
            icon="delete"
            style="margin-left: 10px"
            :disabled="!hasSelected"
            :loading="deleteLoading"
          >
            批量删除
          </a-button>
        </a-popconfirm>
        <a-input
          style="width: 280px; margin-bottom: 15px; margin-left: 10px"
          type="text"
          placeholder="列表过滤"
          v-model="filterKey"
          @change="filterChange"
          allowClear
        >
          <a-tooltip slot="prefix" title="Filter">
            <a-icon type="filter" style="color: rgba(0, 0, 0, 0.45)" />
          </a-tooltip>
        </a-input>
      </div>

      <!-- 镜像管理列表 -->
      <a-table
        rowKey="Id"
        :pagination="false"
        :columns="columns"
        :data-source="imageListData"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
      >
        <a slot="name" slot-scope="text, record" @click="handleToInformation(record.Id.split(':')[1])">
          {{ record.RepoTags[0].split(':')[0] }}
        </a>
        <span slot="createdTime" slot-scope="timestamp">
          {{ timestamp | formatDate }}
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
            <a-button
              type="primary"
              shape="circle"
              size="large"
              icon="caret-right"
              @click="runContainer(record.Id)"
              title="用此镜像创建新容器"
              style="margin-left: 10px"
              :loading="createLoading"
            ></a-button>
          </template>
        </span>
      </a-table>
    </a-card>

    <!-- 镜像搜索与添加 - 抽屉 -->
    <a-drawer
      title="镜像搜索"
      placement="right"
      :width="500"
      :visible="drawerVisible"
      :after-visible-change="afterVisibleChange"
      :destroyOnClose="true"
      @close="onClose"
    >
      <!-- 镜像仓库选择与搜索框 -->
      <a-row>
        <a-col :span="24">
          <a-input-group compact>
            <a-select style="width: 35%" placeholder="请选择" default-value="DockerHub" @change="setSearchRegistry" >
              <a-select-option value="DockerHub"> DockerHub </a-select-option>
              <a-select-option value="Custom"> Custom </a-select-option>
            </a-select>
            <a-input-search
              placeholder="左侧选择仓库后开始搜索"
              style="width: 65%"
              enter-button
              :loading="searchButtonLoading"
              @search="onSearch"
            />
          </a-input-group>
        </a-col>
      </a-row>

      <a-row>
        <a-col> </a-col>
      </a-row>

      <!-- 搜索结果列表 -->
      <a-row>
        <a-col :span="24">
          <a-list item-layout="horizontal" :data-source="searchResults" style="padding: 5px">
            <a-list-item slot="renderItem" slot-scope="item">
              <a-tooltip title="拉取该镜像" slot="actions">
                <a-button
                  type="primary"
                  shape="circle"
                  size="large"
                  icon="cloud-download"
                  @click="showModal(searchRegistry, item.name)"
                />
              </a-tooltip>
              <a-list-item-meta :description="item.description">
                <a slot="title" href="https://www.antdv.com/">
                  {{ item.name }}
                  <a-tag v-if="item.is_official" color="green"> 官方 </a-tag>
                  <a-tag v-if="item.is_automated" color="cyan"> 自动构建 </a-tag>
                  <a-icon type="star" theme="filled" style="color: gold" />
                  {{ item.star_count }}
                </a>
              </a-list-item-meta>
            </a-list-item>
          </a-list>
        </a-col>
      </a-row>
    </a-drawer>

    <!-- 镜像拉取弹出框 -->
    <a-modal
      title="镜像拉取"
      :visible="imagePullModalVisible"
      :confirm-loading="pullLoading"
      :destroyOnClose="true"
      :zIndex="9999"
      @ok="handleOk"
      @cancel="handleCancel"
    >
      <p style="margin-top: 10px">从仓库:</p>
      <a-select style="width: 45%" placeholder="请选择" v-model="modalPullRegistry">
        <a-select-option value="DockerHub"> DockerHub </a-select-option>
        <a-select-option value="Custom"> Custom </a-select-option>
      </a-select>
      <p style="margin-top: 10px">拉取镜像:</p>
      <a-input-group compact>
        <a-input style="width: 35%" placeholder="镜像名称" v-model="modalPullImageName" />
        <a-input style="width: 45%" placeholder="镜像标签" v-model="modalPullImageTag"/>
      </a-input-group>
    </a-modal>

  </page-header-wrapper>
</template>

<script>
import { STable, Ellipsis } from '@/components'

import { ImageList, ImageSearch, ImagePull, ImagesRemove } from '@/api/images'
import { ContainerCreate } from '@/api/containers'

const columns = [
  {
    title: ' ',
    dataIndex: 'Name',
    scopedSlots: { customRender: 'name' }
    // customRender: (repotag) => repotag.split('@')[0]
  },
  {
    title: 'Tag',
    dataIndex: 'RepoTags[0]',
    customRender: (repotag) => repotag.split(':')[1]
  },
  {
    title: 'ID',
    dataIndex: 'Id',
    ellipsis: true,
    customRender: (Id) => Id.split(':')[1]
  },
  {
    title: '创建时间',
    dataIndex: 'Created',
    scopedSlots: { customRender: 'createdTime' }
  },
  {
    title: '存储占用',
    dataIndex: 'Size',
    customRender: (size) => (size / 1000000.0).toFixed(2) + ' MB'
  },
  {
    title: ' ',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  name: 'Images',
  components: {
    STable,
    Ellipsis
  },
  data () {
    this.columns = columns
    return {
      filterKey: '',
      imageListData: [],
      cacheData: [],
      deleteLoading: false,
      imagesCount: 0,
      imagesTotalSize: 0,
      drawerVisible: false,
      selectedRowKeys: [],
      searchResults: [],
      searchButtonLoading: false,
      searchRegistry: 'DockerHub',
      imagePullModalVisible: false,
      pullLoading: false,
      modalPullRegistry: 'DockerHub',
      modalPullImageName: '',
      modalPullImageTag: '',
      createLoading: false
    }
  },
  filters: {
    formatDate: function (value) {
      const date = new Date(parseInt(value) * 1000)
      const y = date.getFullYear()
      let MM = date.getMonth() + 1
      MM = MM < 10 ? '0' + MM : MM
      let d = date.getDate()
      d = d < 10 ? '0' + d : d
      let h = date.getHours() + 8
      h = h < 10 ? '0' + h : h
      let m = date.getMinutes()
      m = m < 10 ? '0' + m : m
      let s = date.getSeconds()
      s = s < 10 ? '0' + s : s
      return y + '-' + MM + '-' + d + ' ' + h + ':' + m + ':' + s
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
        return item.RepoTags[0].split(':')[0].toLowerCase().indexOf(key.toLowerCase()) !== -1
      })
    },
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange
      }
    },
    hasSelected () {
      return this.selectedRowKeys.length > 0
    }
  },
  watch: {},
  methods: {
    // 列表更新
    freshList () {
      ImageList()
        .then((res) => {
          this.cacheData = res.data
          this.imageListData = this.cacheData
          this.imagesCount = res.data.length
          this.imagesTotalSize = (res.data.reduce((sum, e) => sum + Number(e.Size), 0) / 1000000.0).toFixed(2)
        })
        .catch((err) => {
          this.$message.error(`更新镜像列表失败: ${err.message}`)
        })
    },
    filterChange () {
      var key = this.filterKey.trim()
      if (key.length === 0) {
        this.imageListData = this.cacheData
      } else {
        this.imageListData = this.filterList
      }
    },
    // 多选操作
    onSelectChange (selectedRowKeys) {
      // console.log('selectedRowKeys changed: ', selectedRowKeys)
      this.selectedRowKeys = selectedRowKeys
    },
    // 多选删除
    batchRemove () {
      this.deleteLoading = true
      ImagesRemove({ 'array': this.selectedRowKeys })
        .then((res) => {
          if (res.data !== null) {
            this.$message.warning(`成功删除 ${res.data.length} 个镜像，删除镜像的网络有: ${res.data}`)
          } else {
            this.$message.success(`成功删除 ${this.selectedRowKeys.length} 个镜像`)
          }
        })
        .catch((err) => {
          this.$message.error(`镜像删除错误: ${err.response.data.data}`)
        })
        .finally(() => {
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
    },

    // 镜像详情跳转
    handleToInformation (id) {
      this.$router.push({ path: `/image/information?id=${id}` })
    },

    // 镜像创建容器
    runContainer (id) {
      this.createLoading = true
      var data = {
        Image: id
      }
      ContainerCreate(data)
        .then((res) => {
          this.$message.success(`容器创建成功`)
          this.$router.push({ path: `/container/information?id=${res.data}` })
        })
        .catch((err) => {
          this.$message.error(`容器创建错误: ${err.response.data.data}`)
        })
        .finally(() => {
          this.createLoading = false
          this.freshList()
        })
    },

    // 单个镜像删除确认
    // deleteConfirm (id) {
    //   ImageRemove(id)
    //     .then((res) => {
    //       this.$message.success(`镜像删除成功`)
    //       this.freshList()
    //     })
    //     .catch((err) => {
    //       this.$message.error(`镜像删除错误: ${err.response.data.data}`)
    //     })
    //     .finally(() => {
    //     })
    // },
    cancel (e) {
      // console.log(e)
    },

    // 抽屉 - 搜索/添加 镜像
    afterVisibleChange (val) {
      // console.log('visible', val)
    },
    showDrawer () {
      this.drawerVisible = true
    },
    onClose () {
      this.drawerVisible = false
      this.searchResults = []
    },
    setSearchRegistry (value) {
      this.searchRegistry = value
    },
    onSearch (value) {
      this.searchButtonLoading = true
      ImageSearch(value)
        .then((res) => {
          this.searchResults = res.data.sort(function (a, b) {
            return b.star_count - a.star_count
          })
        })
        .catch((err) => {
          this.$message.error(`搜索错误: ${err.message}`)
        })
        .finally(() => {
          this.searchButtonLoading = false
        })
    },

    // 镜像拉取弹出框
    showModal (registry, name) {
      this.modalPullRegistry = registry
      this.modalPullImageName = name
      this.imagePullModalVisible = true
    },
    handleOk (e) {
      // console.log(e)
      this.pullLoading = true
      var image = this.modalPullImageName
      if (this.modalPullImageTag !== '') {
        image += ':' + this.modalPullImageTag
      }
      ImagePull(image)
        .then((res) => {
          this.$message.success('拉取成功')
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`拉取错误: ${err.message}`)
        })
        .finally(() => {
          this.pullLoading = false
          this.imagePullModalVisible = false
          this.modalPullRegistry = 'DockerHub'
          this.modalPullImageName = ''
          this.modalPullImageTag = ''
        })
    },
    handleCancel (e) {
      this.imagePullModalVisible = false
      this.modalPullRegistry = 'DockerHub'
      this.modalPullImageName = ''
      this.modalPullImageTag = ''
    }
  }
}
</script>

<style lang="less" scoped>
.image-use-statu {
  float: left;
  width: 170px;
  font-weight: 600;
  margin-right: 30px;
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
