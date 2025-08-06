<script setup>
import axios from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';

const posts = ref()
const totalPage = ref(1)
const presentPage = ref(1)
const limit = ref(10)
const deletedPostTitle = ref("")
let btn = document.getElementById("delete_post_message")

const router = useRouter()

onMounted(()=>{
    listPostRequest()
    btn = document.getElementById("delete_post_message")
    btn.hidden = true
})

function listPostRequest(){
    const offset = limit.value*(presentPage.value-1)
    axios.get("/api/v1/posts?" + "limit=" + limit.value + "&offset=" + offset)
    .then((response)=>{
        posts.value = response.data.posts
        totalPage.value = parseInt((response.data.count-1)/limit.value + 1)
    })
}

watch(presentPage, (newValue, oldValue, cleanUp) => {
    if(presentPage.value > totalPage.value){
        presentPage.value = oldValue
    }else if(presentPage.value < 1) {
        presentPage.value = 1
    }
})

function nextPage(){
    if(presentPage.value < totalPage.value){
        presentPage.value += 1
        listPostRequest()
    }
}

function beforePage(){
    if(presentPage.value > 1){
        presentPage.value -= 1
        listPostRequest()
    }
}

function firstPage(){
    if(presentPage.value != 1){
        presentPage.value = 1
        listPostRequest()
    }
    
}

function lastPage(){
    if(presentPage.value != totalPage.value){
        presentPage.value = totalPage.value
        listPostRequest()
    }
    
}

function updatePost(postId){
    console.log(postId)
    router.push({
        path: "/editpost",
        query: {"postid": postId}
    })
}

function onDeletePost(post){
    let result = confirm("确定删除" + post.title + "吗？")
    if(result){
        deletePostRequest(post.postId)
    }
}

function deletePostRequest(postId){
    axios.delete("/api/v1/posts/" + postId)
    .then((response=>{
        alert("删除成功")
        listPostRequest()
    })).catch((error)=>{
        alert("删除失败")
        console.log(error)
    })
}
</script>

<template>
    <div class="post_item" v-for="post in posts">
        <div class="post_item_head">
            <div style="font-size:x-large;font-weight: bold;align-self: center;justify-self: center;">{{post.title}}</div>
            <div style="align-self:self-end;">（updated at {{ post.updatedAt }}）</div>
            <div style="flex-grow: 1;"></div>

            <button class="edit_post_button" @click="()=>onDeletePost(post)">delete</button>
            <button class="edit_post_button" @click="()=>updatePost(post.postId)">edit</button>
        </div>
        <p class="overflow_hidden_element">{{ post.content }}</p>
        
    </div>
    <div class="list_post_form">
        <button class="set_page_button" @click="firstPage"><<</button>
        <button class="set_page_button" @click="beforePage"><</button>
        <p style="margin-left: 20px;">page: </p>
        <input name="page" style="width: 25px;flex-grow: 0;height: 18px;text-align: center;font-size: large;" v-model="presentPage" type="number" min="1" :max="totalPage"></input>
        <p style="margin-right: 20px;">/{{ totalPage }}</p>
        <button class="set_page_button" @click="nextPage">></button>
        <button class="set_page_button" @click="lastPage">>></button>
    </div>

    <div id="delete_post_message">
        <p>确定要删除{{deletedPostTitle}}吗？</p>
        <button class="delete_post_message_button">yes</button>
        <button class="delete_post_message_button">no</button>
    </div>
</template>


<style>
#delete_post_message {
    position:fixed !important;
    top:35%;
    border: 2px solid black;
    border-radius: 4px;
    padding: 3px;
    background-color: white;
}

.delete_post_message_button {
    border: 1px solid black;
    border-radius: 1px;
    width: 70px;
    margin-left: 2px;
}

.overflow_hidden_element {
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    margin: 1px;
    text-align: left;
}

.edit_post_button {
    border: 1px solid black;
    border-radius: 1px;
    width: 70px;
    margin-left: 2px;
}

.edit_post_button:hover {
    background-color: black;
    color: white;
}


.set_page_button {
    border: 1px solid black;
    border-radius: 1px;
    width: 25px;
    margin: 1px;
}

.post_item_head {
    display: flex;
    flex-direction: row;
    padding: 2px;
}

.post_item {
    width: 1000px;
    max-width: 83%;
    height: 80px;
    border: 2px solid white;
    box-sizing: border-box;
}

.post_item:hover {
    border: 0.5px solid black;
    box-sizing: border-box;
}

.list_post_form {
    width: 90%;
    height: 50px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    font-size: large;
}


</style>