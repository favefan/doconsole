// eslint-disable-next-line
import { UserLayout, BasicLayout, BootLayout } from '@/layouts'
import { containers, images, dashboard, network, volume, registry, host } from '@/core/icons' // repository, volume, network, host

const RouteView = {
  name: 'RouteView',
  render: h => h('router-view')
}

export const asyncRouterMap = [
  {
    path: '/',
    name: 'index',
    component: BasicLayout,
    meta: { title: 'menu.home' },
    redirect: '/dashboard/index',
    children: [
      // dashboard
      {
        path: '/dashboard',
        name: 'dashboard',
        redirect: '/dashboard/index',
        component: RouteView,
        meta: { title: 'menu.dashboard', keepAlive: true, icon: dashboard },
        children: [
          {
            path: '/dashboard/index',
            name: 'Index',
            component: () => import('@/views/dashboard/Index'),
            meta: { title: 'menu.dashboard.index', keepAlive: true }
          }
        ]
      },
      // 容器管理
      {
        path: '/container',
        name: 'container',
        redirect: '/container/containers',
        component: RouteView,
        meta: { title: 'menu.container', icon: containers, keepAlive: true },
        children: [
          {
            path: '/container/containers',
            name: 'Containers',
            component: () => import('@/views/container/Containers'),
            meta: { title: 'menu.container.containers', keepAlive: true }
          },
          {
            path: '/container/information',
            name: 'ContainersInformation',
            component: () => import('@/views/container/Information'),
            meta: { title: 'menu.container.information', keepAlive: false },
            hidden: true
          },
          {
            path: '/container/creator',
            name: 'ContainersCreator',
            component: () => import('@/views/container/Creator'),
            meta: { title: 'menu.container.creator', keepAlive: false },
            hidden: true
          }
        ]
      },
      // 镜像管理
      {
        path: '/image',
        name: 'image',
        redirect: '/image/images',
        component: RouteView,
        meta: { title: 'menu.image', icon: images, keepAlive: true },
        children: [
          {
            path: '/image/images',
            name: 'Images',
            component: () => import('@/views/image/Images'),
            meta: { title: 'menu.image.images', keepAlive: true }
          },
          {
            path: '/image/information',
            name: 'ImageInformation',
            component: () => import('@/views/image/Information'),
            meta: { title: 'menu.image.information', keepAlive: false },
            hidden: true
          }
        ]
      },
      // 网络管理
      {
        path: '/network',
        name: 'network',
        redirect: '/network/networks',
        component: RouteView,
        meta: { title: 'menu.network', icon: network, keepAlive: true },
        children: [
          {
            path: '/network/networks',
            name: 'Networks',
            component: () => import('@/views/network/Networks'),
            meta: { title: 'menu.network.networks', keepAlive: true }
          },
          {
            path: '/network/information',
            name: 'NetworkInformation',
            component: () => import('@/views/network/Information'),
            meta: { title: 'menu.network.information', keepAlive: false },
            hidden: true
          }
        ]
      },
      // 卷管理
      {
        path: '/volume',
        name: 'volume',
        redirect: '/volume/volumes',
        component: RouteView,
        meta: { title: 'menu.volume', icon: volume, keepAlive: true },
        children: [
          {
            path: '/volume/volumes',
            name: 'Volumes',
            component: () => import('@/views/volume/Volumes'),
            meta: { title: 'menu.volume.volumes', keepAlive: true }
          },
          {
            path: '/volume/information',
            name: 'VolumeInformation',
            component: () => import('@/views/volume/Information'),
            meta: { title: 'menu.volume.information', keepAlive: false },
            hidden: true
          }
        ]
      },
      // 仓库管理
      {
        path: '/registry',
        name: 'registry',
        redirect: '/registry/registries',
        component: RouteView,
        meta: { title: 'menu.registry', icon: registry, keepAlive: true },
        children: [
          {
            path: '/registry/registries',
            name: 'Registries',
            component: () => import('@/views/registry/Registries'),
            meta: { title: 'menu.registry.registries', keepAlive: true }
          }
        ]
      },
      // 主机管理
      {
        path: '/host',
        name: 'host',
        redirect: '/host/hosts',
        component: RouteView,
        meta: { title: 'menu.host', icon: host, keepAlive: true },
        children: [
          {
            path: '/host/hosts',
            name: 'Hosts',
            component: () => import('@/views/host/Hosts'),
            meta: { title: 'menu.host.hosts', keepAlive: true }
          }
        ]
      }
    ]
  },
  {
    path: '*',
    redirect: '/404',
    hidden: true
  }
]

/**
 * 基础路由
 * @type { *[] }
 */
export const constantRouterMap = [
  {
    path: '/boot',
    component: BootLayout,
    redirect: '/boot/setup',
    hidden: true,
    children: [
      {
        path: 'setup',
        name: 'setup',
        component: () => import(/* webpackChunkName: "boot" */ '@/views/boot/Setup')
      },
      {
        path: 'redirect',
        name: 'redirect',
        component: () => import(/* webpackChunkName: "boot" */ '@/views/boot/Redirect')
      }
    ]
  },
  {
    path: '/user',
    component: UserLayout,
    redirect: '/user/login',
    hidden: true,
    children: [
      {
        path: 'login',
        name: 'login',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Login')
      },
      {
        path: 'register',
        name: 'register',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/Register')
      },
      {
        path: 'register-result',
        name: 'registerResult',
        component: () => import(/* webpackChunkName: "user" */ '@/views/user/RegisterResult')
      },
      {
        path: 'recover',
        name: 'recover',
        component: undefined
      }
    ]
  },

  {
    path: '/404',
    component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404')
  }
]
