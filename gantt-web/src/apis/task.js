import request from '@/utils/request'

export function getTaskList() {
    return request({
        url: '/task/list',
        method: 'get'
    })
}

export function addTask() {
    return request({
        url: '/task/add',
        method: 'post'
    })
}

export function updateTask() {
    return request({
        url: '/task/update',
        method: 'put'
    })
}
