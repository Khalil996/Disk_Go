import {defineStore} from "pinia";
import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface UserInfo {
    id: number
    username: string
    name: string
    avatar: string
    email: string
    signature: string
    used: number
    capacity: number
    status: number
    code?: string
}

export const useBaseStore = defineStore('base', () => {
    let token = localStorage.getItem("token") || '',
        user: { data: UserInfo } = {
            data:
                {id: 0, username: '', name: '', avatar: '', email: '', signature: '', used: 0, capacity: 0, status: 0}
        }

    function updateToken(tokenStr: string) {
        token = tokenStr
        localStorage.setItem("token", tokenStr)
    }

    function getToken() {
        if (!token || token === '') {
            token = localStorage.getItem("token") || ''
        }
        return token
    }

    function clearToken() {
        token = ''
        localStorage.removeItem('token')
        window.location.href = '/login'
    }

    async function getUserInfo() {
        if (user.data.id === 0) {
            const resp = await api.get<any, Resp<UserInfo>>(`/user/detail/0`)
            if (resp.code === 0) {
                user.data = resp.data
            }
        }
        return user.data
    }


    async function updateUserInfo(userInfo: UserInfo, post: boolean) {
        if (post) {
            const resp = await api.post<any, Resp<UserInfo>>(`/user/detail`, userInfo)
            if (resp.code === 0) {
                user.data = userInfo
                return resp
            }
        }
        user.data = userInfo
        return {code: 0}
    }

    return {
        updateToken, getToken, clearToken,
        getUserInfo, updateUserInfo
    }
})
