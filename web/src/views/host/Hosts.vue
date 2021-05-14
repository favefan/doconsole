<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
    >
      <div slot="title">
        <a-popconfirm
          placement="top"
          title="确定删除选中的主机?"
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
          @click="showDrawer()"
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

      <!-- 主机列表 -->
      <a-table
        rowKey="Id"
        :pagination="false"
        :columns="columns"
        :data-source="listData"
        :row-selection="rowSelection"
      >
        <span slot="action" slot-scope="text, record">
          <template>
            <a-button
              type="primary"
              icon="edit"
              @click="showDrawer(record)"
              style="margin-left: 10px"
            >
              修改
            </a-button>
          </template>
        </span>
      </a-table>
    </a-card>

    <!-- 添加抽屉 -->
    <a-drawer
      :title="drawerTitle"
      placement="right"
      :width="521"
      :visible="drawerVisible"
      :destroyOnClose="true"
      :maskClosable="true"
      @close="drawerOnClose"
    >
      <a-alert
        message="重要提示"
        description="您可以通过Socket或TCP连接到指定环境Docker。
          你可以在Docker文档中找到更多关于如何通过TCP公开Docker API的信息。
          使用socket时请确保Docker启动flag包含如下
           -v /var/run/docker.sock:/var/run/docker.sock (在Linux中) 或
           -v \.\pipe\docker_engine:\.\pipe\docker_engine (在Windows中)."
        type="info"
        close-text="关闭"
        style="margin-bottom: 15px"
      />
      <a-form :label-col="{ span: 5 }" :wrapper-col="{ span: 19 }" >
        <a-form-item label="名称">
          <a-input name="Name" v-model="formData.Name" placeholder="例如 myHost" />
        </a-form-item>
        <a-form-item label="通过Socket">
          <a-switch name="ViaSocket" v-model="formData.ViaSocket"/>
        </a-form-item>
        <a-form-item label="Endpoint URL">
          <a-input name="DockerEngineURL" v-model="formData.DockerEngineURL" placeholder="例如 10.0.0.1:12345 或 example.com:12345" :disabled="formData.ViaSocket"/>
        </a-form-item>
        <a-form-item label="外部IP">
          <a-input name="HostIP" v-model="formData.HostIP" placeholder="例如  10.0.0.1 或 example.com" />
        </a-form-item>
        <a-form-item label="启用TLS连接">
          <a-switch name="TLS" v-model="formData.TLS"/>
        </a-form-item>
        <a-form-item :wrapper-col="{ span: 12, offset: 5 }">
          <a-button type="primary" @click="drawerSubmit">
            提交
          </a-button>
        </a-form-item>
      </a-form>
    </a-drawer>
  </page-header-wrapper>
</template>

<script>
import { HostList, HostCreate, HostsRemove, HostUpdate } from '@/api/hosts'

const columns = [
  {
    title: '名称',
    dataIndex: 'Name'
  },
  {
    title: '通过Socket连接',
    dataIndex: 'ViaSocket',
    customRender: (is) => is === false ? 'false' : 'true'
  },
  {
    title: 'Docker Endpoint URL',
    dataIndex: 'DockerEngineURL'
  },
  {
    title: '外部IP',
    dataIndex: 'HostIP'
  },
  {
    title: '启用TLS连接',
    dataIndex: 'TLS',
    customRender: (is) => is === false ? 'false' : 'true'
  },
  {
    title: '',
    dataIndex: 'action',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  name: 'Hosts',
  components: {

  },
  data () {
    this.columns = columns
    return {
      formLayout: 'horizontal',
      // form: this.$form.createForm(this, { name: 'hostcreate' }),
      formData: {},
      listData: [],
      hostListData: [],
      drawerVisible: false,
      selectedRowKeys: [],
      deleteLoading: false,
      filterKey: '',
      drawerTitle: '创建主机',
      drawerType: 'create'
    }
  },
  filters: {
  },
  created () {
    this.freshList()
    this.formDataInit()
  },
  computed: {
    filterList: function () {
      var key = this.filterKey
      var list = this.hostListData
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
            disabled: record.Name === this.$store.state.currentHost
          }
        })
      }
    },
    hasSelected () {
      return this.selectedRowKeys.length > 0
    }
  },
  watch: {
  },
  methods: {
    formDataInit () {
      this.formData = {
        Id: 0,
        Name: '',
        ViaSocket: false,
        DockerEngineURL: '',
        HostIP: '',
        TLS: false,
        Default: false
      }
    },
    // 列表更新
    freshList () {
      HostList()
        .then((res) => {
          this.hostListData = res.data
          this.$store.commit('updateHostList', res.data)
          this.listData = this.hostListData
          this.filterKey = ''
        })
        .catch((err) => {
          this.$message.error(`更新主机列表失败: ${err.message}`)
        })
    },

    filterChange () {
      var key = this.filterKey.trim()
      if (key.length === 0) {
        this.listData = this.hostListData
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
      this.deleteLoading = true
      HostsRemove({ 'array': this.selectedRowKeys })
        .then((res) => {
          // console.log(res.data)
          if (res.data !== null) {
            this.$message.warning(`成功删除 ${res.data.length} 个主机，删除失败的主机有: ${res.data}`)
          } else {
            this.$message.success(`成功删除 ${this.selectedRowKeys.length} 个主机`)
          }
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`主机删除错误: ${err.message}`)
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
    },

    // 抽屉 - 创建/修改
    showDrawer (e) {
      console.log(e)
      if (typeof (e) !== 'undefined') {
        this.drawerType = 'update'
        this.drawerTitle = '修改主机信息'
        this.formData = {
          Id: e.Id,
          Name: e.Name,
          ViaSocket: e.ViaSocket,
          DockerEngineURL: e.DockerEngineURL,
          HostIP: e.HostIP,
          TLS: e.TLS,
          Default: false
      }
      }
      this.drawerVisible = true
    },

    drawerOnClose () {
      this.drawerVisible = false
      this.drawerType = 'create'
      this.drawerTitle = '创建主机'
      this.formDataInit()
    },

    // switchAuth (value) {
    //   // console.log(value)
    // },

    drawerSubmit (e) {
      e.preventDefault()
      // console.log(this.formData)
      if (this.formData.Name.trim().length === 0) {
        this.$message.warning('名称不能为空, 且不能包含空格, 请重新输入')
      } else if (this.formData.ViaSocket === false &&
        (this.formData.DockerEngineURL.trim().length === 0)) {
        this.$message.warning('启用Socket连接时，Docker Endpoint URL不能为空, 请重新输入')
      } else if (this.drawerType === 'create') {
        HostCreate(this.formData)
          .then((res) => {
            this.$message.success('创建主机成功')
            // console.log(res)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
          .catch((err) => {
            this.$message.error(`创建主机列表失败: ${err.message}`)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
      } else {
        HostUpdate(this.formData)
          .then((res) => {
            this.$message.success('更新主机成功')
            // console.log(res)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
          .catch((err) => {
            this.$message.error(`更新主机失败: ${err.message}`)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
      }
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
