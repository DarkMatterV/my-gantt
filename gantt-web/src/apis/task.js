import request from '@/utils/request'

export function getTaskList() {
    return request({
        url: '/task/list',
        method: 'get'
    })
}

export function addTask(data) {
    return request({
        url: '/task/add',
        method: 'post',
        data
    })
}

export function updateTask() {
    return request({
        url: '/date/update',
        method: 'put'
    })
}
