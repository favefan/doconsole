import request from '@/utils/request'

const api = {
  base: 'api/v1/registries'
}

export default api

export function RegistryList () {
  return request({
    url: `${api.base}/json`,
    method: 'get'
  })
}

export function RegistryInspect (id) {
  return request({
      url: `${api.base}/json/${id}`,
      method: 'get'
  })
}

export function RegistriesRemove (a) {
  return request({
      url: `${api.base}`,
      method: 'delete',
      data: a
  })
}

export function RegistryCreate (values) {
  return request({
      url: `${api.base}/create`,
      method: 'post',
      data: values
  })
}
