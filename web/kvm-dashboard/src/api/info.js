import request from '@/utils/request.js'

export const getVMInfo = (uuid) => {
    return request(`/api/vm/basic?UUID=${uuid}`,'get')
}

export const getHostInfo = (uuid) => {
    return request(`/api/host/basic?UUID=${uuid}`, 'get')
}

export const getVMThreshold = (uuid) => {
    return request(`/api/machine/threshold?UUID=${uuid}`, 'get')
}