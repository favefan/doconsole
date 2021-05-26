<template>
  <page-header-wrapper
    class="container-name"
    :title="data.Name"
    :tab-list="tabList"
    :tab-active-key="tabActiveKey"
    @tabChange="handleTabChange"
  >

    <template v-slot:content>
      <a-descriptions size="default" :column="1"> <!-- isMobile ? 1 : 2 -->
        <a-descriptions-item label="Id">{{ data.Id | shortString }}</a-descriptions-item>
        <a-descriptions-item label="镜像">
          <a @click="() => {$router.push({path: `/image/information?id=${data.Image}`})}">
            {{ data.Config.Image }} @{{ (data.Image || '').split(':')[1] | shortString }}
          </a>
        </a-descriptions-item>
        <a-descriptions-item label="IP地址">{{ data.NetworkSettings.IPAddress || '无' }}</a-descriptions-item>
        <a-descriptions-item label="创建时间">{{ data.Created | timeFormat }}</a-descriptions-item>
        <a-descriptions-item label="启动时间">{{ data.State.StartedAt | timeFormat }}</a-descriptions-item>
        <a-descriptions-item label="结束时间">{{ data.State.FinishedAt | timeFormat }}</a-descriptions-item>
      </a-descriptions>
    </template>

    <!-- OperationButtonGroup -->
    <template v-slot:extra>
      <operation-button-group :data="{ State: data.State.Status, Id: data.Id }" :call="initContainerDetail" :outpage="true"></operation-button-group>
    </template>

    <!-- Status under OperationButtonGroup -->
    <template v-slot:extraContent>
      <a-row class="status-list">
        <a-col>
          <div class="text">状态</div>
          <div class="heading">{{ data.State | getstatetext }}
            <span v-if="!data.State.Running && data.State.Status !== 'created'">, 退出代码 {{ data.State.ExitCode }}</span>
          </div>
        </a-col>
        <a-col style="margin-top: 10px">
          <div class="text">PID</div>
          <div class="heading">{{ data.State.Pid }}
          </div>
        </a-col>
      </a-row>
    </template>

    <div
      v-show="tabActiveKey === 'detail'"
    >
      <a-card class="card-row" :bordered="false" title="容器信息">
        <a-row>
          <a-col :span="rowHeadSpan">
            端口配置
          </a-col>
          <a-col :span="rowContentSpan">
            <div v-for="portMapping in portBindings" :key="portMapping.host"> {{ portMapping.host }} <a-icon type="arrow-right" /> {{ portMapping.container }} </div>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="rowHeadSpan">
            CMD
          </a-col>
          <a-col :span="rowContentSpan">
            <code>{{ data.Config.Cmd | command }}</code>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="rowHeadSpan">
            入口 (Entrypoint)
          </a-col>
          <a-col :span="rowContentSpan">
            <code>{{ data.Config.Entrypoint | command }}</code>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="rowHeadSpan">
            环境变量 (Env)
          </a-col>
          <a-col :span="rowContentSpan">
            <code v-for="env in data.Config.Env" :key="env">{{ env }} </code>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="rowHeadSpan">
            重启策略
          </a-col>
          <a-col :span="rowContentSpan">
            <a-select v-model="restartPolicy" size="small" style="width: 80%" @change="restartPolicyChange">
              <a-select-option value="no">
                不重启 (no)
              </a-select-option>
              <a-select-option value="on-failure">
                发生错误时 (on-failure)
              </a-select-option>
              <a-select-option value="always">
                总是 (always)
              </a-select-option>
              <a-select-option value="unless-stopped">
                手动停止后 (unless-stopped)
              </a-select-option>
            </a-select>
            <a-button style="margin-left: 15px" size="small" type="primary" @click="updateRestartPolicy(data.Id)" :loading="updateRPLoading">更新</a-button>
          </a-col>
        </a-row>
      </a-card>

      <a-card style="margin-top: 24px" :bordered="false" title="卷信息">
        <table class="volume-table table">
          <thead>
            <tr>
              <th>Host/volume</th>
              <th>Path in container</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="vol in data.Mounts" :key="vol.Name">
              <td v-if="vol.Type === 'bind'">{{ vol.Source }}</td>
              <td v-if="vol.Type === 'volume'">
                <a @click="() => {$router.push({path: `/volume/information?name=${vol.Name}`})}">{{ vol.Name }}</a>
              </td>
              <td>{{ vol.Destination }}</td>
            </tr>
          </tbody>
        </table>
      </a-card>

      <a-card style="margin-top: 24px" :bordered="false" title="加入的网络">
        <a-row style="margin-bottom: 10px">
          <a-col :span="6" style="text-align: right; padding-right: 15px">
            加入一个网络
          </a-col>
          <a-col :span="12">
            <a-select size="small" style="width: 100%" placeholder="选择一个网络" @change="SelectNetwork">
              <a-select-option v-for="net in networkList" :key="net.Id" :value="net.Id">
                {{ net.Name }}
              </a-select-option>
            </a-select>
          </a-col>
          <a-col :span="6">
            <a-button
              style="margin-left: 15px"
              size="small"
              type="primary"
              @click="joinNetwork(data.Id)"
              :loading="joinNetworkLoading"
              :disabled="selectedNetwork === ''"
            >加入网络</a-button>
          </a-col>
        </a-row>
        <a-row>
          <table class="network-table table">
            <thead>
              <tr>
                <th>网络</th>
                <th>IP地址</th>
                <th>网关</th>
                <th>MAC地址</th>
                <th>操作</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(v, k) in data.NetworkSettings.Networks" :key="k">
                <td><a @click="() => {$router.push({path: `/network/information?id=${v.NetworkID}`})}">{{ k }}</a></td>
                <td>{{ v.IPAddress || '-' }}</td>
                <td>{{ v.Gateway || '-' }}</td>
                <td>{{ v.MacAddress || '-' }}</td>
                <td>
                  <a-button
                    type="danger"
                    size="small"
                    @click="leaveNetwork(v.NetworkID, data.Id)"
                    :loading="leaveNetworkLoading"
                  >
                    离开网络
                  </a-button>
                </td>
              </tr>
            </tbody>
          </table>
        </a-row>
      </a-card>

    </div>

    <div
      v-show="tabActiveKey === 'logs'"
    >
      <a-card :bordered="false" >
        <a-row>
          <a-col>
            <pre class="log_viewer">{{ logs }}</pre>
          </a-col>
        </a-row>
      </a-card>
    </div>

    <div
      v-show="tabActiveKey === 'stats'"
    >
      <usage-charts :cid="this.$route.query.id" />
    </div>

    <div
      v-show="tabActiveKey === 'cli'"
    >
      <!-- 终端连接配置 -->
      <a-card class="terminal-tab" :bordered="false" title="终端连接配置">
        <a-row>
          <a-col :span="5">选择终端</a-col>
          <a-col :span="19">
            <a-select v-model="terminalConfig.terminalCommand" size="small" style="width: 33%" :disabled="terminalButtonDisabled">
              <a-select-option value="/bin/bash" v-if="data.Platform == 'linux'">/bin/bash</a-select-option>
              <a-select-option value="/bin/sh" v-if="data.Platform == 'linux'">/bin/sh</a-select-option>
              <a-select-option value="powershell" v-if="data.Platform == 'windows'">powershell</a-select-option>
              <a-select-option value="cmd.exe" v-if="data.Platform == 'windows'">cmd.exe</a-select-option>
            </a-select>
          </a-col>
        </a-row>
        <a-row>
          <a-col :span="5">访问用户 (User)</a-col>
          <a-col :span="19">
            <a-input v-model="terminalConfig.terminalUser" size="small" placeholder="root" :disabled="terminalButtonDisabled"/>
          </a-col>
        </a-row>
        <a-row>
          <a-button type="primary" size="small" @click="terminalConnect" :disabled="terminalButtonDisabled">连接</a-button>
        </a-row>
      </a-card>
      <!-- 终端 -->
      <a-card v-if="terminalRun" :bordered="false" style="margin-top: 25px" title="容器终端">
        <container-terminal :execid="terminalExecID" :connectinfo="terminalConfig" @disconnect="terminalDisconnect"></container-terminal>
      </a-card>
    </div>

  </page-header-wrapper>
