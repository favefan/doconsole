<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
    >
      <div slot="title">
        <a-popconfirm
          placement="top"
          title="确定删除选中的网络?"
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

      <!-- 网络列表 -->
      <a-table
        rowKey="Id"
        :pagination="false"
        :columns="columns"
        :data-source="listData"
        :row-selection="rowSelection"
      >
        <a slot="name" slot-scope="text, record" @click="handleToInformation(record.Id)">
          {{ record.Name }}
        </a>
        <a-tag
          color="cyan"
          slot="issystem"
          slot-scope="text, record"
          v-if="record.Name === 'bridge' || record.Name === 'none' || record.Name === 'host'"
        >
          预设
        </a-tag>
      </a-table>
    </a-card>

    <!-- 网络添加抽屉 -->
    <a-drawer
      title="创建网络"
      placement="right"
      :width="521"
      :visible="drawerVisible"
      :destroyOnClose="true"
      :maskClosable="true"
      @close="drawerOnClose"
    >
      <a-form :form="form" :label-col="{ span: 5 }" :wrapper-col="{ span: 19 }" @submit="createNetworkSubmit">
        <a-form-item label="名称">
          <a-input
            v-decorator="[
              'Name',
              {
                rules: [
                  {
                    required: true,
                    message: '名称是必要的，且不能与已有的名称重复'
                  }
                ]
              }
            ]"
            placeholder="例如 myNetwork"
          />
        </a-form-item>
        <a-form-item label="Driver">
          <a-select
            v-decorator="[
              'Driver',
              {
                initialValue: 'bridge',
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

        <!-- IPV4 Network configuration -->
        <a-divider orientation="left">
          IPV4 Network 配置
        </a-divider>
        <a-form-item label="Subnet">
          <a-input
            v-decorator="[
              'v4subnet'
            ]"
            placeholder="例如 172.20.0.0/16"
          />
        </a-form-item>
        <a-form-item label="Gateway">
          <a-input
            v-decorator="[
              'v4gateway'
            ]"
            placeholder="例如 172.20.10.11"
          />
        </a-form-item>
        <a-form-item label="IP range">
          <a-input
            v-decorator="[
              'v4iprange'
            ]"
            placeholder="例如 172.20.10.128/25"
          />
        </a-form-item>

        <!-- IPV6 Network configuration -->
        <a-divider orientation="left">
          IPV6 Network 配置
        </a-divider>
        <a-form-item label="Subnet">
          <a-input
            v-decorator="[
              'v6subnet'
            ]"
            placeholder="例如 2001:db8::/48"
          />
        </a-form-item>
        <a-form-item label="Gateway">
          <a-input
            v-decorator="[
              'v6gateway'
            ]"
            placeholder="例如 2001:db8::1"
          />
        </a-form-item>
        <a-form-item label="IP range">
          <a-input
            v-decorator="[
              'v6iprange'
            ]"
            placeholder="例如 2001:db8::/64"
          />
        </a-form-item>

        <!-- Advanced configuration -->
        <a-divider orientation="left">
          高级配置
        </a-divider>
        <a-form-item label="Label">
          <a-input
            v-decorator="[
              'label',
              {
                rules: [
                  {
                    required: false,
                    message: ''
                  }
                ]
              }
            ]"
            placeholder="例如 com.example.tag:good"
          />
        </a-form-item>
        <a-form-item label="网络隔离">
          <a-switch v-decorator="['Internal', { initialValue: false, valuePropName: 'checked' }]" />
        </a-form-item>
        <a-form-item label="手动连接容器">
          <a-switch v-decorator="['Attachable', { initialValue: false, valuePropName: 'checked' }]" />
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
import { STable, Ellipsis } from '@/components'

import { NetworkList, NetworkCreate, NetworksRemove } from '@/api/networks'

const columns = [
  {
    title: 'Name',
    dataIndex: 'Name',
    scopedSlots: { customRender: 'name' }
  },
  {
    title: '',
    dataIndex: '',
    scopedSlots: { customRender: 'issystem' }
  },
  {
    title: 'ID',
    dataIndex: 'Id',
    ellipsis: true
  },
  {
    title: 'Created',
    dataIndex: 'Created',
    customRender: (datetime) => datetime.split('.')[0]
  },
  {
    title: 'Driver',
    dataIndex: 'Driver'
  },
  {
    title: 'Attachable',
    dataIndex: 'Attachable',
    customRender: (is) => is === false ? 'false' : 'true'
  },
  {
    title: 'IPAM Driver',
    dataIndex: 'IPAM.Driver'
  },
  {
    title: 'IPAM Subnet',
    dataIndex: 'IPAM.Config[0].Subnet'
  },
  {
    title: 'IPAM Gateway',
    dataIndex: 'IPAM.Config[0] .Gateway'
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
      formLayout: 'horizontal',
      form: this.$form.createForm(this, { name: 'networkcreate' }),
      listData: [],
      networkListData: [],
      drawerVisible: false,
      selectedRowKeys: [],
      deleteLoading: false,
      filterKey: ''
    }
  },
  filters: {
  },
  created () {
    this.freshList()
  },
  computed: {
    filterList: function () {
      var key = this.filterKey
      var list = this.networkListData
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
            disabled: record.Name === 'bridge' || record.Name === 'none' || record.Name === 'host'
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
      NetworkList()
        .then((res) => {
          this.networkListData = res.data
          this.listData = this.networkListData
          this.filterKey = ''
        })
        .catch((err) => {
          this.$message.error(`更新网络列表失败: ${err.message}`)
        })
    },

    filterChange () {
      var key = this.filterKey.trim()
      if (key.length === 0) {
        this.listData = this.networkListData
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
      NetworksRemove({ 'array': this.selectedRowKeys })
        .then((res) => {
          console.log(res.data)
          if (res.data !== null) {
            this.$message.warning(`成功删除 ${res.data.length} 个网络，删除失败的网络有: ${res.data}`)
          } else {
            this.$message.success(`成功删除 ${this.selectedRowKeys.length} 个网络`)
          }
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`网络删除错误: ${err.message}`)
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
    },

    // 镜像详情跳转
    handleToInformation (id) {
      this.$router.push({ path: `/network/information?id=${id}` })
    },

    // 抽屉 - 搜索/添加 镜像
    showDrawer () {
      this.drawerVisible = true
    },

    drawerOnClose () {
      this.drawerVisible = false
    },

    createNetworkSubmit (e) {
      e.preventDefault()
      this.form.validateFields((err, values) => {
        if (!err) {
          NetworkCreate(values)
          .then((res) => {
            this.$message.success('创建网络成功')
            this.freshList()
            this.drawerVisible = false
          })
          .catch((err) => {
            this.$message.error(`更新网络列表失败: ${err.message}`)
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
