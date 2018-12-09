import request from '@/utils/request'

export function login(username, password) {
  return request({
    url: '/api_login',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

export function getInfo(token) {
  return request({
    url: '/api/staffuser/info',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/api_logout',
    method: 'post'
  })
}
