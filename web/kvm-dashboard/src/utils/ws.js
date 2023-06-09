
import {TEMPINFO} from '@/constant/constant'

export default (url, uuid) => {
    // todo: use token to replace time
    let token = new Date().getTime()

    return new WebSocket(
        `${TEMPINFO.ws_url}${url}?UUID=${uuid}&token=${token}`
    )
}