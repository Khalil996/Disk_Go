import {ElMessage} from "element-plus";

const codeOk = 0
const codeError = -1
const msgOk = '操作成功！😻'

function promptSuccess(prompt) {
    ElMessage({
        type: "success",
        message: prompt || msgOk,
    })
}

function promptError(prompt) {
    ElMessage({
        type: "error",
        message: prompt || msgOk,
    })
}

export {
    codeOk, codeError, msgOk,
    promptSuccess, promptError
}