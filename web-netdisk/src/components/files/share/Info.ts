import api from "@/utils/apis/request.ts";
import {Resp} from "@/utils/apis/base.ts";

export interface ShareItem {
    id: number
    // type: number
    name: string
    updated: string
    size: number
    url?: string
    sizeStr?: string
}

export interface ListResp {
    name: string
    created: string
    expired: number
    owner: number
    items: ShareItem[]
}

export interface OwnerInfoResp {
    shareStatus: number
    userId: number
    name: string
    avatar: string
    signature: string
    userStatus: number
}

export function listFilesByShareId(shareId: string, pwd: string) {
    return api.post<any, Resp<ListResp>>(`/file/share-info`, {
        'id': shareId,
        'pwd': pwd,
    })
}

export function getOwnerInfoByShareId(shareId: string) {
    return api.get<any, Resp<OwnerInfoResp>>(`/file/share-user/${shareId}`)
}

export function shareCancel(shareIds: string[]) {
    return api.post<any, Resp<any>>(`/file/share-cancel`, {
        "ids": shareIds
    })
}
export function shareReport(reason: string, shareId: string) {
    return api.post<any, Resp<any>>(`/file/share-report`, {
        "reason": reason,
        "shareId": shareId,
    })
}