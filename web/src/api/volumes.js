import request from '@/utils/request'

const api = {
  base: 'api/v1/volumes'
}

export default api

export function VolumeList () {
  return request({
    url: `${api.base}/json`,
    method: 'get'
  })
}

export function VolumeInspect (id) {
  return request({
      url: `${api.base}/json/${id}`,
      method: 'get'
  })
}

export function VolumesRemove (a) {
  return request({
      url: `${api.base}`,
      method: 'delete',
      data: a
  })
}

export function VolumeCreate (values) {
  return request({
      url: `${api.base}/create`,
      method: 'post',
      data: values
  })
}
