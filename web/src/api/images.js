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

export function ImageRemove (id) {
  return request({
      url: `${api.base}/${id}`,
      method: 'delete'
  })
}

export function ImageHistory (e) {
  return request({
      url: `${api.base}/history/${e}`,
      method: 'get'
  })
}

export function ImageSearch (e) {
  return request({
      url: `${api.base}/search?term=${e}`,
      method: 'get'
  })
}
