<template>
  <div>
    连接主机：
    <a-select :value="this.$store.state.currentHost" @change="switchHost" style="width: 188px; margin-right: 18px">
      <a-icon slot="suffixIcon" :type="IconType" theme="twoTone" :two-tone-color="IconColor" :spin="isSpin"/>
      <a-select-option v-for="host in this.$store.state.hostList" :key="host.Id" :value="host.Name">
        {{ host.Name }}
      </a-select-option>
    </a-select>
  </div>
</template>

<script>
import { HostList } from '@/api/hosts'
import { DockerClientCreate } from '@/api/docker'

export default {
  name: 'HostSelection',
  props: {
  },
  data () {
    return {
      // showMenu: true,
      // currentUser: {}
      // HostSelectionList: [],
      // DefaultSelection: '',
      IconType: 'question',
      IconColor: '#10D269',
      isSpin: false
    }
  },
  created () {
    HostList()
      .then((res) => {
        this.$store.commit('updateHostList', res.data)
        if (res.data.length === 0) {
          this.$router.push({ path: '/boot' })
        } else {
          var isInit = false
          for (const i of res.data) {
            // console.log(i)
            if (i.Name === this.$store.state.currentHost) {
              this.initDockerClient(i)
              isInit = true
            }
          }
          if (isInit !== true) {
            this.initDockerClient(res.data[0])
            this.$message.warning('无默认主机，已自动设置')
          }
        }
      })
      .catch((err) => {
        this.$message.error(`主机选择器: 加载主机列表失败 -> ${err.message}`)
      })
  },
  methods: {
    initDockerClient (host) {
      // console.log(host)
      this.IconType = 'sync'
      this.isSpin = true
      DockerClientCreate(host)
        .then((res) => {
          setTimeout(() => {
            this.isSpin = false
            this.IconType = 'check-circle'
            this.IconColor = '#10D269'
            this.$message.success(`已连接到主机: ${res.data}`)
          }, 1000)
          this.$store.commit('changeHost', res.data)
          this.$store.commit('updateContent')
        })
        .catch((err) => {
          this.isSpin = false
          this.IconType = 'close-circle'
          this.IconColor = '#FF0000'
          this.$message.error(`连接主机失败: ${err.message}`)
          this.$store.commit('changeHost', host.Name)
        })
    },
    switchHost (e) {
      var host = {}
      this.IconType = 'sync'
      this.isSpin = true
      for (const i of this.$store.state.hostList) {
        if (i.Name === e) {
          host = i
        }
      }
      this.initDockerClient(host)
    }
  }
}
</script>
