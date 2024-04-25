import api from "../../utils/apis/request.ts";
import {Resp} from "../../utils/apis/base.ts";

export interface CheckReq {
    folderId: number
    hash: string
    size: number
    name: string
    ext: string
}

export interface CheckResp {
    fileId: number
    status: number
    confirmShard: number
}

export interface CheckRes extends CheckResp {
    success: boolean
    hash: string
}

export interface CheckChunkReq {
    fileId: number
    hash: string
    chunkSeq: number
}

export const uploadConst = {
    shardingSize: 5242880,
    shardingFloor: 10485760,
    codeNeedUpload: 0,
    shardConfirmed: 1
}

export function checkFile(req: CheckReq) {
    return api.post<any, Resp<CheckResp>>('/upload/check', req)
}

export function upload(formData: FormData) {
    return api.post<any, Resp<any>>('/upload', formData, {headers: {'Content-Type': 'multipart/form-data'}})
}

export function checkChunk(req: CheckChunkReq) {
    return api.post<any, Resp<any>>('/upload/check_chunk', req)
}

export function uploadChunk(formData: FormData) {
    return api.post<any, Resp<any>>('/upload/chunk', formData, {headers: {'Content-Type': 'multipart/form-data'}})
}