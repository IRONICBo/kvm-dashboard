
import {TEMPINFO} from '@/constant/constant'

export default (url, uuid) => {
    return new WebSocket(
        `${TEMPINFO.ws_url}${url}?UUID=${uuid}`
    )
}