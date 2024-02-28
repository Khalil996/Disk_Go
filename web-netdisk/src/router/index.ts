import {createRouter, createWebHistory} from 'vue-router'

const routes = [
    {
        path: '/login',
        name: 'login',
        component: () => import('@/views/Login.vue')
    },
    {
        path: '/',
        redirect: 'file/folder/0',
        component: () => import('@/views/Main.vue'),
        children: [
            {
                path: 'file',
                component: () => import('@/views/fileFolder/FileFolder.vue'),
                props: true,
                children: [
                    {
                        path: 'folder/:folderId',
                        component: () => import('@/components/files/Folder.vue'),
                        props: true,
                    },
                    {
                        path: ':fileType',
                        component: () => import('@/components/files/File.vue'),
                        props: true,
                    }
                ]
            },
            {
                path: 'bin',
                component: () => import('@/views/bin/Bin.vue')
            }
        ]
    },
]

const router = createRouter({history: createWebHistory(), routes})

export default router