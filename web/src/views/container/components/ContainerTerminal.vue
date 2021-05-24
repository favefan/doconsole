<template>
  <div>
    <div style="widh: 100%; margin-bottom: 15px">
      以用户 <code>{{ connectinfo.terminalUser }}</code> 执行 <code>{{ connectinfo.terminalCommand }}</code> 命令
      <a-button type="danger" size="small" @click="disconnect" style="margin-left: 15px">断开连接</a-button>
    </div>
    <div id="terminal" ref="terminal" style="width: 100%" />
  </div>
</template>

<script>
import 'xterm/css/xterm.css'
import { Terminal } from 'xterm'
import { WebLinksAddon } from 'xterm-addon-web-links'
import { AttachAddon } from 'xterm-addon-attach'
import { FitAddon } from 'xterm-addon-fit'
// import { SearchAddon } from 'xterm-addon-search'

export default {
  name: 'ContainerTerminal',
  data () {
    return {
      socket: null,
      terminal: null // 保存terminal实例
    }
  },
  props: {
    execid: {
      type: String,
      required: true,
      default: ''
    },
    connectinfo: {
      type: Object,
      required: true,
      default: null
    }
  },
  created () {
    // const webSocket = new WebSocket('ws://localhost:8192/api/v1/containers/exec/' + this.execid)
    // const attachAddon = new AttachAddon(webSocket)
    // this.terminal.loadAddon(attachAddon)
  },
  // mounted
  mounted () {
    this.terminal = new Terminal({
        useStyle: true,
        cursorBlink: true,
        rows: 35
    })
    this.terminal.loadAddon(new WebLinksAddon())
    const fitAddon = new FitAddon()
    this.terminal.loadAddon(fitAddon)
    // 这里的wx链接就是上面后端执行接口
    this.socket = new WebSocket('ws://localhost:8192/api/v1/containers/exec/' + this.execid)
    const attachAddon = new AttachAddon(this.socket)
    this.terminal.loadAddon(attachAddon)
    // const searchAddon = new SearchAddon()
    // terminal.loadAddon(searchAddon)

    this.terminal.open(document.getElementById('terminal'))

    fitAddon.fit()
  },

  methods: {
    disconnect () {
      var _this = this
      if (_this.socket) {
        _this.socket.close()
        _this.$emit('disconnect')
      }
    }
  }
}
</script>
