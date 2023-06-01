import axios from 'axios'
import { ElMessage } from 'element-plus'
import router from "../router";

axios.interceptors.response.use(success => {
    // if (success.states && success.states == 200) {
    //     if (success.data.code != 200) {
    //         ElMessage.error({message: success.data.message});
    //         return;
    //     }
    //     if (success.data.message) {
    //         ElMessage.success({message: success.data.message});
    //     }
    // }
    return success.data;
}, error => {
    // if (error.response.code == 504 || error.response.code == 404) {
    //     ElMessage.error({message: '服务器被吃了/(ㄒoㄒ)/~~'});
    // } else if (error.response.code == 403) {
    //     ElMessage.error({message: '权限不足，请联系管理员!'});
    // } else if (error.response.code == 401) {
    //     ElMessage.error({message: '尚未登陆，请登录!'});
    router.replace('/login').catch(err => {err});
    // } else {
    //     if (error.response.data.message) {
    //         ElMessage.error({message: error.response.data.message});
    //     } else {
    //         ElMessage.error({message: '未知错误!'});
    //     }
    // }
    Promise.reject(error)
});

let base = '/web';

//传送json格式的post请求
export const postRequest = (url, params) => {
    return axios({
        method: 'post',
        url: `${base}${url}`,
        data: params
    })
}

//传送json格式的put请求
export const putRequest = (url, params) => {
    return axios({
        method: 'put',
        url: `${base}${url}`,
        data: params
    })
}

//传送json格式的get请求
export const getRequest = (url, params) => {
    return axios({
        method: 'get',
        url: `${base}${url}`,
        data: params
    })
}

//传送json格式的delete请求
export const deleteRequest = (url, params) => {
    return axios({
        method: 'delete',
        url: `${base}${url}`,
        data: params
    })
}