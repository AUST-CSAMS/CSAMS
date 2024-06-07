import {type baseResponse, type listDataType, type paramsType, useAxios} from "@/api/index";

export interface associationInfoType {
    id: number
    association_name: string
    create_at: string
    teacher_name: string
    president: string
    introduction: string
}

export function associationInfoApi(): Promise<baseResponse<associationInfoType>> {
    return useAxios.get("/api/association_info")
}


export interface associationInfoUpdateType {
    name: string
    teacher: string
    president: string
    introduction: string
}


export function associationListApi(params: paramsType): Promise<baseResponse<listDataType<associationInfoType>>> {
    return useAxios.get("/api/associations", {params})
}

export interface associationCreateType {
    id: number
    association_name: string
    introduction: string
}

export interface associationCreateFormType {
    id: string
    association_name: string
    introduction: string
}

export function associationCreateApi(data: associationCreateType): Promise<baseResponse<string>> {
    return useAxios.post("/api/associations", data)
}

export interface associationMemberType {
    user_id: number
    association_id: number
    posts: string
    joining_time: string

}


export function associationMemberListApi(params?: paramsType): Promise<baseResponse<listDataType<associationMemberType>>> {
    return useAxios.get("/api/associations/member", {params: params})
}

export interface associationDeleteType {
    id: number[]
}

export function associationMemberDeleteApi(data: associationDeleteType): Promise<baseResponse<string>> {
    return useAxios.delete("/api/associations/member", {data})
}

