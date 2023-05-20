import axios from 'axios';
import { getCookie } from './token.js';

const service = axios.create({
    baseURL: '/golang_web',
    timeout: 5000,
})

// request interceptor
service.interceptors.request.use(config => {
    const token = getCookie('token')
    if(token) {
        config.headers.token = token;
    }
    return config
}, error => {
    Promise.reject(error)
}) 

// response interceptor
service.interceptors.response.use(({ data }) => {
    return data
}, error => {
    Promise.reject(error)
})

// used for axios
export default (url, method, submitData) => {
    return service({
        url,
        method,
        // change the data to params when the method is get
        [method === 'get' ? 'params' : 'data']: submitData
    })
}