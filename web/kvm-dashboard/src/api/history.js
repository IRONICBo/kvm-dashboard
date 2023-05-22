import request from '@/utils/request.js'

export const postSimpleHistory = (data) => {
    // {
    //     "UUID": "5acefc92-b06d-48c8-a862-889c74d61c23",
    //     "Fields": [
    //         "cpu_usage",
    //         "mem_usage",
    //         "disk_usage",
    //         "net_rx_rate",
    //         "net_tx_rate"
    //     ]
    // }
    return request(`/api/vm/graph/workload/history`,'post', data)
}