import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface activitySearchType {
    id: string
    key: string
    title: string
    slug: string
}

export function activitySearchApi(params: paramsType): Promise<baseResponse<listDataType<activitySearchType>>> {
    return useAxios.get("/api/activities", {params})
}

export interface activityRequest {
    id: number
    title: string
    create_at: string
    place: string
    content: string
}

export function activityCreateApi(data: activityRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/activities", data)
}

export interface activityUpdateRequest {
    id: number
    title: string
    create_at: string
    place: string
    content: string
}

export function activityUpdateApi(data: activityUpdateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/activities", data)
}

export function activityDetailApi(id: string): Promise<baseResponse<activityRequest>> {
    return useAxios.get("/api/activities/" + id)
}


export function activityListApi(params: paramsType): Promise<baseResponse<listDataType<activityRequest>>> {
    return useAxios.get("/api/activities", {params: params})
}

