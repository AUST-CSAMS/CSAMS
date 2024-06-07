import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface activitySearchType {
    id: string
    key: string
    title: string
    slug: string
}

export function activitySearchApi(params: paramsType): Promise<baseResponse<listDataType<activitySearchType>>> {
    return useAxios.post("/api/activities", {params})
}

export interface activityRequest {
    id: number
    activity_name: string
    startTime: string
    endTime: string
    location: string
    introduction: string
    image: string
    responsible_person: string
    tel: number
    score: number
    limit: number[]
}

export interface activityCreateRequest {
    id: number
    activity_name: string
    startTime: string
    endTime: string
    location: string
    introduction: string
    image: string
    score: number
    limit: string[]
}

export interface activityCreateFormRequest {
    id: string
    activity_name: string
    startTime: string
    endTime: string
    location: string
    introduction: string
    image: string
    score: string
    limit: string
}

export function activityCreateApi(data: activityCreateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/activities", data)
}


export function activityFinishedListApi(params: paramsType): Promise<baseResponse<listDataType<activityRequest>>> {
    return useAxios.get("/api/activities/finished", {params: params})
}

export function activityOngoingListApi(params: paramsType): Promise<baseResponse<listDataType<activityRequest>>> {
    return useAxios.get("/api/activities/ongoing", {params: params})
}

export function activityUpcomingListApi(params: paramsType): Promise<baseResponse<listDataType<activityRequest>>> {
    return useAxios.get("/api/activities/upcoming", {params: params})
}

export function activityListApi(params: paramsType): Promise<baseResponse<listDataType<activityRequest>>> {
    return useAxios.get("/api/activities", {params: params})
}

export function activityListTeacherApi(params: paramsType): Promise<baseResponse<listDataType<activityRequest>>> {
    return useAxios.get("/api/activities_create", {params: params})
}

export interface activityJoinType {
    id: number
}

export function activityJoinApi(data: activityJoinType): Promise<baseResponse<string>> {
    return useAxios.post("/api/activities/join", data)
}


export function activityInfoApi(id: number): Promise<baseResponse<activityRequest>> {
    return useAxios.get("/api/activity_info/" + id)
}

export function activityEndApi(data: activityJoinType): Promise<baseResponse<activityRequest>> {
    return useAxios.put("/api/activities/end", data)
}



