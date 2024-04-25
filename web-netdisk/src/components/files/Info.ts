import api from '../../utils/apis/request.ts';
import type {Resp} from '../../utils/apis/base.ts';
import type {File} from './file.ts'

export function search(str: string) {
    const encodedStr = encodeURIComponent(str);
    return api.get<any, Resp<any>>(`/file/search?phrase=${str}`)
}

export function getFileDetailById(id: number) {
    return api.get<any, Resp<File>>(`/file/${id}`)
}

export function listFolderDetailById(id: number) {
    return api.get<any, Resp<File>>(`/file/list/${id}`)
}
