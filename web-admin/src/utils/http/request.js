import {ElMessage} from 'element-plus'
import axios from 'axios'
import router from "../../router";
import {useBaseStore} from "../../../store/index.js";


function jumpLogin() {
    useBaseStore().updateToken('')
    router.push("/login")
}

// 1.axios二次封装
// 相当于我们创建了一个带配置的axios，然后用这个
const api = axios.create({
    // 基础路径
    baseURL: '/api',
    // import.meta.env.BASE_URL.VITE_APP_BASE_API, //请求路径上都会写带上这个url  但是这里不知道为什么读取不到
    timeout: 60 * 1000, // 请求超时时间
})

// 2.request实例添加请求和拦截器
api.interceptors.request.use(
    config => {
        const token = useBaseStore().getToken()
        if (token && typeof window !== "undefined") {
            config.headers.Authorization = token
        }
        // config配置对象有headers请求头
        return config
    },
    error => {
        error.data = {}
        error.data.msg = "请求异常"
        return Promise.reject(error)
    })

// 3.响应拦截器
api.interceptors.response.use(
    resp => {
        const {code: code, msg: message} = resp.data

        if (code === 401 || code === 403) {
            localStorage.clear()
            jumpLogin()
        }

        if (
            code.toString() &&
            code.toString() !== "0" &&
            code.toString() !== "401" &&
            code.toString() !== "403"
        ) {
            promptError(message)
            return Promise.reject(resp.data)
        }

        if (resp.data.code && resp.data.code !== 0) {
            ElMessage({
                type: 'error',
                message: resp.data.msg,
            })
        }

        // 这里是简化了数据
        return Promise.resolve(resp.data)
    },
    error => {

        const status = error.response.status
        let msg
        switch (status) {
            case 201:
                msg = '用户名或密码不对'
                break
            case 401:
                // 401 一般是token过期
                msg = 'TOKEN过期'
                jumpLogin()
                break
            case 403:
                msg = '无权访问'
                break
            case 404:
                msg = '没有这个资源'
                break
            case 500:
                msg = '服务器嗝屁了，哈哈哈'
                break
            default:
                msg = '网络出现问题'
                break
        }
        promptError(msg)
        return Promise.resolve(error)
    },
)

export default api