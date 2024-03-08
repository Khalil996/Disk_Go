import {Resp} from "@/utils/apis/base.ts";
import api from "@/utils/apis/request.ts";

export const extM = {
    '.jpg': true,
    '.png': true,
    '.jpeg': true,
}

export async function updateAvatar(formData: FormData) {
    return await api.post<any, Resp<any>>('/user/avatar', formData, {headers: {'Content-Type': 'multipart/form-data'}});
}

export async function updateDetail(formData: FormData) {
    return await api.post<any, Resp<any>>('/user/detail', formData, {headers: {'Content-Type': 'multipart/form-data'}});
}