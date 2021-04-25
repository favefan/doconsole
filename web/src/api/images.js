import request from '@/utils/request'

const api = {
  base: 'api/v1/images'
}

export default api

export function ImageList () {
  return request({
    url: `${api.base}/json`,
    method: 'get'
  })
}

export function ImageInspect (id) {
  return request({
      url: `${api.base}/json/${id}`,
      method: 'get'
  })
}
