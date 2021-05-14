import request from '@/utils/request'

const api = {
  base: 'api/v1/docker'
}

export default api

export function DockerClientCreate (value) {
  return request({
      url: `${api.base}/client`,
      method: 'put',
      data: value
  })
}
