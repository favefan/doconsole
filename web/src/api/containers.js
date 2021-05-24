import request from '@/utils/request'

const api = {
  base: 'api/v1/containers'
}

export default api

export function ContainerCreate (data) {
  return request({
    url: `${api.base}/create`,
    method: 'post',
    data: data
  })
}

export function ContainerExecCreate (id, data) {
  return request({
    url: `${api.base}/exec/${id}`,
    method: 'post',
    data: data
  })
}

export function ContainerList () {
  return request({
    url: `${api.base}/json`,
    method: 'get'
  })
}

export function ContainerInspect (id) {
  return request({
      url: `${api.base}/json/${id}`,
      method: 'get'
  })
}

export function ContainerStart (id) {
  return request({
    url: `${api.base}/start?id=${id}`,
    method: 'get'
  })
}

export function ContainerStop (id) {
  return request({
    url: `${api.base}/stop?id=${id}`,
    method: 'get'
  })
}

export function ContainerPause (id) {
  return request({
    url: `${api.base}/pause?id=${id}`,
    method: 'get'
  })
}

export function ContainerRestart (id) {
  return request({
    url: `${api.base}/restart?id=${id}`,
    method: 'get'
  })
}

export function ContainerRemove (id) {
  return request({
    url: `${api.base}/${id}`,
    method: 'delete'
  })
}

export function ContainerUpdate (id, data) {
  return request({
    url: `${api.base}/update/${id}`,
    method: 'post',
    data: data
  })
}

// id == 0 add     post
// id != 0 update  put
// export function saveService(parameter) {
//   return request({
//     url: api.service,
//     method: parameter.id === 0 ? 'post' : 'put',
//     data: parameter
//   })
// }

// export function saveSub(sub) {
//   return request({
//     url: '/sub',
//     method: sub.id === 0 ? 'post' : 'put',
//     data: sub
//   })
// }
