import request from '@/utils/request.js'


// vm
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
export const postAlertHistory = (data) => {
    // {
    //     "UUID": "5acefc92-b06d-48c8-a862-889c74d61c23",
    //     "PageSize": 10,
    //     "Page": 1
    // }
    return request(`/api/vm/table/alert/history`,'post', data)
}
export const postMetricHistory = (data) => {
    // {
    //     "UUID": "5acefc92-b06d-48c8-a862-889c74d61c23",
    //     "Fields": ["block_stats.rd_bytes"],
    //     "Period": "1h",
    //     "Func": "mean"
    // }
    return request(`/api/vm/graph/metric/history`,'post', data)
}
export const getMetricList = () => {
    return request(`/api/vm/graph/metric/list`,'get')
}

// host
export const postHostSimpleHistory = (data) => {
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
    return request(`/api/host/graph/workload/history`,'post', data)
}
export const postHostAlertHistory = (data) => {
    // {
    //     "UUID": "5acefc92-b06d-48c8-a862-889c74d61c23",
    //     "PageSize": 10,
    //     "Page": 1
    // }
    return request(`/api/host/table/alert/history`,'post', data)
}