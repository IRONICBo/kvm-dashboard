import request from '@/utils/request.js'

export const getVMInfo = (uuid) => {
    return request(`/api/vm/basic?UUID=${uuid}`,'get')
}

export const getHostInfo = () => {
    return request(`/api/host/basic`, 'get')
}