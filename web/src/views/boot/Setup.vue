<template>
  <div class="main">
    <a-alert
      message="添加主机"
      description="请先添加主机信息连接Docker Engine, 以开始我们的工作"
      type="info"
      close-text="关闭"
      style="margin-bottom: 10px"
    />
    <a-form
      id="formSetup"
      class="boot-layout-setup"
      ref="formSetup"
    >
      <a-alert v-if="isSetupError" type="error" showIcon style="margin-bottom: 24px;" :message="message" />
      <a-form-item label="名称">
        <a-input name="Name" v-model="formData.Name" placeholder="例如 myHost" />
      </a-form-item>
      <a-form-item label="通过Socket (本地客户端)">
        <a-switch name="ViaSocket" v-model="formData.ViaSocket"/>
      </a-form-item>
      <a-form-item label="Endpoint URL">
        <a-input name="DockerEngineURL"
                 v-model="formData.DockerEngineURL"
                 placeholder="例如 10.0.0.1:12345 或 example.com:12345"
                 :disabled="formData.ViaSocket"/>
      </a-form-item>
      <a-form-item label="外部IP">
        <a-input name="HostIP" v-model="formData.HostIP" placeholder="例如 10.0.0.1 或 example.com" />
      </a-form-item>
      <a-form-item label="启用TLS连接">
        <a-switch name="TLS" v-model="formData.TLS"/>
      </a-form-item>
      <a-form-item>
        <a-button
          size="large"
          type="primary"
          htmlType="submit"
          class="setup-button"
          :loading="state.setupBtn"
          @click="submit"
        >
          创建
        </a-button>
      </a-form-item>
    </a-form>

  </div>
</template>

<script>
// import md5 from 'md5'
// import TwoStepCaptcha from '@/components/tools/TwoStepCaptcha'
import { mapActions } from 'vuex'
import { timeFix } from '@/utils/util'
// import { getSmsCaptcha, get2step } from '@/api/setup'
import { HostCreate } from '@/api/hosts'

export default {
  components: {
    // TwoStepCaptcha
  },
  data () {
    return {
      message: '',
      setupBtn: false,
      isSetupError: false,
      formData: {
        Name: '',
        ViaSocket: false,
        DockerEngineURL: '',
        HostIP: '',
        TLS: false
      },
      state: {
        time: 60,
        setupBtn: false
      }
    }
  },
  created () {
  },
  methods: {
    ...mapActions(['Setup']),
    // handler
    submit (e) {
      e.preventDefault()
      // console.log(this.formData)
      if (this.formData.Name.trim().length === 0) {
        this.message = '名称不能为空, 且不能包含空格, 请重新输入'
      } else if (this.formData.ViaSocket === false &&
        (this.formData.DockerEngineURL.trim().length === 0)) {
        this.message = 'Docker Endpoint URL不能为空, 请重新输入'
      } else {
        HostCreate(this.formData)
          .then((res) => {
            this.setupSuccess(res)
          })
          .catch((err) => {
            this.requestFailed(err)
          })
      }
    },
    setupSuccess (res) {
      console.log(res)
      this.$store.commit('saveCurrentHost', res.data)
      this.$router.push({ path: '/' })
      // 延迟 1 秒显示欢迎信息
      setTimeout(() => {
        this.$notification.success({
          message: '欢迎',
          description: `${timeFix()}，欢迎使用doconsole`
        })
      }, 1000)
      this.isSetupError = false
    },
    requestFailed (err) {
      console.log(err)
      this.isSetupError = true
      this.$notification['error']({
        message: '错误',
        description: ((err.response || {}).data || {}).message || '请求出现错误，请稍后再试',
        duration: 4
      })
    }
  }
}
</script>

<style lang="less" scoped>
.boot-layout-setup {
  label {
    font-size: 14px;
  }

  button.setup-button {
    padding: 0 15px;
    font-size: 16px;
    height: 40px;
    width: 100%;
  }
}
</style>
