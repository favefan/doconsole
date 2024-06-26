<template>
  <page-header-wrapper>
    <a-card
      style="margin-top: 24px"
      :bordered="false"
    >
      <div slot="title">
        <a-popconfirm
          placement="top"
          title="确定删除选中的仓库?"
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

      <!-- 仓库列表 -->
      <a-table
        rowKey="Id"
        :pagination="false"
        :columns="columns"
        :data-source="listData"
        :row-selection="rowSelection"
      >
        <a-tag
          color="cyan"
          slot="issystem"
          slot-scope="text, record"
          v-if="record.Name === 'DockerHub'"
        >
          官方
        </a-tag>
        <span slot="action" slot-scope="text, record">
          <template>
            <a-button
              type="primary"
              :icon="record.Name === 'DockerHub' ? 'key' : 'edit'"
              @click="showDrawer(record)"
              style="margin-left: 10px"
            >
              {{ record.Name === 'DockerHub' ? '认证' : '修改' }}
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
      <a-form :label-col="{ span: 5 }" :wrapper-col="{ span: 19 }" >
        <a-form-item label="名称">
          <a-input name="Name" v-model="formData.Name" placeholder="例如 myRegistry" :disabled="formData.Name === 'DockerHub' ? true : false" required/>
        </a-form-item>
        <a-form-item label="地址">
          <a-input name="URL" v-model="formData.URL" placeholder="例如 example.com:12345" :disabled="formData.Name === 'DockerHub' ? true : false" required/>
        </a-form-item>
        <a-form-item label="启用认证">
          <a-switch name="NeedAuth" v-model="formData.NeedAuth"/>
        </a-form-item>
        <a-form-item label="用户名">
          <a-input name="username" v-model="formData.Username" :disabled="formData.NeedAuth ? false : true"/>
        </a-form-item>
        <a-form-item label="密码">
          <a-input name="password" type="password" v-model="formData.Password" :disabled="formData.NeedAuth ? false : true"/>
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
import { RegistryList, RegistryCreate, RegistriesRemove, RegistryUpdate } from '@/api/registries'

const columns = [
  // {
  //   title: 'Id',
  //   dataIndex: 'Id',
  //   scopedSlots: { customRender: 'Id' }
  // },
  {
    title: '名称',
    dataIndex: 'Name'
  },
  {
    title: '',
    // dataIndex: 'Name',
    scopedSlots: { customRender: 'issystem' }
  },
  {
    title: '地址',
    dataIndex: 'URL'
  },
  {
    title: '认证',
    dataIndex: 'NeedAuth',
    customRender: (is) => is === true ? '是' : '否'
  },
  {
    title: '创建时间',
    dataIndex: 'CreatedAt',
    customRender: (datetime) => datetime.split('T')[0] + ' ' + datetime.split('T')[1].split('.')[0]
  },
  {
    title: ' ',
    dataIndex: 'action',
    width: '150px',
    scopedSlots: { customRender: 'action' }
  }
]

export default {
  name: 'Registries',
  components: {

  },
  data () {
    this.columns = columns
    return {
      formLayout: 'horizontal',
      // form: this.$form.createForm(this, { name: 'registrycreate' }),
      formData: {},
      listData: [],
      registryListData: [],
      drawerVisible: false,
      selectedRowKeys: [],
      deleteLoading: false,
      filterKey: '',
      drawerTitle: '创建仓库',
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
      var list = this.registryListData
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
            disabled: record.Name === 'DockerHub'
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
    // 列表更新
    freshList () {
      RegistryList()
        .then((res) => {
          this.registryListData = res.data
          this.listData = this.registryListData
          this.filterKey = ''
        })
        .catch((err) => {
          this.$message.error(`更新仓库列表失败: ${err.message}`)
        })
    },

    filterChange () {
      var key = this.filterKey.trim()
      if (key.length === 0) {
        this.listData = this.registryListData
      } else {
        this.listData = this.filterList
      }
    },

    // 多选操作
    onSelectChange (selectedRowKeys) {
      // // console.log('selectedRowKeys changed: ', selectedRowKeys)
      this.selectedRowKeys = selectedRowKeys
    },
    // 多选删除
    batchRemove () {
      this.deleteLoading = true
      RegistriesRemove({ 'array': this.selectedRowKeys })
        .then((res) => {
          // console.log(res.data)
          if (res.data !== null) {
            this.$message.warning(`成功删除 ${res.data.length} 个仓库，删除失败的仓库有: ${res.data}`)
          } else {
            this.$message.success(`成功删除 ${this.selectedRowKeys.length} 个仓库`)
          }
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
        .catch((err) => {
          this.$message.error(`仓库删除错误: ${err.message}`)
          this.selectedRowKeys.splice(0, this.selectedRowKeys.length)
          this.deleteLoading = false
          this.freshList()
        })
    },

    // 镜像详情跳转
    handleToInformation (id) {
      this.$router.push({ path: `/registry/information?name=${id}` })
    },

    // 抽屉 - 搜索/添加 镜像
    showDrawer (e) {
      console.log(e)
      if (typeof (e) !== 'undefined') {
        this.drawerType = 'update'
        this.drawerTitle = e.Name === 'DockerHub' ? 'DockerHub认证' : '修改仓库信息'
        this.formData = {
        Id: e.Id,
        Name: e.Name,
        URL: e.URL,
        NeedAuth: e.NeedAuth,
        Username: e.Username,
        Password: e.Password,
        Comment: ''
      }
      }
      this.drawerVisible = true
    },

    drawerOnClose () {
      this.drawerVisible = false
      this.drawerType = 'create'
      this.drawerTitle = '创建仓库'
      this.formDataInit()
    },

    // switchAuth (value) {
    //   // console.log(value)
    // },

    drawerSubmit (e) {
      e.preventDefault()
      // console.log(this.formData)
      if (this.drawerType === 'create' && this.formData.Name === 'DockerHub') {
        this.$message.error('DockerHub是保留值,请输入其他名称')
      } else if (this.formData.Name.trim().length === 0 || this.formData.URL.trim().length === 0) {
        this.$message.warning('名称和地址字段不能为空, 且不能包含空格, 请重新输入')
      } else if (this.formData.NeedAuth === true &&
        (this.formData.Username.split(' ').length !== 1 ||
          this.formData.Username.trim().length === 0 ||
            this.formData.Passsword === '')) {
        this.$message.warning('启用认证时，用户名和地址字段不能为空, 且不能包含空格, 请重新输入')
      } else if (this.drawerType === 'create') {
        RegistryCreate(this.formData)
          .then((res) => {
            this.$message.success('创建仓库成功')
            // console.log(res)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
          .catch((err) => {
            this.$message.error(`创建仓库列表失败: ${err.message}`)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
      } else {
        RegistryUpdate(this.formData)
          .then((res) => {
            this.$message.success('更新仓库成功')
            // console.log(res)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
          .catch((err) => {
            this.$message.error(`更新仓库失败: ${err.message}`)
            this.freshList()
            this.drawerVisible = false
            this.formDataInit()
          })
      }
    },

    formDataInit () {
      this.formData = {
        Id: 0,
        Name: '',
        URL: '',
        NeedAuth: false,
        Username: '',
        Password: '',
        Comment: ''
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
