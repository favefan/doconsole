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
    currentHost: JSON.parse(localStorage.getItem('currentHost')) || '',
    // isConnected: false,
    // contentUpdate: true,
    toRender: false
  },
  mutations: {
    updateHostList (state, data) {
      state.hostList = data
    },
    // saveCurrentHost (state, host) {
    //   state.currentHost = host
    // },
    isRenderContent (state, is) {
      state.toRender = is
    },
    saveCurrentHost (state, host) {
      localStorage.setItem('currentHost', JSON.stringify(host))
      state.currentHost = host
    }
    // changingConnectionState (state, is) {
    //   state.isConnected = is
    // }
  },
  actions: {

  },
  getters
})
