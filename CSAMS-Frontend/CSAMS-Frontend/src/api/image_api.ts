import {useAxios} from "@/api/index";
import type {baseResponse, listDataType, paramsType} from "@/api/index";
import {Message} from "@arco-design/web-vue";

export interface imageIdType {
    id: number
    path: string
    name: string
}


export function imageIdListApi(): Promise<baseResponse<imageIdType[]>> {
    return useAxios.get("/api/image_names")
}


export interface imageType {
    id: number
    created_at: string
    path: string
    hash: string
    name: string
    image_type: "本地" | "七牛云",
    bannerCount: number // 关联banner的个数
    articleCount: number // 关联文章的个数
}

export function imageListApi(params: paramsType): Promise<baseResponse<listDataType<imageType>>> {
    return useAxios.get("/api/images", {params: params})
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


export async function onUploadImg(files: Array<File>, callback: (urls: Array<string>) => void): Promise<void> {
    let resList: baseResponse<string>[] = []

    try {
        //使用map方法遍历files数组中的每个元素file，对每个file调用uploadImageApi(file)函数，返回一个Promise数组
        //Promise.all()方法接收一个Promise数组作为参数，并返回一个新的Promise。这个新的Promise在传入的所有Promise
        //都变为fulfilled后才会变为fulfilled，或者只要有一个 Promise变为rejected，那么这个新的Promise就会变为rejected
        resList = await Promise.all(files.map(file => uploadImageApi(file)))
    } catch (e) {
        return
    }

    const urlList: string[] = []
    //遍历每一个文件
    resList.forEach(res => {
        if (res.code) {
            Message.error(res.msg)
            return
        }
        urlList.push(res.data)
    })
    callback(urlList)
}
