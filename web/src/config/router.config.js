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
    redirect: '/dashboard/workplace',
    children: [
      // dashboard
      {
        path: '/dashboard',
        name: 'dashboard',
        redirect: '/dashboard/workplace',
        component: RouteView,
        meta: { title: 'menu.dashboard', keepAlive: true, icon: dashboard },
        children: [
          {
            path: '/dashboard/analysis/:pageNo([1-9]\\d*)?',
            name: 'Analysis',
            component: () => import('@/views/dashboard/Analysis'),
            meta: { title: 'menu.dashboard.analysis', keepAlive: false }
          },
          // 外部链接
          {
            path: 'https://www.baidu.com/',
            name: 'Monitor',
            meta: { title: 'menu.dashboard.monitor', target: '_blank' }
          },
          {
            path: '/dashboard/workplace',
            name: 'Workplace',
            component: () => import('@/views/dashboard/Workplace'),
            meta: { title: 'menu.dashboard.workplace', keepAlive: true }
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
      },
      // 仓库管理
      // {
      //   path: '/repository',
      //   name: 'repository',
      //   redirect: '/repository/repositories',
      //   component: RouteView,
      //   meta: { title: 'menu.repository', icon: repository, keepAlive: true },
      //   children: [
      //     {
      //       path: '/repository/repositories',
      //       name: 'Images',
      //       component: () => import('@/views/repository/Repositories'),
      //       meta: { title: 'menu.repository.repositories', keepAlive: true }
      //     }
      //   ]
      // },
      // {
      //   path: '/repository',
      //   name: 'repository',
      //   redirect: '/repository/repositories',
      //   component: RouteView,
      //   meta: { title: 'menu.repository', icon: volume, keepAlive: true },
      //   children: [
      //     {
      //       path: '/repository/repositories',
      //       name: 'Images',
      //       component: () => import('@/views/repository/Repositories'),
      //       meta: { title: 'menu.repository.repositories', keepAlive: true }
      //     }
      //   ]
      // },
      // {
      //   path: '/repository',
      //   name: 'repository',
      //   redirect: '/repository/repositories',
      //   component: RouteView,
      //   meta: { title: 'menu.repository', icon: network, keepAlive: true },
      //   children: [
      //     {
      //       path: '/repository/repositories',
      //       name: 'Images',
      //       component: () => import('@/views/repository/Repositories'),
      //       meta: { title: 'menu.repository.repositories', keepAlive: true }
      //     }
      //   ]
      // },
      // {
      //   path: '/repository',
      //   name: 'repository',
      //   redirect: '/repository/repositories',
      //   component: RouteView,
      //   meta: { title: 'menu.repository', icon: host, keepAlive: true },
      //   children: [
      //     {
      //       path: '/repository/repositories',
      //       name: 'Images',
      //       component: () => import('@/views/repository/Repositories'),
      //       meta: { title: 'menu.repository.repositories', keepAlive: true }
      //     }
      //   ]
      // },
      // profile
      // {
      //   path: '/profile',
      //   name: 'profile',
      //   component: RouteView,
      //   redirect: '/profile/basic',
      //   meta: { title: 'menu.profile', icon: 'setting', permission: ['profile'] },
      //   children: [
      //     {
      //       path: '/profile/basic',
      //       name: 'ProfileBasic',
      //       component: () => import('@/views/profile/basic'),
      //       meta: { title: 'menu.profile.basic', permission: ['profile'] }
      //     },
      //     {
      //       path: '/profile/advanced',
      //       name: 'ProfileAdvanced',
      //       component: () => import('@/views/profile/advanced/Advanced'),
      //       meta: { title: 'menu.profile.advanced', permission: ['profile'] }
      //     }
      //   ]
      // }

      // // result
      // {
      //   path: '/result',
      //   name: 'result',
      //   component: RouteView,
      //   redirect: '/result/success',
      //   meta: { title: 'menu.result', icon: 'check-circle-o', permission: ['result'] },
      //   children: [
      //     {
      //       path: '/result/success',
      //       name: 'ResultSuccess',
      //       component: () => import(/* webpackChunkName: "result" */ '@/views/result/Success'),
      //       meta: { title: 'menu.result.success', keepAlive: false, hiddenHeaderContent: true, permission: ['result'] }
      //     },
      //     {
      //       path: '/result/fail',
      //       name: 'ResultFail',
      //       component: () => import(/* webpackChunkName: "result" */ '@/views/result/Error'),
      //       meta: { title: 'menu.result.fail', keepAlive: false, hiddenHeaderContent: true, permission: ['result'] }
      //     }
      //   ]
      // },

      // // Exception
      // {
      //   path: '/exception',
      //   name: 'exception',
      //   component: RouteView,
      //   redirect: '/exception/403',
      //   meta: { title: 'menu.exception', icon: 'warning', permission: ['exception'] },
      //   children: [
      //     {
      //       path: '/exception/403',
      //       name: 'Exception403',
      //       component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/403'),
      //       meta: { title: 'menu.exception.not-permission', permission: ['exception'] }
      //     },
      //     {
      //       path: '/exception/404',
      //       name: 'Exception404',
      //       component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/404'),
      //       meta: { title: 'menu.exception.not-find', permission: ['exception'] }
      //     },
      //     {
      //       path: '/exception/500',
      //       name: 'Exception500',
      //       component: () => import(/* webpackChunkName: "fail" */ '@/views/exception/500'),
      //       meta: { title: 'menu.exception.server-error', permission: ['exception'] }
      //     }
      //   ]
      // },

      // account
      {
        path: '/account',
        component: RouteView,
        redirect: '/account/center',
        name: 'account',
        meta: { title: 'menu.account', icon: 'setting', keepAlive: true, permission: ['user'] },
        children: [
          {
            path: '/account/center',
            name: 'center',
            component: () => import('@/views/account/center'),
            meta: { title: 'menu.account.center', keepAlive: true, permission: ['user'] }
          },
          {
            path: '/account/settings',
            name: 'settings',
            component: () => import('@/views/account/settings/Index'),
            meta: { title: 'menu.account.settings', hideHeader: true, permission: ['user'] },
            redirect: '/account/settings/basic',
            hideChildrenInMenu: true,
            children: [
              {
                path: '/account/settings/basic',
                name: 'BasicSettings',
                component: () => import('@/views/account/settings/BasicSetting'),
                meta: { title: 'account.settings.menuMap.basic', hidden: true, permission: ['user'] }
              },
              {
                path: '/account/settings/security',
                name: 'SecuritySettings',
                component: () => import('@/views/account/settings/Security'),
                meta: {
                  title: 'account.settings.menuMap.security',
                  hidden: true,
                  keepAlive: true,
                  permission: ['user']
                }
              },
              {
                path: '/account/settings/custom',
                name: 'CustomSettings',
                component: () => import('@/views/account/settings/Custom'),
                meta: { title: 'account.settings.menuMap.custom', hidden: true, keepAlive: true, permission: ['user'] }
              },
              {
                path: '/account/settings/binding',
                name: 'BindingSettings',
                component: () => import('@/views/account/settings/Binding'),
                meta: { title: 'account.settings.menuMap.binding', hidden: true, keepAlive: true, permission: ['user'] }
              },
              {
                path: '/account/settings/notification',
                name: 'NotificationSettings',
                component: () => import('@/views/account/settings/Notification'),
                meta: {
                  title: 'account.settings.menuMap.notification',
                  hidden: true,
                  keepAlive: true,
                  permission: ['user']
                }
              }
            ]
          }
        ]
      }

      // other
      /*
      {
        path: '/other',
        name: 'otherPage',
        component: PageView,
        meta: { title: '其他组件', icon: 'slack', permission: [ 'dashboard' ] },
        redirect: '/other/icon-selector',
        children: [
          {
            path: '/other/icon-selector',
            name: 'TestIconSelect',
            component: () => import('@/views/other/IconSelectorView'),
            meta: { title: 'IconSelector', icon: 'tool', keepAlive: true, permission: [ 'dashboard' ] }
          },
          {
            path: '/other/list',
            component: RouteView,
            meta: { title: '业务布局', icon: 'layout', permission: [ 'support' ] },
            redirect: '/other/list/tree-list',
            children: [
              {
                path: '/other/list/tree-list',
                name: 'TreeList',
                component: () => import('@/views/other/TreeList'),
                meta: { title: '树目录表格', keepAlive: true }
              },
              {
                path: '/other/list/edit-table',
                name: 'EditList',
                component: () => import('@/views/other/TableInnerEditList'),
                meta: { title: '内联编辑表格', keepAlive: true }
              },
              {
                path: '/other/list/user-list',
                name: 'UserList',
                component: () => import('@/views/other/UserList'),
                meta: { title: '用户列表', keepAlive: true }
              },
              {
                path: '/other/list/role-list',
                name: 'RoleList',
                component: () => import('@/views/other/RoleList'),
                meta: { title: '角色列表', keepAlive: true }
              },
              {
                path: '/other/list/system-role',
                name: 'SystemRole',
                component: () => import('@/views/role/RoleList'),
                meta: { title: '角色列表2', keepAlive: true }
              },
              {
                path: '/other/list/permission-list',
                name: 'PermissionList',
                component: () => import('@/views/other/PermissionList'),
                meta: { title: '权限列表', keepAlive: true }
              }
            ]
          }
        ]
      }
      */
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
