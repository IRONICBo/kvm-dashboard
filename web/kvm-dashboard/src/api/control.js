import ws from '@/utils/ws'

// vm control
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

// host control
export const HostStartWebSocket = (uuid) =>  {
    return ws(
        '/api/host/control/start_report/ws',
        uuid
    )
}
export const HostStopWebSocket = (uuid) => {
    return ws(
        '/api/host/control/stop_report/ws',
        uuid
    )
}