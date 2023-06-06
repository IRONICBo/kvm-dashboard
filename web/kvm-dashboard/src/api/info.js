import request from '@/utils/request.js'

export const getVMInfo = (uuid) => {
    return request(`/api-go/vm/basic?UUID=${uuid}`,'get')
}

export const getHostInfo = (uuid) => {
    return request(`/api-go/host/basic?UUID=${uuid}`, 'get')
}

export const getVMThreshold = (uuid) => {
    return request(`/api-go/machine/threshold?UUID=${uuid}`, 'get')
}