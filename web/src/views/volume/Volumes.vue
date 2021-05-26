<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
    >
      <div slot="title">
        <a-popconfirm
          placement="top"
          title="确定删除选中的卷?"
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
            删除
          </a-button>
        </a-popconfirm>
        <a-button
          type="primary"
          icon="plus"
          style="margin-left: 10px"
          @click="showDrawer"
        >
          新建
        </a-button>
        <!-- 列表过滤框 -->
        <a-input
          style="width: 280px; margin-bottom: 15px; margin-left: 10px"
          type="text"
          allowClear
          placeholder="列表过滤"
          v-model="filterKey"
          @change="filterChange"
        >
          <a-icon type="filter" style="color: rgba(0, 0, 0, 0.45)" />
        </a-input>
      </div>

      <!-- 卷列表 -->
      <a-table
        rowKey="Name"
        :pagination="false"
        :columns="columns"
        :data-source="listData"
        :row-selection="rowSelection"
      >
        <a slot="name" slot-scope="text, record" @click="handleToInformation(record.Name)">
          {{ record.Name | shortName }}
        </a>
        <p slot="mountpoint" slot-scope="text, record">
          {{ record.Mountpoint | midString }}
        </p>
        <!-- <a-tag
          color="cyan"
          slot="issystem"
          slot-scope="text, record"
          v-if="record.Name === 'bridge' || record.Name === 'none' || record.Name === 'host'"
        >
          预设
        </a-tag> -->
      </a-table>
    </a-card>

    <!-- 添加抽屉 -->
    <a-drawer
      title="创建卷"
      placement="right"
      :width="521"
      :visible="drawerVisible"
      :destroyOnClose="true"
      :maskClosable="true"
      @close="drawerOnClose"
    >
      <a-form :form="form" :label-col="{ span: 5 }" :wrapper-col="{ span: 19 }" @submit="createSubmit">
        <a-form-item label="名称">
          <a-input
            v-decorator="[
              'Name',
              {
                rules: [
                  {
                    required: true,
                    message: '名称可能是必要的，填就对了'
                  }
                ]
              }
            ]"
            placeholder="例如 myVolume"
          />
        </a-form-item>
        <a-form-item label="Driver">
          <a-select
            v-decorator="[
              'Driver',
              {
                initialValue: 'local',
                rules: [
                  {
                    required: true,
                    message: 'Driver是必需的'
                  }
                ]
              },
            ]"
            placeholder="请选择"
          >
            <a-select-option value="bridge">
              bridge
            </a-select-option>
            <a-select-option value="etc">
              etc
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
          <a-button type="primary" html-type="submit">
            创建
          </a-button>
        </a-form-item>
      </a-form>
    </a-drawer>
  </page-header-wrapper>
</template>

<script>
import { VolumeList, VolumeCreate, VolumesRemove } from '@/api/volumes'

const columns = [
  {
    title: 'Name',
    dataIndex: 'Name',
    scopedSlots: { customRender: 'name' }
  },
  {
    title: 'CreatedAt',
    dataIndex: 'CreatedAt',
    customRender: (datetime) => datetime.split('T')[0] + ' ' + datetime.split('T')[1].split('Z')[0]
  },
  {
    title: 'Driver',
    dataIndex: 'Driver'
  },
  {
    title: 'Mountpoint',
    dataIndex: 'Mountpoint',
    scopedSlots: { customRender: 'mountpoint' }
  }
]

export default {
  name: 'Images',
  components: {

  },
  data () {
    this.columns = columns
    return {
      formLayout: 'horizontal',
      form: this.$form.createForm(this, { name: 'volumecreate' }),
      listData: [],
      volumeListData: [],
      drawerVisible: false,
      selectedRowKeys: [],
      deleteLoading: false,
      filterKey: ''
    }
  },
  filters: {
    shortName: function (val) {
      var short = val
      if (val.length >= 10) {
        short = val.substr(0, 10)
      }
      return short
    },
    midString: function (val) {
      var mid = val
      if (val.length >= 50) {
        mid = val.substr(0, 50)
      }
      return mid
    }
  },
  created () {
    this.freshList()
  },
  computed: {
    filterList: function () {
      var key = this.filterKey
      var list = this.volumeListData
      return list.filter(function (item) {
        return item.Name.toLowerCase().indexOf(key.toLowerCase()) !== -1
      })
    },
    rowSelection () {
      return {
        selectedRowKeys: this.selectedRowKeys,
        onChange: this.onSelectChange,
        getCheckboxProps: record => ({
          props: {
          }
        })
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
      VolumeList()
        .then((res) => {
          this.volumeListData = res.data.Volumes
          this.listData = this.volumeListData
          this.filterKey = ''
        })
        .catch((err) => {
          this.$message.error(`更新卷列表失败: ${err.message}`)
        })
    },

    filterChange () {
      var key = this.filterKey.trim()
      if (key.length === 0) {
        this.listData = this.volumeListData
      } else {
        this.listData = this.filterList
      }
    },

    // 多选操作
    onSelectChange (selectedRowKeys) {
      // console.log('selectedRowKeys changed: ', selectedRowKeys)
      this.selectedRowKeys = selectedRowKeys
    },
    // 多选删除
    batchRemove () {
      console.log('selectedRowKeys selected: ', this.selectedRowKeys)
      this.deleteLoading = true
      VolumesRemove({ 'array': this.selectedRowKeys })
        .then((res) => {
          console.log(res.data)
          if (res.data !== null) {
            this.$message.warning(`成功删除 ${res.data.length} 个卷，删除失败的卷有: ${res.data}`)
          } else {
            this.$message.success(`成功删除 ${this.selectedRowKeys.length} 个卷`)
          }
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`卷删除错误: ${err.message}`)
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
    },

    // 镜像详情跳转
    handleToInformation (id) {
      this.$router.push({ path: `/volume/information?name=${id}` })
    },

    // 抽屉 - 搜索/添加 镜像
    showDrawer () {
      this.drawerVisible = true
    },

    drawerOnClose () {
      this.drawerVisible = false
    },

    createSubmit (e) {
      e.preventDefault()
      this.form.validateFields((err, values) => {
        if (!err) {
          VolumeCreate(values)
          .then((res) => {
            this.$message.success('创建卷成功')
            console.log(res)
            this.freshList()
            this.drawerVisible = false
          })
          .catch((err) => {
            this.$message.error(`创建s卷列表失败: ${err.message}`)
            this.drawerVisible = false
          })
        }
      })
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
