<script setup>
import { useRoute, useRouter, RouterLink, RouterView } from 'vue-router'
import { useAuthStore } from '../stores/user'
import { onBeforeMount, onMounted } from 'vue'


const router = useRouter()
const authStore = useAuthStore()


onMounted(() => {
    if(!authStore.isAuthenticated){
        router.push("/")
    }
})

function logout(){
    authStore.clearUserInfo()
}
</script>

<template>
<div style="width: 100%;display:flex;flex-direction: row;font-size: large;">
    <div class = "user_tool_bar">
        <div class="user_tool_bar_top">
            <RouterLink to="./userinfo"><button>profile</button></RouterLink>
            <RouterLink to="./postsinfo"><button>posts</button></RouterLink>
        </div>
        <RouterLink to="./login"><button @click="logout" class="user_tool_bar_logout_button">logout</button></RouterLink>
    </div>
    <div style="height: 100%;background-color: black;width: 1px;"></div>
    <div class = "userpage_content">
        <RouterView></RouterView>
    </div>
</div>

</template>


<style>

.user_tool_bar_top {
    display: flex;
    flex-direction: column;
    justify-content: start;
}

.user_tool_bar_logout_button {
    background-color: rgb(255, 191, 191) !important;
}

.user_tool_bar_logout_button:hover {
    background-color: rgb(0, 0, 0) !important;
}

.user_tool_bar {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    /* flex-grow: 1; */
    width: 17%;
}

.user_tool_bar button{
    background-color: rgb(255, 255, 255);
    color: rgb(0, 0, 0);
    height: 80px;
    width: 100%;
    /* box-sizing: border-box;
    border-color: black;
    border-style: solid;
    border-width: 3px; */
}

.user_tool_bar button:hover{
    background-color: black;
    color: white;
}

.userpage_content {
    width: 83%;
    display: flex;
    flex-direction: column;
    flex-grow: 5;
    align-items: center;
    background-color: rgb(255, 255, 255);
    padding: 20px;
    
}


</style>