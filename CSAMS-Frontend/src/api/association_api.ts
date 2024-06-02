import {type baseResponse, useAxios} from "@/api/index";

export interface associationInfoType {
    name: string
    avatar: string
    create_at: string
    teacher: string
    president: string
    introduction: string
}

export function associationInfoApi(): Promise<baseResponse<associationInfoType>> {
    return useAxios.get("/api/association")
}


export interface associationInfoUpdateType {
    name: string
    avatar: string
    teacher: string
    president: string
    introduction: string
}

export function associationInfoUpdateApi(data: associationInfoUpdateType): Promise<baseResponse<string>> {
    return useAxios.put("/api/association", data)
}
