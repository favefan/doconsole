<template>
  <page-header-wrapper v-if="loadData" :title="loadData.Name">
    <template v-slot:content>
      <a-descriptions size="default" :column="2">
        <a-descriptions-item label="Driver">{{ loadData.Driver }}</a-descriptions-item>
        <a-descriptions-item label="Scope">{{ loadData.Scope }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ loadData.CreatedAt }}</a-descriptions-item>
        <a-descriptions-item label="挂载点" v-if="loadData.Mountpoint">{{ loadData.Mountpoint }}</a-descriptions-item>
      </a-descriptions>
    </template>

    <!-- OperationButtonGroup -->
    <template v-slot:extra>
      <a-popconfirm
        placement="top"
        title="确定删除这个卷?"
        ok-text="是"
        cancel-text="我再想想"
        @confirm="removeVolume"
      >
        <a-button size="large" type="danger" shape="circle" icon="delete" title="删除容器"
        />
      </a-popconfirm>
    </template>

    <!-- <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col :xs="8" :sm="8">
          <div class="text">Attachable</div>
          <div class="heading">{{ loadData.Attachable }}</div>
        </a-col>
      </a-row>
    </template> -->

    <a-card style="margin-top: 24px" :bordered="false" title="Labels">
      <a-list size="small">
        <a-list-item v-for="(value, name) in loadData.Labels" :key="name">
          <a-list-item-meta>
            <p slot="title">{{ name }}</p>
          </a-list-item-meta>
          <div>{{ value }}</div>
        </a-list-item>
      </a-list>
    </a-card>

    <a-card style="margin-top: 24px" :bordered="false" title="Options">
      <a-list size="small">
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
// import { Ellipsis } from '@/components'
import { VolumeInspect, VolumesRemove } from '@/api/volumes'

export default {
  name: 'ImageInformation',
  // mixins: [baseMixin],
  components: {
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
    VolumeInspect(this.$route.query.name)
      .then((res) => {
        this.loadData = res.data
      })
      .catch((err) => {
        this.$message.error(`获取卷信息失败: ${err.message}`)
      })
  },
  filters: {

  },
  methods: {
    removeVolume () {
      VolumesRemove({ array: [this.loadData.Name] })
        .then((res) => {
        this.$message.success('删除成功')
        this.$router.push('/volume/volumes')
      })
      .catch((err) => {
        this.$message.error(`获取卷信息失败: ${err.message}`)
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
