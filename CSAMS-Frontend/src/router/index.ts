import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    //createWebHistory 创建HTML5历史记录
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: []
})

export default router
