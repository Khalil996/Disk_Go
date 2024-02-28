import {ElMessage} from "element-plus";

interface Resp<T> {
    code: number
    msg: string
    data: T
}

const codeOk = 0
const codeError = -1
const msgOk = '操作成功！😻'

function promptSuccess(prompt?: string) {
    ElMessage({
        type: "success",
        message: prompt || msgOk,
    })
}

function promptError(prompt?: string) {
    ElMessage({
        type: "error",
        message: prompt || msgOk,
    })
}

export type {
    Resp
}

export {
    codeOk, codeError, msgOk,
    promptSuccess, promptError
}