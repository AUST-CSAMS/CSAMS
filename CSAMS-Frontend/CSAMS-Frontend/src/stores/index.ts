//pinia 前端状态管理库
import {defineStore} from 'pinia'
import {Message} from "@arco-design/web-vue";
import {parseToken} from "@/utils/jwt";
import {logoutApi, userInfoApi} from "@/api/user_api";
import {associationInfoApi, type associationInfoType} from "@/api/association_api";

export interface userStoreInfoType {
    user_name: string
    role: number
    user_id: number
    avatar: string
    token: string
    exp: number // 过期时间(毫秒)
}

const userInfo: userStoreInfoType = {
    user_name: "",
    role: 0,
    user_id: 0,
    avatar: "",
    token: "",
    exp: 0,
}

const associationInfo: associationInfoType = {
    name: "",
    avatar: "",
    create_at: "",
    teacher: "",
    president: "",
    introduction: ""
}


export const useStore = defineStore('counter', {
    state() {
        return {
            userInfo: userInfo,
            associationInfo: associationInfo,
        }
    },
    actions: {
        async setToken(token: string) {
            this.userInfo.token = token
            let res = await userInfoApi()
            let info = parseToken(token)
            let data = res.data
            this.userInfo = {
                user_name: data.user_name,
                role: info.role,
                user_id: info.user_id,
                avatar: data.avatar,
                token: token,
                exp: info.exp,
            }
            localStorage.setItem("userInfo", JSON.stringify(this.userInfo))
        },
        setUserInfo(key: "avatar", val: string) {
            this.userInfo[key] = val
            localStorage.setItem("userInfo", JSON.stringify(this.userInfo))
        },
        loadToken() {
            let val = localStorage.getItem("userInfo")
            if (val === null) {
                return
            }
            try {
                this.userInfo = JSON.parse(val)
            } catch (e) {
                this.clearUserInfo()
                return;
            }
            let exp = Number(this.userInfo.exp) * 1000
            let nowTime = new Date().getTime()
            if (exp - nowTime <= 0) {
                Message.warning("登录已过期")
                this.clearUserInfo()
                return;
            }
        },
        async logout() {
            await logoutApi()
            this.clearUserInfo()
        },
        clearUserInfo() {
            this.userInfo = userInfo
            localStorage.removeItem("userInfo")
        },
        async loadAssociationInfo() {
            const val = sessionStorage.getItem("associationInfo")
            if (val !== null) {
                try {
                    this.associationInfo = JSON.parse(val)
                    return
                } catch (e) {

                }
            }
            let res = await associationInfoApi()
            this.associationInfo = res.data
            sessionStorage.setItem("associationInfo", JSON.stringify(this.associationInfo))
        }

    },
    getters: {
        isLogin(): boolean {
            return this.userInfo.role !== 0
        },
        isAdmin(): boolean {
            return this.userInfo.role == 1
        },
        isGeneral(): boolean {
            return this.userInfo.role == 2
        },
        isTeacher(): boolean {
            return this.userInfo.role == 3
        }
    }
})
