import request from '@/utils/request.js'

export const getMachineList = () => {
    return request(`/api-go/machine/list`,'get')
}