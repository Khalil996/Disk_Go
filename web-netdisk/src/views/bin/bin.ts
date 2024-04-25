import api from "../../utils/apis/request.ts";
import {Resp} from "@/utils/apis/base.ts";

export interface DeleteFile {
    id: number
    name: string
    folderId?: number
    folderName?: string
    delTime: number
    size: number
    delTimeStr?: string
    sizeStr?: string
    left?: string
}

export function getDeletedFiles() {
    return api.get<any, Resp<DeleteFile[]>>('/file/delete')
}

export function deleteFilesTruly(ids: number[]) {
    return api.post<any, Resp<any>>('/file/delete', {ids})
}

export function deleteAllFilesTruly() {
    return api.post<any, Resp<DeleteFile[]>>('/file/clear')
}

export function recoverFiles(obj: {}) {
    return api.put<any, Resp<DeleteFile[]>>('/file/recover', obj)
}