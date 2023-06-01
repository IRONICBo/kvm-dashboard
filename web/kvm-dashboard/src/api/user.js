import request from '@/utils/request.js'

export const login = (username, password) => {
    return request('/api/user/login', 'post', {username, password})
}

export const logout = () => {
    return request('/api/user/logout', 'get')
}