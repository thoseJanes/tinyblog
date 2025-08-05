import { createApp, ref } from 'vue'
import './style.css'
import App from './App.vue'
import { createRouter, createWebHashHistory } from 'vue-router';
import Login from './components/Login.vue';
import UserPage from './components/UserPage.vue';
import UserInfo from './components/UserInfo.vue';
import PostsInfo from './components/PostsInfo.vue';
import MainPage from './components/MainPage.vue';

import { createPinia } from 'pinia';

const app = createApp(App)

const routes = [
    { path: '/', component: MainPage},
    { path: '/login', component: Login},
    { path: '/userpage', component: UserPage,
        props: ['login'],
        children: [
            { path: '/userinfo', component: UserInfo },
            { path: '/postsinfo', component: PostsInfo },
            { path: '/:pathMatch(.*)*', redirect: '/userinfo'},
        ]
    }
]
   

const router = createRouter({
    history: createWebHashHistory('/index.html'),
    routes,
})

app.use(router)
app.use(createPinia())

app.mount('#app')

