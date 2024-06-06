import {useAxios} from "@/api/index";
import type {baseResponse} from "@/api/index";

export interface imageType {
    path: string
}

export interface imagesUploadResponse {
    file_name: string
    is_success: boolean
    msg: string
}

export function uploadImageApi(file: File): Promise<baseResponse<string>> {
    return new Promise((resolve, reject) => {
        const form = new FormData()
        form.set("image", file)
        useAxios.post("/api/image", form, {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        }).then((res: any) => resolve(res)).catch(err => reject(err))
    })
}


