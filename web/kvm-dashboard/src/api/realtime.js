import ws from '@/utils/ws'

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