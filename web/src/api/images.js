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

export function ImagesRemove (a) {
  return request({
      url: `${api.base}`,
      method: 'delete',
      data: a
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

export function ImagePull (ref) {
  return request({
      url: `${api.base}/pull?ref=${ref}`,
      method: 'get'
  })
}
