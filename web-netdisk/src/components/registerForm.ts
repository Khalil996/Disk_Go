import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";

export interface RegisterReq {
    name: string,
    password: string,
    email: string,
    code: string,
}

export interface RegisterResp {
}

export const registerPost = (registerReq: RegisterReq): Resp<RegisterResp> => {
    return api.post<any, Resp<RegisterResp>>("/register", registerReq)
}
export const sendEmailPost = (email: string): Resp<RegisterResp> => {
    return api.post<any, Resp<RegisterResp>>("/EmailSend", {email})
}