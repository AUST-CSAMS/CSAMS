import {type baseResponse, type listDataType, useAxios} from "@/api/index";
import type {associationInfoType} from "@/api/association_api";


export interface assignmentInfoType {
    id: number
    content: string
}

export function assignmentListApi(): Promise<baseResponse<listDataType<assignmentInfoType>>> {
    return useAxios.get("/api/assignment")
}

export function assignmentSubmitApi(data: assignmentInfoType): Promise<baseResponse<string>> {
    return useAxios.post("/api/assignment", data)
}
