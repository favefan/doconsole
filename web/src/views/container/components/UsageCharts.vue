<template>
  <div>
    <a-card title="更新设置">
      <a-row>
        <a-col :span="6">刷新间隔(秒)</a-col>
        <a-col :span="18">
          <a-select v-model="refreshRateChange" size="small" style="width: 33%">
            <a-select-option value="1000">1</a-select-option>
            <a-select-option value="3000">3</a-select-option>
            <a-select-option value="5000">5</a-select-option>
          </a-select>
        </a-col>
      </a-row>
    </a-card>
    <a-row style="margin-top: 25px">
      <a-col :sm="24" :md="8">
        <a-card title="CPU利用率">
          <area-chart v-bind="cpu" />
        </a-card>
      </a-col>
      <a-col :sm="24" :md="8">
        <a-card title="内存利用率">
          <area-chart v-bind="mem" />
        </a-card>
      </a-col>
      <a-col :sm="24" :md="8">
        <a-card title="网络利用率(总计)">
          <area-chart v-bind="net" />
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script>
// import Vue from 'vue'
import _ from 'lodash-es'
import moment from 'moment'
// import { Area } from '@antv/g2plot'
import { AreaChart } from '@opd/g2plot-vue'
import { ContainerStats } from '@/api/containers'

// Vue.use(AreaChart)

export default {
  name: 'UsageCharts',
  props: {
    cid: {
      type: String,
      required: true,
      default: null
    }
  },
  components: {
    AreaChart
  },
  data () {
    return {
      data: [],
      tryTime: 3,
      refreshRateChange: 3000,
      yMax: 30,
      // preTotalTX: 0,
      // preTotalRX: 0,
      cpu: {
        height: 233,
        xField: 'label',
        yField: 'value',
        yAxis: {
          max: 1
        },
        meta: {
          value: {
            formatter: (val) => {
              return val * 100 + '%'
            }
          }
        },
        xAxis: false,
        smooth: true,
        data: []
      },
      mem: {
        height: 233,
        xField: 'label',
        yField: 'value',
        seriesField: 'type',
        meta: {
          value: {
            formatter: (val) => {
              return val / 1000000 + 'MB'
            }
          }
        },
        xAxis: false,
        smooth: true,
        data: []
      },
      net: {
        height: 233,
        xField: 'label',
        yField: 'value',
        seriesField: 'type',
        meta: {
          value: {
            formatter: (val) => {
              return val / 1000000 + 'MB'
            }
          }
        },
        xAxis: false,
        smooth: true,
        data: []
      }
    }
  },
  mounted () {
  },
  watch: {
    data () {
      this.timer(this.refreshRateChange)
    }
  },
  created () {
    this.updateStats()
  },
  destroyed () {
    clearTimeout(this.timer)
  },
  methods: {
    timer (rate) {
      return setTimeout(() => {
        this.updateStats()
      }, rate)
    },
    updateStats () {
      var _this = this
      ContainerStats(_this.cid)
        .then((res) => {
          _this.data.push(res.data)
          const stats = _this.containerStatsPrase(res.data)
          _this.updateCPUChart(stats)
          _this.updateNetworkChart(stats)
          _this.updateMemoryChart(stats)
        })
        .catch((err) => {
          _this.tryTime -= 1
          if (_this.tryTime === 0) {
            _this.$message.error(`容器状态获取失败: ${err.message}`)
            _this.tryTime = 3
          }
        })
    },
    updateNetworkChart (stats) {
      var _this = this
      if (stats.Networks.length > 0) {
        var rx = stats.Networks[0].rx_bytes
        var tx = stats.Networks[0].tx_bytes
        var label = moment(stats.read).format('HH:mm:ss')
        _this.net.data.push({ label: label, value: rx, type: 'rx' })
        _this.net.data.push({ label: label, value: tx, type: 'tx' })
        if (_this.net.data.length > _this.yMax * 2) {
          _this.net.data.splice(0, 2)
        }
      }
    },

    updateMemoryChart (stats) {
      var _this = this
      var label = moment(stats.read).format('HH:mm:ss')
      this.mem.data.push({ label: label, value: stats.MemoryUsage, type: '使用率' })
      this.mem.data.push({ label: label, value: stats.MemoryCache, type: '缓存' })
      if (_this.mem.data.length > _this.yMax * 2) {
        _this.mem.data.splice(0, 2)
      }
    },

    updateCPUChart (stats) {
      var _this = this
      var label = moment(stats.read).format('HH:mm:ss')
      var value = stats.isWindows ? this.calculateCPUPercentWindows(stats) : this.calculateCPUPercentUnix(stats)
      this.cpu.data.push({ label: label, value: value })
      if (_this.cpu.data.length > _this.yMax) {
        _this.cpu.data.splice(0, 1)
      }
    },

    calculateCPUPercentUnix (stats) {
      var cpuPercent = 0.0
      var cpuDelta = stats.CurrentCPUTotalUsage - stats.PreviousCPUTotalUsage
      var systemDelta = stats.CurrentCPUSystemUsage - stats.PreviousCPUSystemUsage

      if (systemDelta > 0.0 && cpuDelta > 0.0) {
        cpuPercent = (cpuDelta / systemDelta) * stats.CPUCores * 100.0
      }

      return cpuPercent
    },

    calculateCPUPercentWindows (stats) {
      var possIntervals =
        stats.NumProcs * parseFloat(moment(stats.read, 'YYYY-MM-DDTHH:mm:ss.SSSSSSSSSZ').valueOf() - moment(stats.preread, 'YYYY-MM-DDTHH:mm:ss.SSSSSSSSSZ').valueOf())
      var windowsCpuUsage = 0.0
      if (possIntervals > 0) {
        windowsCpuUsage = parseFloat(stats.CurrentCPUTotalUsage - stats.PreviousCPUTotalUsage) / parseFloat(possIntervals * 100)
      }
      return windowsCpuUsage
    },
    containerStatsPrase (data) {
      var stats = {}
      stats.read = data.read
      stats.preread = data.preread
      if (data.memory_stats.privateworkingset !== undefined) {
        // Windows
        stats.MemoryUsage = data.memory_stats.privateworkingset
        stats.MemoryCache = 0
        stats.NumProcs = data.num_procs
        stats.isWindows = true
      } else {
        // Linux
        if (data.memory_stats.stats === undefined || data.memory_stats.usage === undefined) {
          stats.MemoryUsage = stats.MemoryCache = 0
        } else {
          stats.MemoryUsage = data.memory_stats.usage - data.memory_stats.stats.cache
          stats.MemoryCache = data.memory_stats.stats.cache
        }
      }
      stats.PreviousCPUTotalUsage = data.precpu_stats.cpu_usage.total_usage
      stats.PreviousCPUSystemUsage = data.precpu_stats.system_cpu_usage
      stats.CurrentCPUTotalUsage = data.cpu_stats.cpu_usage.total_usage
      stats.CurrentCPUSystemUsage = data.cpu_stats.system_cpu_usage
      if (data.cpu_stats.cpu_usage.percpu_usage) {
        stats.CPUCores = data.cpu_stats.cpu_usage.percpu_usage.length
      }
      stats.Networks = _.values(data.networks)
      return stats
    }
  }
}
</script>

<style lang="less" scoped>
.action-button-list * {
  margin-left: 10px;
}
</style>
