import {Resp} from "@/utils/apis/base.ts";
import api from "../../utils/apis/request.ts";

export interface Folder {
    id: number
    name: string
    updated: string
}

// folder
export function listFoldersByParentFolderId(parentFolderId: number) {
    return api.get<any, Resp<Folder[]>>(`/file/folder-list/${parentFolderId}`)
}

export function createFolder(parentFolderId: number, name: string) {
    return api.post<any, Resp<any>>('/file/folder', {
        'parentFolderId': parentFolderId,
        'name': name
    })
}

export function updateFolderName(folder: Folder) {
    return api.put<any, Resp<any>>('/file/folder', {
        'id': folder.id,
        'name': folder.name
    })
}

export function listFolderMovableFolders(parentFolderId: number, selectedFolderIds: number[]) {
    return api.post<any, Resp<{ id: number, name: string }[]>>(`/file/folder-move`, {
        'parentFolderId': parentFolderId,
        'selectedFolderIds': selectedFolderIds
    })
}

export function moveFolders(parentFolderId: number, folderIds: number[]) {
    return api.put<any, Resp<any>>('/folder-move', {
        'parentFolderId': parentFolderId,
        'folderIds': folderIds
    })
}

export function copyFolders(parentFolderId: number, folderIds: number[]) {
    return api.post<any, Resp<any>>('/folder-copy', {
        'parentFolderId': parentFolderId,
        'folderIds': folderIds
    })
}

export function deleteFolders(ids: number[]) {
    return api.put<any, Resp<any>>('/file/folder-delete', {'ids': ids})
}

export function downloadFolder(ids: number[]) {
    return api.post<any, Resp<string[]>>('/file/folder-download', {'folderIds': ids})
}