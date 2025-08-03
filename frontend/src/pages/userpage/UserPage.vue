<script setup>
import {ref} from 'vue'
import axios from 'axios';
import { useRoute, useRouter, RouterLink, RouterView } from 'vue-router'

const token = localStorage.getItem("jwt_token")
const username=ref("")
const nickname=ref("")
const phone=ref("")
const email=ref("")

function getRequest(){
    axios.get("/", token)
    .then(function(response){
        if(response.status == axios.HttpStatusCode.Ok){
            username = response.data.username
            nickname = response.data.nickname
            phone = response.data.phone
            email = response.data.email
        }else{
            localStorage.removeItem("jwt_token")
            alert("网络错误，请重新登陆")
            useRouter().push("/login")
        }
    })
    .catch(function(err){
        console.log(err)
    })
}
</script>

<template>
<div style="min-width: 100%;width: max-content;min-height: 100%;height: max-content;display:flex;flex-direction: row;font-size: large;">
    <div class = "user_tool_bar">
        <RouterLink to="./userinfo"><button @click="getRequest">profile</button></RouterLink>
        <RouterLink to="./postsinfo"><button @click="getRequest">posts</button></RouterLink>
    </div>
    
    <RouterView></RouterView>
</div>

</template>


<style>



.user_tool_bar {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
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


</style>