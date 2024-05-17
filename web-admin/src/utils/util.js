// 负责用户登入进来后，说 上午|中午|下午|晚上|半夜|凌晨|早上 好

import {promptSuccess} from "@/utils/http/base.js";

export const util = () => {
    const hours = new Date().getHours()
    if (hours < 3)
        return '深夜'
    else if (hours < 6)
        return '凌晨'
    else if (hours < 9)
        return '早上'
    else if (hours < 12)
        return '上午'
    else if (hours < 14)
        return '中午'
    else if (hours < 18)
        return '下午'
    else
        return '晚上'
}

export function formatSize(size) {
    const units = ['B', 'KB', 'MB', 'GB', 'TB', 'PB']
    while (size > 1024 && units.length > 0) {
        size /= 1024
        units.shift()
    }
    return (units[0] === 'B' ? size : size.toFixed(2)) + ' ' + units[0]
}

export function formatLeft(expiration) {
    const now = Math.floor(Date.now() / 1000) // 当前时间的 UNIX 时间戳

    if (now > expiration) {
        return "已过期";
    }

    const remainingSeconds = expiration - now;
    const days = Math.floor(remainingSeconds / (24 * 60 * 60));
    const hours = Math.floor((remainingSeconds % (24 * 60 * 60)) / (60 * 60));
    const minutes = Math.floor((remainingSeconds % (60 * 60)) / 60);

    let remainingTime = "";
    if (days > 0) {
        remainingTime += `${days}天`;
    }
    if (hours > 0) {
        remainingTime += `${hours}小时`;
    }
    if (minutes > 0) {
        remainingTime += `${minutes}分钟`;
    }

    return remainingTime.trim();
}

export function formatTime(unix) {
    const date = new Date(unix * 1000), // 将秒转换为毫秒
        year = date.getFullYear(),
        month = date.getMonth() + 1, // 月份从 0 开始，需要加 1
        day = date.getDate(),
        hours = date.getHours(),
        minutes = date.getMinutes(),
        seconds = date.getSeconds()

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

export function formatState(expired) {
    if (expired === 0) {
        return '永久有效'
    }
    const now = new Date().getTime() / 1000
    if (now >= expired - 10) {
        return '已过期'
    }
    return formatLeft(expired) + '后过期'
}

