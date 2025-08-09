<script setup>
import axios from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const posts = ref()
const totalPage = ref(1)
const presentPage = ref(1)
const limit = ref(10)
const deletedPostTitle = ref("")
let btn = document.getElementById("delete_post_message")

const router = useRouter()
const route = useRoute()

onMounted(()=>{
    const searchValue = route.query.text
    if(searchValue){
        searchPostRequest(searchValue)
    }
})

function searchPostRequest(text){
    const offset = limit.value*(presentPage.value-1)
    const url = "/api/v1/posts/search?text=" + text + "&limit=" + limit.value + "&offset=" + offset
    axios.get(url)
    .then((response)=>{
        posts.value = response.data.posts
        totalPage.value = parseInt((response.data.totalCount-1)/limit.value + 1)
    }).catch(err=>{
        console.log(err)
        alert(err.response.data.message)
        router.push("/")
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

</script>

<template>
    <div class="search_post_content">
        <div class="post_item" v-for="post in posts">
            <div class="post_item_head">
                <div style="font-size:x-large;font-weight: bold;align-self: center;justify-self: center;">{{post.title}}</div>
                <div style="align-self:self-end;">（updated at {{ post.updatedAt }}）</div>
                <div style="flex-grow: 1;"></div>

                <!-- <button class="edit_post_button" @click="()=>onDeletePost(post)">delete</button>
                <button class="edit_post_button" @click="()=>updatePost(post.postId)">edit</button> -->
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
    </div>
</template>


<style>
.search_post_content {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
}
</style>