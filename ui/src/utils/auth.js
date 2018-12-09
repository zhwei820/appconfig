import ls from 'local-storage'

const TokenKey = 'Admin-Token'

export function getToken() {
  return ls.get(TokenKey)
}

export function setToken(token) {
  return ls.set(TokenKey, token)
}

export function removeToken() {
  return ls.remove(TokenKey)
}
