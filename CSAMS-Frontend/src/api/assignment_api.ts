import {type baseResponse, type listDataType, useAxios} from "@/api/index";


export interface assignmentType {
    id: number
    content: string
    activity_id: number
    user_id: number
    is_submit: boolean
    is_finish: boolean
    is_correct: boolean
}

export function assignmentListApi(): Promise<baseResponse<listDataType<assignmentType>>> {
    return useAxios.get("/api/assignments")
}

export interface assignmentSubmitType {
    id: number
    content: string
}

export interface assignmentSubmitFormType {
    id: string
    content: string
}


export function assignmentSubmitApi(data: assignmentSubmitType): Promise<baseResponse<string>> {
    return useAxios.put("/api/assignments/submit", data)
}

export interface activityCorrectType {
    id: number
    is_finish: boolean
}

export function assignmentCorrectApi(data: activityCorrectType): Promise<baseResponse<string>> {
    return useAxios.put("assignments/correct", data)
}
