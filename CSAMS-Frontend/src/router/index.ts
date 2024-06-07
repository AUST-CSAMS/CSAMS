import {createRouter, createWebHistory} from 'vue-router'

const router = createRouter({
    //createWebHistory 创建HTML5历史记录
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'web',
            component: () => import('../views/web/web.vue'),
            children: [
                {
                    path: '',
                    name: 'index',
                    component: () => import('../views/web/index.vue'),
                }, {
                    path: 'association',
                    name: 'association',
                    component: () => import('../views/web/association.vue'),
                },
                {
                    path: "activities/:id",
                    name: "activities",
                    component: () => import('../views/web/activity.vue'),
                },
            ]
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/login/loginindex.vue')
        },
        {
            path: '/register',
            name: 'register',
            component: () => import('../views/login/registerindex.vue'),
        },
        {
            path: '/admin',
            name: 'admin',
            component: () => import('../views/admin/index.vue'),
            meta: {
                title: "首页",
                isLogin: true,
            },
            children: [
                {
                    path: "",
                    name: "home",
                    meta: {
                        title: "Home"
                    },
                    component: () => import('../views/admin/home/index.vue'),
                },
                {
                    path: "user_center",
                    name: "user_center",
                    meta: {
                        title: "个人中心"
                    },
                    children: [
                        {
                            path: "user_info",
                            name: "user_info",
                            meta: {
                                title: "我的信息"
                            },
                            component: () => import('../views/admin/user_center/user_info.vue'),
                        },
                    ]
                },
                {
                    path: "associations",
                    name: "my_association",
                    meta: {
                        title: "我的协会",
                        isAdmin: true,
                    },
                    children: [
                        {
                            path: "association_info",
                            name: "association_info",
                            meta: {
                                title: "协会信息"
                            },
                            component: () => import('../views/admin/my_association/association_info.vue'),
                        },
                        {
                            path: "member",
                            name: "association_members",
                            meta: {
                                title: "协会成员"
                            },
                            component: () => import('../views/admin/my_association/association_members.vue'),
                        }
                    ]
                },
                {
                    path: "activities",
                    name: "my_activities",
                    meta: {
                        title: "我的活动",
                        isAdmin: true,
                    },
                    children: [
                        {
                            path: "get_activities",
                            name: "get_activities",
                            meta: {
                                title: "参与活动"
                            },
                            component: () => import('../views/admin/my_activities/get_activities.vue'),
                        },
                    ]
                },
                {
                    path: "assignments",
                    name: "assignments",
                    meta: {
                        title: "活动作业",
                        isAdmin: true,
                    },
                    children: [
                        {
                            path: "assignment_list",
                            name: "assignment_list",
                            meta: {
                                title: "作业列表"
                            },
                            component: () => import('../views/admin/assignment/assignment_list.vue'),
                        }
                    ]
                },
            ]
        }
    ]
})

export default router

