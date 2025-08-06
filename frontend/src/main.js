import { createApp, ref } from 'vue'
import './style.css'
import App from './App.vue'
import { createRouter, createWebHashHistory } from 'vue-router';

import Login from './components/Login.vue';
import UserPage from './components/UserPage.vue';
import UserInfo from './components/UserInfo.vue';
import PostsInfo from './components/PostsInfo.vue';
import MainPage from './components/MainPage.vue';
import EditPost from './components/EditPost.vue';

import { useAuthStore } from './stores/user';
import { createPinia } from 'pinia';
import axios from 'axios';

const app = createApp(App)
const pinia = createPinia()

const routes = [
    { path: '/', component: MainPage},
    { path: '/login', component: Login},
    { path: '/editpost', component: EditPost},
    { path: '/userpage', component: UserPage,
        children: [
            { path: '/userinfo', component: UserInfo },
            { path: '/postsinfo', component: PostsInfo },
            { path: '/:pathMatch(.*)*', redirect: '/userinfo'},
        ]
    },
]


const router = createRouter({
    history: createWebHashHistory('/index.html'),
    routes,
})

app.use(router)

app.use(pinia)
const authStore = useAuthStore()

if (authStore.token) {
  axios.defaults.headers.common['Authorization'] = `Bearer ${authStore.token}`
}

app.mount('#app')

