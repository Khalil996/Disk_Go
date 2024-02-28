import api from "../../utils/apis/request.ts";
import {Resp} from "../../utils/apis/base.ts";

export interface File {
    id: number
    name: string
    size: number
    url: string
    type: number
    status: number
    updated: string
    ext?: string
    sizeStr?: string
}

export function listFilesByFolderId(parentFolderId: number) {
    return api.get<any, Resp<File[]>>(`/file/list/${parentFolderId}`)
}

export function listFilesByFileType(fileType: number) {
    return api.get<any, Resp<File[]>>(`/file/type/${fileType}`)
}

export function updateFileName(file: File) {
    return api.put<any, Resp<any>>('/file', {
        'id': file.id,
        'name': file.name
    })
}

export function listFileMovableFolders(folderId: number) {
    return api.get<any, Resp<{ id: number, name: string, updated: string | '' }[]>>(`/file/move/${folderId}`)
}

export function moveFiles(parentFolderId: number, fileIds: number[]) {
    return api.put<any, Resp<any>>('/file/move', {
        'folderId': parentFolderId,
        'fileIds': fileIds
    })
}

export function copyFiles(parentFolderId: number, fileIds: number[]) {
    return api.post<any, Resp<any>>('/file/copy', {
        'folderId': parentFolderId,
        'fileIds': fileIds
    })
}

export function deleteFiles(fileIds: number[], folderId: number) {
    return api.put<any, Resp<any>>('/file/delete', {
        'fileIds': fileIds,
        'folderId': folderId
    })
}