</template>

<script>
import _ from 'lodash-es'
import { baseMixin } from '@/store/app-mixin'
import { ContainerInspect, ContainerUpdate, ContainerExecCreate, ContainerLogs } from '@/api/containers'
import { NetworkList, NetworkConnect, NetworkDisconnect } from '@/api/networks'
import OperationButtonGroup from './components/OperationButtonGroup'
import ContainerTerminal from './components/ContainerTerminal'
import UsageCharts from './components/UsageCharts'

export default {
  name: 'ContainerInformation',
  mixins: [baseMixin],
  components: {
    OperationButtonGroup,
    ContainerTerminal,
    UsageCharts
  },
  data () {
    return {
      rowHeadSpan: 6,
      rowContentSpan: 18,
      logs: null,
      data: {
        Args: [],
        State: {},
        HostConfig: {
          RestartPolicy: {
            Name: ''
          }
        },
        Mounts: [],
        Config: {},
        NetworkSettings: {}
      },
      tabList: [
        { key: 'detail', tab: '详细信息' },
        { key: 'logs', tab: '日志' },
        { key: 'stats', tab: '资源监控' },
        { key: 'cli', tab: '远程连接' }
        // { key: 'attach', tab: '输出跟踪' }
      ],
      tabActiveKey: 'detail',
      portBindings: [],
      networkList: [],
      selectedNetwork: '',
      restartPolicy: '',
      updateRPLoading: false,
      joinNetworkLoading: false,
      leaveNetworkLoading: false,
      terminalConfig: {
        terminalCommand: '/bin/bash',
        terminalUser: 'root'
      },
      terminalRun: false,
      terminalExecID: '',
      terminalButtonDisabled: false
    }
  },
  created () {
    this.initContainerDetail()
    this.initConnectedNetworksList()
  },
  filters: {
    getstatetext: function (value) {
      if (value === undefined) {
        return '无'
      }
      if (value.Dead) {
        return 'Dead'
      }
      if (value.Ghost && value.Running) {
        return 'Ghost'
      }
      if (value.Running && value.Paused) {
        return '运行中(暂停)'
      }
      if (value.Running) {
        return '运行中'
      }
      if (value.Status === 'created') {
        return '已创建'
      }
      return '已停止'
    },
    timeFormat: function (value) {
      var byT = (value || '').split('T')
      var byD = (byT[1] || '').split('.')
      return byT[0] + ' ' + byD[0]
    },
    shortString: function (value) {
      return (value || '').slice(0, 12)
    },
    command: function (value) {
      if (value) {
        return value.join(' ')
      }
    }
  },
  methods: {
    initContainerDetail () {
      var _this = this
      ContainerInspect(this.$route.query.id)
        .then((res) => {
          this.data = res.data
          this.portBindings = []
          if (_this.data.NetworkSettings.Ports) {
            _.forEach(Object.keys(_this.data.NetworkSettings.Ports), function (key) {
              if (_this.data.NetworkSettings.Ports[key]) {
                _.forEach(_this.data.NetworkSettings.Ports[key], (portMapping) => {
                  const mapping = {}
                  mapping.container = key
                  mapping.host = `${portMapping.HostIp}:${portMapping.HostPort}`
                  _this.portBindings.push(mapping)
                })
              }
            })
          }
          _this.restartPolicy = _this.data.HostConfig.RestartPolicy.Name
        })
        .catch((err) => {
          this.$message.error(`获取容器信息失败: ${err.message}`)
        })
    },
    initConnectedNetworksList () {
      NetworkList()
        .then((res) => {
          this.networkList = res.data
        })
        .catch((err) => {
          this.$message.error(`获取网络列表失败: ${err.message}`)
        })
    },
    getContainerLogs () {
      ContainerLogs(this.data.Id)
        .then((res) => {
          this.logs = res.data
        })
        .catch((err) => {
          this.$message.error(`获取日志失败: ${err.message}`)
        })
    },
    handleTabChange (key) {
      var _this = this
      this.tabActiveKey = key
      if (key === 'logs') {
        _this.getContainerLogs()
      }
    },
    restartPolicyChange (val) {
      this.restartPolicy = val
    },
    updateRestartPolicy (id) {
      this.updateRPLoading = true
      ContainerUpdate(id, { RestartPolicy: this.restartPolicy })
        .then((res) => {
          this.$message.success(`更新容器重启策略成功`)
        })
        .catch((err) => {
          this.$message.error(`更新容器重启策略失败: ${err.message}`)
        })
        .finally(() => {
          this.updateRPLoading = false
          this.initContainerDetail()
        })
    },
    SelectNetwork (val) {
      this.selectedNetwork = val
    },
    joinNetwork (containerID) {
      this.joinNetworkLoading = true
      NetworkConnect(this.selectedNetwork, containerID)
        .then((res) => {
          this.$message.success(`加入网络成功`)
        })
        .catch((err) => {
          this.$message.error(`加入网络失败: ${err.message}`)
        })
        .finally(() => {
          this.joinNetworkLoading = false
          this.initContainerDetail()
        })
    },
    leaveNetwork (networkID, containerID) {
      this.leaveNetworkLoading = true
      NetworkDisconnect(networkID, containerID)
        .then((res) => {
          this.$message.success(`离开网络成功`)
        })
        .catch((err) => {
          this.$message.error(`离开网络失败: ${err.message}`)
        })
        .finally(() => {
          this.leaveNetworkLoading = false
          this.initContainerDetail()
        })
    },
    // 终端操作
    terminalConnect () {
      var execConfig = {
        User: this.terminalConfig.terminalUser,
        Tty: true,
        AttachStdin: true,
        AttachStderr: true,
        AttachStdout: true,
        Cmd: [this.terminalConfig.terminalCommand]
      }
      ContainerExecCreate(this.data.Id, execConfig)
        .then((res) => {
          this.terminalExecID = res.data.Id
          this.$message.success(`终端进程创建成功`)
          this.terminalRun = true
          this.terminalButtonDisabled = true
        })
        .catch((err) => {
          this.$message.error(`终端进程创建错误: ${err.message}`)
        })
    },
    terminalDisconnect () {
      this.terminalExecID = ''
      this.terminalConfig = {
        terminalCommand: '/bin/bash',
        terminalUser: 'root'
      }
      this.terminalRun = false
      this.terminalButtonDisabled = false
    }
  }
}
</script>

