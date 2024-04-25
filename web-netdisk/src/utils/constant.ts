export const
    TypeDocs = 0,
    typeImage = 1,
    typeVideo = 2,
    typeAudio = 3,
    typeOther = 4,
    typeMulti = 5

export const typeMap = {
    0: '文档',
    1: '图片',
    2: '视频',
    3: '音频',
    4: '其他',
    5: '多文件',
}

export const
    shareNotExpired = 0,
    shareExpired = 1,
    shareIllegal = 2,
    shareNotExistOrDeleted = 3,

    userStatus = {
        ok: 0,
        bannedByAvatar: 1,
        bannedByUsername: 2,
        bannedByName: 3,
        bannedBySignature: 4,
        bannedByShare: 5
    },

    expireType = {
        day: 0,
        day7: 1,
        month: 2,
        forever: 3
    }


export const userMap = {
    0: '正常',
    1: '头像违规，暂时封禁',
    2: '帐号包含违规信息，暂时封禁',
    3: '昵称包含违规信息，暂时封禁',
    4: '签名包含违规信息，暂时封禁',
    5: '分享内容违规，暂时封禁',
}

export const fileStatus = {
    unuploaded: 0,
    ok: 1,
    undeleted: 2,
    deleted: 3,
    banned: 4,
    needMerge: 5
}

export const fileStatusMap = {
    0: '未完成上传',
    1: '可下载',
    2: '未删除',
    3: '已被删除',
    4: '已封禁',
    5: '服务器处理中'
}
