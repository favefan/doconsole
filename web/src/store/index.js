import Vue from 'vue'
import Vuex from 'vuex'

import app from './modules/app'
import user from './modules/user'

// default router permission control
import permission from './modules/permission'

// dynamic router permission control (Experimental)
// import permission from './modules/async-router'
import getters from './getters'

// import { HostList } from '@/api/hosts'

Vue.use(Vuex)

export default new Vuex.Store({
  modules: {
    app,
    user,
    permission
  },
  state: {
    hostList: [],
    currentHost: '',
    contentUpdate: true
  },
  mutations: {
    updateHostList (state, data) {
      state.hostList = data
    },
    changeHost (state, host) {
      state.currentHost = host
    },
    updateContent (state) {
      if (state.contentUpdate === true) {
        state.contentUpdate = false
      } else {
        state.contentUpdate = true
      }
    }
  },
  actions: {

  },
  getters
})
