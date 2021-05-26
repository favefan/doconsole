<template>
  <div class="dashboard-cards">
    <a-page-header
      style="border: 1px solid rgb(235, 237, 240)"
      title="系统信息"
      :ghost="false"
      :sub-title="info.SystemTime | timeFormat"
    >
      <div class="content">
        <div class="main">
          <a-descriptions size="small" :column="3">
            <a-descriptions-item label="操作系统">
              {{ info.OperatingSystem || 'None' }}
            </a-descriptions-item>
            <a-descriptions-item label="系统版本">
              {{ info.OSVersion || 'None' }}
            </a-descriptions-item>
            <a-descriptions-item label="系统类型">
              {{ info.OSType || 'None' }}
            </a-descriptions-item>
            <a-descriptions-item label="内核">
              {{ info.KernelVersion || 'None' }}
            </a-descriptions-item>
            <a-descriptions-item label="体系架构">
              {{ info.Architecture || 'None' }}
            </a-descriptions-item>
          </a-descriptions>
          <a-descriptions size="small" :column="2" style="border-top: 1px solid rgb(234, 234, 234); margin-top: 5px; padding-top: 10px">
            <a-descriptions-item label="Docker版本">
              {{ info.ServerVersion || 'None' }}
            </a-descriptions-item>
            <a-descriptions-item label="DockerHub镜像">
              <code v-for="m in info.RegistryConfig.Mirrors" :key="m">
                {{ m }}
              </code>
            </a-descriptions-item>
          </a-descriptions>
        </div>
        <div class="extra">
          <div style="display: flex; width: max-content; justifyContent: flex-end">
            <a-statistic title="CPU" prefix="core" :value="info.NCPU" style="margin-right: 32px"/>
            <a-statistic title="运行内存" prefix="GB" :value="(info.MemTotal / 1000000000).toFixed(2)" />
          </div>
        </div>
      </div>
    </a-page-header>
    <a-row>
      <a-col :sm="24" :md="12">
        <a-card :bordered="false" @click="goTo('container')" hoverable>
          <a-card-meta title="容器">
            <img
              slot="avatar"
              :style="{ height: imageSize }"
              :src="containersPNG"
            />
            <div slot="description">
              <span style="font-size: 188%; font-weight: 600">{{ containers.length + ' ' }}</span>
              <span style="color: #10D269">{{ containers | runningcontainers }} 运行中 </span>
              <span style="color: #FF4500">{{ containers | stoppedcontainers }} 已停止 </span>
              <span style="color: #00BFFF">{{ containers | healthycontainers }} 健康 </span>
              <span style="color: #FFD700">{{ containers | unhealthycontainers }} 不健康</span>
            </div>
          </a-card-meta>
        </a-card>
      </a-col>
      <a-col :sm="24" :md="12">
        <a-card :bordered="false" @click="goTo('image')" hoverable>
          <a-card-meta title="镜像">
            <img
              slot="avatar"
              :style="{ height: imageSize }"
              :src="imagesPNG"
            />
            <div slot="description">
              <span style="font-size: 188%; font-weight: 600">{{ images.length + ' ' }}</span>
              <span style="color: #10D269">总大小: {{ images | imagestotalsize }}</span>
            </div>
          </a-card-meta>
        </a-card>
      </a-col>
    </a-row>
    <a-row>
      <a-col :sm="24" :md="12">
        <a-card :bordered="false" @click="goTo('network')" hoverable>
          <a-card-meta title="网络">
            <img
              slot="avatar"
              :style="{ height: imageSize }"
              :src="networksPNG"
            />
            <div slot="description">
              <span style="font-size: 188%; font-weight: 600">{{ networkCount }}</span>
            </div>
          </a-card-meta>
        </a-card>
      </a-col>
      <a-col :sm="24" :md="12">
        <a-card :bordered="false" @click="goTo('volume')" hoverable>
          <a-card-meta title="卷">
            <img
              slot="avatar"
              :style="{ height: imageSize }"
              :src="volumesPNG"
            />
            <div slot="description">
              <span style="font-size: 188%; font-weight: 600">{{ volumeCount }}</span>
            </div>
          </a-card-meta>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script>
