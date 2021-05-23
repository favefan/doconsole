import request from '@/utils/request'

const api = {
  base: 'api/v1/networks'
}

export default api

export function NetworkList () {
  return request({
    url: `${api.base}/json`,
    method: 'get'
  })
}

export function NetworkInspect (id) {
  return request({
      url: `${api.base}/json/${id}`,
      method: 'get'
  })
}

export function NetworksRemove (a) {
  return request({
      url: `${api.base}`,
      method: 'delete',
      data: a
  })
}

export function NetworkCreate (values) {
  return request({
      url: `${api.base}/create`,
      method: 'post',
      data: values
  })
}

export function NetworkConnect (imageID, containerID) {
  return request({
      url: `${api.base}/connect/${imageID}`,
      method: 'post',
      data: {
        Container: containerID
      }
  })
}

export function NetworkDisconnect (imageID, containerID) {
  return request({
      url: `${api.base}/disconnect/${imageID}`,
      method: 'post',
      data: {
        Container: containerID
      }
  })
}
