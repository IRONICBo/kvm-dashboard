import ws from '@/utils/ws'

// vm
export const SimpleWebSocket = (uuid) =>  {
    return ws(
        '/api/vm/graph/workload/realtime/ws',
        uuid
    )
}
export const ProcessWebSocket = (uuid) => {
    return ws(
        '/api/vm/table/process/realtime/ws',
        uuid
    )
}

// host
export const HostSimpleWebSocket = (uuid) =>  {
    return ws(
        '/api/host/graph/workload/realtime/ws',
        uuid
    )
}
export const HostProcessWebSocket = (uuid) => {
    return ws(
        '/api/host/table/process/realtime/ws',
        uuid
    )
}