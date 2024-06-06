/// <reference types="vite/client" />

//使env可以点出VITE_SERVER_URL console.log(import.meta.env.VITE_SERVER_URL)
export interface ImportMetaEnv {
    VITE_SERVER_URL: string
}

//使在router/index.ts能拿到meta
import 'vue-router';

declare module 'vue-router' {
    interface RouteMeta {
        isLogin?: boolean
        isAdmin?: boolean
        title?: string
    }
}
