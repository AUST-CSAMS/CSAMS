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
                }
            ]
        },
        {
            path: '/login',
            name: 'login',
            component: () => import('../views/login/index.vue')
        },
        {
            path: '/enroll',
            name: 'enroll',
            component: () => import('../views/login/enroll.vue'),
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
                    path: "association",
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
                            path: "association_members",
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
                // {
                //     path: "assignments",
                //     name: "active_assignments",
                //     meta: {
                //         title: "活动作业",
                //         isAdmin: true,
                //     },
                //     children: [
                //         {
                //             path: "submit_assignments",
                //             name: "submit_assignments",
                //             meta: {
                //                 title: "提交作业"
                //             },
                //             component: () => import('../views/admin/active_assignments/submit_assignments.vue'),
                //         }
                //     ]
                // },
            ]
        }
    ]
})

export default router

