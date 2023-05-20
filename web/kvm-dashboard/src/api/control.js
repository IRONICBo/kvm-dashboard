import ws from '@/utils/ws'

export const StartWebSocket = (uuid) =>  {
    return ws(
        '/api/vm/control/start/ws',
        uuid
    )
}

export const StopWebSocket = (uuid) => {
    return ws(
        '/api/vm/control/stop/ws',
        uuid
    )
}

export const SuspendWebSocket = (uuid) =>  {
    return ws(
        '/api/vm/control/suspend/ws',
        uuid
    )
}

export const ResumeWebSocket = (uuid) => {
    return ws(
        '/api/vm/control/resume/ws',
        uuid
    )
}