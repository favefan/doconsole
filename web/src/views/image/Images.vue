<template>
  <page-header-wrapper>
    <a-card style="margin-top: 24px" :bordered="false" :title="'共' + imagesCount + '个镜像 / 占用空间 ' + (imagesTotalSize / 1000000.0).toFixed(2) + ' MB'">
      <div slot="extra">
        <a-tooltip placement="top" class="image-use-statu">
          <template slot="title">
            <span>27.94MB/27.94MB</span>
          </template>
          <a-progress :percent="100" :show-info="false" size="small" />
          <div style="color: #87d068; float: left">
            使用中
          </div>
          <div style="color: gray; float: right">
            未使用
          </div>
        </a-tooltip>
        <a-input-search style="margin-left: 10px; width: 272px;" />
      </div>

      <a-table
        rowKey="Id"
        :pagination="false"
        :columns="columns"
        :data-source="data"
        :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
      >
        <a slot="name" slot-scope="text, record" @click="handleToInformation(record.Id.split(':')[1])">
          {{ record.RepoTags[0].split(':')[0] }}
        </a>
        <span slot="isinuse" slot-scope="flag">
          <a-tag
            :color="flag === -1 ? 'gray' : '#10D269'"
          >
            {{ flag === -1 ? 'UN USE' : 'IN USE' }}
          </a-tag>
        </span>
        <span slot="createdTime" slot-scope="timestamp">
          {{ timestamp | formatDate }}
        </span>
        <span slot="action" slot-scope="text, record">
          <template>
            <a-popconfirm
              placement="top"
              title="确定删除这个镜像?"
              ok-text="是"
              cancel-text="我再想想"
              @confirm="delete(item.Id)"
              @cancel="cancel"
            >
              <a-button type="circle" size="large" icon="delete" title="删除"></a-button>
            </a-popconfirm>
            <a-button
              type="circle"
              size="large"
              icon="caret-right"
              @click="start(record.Id)"
              title="用此镜像创建新容器"
              style="margin-left: 10px"
            ></a-button>
          </template>
        </span>
      </a-table>
    </a-card>
  </page-header-wrapper>
</template>

<script>
import { STable, Ellipsis } from '@/components'
// 演示如何使用 this.$dialog 封装 modal 组件
import TaskForm from './modules/TaskForm'
import Info from './components/Info'

import {
  ImageList
} from '@/api/images'

const columns = [
  {
    title: ' ',
    dataIndex: 'Name',
    scopedSlots: { customRender: 'name' }
    // customRender: (repotag) => repotag.split('@')[0]
  },
  {
    title: ' ',
    dataIndex: 'Containers',
    scopedSlots: { customRender: 'isinuse' }
    // customRender: (isinuse) => isinuse === -1 ? 'UN USE' : 'IN USE'
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
    Ellipsis,
    TaskForm,
    Info
  },
  data () {
    this.columns = columns
    return {
      data: [],
      imagesCount: 0,
      imagesTotalSize: 0,
      status: 'all',
      selectedRowKeys: []
    }
  },
  filters: {
    // stateFilter (state) {
    //   // console.log(originVal)
    //   if (state === 'running') {
    //     return 'pause'
    //   } else {
    //     return 'caret-right'
    //   }
    // }
    formatDate: function (value) {
        const date = new Date(parseInt(value) * 1000)
        const y = date.getFullYear()
        let MM = date.getMonth() + 1
        MM = MM < 10 ? ('0' + MM) : MM
        let d = date.getDate()
        d = d < 10 ? ('0' + d) : d
        let h = date.getHours() + 8
        h = h < 10 ? ('0' + h) : h
        let m = date.getMinutes()
        m = m < 10 ? ('0' + m) : m
        let s = date.getSeconds()
        s = s < 10 ? ('0' + s) : s
        return y + '-' + MM + '-' + d + ' ' + h + ':' + m + ':' + s
      }
  },
  created () {
    this.freshList()
  },
  computed: {
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
  methods: {
    freshList () {
      ImageList()
        .then((res) => {
          this.data = res.data
          this.imagesCount = res.data.length
          this.imagesTotalSize = res.data.reduce((sum, e) => sum + Number(e.Size), 0)
        })
        .catch((err) => {
          this.$message.error(`更新镜像列表失败: ${err.message}`)
        })
    },
    onSelectChange (selectedRowKeys) {
      console.log('selectedRowKeys changed: ', selectedRowKeys)
      this.selectedRowKeys = selectedRowKeys
    },
    handleToInformation (id) {
      this.$router.push({ path: `/image/information?id=${id}` })
    },
    add () {
      this.$dialog(
        TaskForm,
        // component props
        {
          record: {},
          on: {
            ok () {
              console.log('ok 回调')
            },
            cancel () {
              console.log('cancel 回调')
            },
            close () {
              console.log('modal close 回调')
            }
          }
        },
        // modal props
        {
          title: '新增',
          width: 700,
          centered: true,
          maskClosable: false
        }
      )
    },
    cancel (e) {
      console.log(e)
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

.image-use-statu{
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