<style lang="less" scoped>

  .container-name /deep/ .ant-page-header-heading-title {
    display: block;
    float: left;
    margin-bottom: 5px;
    padding-right: 12px;
    color: rgba(0, 0, 0, 0.85);
    font-weight: 600;
    font-size: 35px;
    line-height: 45px;
  }

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

  .card-row /deep/ .ant-row {
    border-bottom: 1px solid rgb(234, 234, 234);
    padding: 15px 0 15px;
  }

  .table {
    width: 100%;
  }
  // .table tr {
  //   height: 30px;
  // }
  .table tr,th,td {
    padding: 10px 0 10px;
  }
  .volume-table th {
    width: 50%;
    font-size: 110%;
    border-bottom: 2px solid rgb(234, 234, 234);
  }
  .volume-table td {
    width: 50%;
  }
  .network-table th {
    border-bottom: 2px solid rgb(234, 234, 234);
  }

  // log-view
  .log_viewer {
    height: 666px;
    overflow-y: scroll;
    color: black;
    font-size: 0.85em;
    background-color: white;
  }

  .log_viewer.wrap_lines {
    white-space: pre-wrap;
  }

  .log_viewer .inner_line {
    padding: 0 15px;
    cursor: pointer;
    margin-bottom: 0;
  }

  .log_viewer .inner_line:empty::after {
    content: '.';
    visibility: hidden;
  }

  .log_viewer .line_selected {
    background-color: #c5cae9;
  }

  .terminal-tab /deep/ .ant-row {
    margin-bottom: 15px;
  }
</style>
