import api from "@/utils/apis/request.ts";
import {Resp} from "@/utils/apis/base.ts";

export interface Share {
    id: number
    name: string
    created: string
    expired: number
    status: number
    downloadNum: number
    clickNum: number
    Url?: string
    state?: string
}

export function listShareFiles() {
    return api.get<any, Resp<any>>('/file/share')
}