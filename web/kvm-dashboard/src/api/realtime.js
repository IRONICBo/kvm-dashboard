import ws from '@/utils/ws'

// vm
export const SimpleWebSocket = (uuid) =>  {
    return ws(
        '/api-go/vm/graph/workload/realtime/ws',
        uuid
    )
}
export const ProcessWebSocket = (uuid) => {
    return ws(
        '/api-go/vm/table/process/realtime/ws',
        uuid
    )
}

// host
export const HostSimpleWebSocket = (uuid) =>  {
    return ws(
        '/api-go/host/graph/workload/realtime/ws',
        uuid
    )
}
export const HostProcessWebSocket = (uuid) => {
    return ws(
        '/api-go/host/table/process/realtime/ws',
        uuid
    )
}