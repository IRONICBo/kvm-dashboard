import axios from 'axios';
import { getCookie } from './token.js';
import { startLoading, stopLoading } from './loading.js';

const service = axios.create({
    baseURL: 'http://127.0.0.1:5001/',
    withCredentials: true,
    timeout: 5000,
})


// request interceptor
service.interceptors.request.use(config => {
    // start loading
    startLoading();

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
    // stop loading
    stopLoading();

    return data
}, error => {
    if (error.response.status === 401) {
        // redirect to login page
        window.location.href = '/login';
    }
    
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