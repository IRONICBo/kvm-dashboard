import request from '@/utils/request.js'

export const getMachineList = () => {
    return request(`/api/machine/list`,'get')
}