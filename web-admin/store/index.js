import {defineStore} from "pinia";
import api from "@/utils/http/request.js";


export const useBaseStore = defineStore('base', () => {
    let token = localStorage.getItem("token") || '',
        user = {
            data:
                {id: 0, username: '', name: '', avatar: '', email: '', signature: '', used: 0, capacity: 0, status: 0}
        }

    function updateToken(tokenStr) {
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
            const resp = await api.get(`/user/detail/0`)
            if (resp.data.code === 0) {
                user.data = resp.data
            }
        }
        return user.data
    }

    async function updateUserInfo(userInfo, post) {
        if (post) {
            const resp = await api.post(`/user/detail`, userInfo)
            if (resp.data.code === 0) {
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
