import api from "../../utils/apis/request.ts";
import {Resp} from "../../utils/apis/base.ts";

export interface DeleteFile {
    id: number
    name: string
    folderId: number
    folderName: string
    deleted: string
    size: number
}

export function getDeletedFiles() {
    return api.get<any, Resp<DeleteFile[]>>('/file/delete')
}

export function deleteFilesTruly() {
    return api.get<any, Resp<DeleteFile[]>>('/file/delete')
}

export function deleteAllFilesTruly() {
    return api.post<any, Resp<DeleteFile[]>>('/file/clear')
}