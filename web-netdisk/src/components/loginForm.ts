import api from "../utils/apis/request.ts";
import {Resp} from "../utils/apis/base.ts";
import {UserInfo} from "../store";

export interface LoginReq {
    username: string,
    password: string
}

export interface LoginResp {
    token: string
    userInfo: UserInfo
}

export const loginPost = (loginReq: LoginReq) => {
    return api.post<any, Resp<LoginResp>>("/login", loginReq)
}