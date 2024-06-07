import type {baseResponse, paramsType, listDataType} from "@/api/index";
import {useAxios} from "@/api/index";

export interface loginType {
    id: string
    password: string
}

export interface loginFormType {
    id: number
    password: string
}

export function loginApi(request: loginFormType): Promise<baseResponse<string>> {
    return useAxios.post("/api/login", request)
}

export interface registerType {
    "id": string,
    "password": string,
    "name": string,
    "age": string,
    "gender": string,
    "role": string,
    "major": string,
    "tel": string
}

export interface registerFormType {
    "id": number,
    "password": string,
    "name": string,
    "age": number,
    "gender": string,
    "role": string,
    "major": string,
    "tel": number
}

export function registerApi(request: registerFormType): Promise<baseResponse<string>> {
    return useAxios.post("/api/register", request)
}


export function logoutApi(): Promise<baseResponse<string>> {
    return useAxios.post("/api/logout")
}

export interface userInfoType {
    id: number
    password: string
    name: string
    age: number
    gender: string
    avatar: string
    role: string
    major: number
    tel: number
    integrity: number
    score: number
    associationID: number
}

export function userInfoApi(): Promise<baseResponse<userInfoType>> {
    return useAxios.get("/api/user_info")
}


export function userListApi(params: paramsType): Promise<baseResponse<listDataType<userInfoType>>> {
    return useAxios.get("/api/users", {params})
}

export interface memberCreateRequest {
    id: number
}

export interface memberCreateFormRequest {
    id: string
}

export function memberCreateApi(data: memberCreateRequest): Promise<baseResponse<string>> {
    return useAxios.post("/api/associations/member", data)
}


export interface memberUpdateRequest {
    id: number
    posts: string
}

export function memberUpdateApi(data: memberUpdateRequest): Promise<baseResponse<string>> {
    return useAxios.put("/api/associations/member", data)
}

export interface userInfoUpdateType {
    avatar: string
}

export function userInfoUpdateApi(data: userInfoUpdateType): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_avatar", data)
}

export interface userUpdatePasswordType {
    old_pwd: string
    pwd: string
    re_pwd: string
}

export function userUpdatePasswordApi(data: userUpdatePasswordType): Promise<baseResponse<string>> {
    return useAxios.put("/api/user_password", data)
}


