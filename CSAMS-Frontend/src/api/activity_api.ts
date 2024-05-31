import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface activitySearchType {
    title: string
    slug: string
}

export function activitySearchApi(params: paramsType): Promise<baseResponse<listDataType<activitySearchType>>> {
    return useAxios.get("/api/activities", {params})
}

export interface activityCreateRequest {
    name: string
    create_at: string
    place: string
    content: string
}

export function activityCreateApi(data: activityCreateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/association/member", data)
}

export interface activityUpdateRequest {
    name: string
    create_at: string
    place: string
    content: string
}

export function activityUpdateApi(data: activityUpdateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/activity", data)
}

export function activityDetailApi(id: string): Promise<baseResponse<activityCreateRequest>> {
    return useAxios.get("/api/articles/" + id)
}
