import request from '@/utils/request'

export function getDateList() {
    return request({
        url: '/date/list',
        method: 'get'
    })
}

export function addDate() {
    return request({
        url: '/date/add',
        method: 'post'
    })
}

export function updateDate() {
    return request({
        url: '/date/update',
        method: 'put'
    })
}