// import moment from 'moment'
import filesize from 'filesize'
import { baseMixin } from '@/store/app-mixin'
import { ContainerList } from '@/api/containers'
import { ImageList } from '@/api/images'
import { VolumeList } from '@/api/volumes'
import { NetworkList } from '@/api/networks'
import { Info } from '@/api/docker'

import containersPNG from '../../assets/icons/containers.png'
import imagesPNG from '../../assets/icons/images.png'
import networksPNG from '../../assets/icons/networks.png'
import volumesPNG from '../../assets/icons/volumes.png'

export default {
  name: 'Index',
  mixins: [baseMixin],
  components: {
  },
  data () {
    return {
      imageSize: '50px',
      containersPNG,
      imagesPNG,
      networksPNG,
      volumesPNG,
      containers: [],
      images: [],
      volumeCount: 0,
      networkCount: 0,
      info: {
        SystemTime: '',
        OperatingSystem: '',
        Architecture: '',
        OSVersion: '',
        OSType: '',
        KernelVersion: '',
        RegistryConfig: { Mirrors: [] }
      }
    }
  },
  computed: {

  },
  created () {
    this.initView()
  },
  filters: {
    timeFormat: function (value) {
      var byT = (value || '').split('T')
      var byD = (byT[1] || '').split('.')
      return byT[0] + ' ' + byD[0]
    },
    runningcontainers: function (value) {
      return value.filter(function (container) {
        return container.State === 'running'
      }).length
    },
    stoppedcontainers: function (value) {
      return value.filter(function (container) {
        return container.State === 'exited'
      }).length
    },
    healthycontainers: function (value) {
      return value.filter(function (container) {
        return container.State === 'healthy'
      }).length
    },
    unhealthycontainers: function (value) {
      return value.filter(function (container) {
        return container.State === 'unhealthy'
      }).length
    },
    imagestotalsize: function (images) {
      var totalImageSize = 0
      for (var i = 0; i < images.length; i++) {
        var item = images[i]
        totalImageSize += item.VirtualSize
      }
      return filesize(totalImageSize, { base: 10, round: 1 })
    }
  },
  methods: {
    initView () {
      ContainerList()
      .then((res) => {
        this.containers = res.data
      })
      .catch((err) => {
        this.$message.error(`获取容器数据失败: ${err.response}`)
      })
      ImageList()
        .then((res) => {
          this.images = res.data
        })
        .catch((err) => {
          this.$message.error(`获取镜像数据失败: ${err.response}`)
        })
      VolumeList()
        .then((res) => {
          this.volumeCount = res.data.Volumes.length
        })
        .catch((err) => {
          this.$message.error(`获取卷数据失败: ${err.response}`)
        })
      NetworkList()
        .then((res) => {
          this.networkCount = res.data.length
        })
        .catch((err) => {
          this.$message.error(`获取网络数据失败: ${err.response}`)
        })
      Info()
        .then((res) => {
          this.info = res.data
        })
        .catch((err) => {
          this.$message.error(`获取系统信息失败: ${err.response}`)
        })
    },
    goTo (tag) {
      switch (tag) {
        case 'container':
            this.$router.push({ path: `/container/containers` })
            break
        case 'image':
            this.$router.push({ path: `/image/images` })
            break
        case 'network':
            this.$router.push({ path: `/network/networks` })
            break
        case 'volume':
            this.$router.push({ path: `/volume/volumes` })
            break
        default:
            this.$message.warning('Target error!')
      }
    }
  }
}
</script>

<style lang="less" scoped>
.dashboard-cards /deep/ .ant-card, .ant-page-header {
  margin: 5px
}
tr:last-child td {
  padding-bottom: 0;
}
.content {
  display: flex;
}
.ant-statistic-content {
  font-size: 20px;
  line-height: 28px;
}
@media (max-width: 576px) {
  .content {
    display: block;
  }

  .main {
    width: 100%;
    margin-bottom: 12px;
  }

  .extra {
    width: 100%;
    margin-left: 0;
    text-align: left;
  }
}
</style>
