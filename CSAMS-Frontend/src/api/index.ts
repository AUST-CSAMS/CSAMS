import axios from "axios";
import {useStore} from "@/stores";
import {Message} from "@arco-design/web-vue";

export const useAxios = axios.create({
    baseURL: ""
})

export interface baseResponse<T> {
    code: number
    data: T
    msg: string
}

export interface listDataType<T> {
    count: number
    list: T[]
}

export interface paramsType {
    page?: number
    limit?: number
    key?: string
    sort?: string
}

useAxios.interceptors.request.use((config) => {
    const store = useStore()
    config.headers["token"] = store.userInfo.token
    return config
})

useAxios.interceptors.response.use((response) => {
    if (response.status !== 200) {
        console.log("服务失败", response.status)
        Message.error(response.statusText)
        return Promise.reject(response.statusText)
    }
    return response.data
}, (err) => {
    console.log("服务错误", err)
    Message.error(err.message)
    return Promise.reject(err.message)
})

export function defaultDeleteApi<T extends number | string>(url: string, idList: T[]): Promise<baseResponse<string>> {
    return useAxios.delete(url, {
        data: {
            id: idList,
        }
    })
}




