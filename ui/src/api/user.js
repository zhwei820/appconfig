import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/api/staffuser/staffuser/',
    method: 'get',
    params: query
  })
}

export function fetchUser(id) {
  return request({
    url: '/api/staffuser/staffuser/' + id,
    method: 'get',
    params: { id }
  })
}

export function createUser(data) {
  return request({
    url: '/api/staffuser/staffuser/',
    method: 'post',
    data
  })
}

export function updateUser(id, data) {
  return request({
    url: '/api/staffuser/staffuser/' + id,
    method: 'put',
    data
  })
}
