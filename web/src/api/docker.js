import request from '@/utils/request'

const api = {
  base: 'api/v1'
}

export default api

export function DockerClientCreate (value) {
  return request({
      url: `${api.base}/docker/client`,
      method: 'put',
      data: value
  })
}

export function Info () {
  return request({
      url: `${api.base}/info`,
      method: 'get'
  })
}
