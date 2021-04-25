<template>
  <page-header-wrapper v-if="loadData.RepoTags" :title="loadData.RepoTags[0].split(':')[0]">
    <template v-slot:content>
      <a-descriptions size="default" :column="1">
        <a-descriptions-item label="Id">{{ loadData.Id }}</a-descriptions-item>
        <a-descriptions-item label="镜像大小">{{ (loadData.Size / 1000000.0).toFixed(2) + " MB" }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ loadData.Created }}</a-descriptions-item>
        <a-descriptions-item label="无">
          <a href="">无</a>
        </a-descriptions-item>
      </a-descriptions>
    </template>

    <!-- OperationButtonGroup -->
    <template v-slot:extra>
      <a-button size="large" shape="circle" icon="export" title="导出容器" />
      <a-button size="large" type="danger" shape="circle" icon="delete" title="删除容器" />
      <a-button size="large" type="primary" shape="circle" icon="caret-right" title="创建容器" />
    </template>

    <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col :xs="12" :sm="12">
          <div class="text">Tag</div>
          <div class="heading">Latest</div>
        </a-col>
        <a-col :xs="12" :sm="12">
          <div class="text">使用状态</div>
          <div class="heading">被使用</div>
        </a-col>
      </a-row>
    </template>

    <a-card style="margin-top: 24px" :bordered="false" title="Image history">
      <div class="no-data"><a-icon type="frown-o"/>暂无数据</div>
    </a-card>

  </page-header-wrapper>
</template>

<script>
import { baseMixin } from '@/store/app-mixin'
import { Ellipsis } from '@/components'
import { ImageInspect } from '@/api/images'

export default {
  name: 'ImageInformation',
  mixins: [baseMixin],
  components: {
        Ellipsis
  },
  data () {
    return {
      loadData: {
        ContainerConfig: {},
        Config: {
          ExposedPorts: {},
          Labels: {}
        },
        GraphDriver: {
          Data: {}
        },
        RootFS: {},
        Metadata: {}
      }
    }
  },
  created () {
    ImageInspect(this.$route.query.id)
      .then((res) => {
        this.loadData = res.data
      })
      .catch((err) => {
        this.$message.error(`获取镜像信息失败: ${err.message}`)
      })
  },
  filters: {

  },
  methods: {

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
