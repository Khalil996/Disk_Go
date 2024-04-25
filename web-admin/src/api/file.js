import http from "@/utils/http/http.js";

function setStatus(obj) {
    return http.post('/admin/file-status', {
        'ids': obj.ids,
        'status': obj.status
    })
}

export default {
    setStatus
}