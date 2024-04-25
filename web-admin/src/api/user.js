import http from '../utils/http/http.js'

const login = (data) => {
    return http.post("/admin/login", data);
};
const getUserList = (data) => {
    return http.post("/admin/user-list", data);
};
const saveUser = (data) => {
    return http.post("/user/save", data);
};
const delUser = (data) => {
    return http.del("/user/delete", data);
};
const getUserDetail = (data) => {
    return http.get("/user/detail", data);
}

function setStatus(id, status) {
    return http.post('/admin/user-status', {
        'id': id,
        'status': status
    })
}

export default {
    login, getUserList, saveUser, delUser, getUserDetail,
    setStatus
}