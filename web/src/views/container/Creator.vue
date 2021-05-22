<template>
  <page-header-wrapper>
    <a-card class="card" :bordered="false">
      <a-form layout="horizontal" >
        <a-form-item label="名称" required>
          <a-input name="name" v-model="formData.Name" placeholder="例如 myContainer" />
        </a-form-item>
        <a-divider orientation="left">
          镜像选择
        </a-divider>
        <a-form-item label="镜像" required>
          <a-alert message="当前无镜像可选，请在镜像管理中创建镜像" type="warning" style="margin-bottom: 5px" :hidden="noImages" closable show-icon />
          <a-select
            name="image"
            placeholder="请选择镜像"
            @change="imageChange"
          >
            <a-select-option v-for="i in formValue.imageList" :key="i.Id" :value="i.RepoTags[0]">
              {{ i.RepoTags[0] }}
            </a-select-option>
          </a-select>
        </a-form-item>
        <a-divider orientation="left">
          网络端口配置
        </a-divider>
        <a-form-item label="将所有容器开放端口发布到随机的主机端口">
          <a-switch name="publishPortRandom" v-model="formData.HostConfig.PublishAllPorts"/>
        </a-form-item>
        <a-form-item label="手动设置开放端口">
          <div v-for="(item, index) in formValue.ports" :key="index">
            <a-input
              v-model="item.host"
              addon-before="Host"
              style="width: 30%; margin-right: 15px"
              placeholder="例如 8888"
            />
            <a-input
              v-model="item.container"
              addon-before="Container"
              style="width: 30%; margin-right: 15px"
              placeholder="例如 80"
            />
            <a-radio-group
              v-model="item.protocol"
              default-value="tcp"
              button-style="solid"
              style="width: 15%; margin-right: 15px"
            >
              <a-radio-button value="tcp">
                TCP
              </a-radio-button>
              <a-radio-button value="udp">
                UDP
              </a-radio-button>
            </a-radio-group>
            <a-button
              type="danger"
              shape="round"
              icon="delete"
              title="删除此项"
              @click="delPort(index)"
            />
          </div>
          <a-button
            type="round"
            @click="addPort">
            <a-icon type="plus" /> 添加新的端口绑定
          </a-button>
        </a-form-item>
        <a-divider orientation="left">
          其他动作
        </a-divider>
        <a-form-item label="自动移除">
          <a-switch name="autoRemove" v-model="formData.HostConfig.AutoRemove"/>
        </a-form-item>
      </a-form>
    </a-card>
    <a-card class="card" title="高级选项" :bordered="false">
      <a-tabs default-active-key="1" @change="tabChange">
        <a-tab-pane key="1" tab="启动配置">
          <a-form >
            <a-form-item label="自动重启策略">
              <a-radio-group v-model="formData.HostConfig.RestartPolicy.Name" default-value="no" button-style="solid">
                <a-radio-button value="no">
                  从不
                </a-radio-button>
                <a-radio-button value="always">
                  总是
                </a-radio-button>
                <a-radio-button value="on-failure">
                  非正常退出时
                </a-radio-button>
                <a-radio-button value="unless-stopped">
                  未手动停止时
                </a-radio-button>
              </a-radio-group>
            </a-form-item>
            <a-form-item label="环境变量">
              <div v-for="(item, index) in formData.Env" :key="index">
                <a-input v-model="formData.Env[index]" style="width: 38%; margin-right: 15px" placeholder="例如 AmIAHandsomeMan=yeah" />
                <a-button slot="addonAfter" type="danger" shape="round" icon="delete" title="删除此项" @click="delEnv(index)" />
              </div>
              <a-button type="round" @click="addEnv">
                <a-icon type="plus" /> 添加环境变量
              </a-button>
            </a-form-item>
            <a-form-item label="启动命令">
              <a-input name="command" v-model="formValue.commands" placeholder="例如 '-logtostderr' '--housekeeping_interval=5s' or /usr/bin/nginx -t -c /mynginx.conf"/>
            </a-form-item>
            <a-form-item label="入口">
              <a-input name="entrypoint" v-model="formData.Entrypoint" placeholder="例如 /bin/sh -c"/>
            </a-form-item>
            <a-form-item label="工作目录">
              <a-input name="workingDir" v-model="formData.WorkingDir" placeholder="例如 /myapp"/>
            </a-form-item>
            <a-form-item label="用户">
              <a-input name="entrypoint" v-model="formData.User" placeholder="例如 nginx"/>
            </a-form-item>
            <a-form-item label="控制台">
              <a-radio-group v-model="formValue.consoles" @change="radioChange">
                <a-radio value="it">
                  交互的TTY (-i -t)
                </a-radio>
                <a-radio value="i">
                  可交互 (-i)
                </a-radio>
                <a-radio value="t">
                  TTY (-t)
                </a-radio>
                <a-radio value="none">
                  无
                </a-radio>
              </a-radio-group>
            </a-form-item>
          </a-form>
        </a-tab-pane>
        <a-tab-pane key="2" tab="卷配置" force-render>
          <a-form>
            <a-form-item label="卷映射">
              <a-alert message="当前无卷可选，请在卷管理中创建新卷" type="warning" style="margin-bottom: 5px" :hidden="noVolumes" closable show-icon />
              <div v-for="(item, index) in formValue.volumeMap" :key="index">
                <a-input v-model="item.container" addon-before="容器位置" style="width:38%; margin-right: 15px" placeholder="例如 /path/to/container" />
                <a-input-group style="width: 38%" compact>
                  <a-button style="background-color: #F5F5F5">卷</a-button>
                  <a-select style="width:66%; margin-right: 15px" @change="volumeChange($event, index)" placeholder="请选择一个卷">
                    <a-select-option v-for="vol in formValue.volumeList" :key="vol.Name" :value="vol.Name">
                      {{ vol.Name }}
                    </a-select-option>
                  </a-select>
                </a-input-group>
                <a-button type="danger" shape="round" icon="delete" title="删除此项" @click="delVolumeMap(index)" />
              </div>
              <a-button type="round" @click="addVolumeMap">
                <a-icon type="plus" /> 添加卷映射
              </a-button>
            </a-form-item>
          </a-form>
        </a-tab-pane>
        <a-tab-pane key="3" tab="网络配置" force-render>
          <a-form>
            <a-form-item label="网络">
              <a-select @change="networkChange" placeholder="请选择一个网络">
                <a-select-option v-for="net in formValue.networkList" :key="net.Id" :value="net.Name">
                  {{ net.Name }}
                </a-select-option>
              </a-select>
            </a-form-item>
            <a-form-item label="主机名(Hostname)">
              <a-input name="hostname" v-model="formData.Hostname" placeholder="例如 first" />
            </a-form-item>
            <a-form-item label="域名(Domain Name)">
              <a-input name="domainName" v-model="formData.Domainname" placeholder="例如 example.com" />
            </a-form-item>
            <a-form-item label="Mac地址">
              <a-input name="mac" v-model="formData.MacAddress" placeholder="例如 12-34-56-78-9a-bc" />
            </a-form-item>
            <a-form-item label="IPv4地址">
              <a-input name="ipv4" v-model="formValue.IPv4Address" placeholder="例如 172.20.0.7" />
            </a-form-item>
            <a-form-item label="IPv6地址">
              <a-input name="ipv6" v-model="formValue.IPv6Address" placeholder="例如 a:b:c:d::1234" />
            </a-form-item>
          </a-form>
        </a-tab-pane>
      </a-tabs>
    </a-card>

    <!-- fixed footer toolbar -->
    <footer-tool-bar :is-mobile="false" :collapsed="true">
      <span class="popover-wrapper">
        <!-- <a-popover title="表单校验信息" overlayClassName="antd-pro-pages-forms-style-errorPopover" trigger="click" :getPopupContainer="trigger => trigger.parentNode">
          <template slot="content">
            <li v-for="item in errors" :key="item.key" @click="scrollToField(item.key)" class="antd-pro-pages-forms-style-errorListItem">
              <a-icon type="cross-circle-o" class="antd-pro-pages-forms-style-errorIcon" />
              <div class="">{{ item.message }}</div>
              <div class="antd-pro-pages-forms-style-errorField">{{ item.fieldLabel }}</div>
            </li>
          </template>
          <span class="antd-pro-pages-forms-style-errorIcon" v-if="errors.length > 0">
            <a-icon type="exclamation-circle" />{{ errors.length }}
          </span>
        </a-popover> -->
      </span>
      <a-button type="primary" @click="submit" :loading="loading">部署</a-button>
    </footer-tool-bar>
  </page-header-wrapper>
