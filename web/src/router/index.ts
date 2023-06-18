import { createRouter, createWebHistory } from "vue-router"

const routes = [
    {
        name: "登录",
        path: "/login",
        component: () => import("../views/login/Login.vue")
    },
    {
        name: "home",
        path: "/",
        redirect: "/home",
        component: () => import("../layout/index.vue"),
        children: [{
            name: "操作历史",
            path: "/operationLog",
            component: () => import("../views/operationLog/OperationLog.vue")
        },
        {
            name: "首页",
            path: "/home",
            component: () => import("../views/dashboard/index.vue")
        },
        {
            name: "设置",
            path: "/setting",
            component: () => import("../views/system/Setting.vue")
        },
        {
            name: "桶",
            path: "/bucket",
            component: () => import("../views/bucket/Bucket.vue")
        },
        {
            name: "桶视图",
            path: "/bucketView",
            component: () => import("../views/bucket/BucketView.vue")
        },
        {
            name: "照片视图",
            path: "/photoView",
            component: () => import("../views/photo/PhotoView.vue")
        },
        {
            name: "更新日志",
            path: "/uplog",
            component: () => import("../views/system/UpLog.vue")
        },




        {
            name: "测试",
            path: "/test",
            component: () => import("../views/test/Test.vue")
        },
        ]
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router

