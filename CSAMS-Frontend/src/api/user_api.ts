import type {baseResponse, paramsType, listDataType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface loginType {
    user_name: string
    password: string
}

export function loginApi(request: loginType): Promise<baseResponse<string>> {
    return useAxios.post("/api/login", request)
}

export function enrollApi(request: loginType): Promise<baseResponse<string>> {
    return useAxios.post("/api/enroll", request)
}

export function logoutApi(): Promise<baseResponse<string>> {
    return useAxios.post("/api/logout")
}

export interface userInfoType {
    id: number
    user_name: string
    password: string
    age: number
    gender: string
    major: string
    avatar: string
    tel: string
    token: string
    role: number
    integrity: number
    score: number
}

export function userInfoApi(): Promise<baseResponse<userInfoType>> {
    return useAxios.get("/api/user_info")
}


export function userListApi(params: paramsType): Promise<baseResponse<listDataType<userInfoType>>> {
    return useAxios.get("/api/users", {params})
}

export interface memberCreateRequest {
    role: number
    user_name: string
}

export function memberCreateApi(data: memberCreateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/association/member", data)
}


export interface memberUpdateRequest {
    role: number
}

export function memberUpdateApi(data: memberUpdateRequest): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_role", data)
}

export interface userInfoUpdateType {
    avatar: string
}

export function userInfoUpdateApi(data: userInfoUpdateType): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_info", data)
}

export interface userUpdatePasswordType {
    old_pwd: string
    pwd: string
    re_pwd: string
}

export function userUpdatePasswordApi(data: userUpdatePasswordType): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_password", data)
}