</template>

<script>
import { VolumeList } from '@/api/volumes'
import { NetworkList } from '@/api/networks'
import { ImageList } from '@/api/images'
import { ContainerCreate } from '@/api/containers'
import FooterToolBar from '@/components/FooterToolbar'
import { baseMixin } from '@/store/app-mixin'
import splitargs from 'splitargs'

export default {
  name: 'Creator',
  mixins: [baseMixin],
  components: {
    FooterToolBar
  },
  data () {
    return {
      loading: false,
      noImages: false,
      noVolumes: false,
      formValue: {
        alwaysPull: false, //
        ports: [],
        consoles: 'none',
        volumeMap: [],
        volumeList: [],
        networkList: [],
        imageList: [],
        commands: '', //
        IPv4Address: '',
        IPv6Address: ''
      },
      formData: {
        Name: '',
        Image: null, // image
        ExposedPorts: {}, //
        Env: [], // env
        Cmd: [], // command
        Entrypoint: '',
        WorkingDir: '',
        User: '',
        Hostname: '',
        Domainname: '',
        MacAddress: '',
        Tty: false, // -t
        OpenStdin: false, // -i
        HostConfig: {
          Binds: [], // volumes
          PortBindings: [],
          PublishAllPorts: false,
          AutoRemove: false,
          RestartPolicy: {
            Name: 'no'
          },
          NetworkMode: ''
        },
        Volumes: {}, // volumes
        NetworkingConfig: {
          EndpointsConfig: {}
        }
      }
    }
  },
  created () {
    this.initForm()
  },
  watch: {
    // '$store.state.isConnected': function (val) {
    //   if (val === true) {
    //     console.log('init start.')
    //     setTimeout(() => {
    //         this.initForm()
    //     }, 1000)
    //   }
    // }
  },
  computed: {},
  methods: {
    initForm () {
        NetworkList()
        .then((res) => {
          this.formValue.networkList = res.data
        })
        .catch((err) => {
          this.$message.error(`获取网络列表失败: ${err.message}`)
        })
      VolumeList()
        .then((res) => {
          this.formValue.volumeList = res.data.Volumes
          if (res.data.Volumes.length !== 0) {
            this.noVolumes = true
          }
        })
        .catch((err) => {
          this.$message.error(`获取卷列表失败: ${err.message}`)
        })
      ImageList()
        .then((res) => {
          // Vue.set(this.formValue, 'imageList', res.data)
          // this.formValue.imageList = [...res.data]
          this.formValue.imageList = res.data
          if (res.data.length !== 0) {
            this.noImages = true
          }
        })
        .catch((err) => {
          this.$message.error(`获取镜像列表失败: ${err.message}`)
        })
        console.log('init end.')
    },
    // 提交动作
    submit () {
      this.praseCommands(this.formValue.commands)
      this.praseVolumeMaps(this.formValue.volumeMap)
      this.praseConsoleConfigs(this.formValue.consoles)
      this.prasePortBindings(this.formValue.ports)
      this.praseNetworkConfigs()
      ContainerCreate(this.formData)
        .then((res) => {
          console.log(res)
        })
        .catch((err) => {
          console.log(err)
        })
      console.log(this.formData)
    },
    // 表单数据整理
    praseNetworkConfigs () {
      var net = this.formData.HostConfig.NetworkMode
      this.formData.NetworkingConfig.EndpointsConfig[net] = {
        IPAMConfig: {
          IPv4Address: this.formValue.IPv4Address,
          IPv6Address: this.formValue.IPv6Address
        }
      }
    },
    praseCommands (val) {
      this.formData.Cmd = splitargs(val)
    },
    praseConsoleConfigs (val) {
      var openStdin = true
      var tty = true
      if (val === 't') {
        openStdin = false
      } else if (val === 'i') {
        tty = false
      } else if (val === 'none') {
        openStdin = false
        tty = false
      }
      this.formData.OpenStdin = openStdin
      this.formData.Tty = tty
    },
    praseVolumeMaps (maps) {
      var binds = []
      var volumes = {}

      maps.forEach(function (map) {
        var volume = map.volume
        var container = map.container
        if (volume && container) {
          var bind = volume + ':' + container
          volumes[container] = {}
          // if (map.readOnly) {
          //   bind += ':ro'
          // }
          binds.push(bind)
        }
      })
      this.formData.HostConfig.Binds = binds
      this.formData.Volumes = volumes
    },
    prasePortBindings (ports) {
      var expose = {}
      var bindings = {}
      ports.forEach(function (p) {
        if (p.host && p.container) {
          const key = p.container + '/' + p.protocol
          expose[key] = {}
          bindings[key] = [{ HostPort: p.host }]
        }
      })
      this.formData.ExposedPorts = expose
      this.formData.HostConfig.PortBindings = bindings
    },

    // 表单动态增减
    addPort () {
      this.formValue.ports.push({
        host: '',
        container: '',
        protocol: 'tcp'
      })
    },
    delPort (index) {
      this.formValue.ports.splice(index, 1)
    },
    addEnv () {
      this.formData.Env.push('')
    },
    delEnv (index) {
      this.formData.Env.splice(index, 1)
    },
    addVolumeMap () {
      this.formValue.volumeMap.push({
        container: '',
        volume: ''
      })
    },
    delVolumeMap (index) {
      this.formValue.volumeMap.splice(index, 1)
    },
    radioChange (e) {
      // console.log(e)
    },
    tabChange (e) {
      // console.log(e)
    },

    // 下拉框数据选择
    imageChange (val) {
      this.formData.Image = val
    },
    networkChange (val) {
      this.formData.HostConfig.NetworkMode = val
    },
    volumeChange (e, index) {
      this.formValue.volumeMap[index].volume = e
      console.log(e, index)
      console.log(this.formValue.volumeMap)
    }
  }
}
</script>

<style lang="less" scoped>
  .card{
    margin-bottom: 24px;
  }
  .popover-wrapper {
    /deep/ .antd-pro-pages-forms-style-errorPopover .ant-popover-inner-content {
      min-width: 256px;
      max-height: 290px;
      padding: 0;
      overflow: auto;
    }
  }
  .antd-pro-pages-forms-style-errorIcon {
    user-select: none;
    margin-right: 24px;
    color: #f5222d;
    cursor: pointer;
    i {
          margin-right: 4px;
    }
  }
  .antd-pro-pages-forms-style-errorListItem {
    padding: 8px 16px;
    list-style: none;
    border-bottom: 1px solid #e8e8e8;
    cursor: pointer;
    transition: all .3s;
    &:hover {
      background: #e6f7ff;
    }
    .antd-pro-pages-forms-style-errorIcon {
      float: left;
      margin-top: 4px;
      margin-right: 12px;
      padding-bottom: 22px;
      color: #f5222d;
    }
    .antd-pro-pages-forms-style-errorField {
      margin-top: 2px;
      color: rgba(0,0,0,.45);
      font-size: 12px;
    }
  }
</style>
