<script setup>
import axios from 'axios';
import { onMounted, ref } from 'vue';
import { useAuthStore } from '../stores/user';
import { useRoute, useRouter } from 'vue-router';


const authStore=useAuthStore()
const router=useRouter()
const route=useRoute()

const title = ref("")
const content = ref("")
const postId = ref(undefined)

onMounted(() => {
    postId.value = route?.query?.postid
    if(postId.value){
        getPostRequest(postId.value)
    }
    if(!authStore.isAuthenticated){
        router.push("/")
    }
    document.getElementById("blog_edit_form").addEventListener("submit", function(event) {
        event.preventDefault();
    });
})

function submitPost(){
    if(postId.value){
        updatePostRequest(postId.value)
    }else{
        createPostRequest()
    }
    router.push("/")
}

function createPostRequest(){
    const form = document.getElementById("blog_edit_form")
    if(!form.checkValidity()){
        return
    }
    const data = new FormData(form)
    axios.post("/api/v1/posts", Object.fromEntries(data.entries()))
    .then((response)=>{
        console.log(response)
    }).catch((error)=>{
        console.log(error)
    })
}

function updatePostRequest(postId){
    const form = document.getElementById("blog_edit_form")
    if(!form.checkValidity()){
        return
    }
    const data = new FormData(form)
    axios.put("/api/v1/posts/" + postId, Object.fromEntries(data.entries()))
    .then((response)=>{
        console.log(response)
    }).catch((error)=>{
        console.log(error)
    })
}

function getPostRequest(postId){
    axios.get("/api/v1/posts/" + postId)
    .then((response => {
        title.value = response.data.title
        content.value = response.data.content
    })).catch(error=>{
        console.log(error)
        alert("获取post内容失败")
    })
}
</script>


<template>
    <form id="blog_edit_form">
        <input class="blog_title_editor" v-model="title" name="title" placeholder="title" minlength="5"></input>
        <textarea class="blog_content_editor" v-model="content" name="content" placeholder="content"></textarea>
        <button class="create_blog_button" type="submit" @click="submitPost">Create Blog</button>
    </form>
</template>

<style>
#blog_edit_form {
    display: flex;
    flex-direction: column;
    flex-grow: 1;
    padding-left: 10%;
    padding-right: 10%;
    align-items: center;
    justify-content: center;
    font-size: large;
    width: 80%;
}

.blog_title_editor {
    font-size:x-large;
    width: 100%;
}

.blog_content_editor {
    height: 70%;
    width: 100%;
    margin: 10px;
}

.create_blog_button {
    background-color: white;
    color: black;
    border-color: black;
    border-width: 1px;
    border-radius: 0px;
    border-style: solid;
    text-align: center;
    width: 100%;
    height: 35px;
}

.create_blog_button:hover{
    background-color: black;
    color: white;
}

</style>