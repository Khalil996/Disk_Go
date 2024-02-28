import {defineStore} from "pinia";
import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface UserInfo {
    id: number
    name: string
    avatar: string
    email: string
    signature: string
    status: number
}

export const useBaseStore = defineStore('base', () => {
    let token = localStorage.getItem("token") || ''
    let user: { data: UserInfo } = {data: {id: 0, name: '', avatar: '', email: '', signature: '', status: 0}}

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

    async function getUserInfo() {
        if (user.data.id == 0) {
            const resp = await api.get<any, Resp<UserInfo>>(`/user/detail/0`)
            if (resp && resp.code === 0) {
                user.data = resp.data
            }
        }
        return user
    }

    function updateUserInfo(userInfo: UserInfo) {
        user.data = userInfo
    }

    return {
        updateToken, getToken,
        getUserInfo, updateUserInfo
    }
})
