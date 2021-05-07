<template>
  <page-header-wrapper v-if="loadData" :title="loadData.Name">
    <template v-slot:content>
      <a-descriptions size="default" :column="2">
        <a-descriptions-item label="Id">
          <ellipsis :length="15" tooltip>
            {{ loadData.Id }}
          </ellipsis>
        </a-descriptions-item>
        <a-descriptions-item label="Driver">{{ loadData.Driver }}</a-descriptions-item>
        <a-descriptions-item label="Scope">{{ loadData.Scope }}</a-descriptions-item>
        <a-descriptions-item label="创建时间" v-if="loadData.Created">{{ loadData.Created.split('.')[0] }}</a-descriptions-item>
        <a-descriptions-item label="IPvSubnet" v-for="item in loadData.IPAM.Config" :key="item.Subnet">{{ item.Subnet }}</a-descriptions-item>
        <a-descriptions-item label="IPvGateway" v-for="item in loadData.IPAM.Config" :key="item.Subnet">{{ item.Gateway }}</a-descriptions-item>
      </a-descriptions>
    </template>

    <!-- OperationButtonGroup -->
    <template v-slot:extra>
      <a-popconfirm
        placement="top"
        title="确定删除这个网络?"
        ok-text="是"
        cancel-text="我再想想"
        :disabled="loadData.Name === 'bridge' || loadData.Name === 'none' || loadData.Name === 'host'"
        @confirm="removeNetwork"
      >
        <a-button size="large" type="danger" shape="circle" icon="delete" title="删除容器"
                  :disabled="loadData.Name === 'bridge' || loadData.Name === 'none' || loadData.Name === 'host'"
        />
      </a-popconfirm>
    </template>

    <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col :xs="8" :sm="8">
          <div class="text">Attachable</div>
          <div class="heading">{{ loadData.Attachable }}</div>
        </a-col>
        <a-col :xs="8" :sm="8">
          <div class="text">Internal</div>
          <div class="heading">{{ loadData.Internal }}</div>
        </a-col>
        <a-col :xs="8" :sm="8">
          <div class="text">Ingress</div>
          <div class="heading">{{ loadData.Ingress }}</div>
        </a-col>
      </a-row>
    </template>

    <a-card style="margin-top: 24px" :bordered="false" title="Network options">
      <a-list size="small" v-if="loadData.Options">
        <a-list-item v-for="(value, name) in loadData.Options" :key="name">
          <a-list-item-meta>
            <p slot="title">{{ name }}</p>
          </a-list-item-meta>
          <div>{{ value }}</div>
        </a-list-item>
      </a-list>
    </a-card>

  </page-header-wrapper>
</template>

<script>
// import { baseMixin } from '@/store/app-mixin'
import { Ellipsis } from '@/components'
import { NetworkInspect, NetworksRemove } from '@/api/networks'

export default {
  name: 'ImageInformation',
  // mixins: [baseMixin],
  components: {
        Ellipsis
  },
  data () {
    return {
      loadData: {
        IPAM: {
          Config: []
        },
        ConfigFrom: {
        },
        Options: {}
      }
    }
  },
  created () {
    NetworkInspect(this.$route.query.id)
      .then((res) => {
        this.loadData = res.data
      })
      .catch((err) => {
        this.$message.error(`获取网络信息失败: ${err.message}`)
      })
  },
  filters: {

  },
  methods: {
    removeNetwork () {
      NetworksRemove({ array: [this.loadData.Id] })
        .then((res) => {
        this.$message.success('删除成功')
        this.$router.push('/network/networks')
      })
      .catch((err) => {
        this.$message.error(`获取网络信息失败: ${err.message}`)
      })
    }
  }
}
</script>

<style lang="less" scoped>
  .detail-layout {
    margin-left: 44px;
  }
  .text {
    font-size: 120%;
    color: rgba(0, 0, 0, .45);
  }

  .heading {
    color: rgba(0, 0, 0, .85);
    font-size: 25px;
  }

  .no-data {
    color: rgba(0, 0, 0, .25);
    text-align: center;
    line-height: 64px;
    font-size: 16px;

    i {
      font-size: 24px;
      margin-right: 16px;
      position: relative;
      top: 3px;
    }
  }

  .mobile {
    .detail-layout {
      margin-left: unset;
    }
    .status-list {
      text-align: right;
    }
  }
</style>
