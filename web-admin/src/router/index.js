import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/',
        redirect: '/login'
    },
    {
        path: '/login',
        name: 'login',
        meta: {
            title: '登录'
        },
        component: () => import('../view/Login.vue')
    },
    {
        path: '/home',
        name: '主页',
        meta: {
            title: '主页'
        },
        component: () => import('../view/Home.vue'),
        redirect: '/index',
        children: [
            {
                path: '/index',
                meta: {
                    title: '首页'
                },
                component: () => import('../view/Welcome.vue')
            },
            {
                path: '/user/list',
                meta: {
                    title: '用户管理'
                },
                component: () => import('../view/user/Index.vue'),
            },
            {
                path: '/user/detail',
                meta: {
                    title: '用户详情'
                },
                component: () => import('../view/user/Detail.vue'),
            },
            {
                path: '/share',
                meta: {
                    title: '分享列表'
                },
                component: () => import('../view/share/Index.vue'),
            },
            {
                path: '/share/:shareId',
                meta: {
                    title: '分享详情'
                },
                component: () => import('../view/share/Info.vue'),
                props: true
            },
            {
                path: '/log',
                meta: {
                    title: '日志'
                },
                component: () => import('../view/log/Index.vue'),
                props: true
            },
            {
                path: '/admin',
                meta: {
                    title: '管理员列表'
                },
                component: () => import('../view/admin/Index.vue'),
                props: true
            }
        ]
    },
]
const router = createRouter({
    history: createWebHistory(),
    routes
})

// 挂载路由导航守卫：to表示将要访问的路径，from表示从哪里来，next是下一个要做的操作
router.beforeEach((to, from, next) => {
    // 修改页面 title
    if (to.meta.title) {
        document.title = '咪咪网盘后台管理系统 - ' + to.meta.title
    }
    // 放行登录页面
    if (to.path === '/login') {
        return next()
    }
    // 获取token
    // const token= sessionStorage.getItem('token')
    // if (!token) {
    //   return next('/login')
    // } else {
    //   next()
    // }
    return next()
})

// 导出路由
export default router