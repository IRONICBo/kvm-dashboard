import request from '@/utils/request.js'

export const login = (username, password) => {
    return request('/api/user/login', 'post', {username, password})
}