import { createApp, ref } from 'vue'
import './style.css'
import { createRouter, createWebHashHistory } from 'vue-router';

import UserPage from './pages/userpage/UserPage.vue'
import PostsInfo from './pages/userpage/PostsInfo.vue'
import UserInfo from './pages/userpage/UserInfo.vue'


const app = createApp(UserPage)

const routes = [
    { path: '/userinfo', component: UserInfo },
    { path: '/postsinfo', component: PostsInfo },
    { path: '/:pathMatch(.*)*', redirect: '/userinfo'},
    // { path: '/userpage.html/', component: UserInfo},
]
   
// // 3. 创建路由实例并传递 `routes` 配置
// // 你可以在这里输入更多的配置，但我们在这里
// // 暂时保持简单

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory('/userpage.html'),
    routes, // `routes: routes` 的缩写
})

app.use(router)

app.mount('#app')

