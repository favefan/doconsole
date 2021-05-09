import request from '@/utils/request'

const api = {
  base: 'api/v1/hosts'
}

export default api

export function HostList () {
  return request({
    url: `${api.base}/json`,
    method: 'get'
  })
}

export function HostInspect (id) {
  return request({
      url: `${api.base}/json/${id}`,
      method: 'get'
  })
}

export function HostsRemove (a) {
  return request({
      url: `${api.base}`,
      method: 'delete',
      data: a
  })
}

export function HostCreate (values) {
  return request({
      url: `${api.base}/create`,
      method: 'post',
      data: values
  })
}

export function HostUpdate (values) {
  return request({
      url: `${api.base}`,
      method: 'put',
      data: values
  })
}
