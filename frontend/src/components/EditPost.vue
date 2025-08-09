<script setup>
import axios from 'axios';
import { onMounted, ref } from 'vue';
import { useAuthStore } from '../stores/user';
import { useRoute, useRouter } from 'vue-router';
import { EventSourcePolyfill } from 'event-source-polyfill';

const authStore=useAuthStore()
const router=useRouter()
const route=useRoute()

const title = ref("")
const content = ref("")
const polishingContent = ref("")
const polishingPrompt = ref("")
const postId = ref(undefined)

const operationType = ref("Create Blog")

const isPolishing = ref(false)
var eventSource = null
onMounted(() => {
    postId.value = route?.query?.postid
    if(postId.value){
        getPostRequest(postId.value)
        operationType.value = "Update Blog"
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
function polishContent(){
    isPolishing.value = true
    polishingContent.value = content.value
}
function cancelPolishing(){
    isPolishing.value = false
    if(eventSource){
        eventSource.close()
    }
}
function confirmPolishing(){
    isPolishing.value = false
    content.value = polishingContent.value
    polishingContent.value = ""
    if(eventSource){
        eventSource.close()
    }
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
function polishContentRequest(){
    polishingContent.value = ""
    const url = "/api/v1/ai/polish-content?content=" + content.value + "&prompt=" + polishingPrompt.value
    eventSource = new EventSourcePolyfill(url, {
        heartbeatTimeout: 30000,
        headers:{
            "Authorization" : "Bearer " + authStore.token,
            'X-Custom-Header': 'value'
        }
    })

    eventSource.onmessage = (event) => {
        try {
            // const data = JSON.parse(event.data)
            console.log(event)
            polishingContent.value += event.data
            // if(event.type == "end"){
            //     eventSource.close()
            // }
            //polishingContent.value.push(data) // 自动触发 Vue 响应式更新
        } catch (err) {
            console.error('SSE 数据解析失败:', err)
        }
    }

    eventSource.onerror = (err) => {
        console.error('SSE 连接错误:', err)
        if(eventSource){
            eventSource.close()
            eventSource = null
        }
        // 自动重连逻辑可在此添加
    }
}
function generateTitleRequest(){
    const url = "/api/v1/ai/generate-title?content=" + content.value
    axios.get(url)
    .then(response=>{
        const t = response.data.title
        title.value = t
    }).catch(err=>{
        console.log(err)
    })
}
</script>


<template>
    <form id="blog_edit_form">
        <div class="blog_edit_hbox">
            <input class="blog_title_editor" v-model="title" name="title" placeholder="title" minlength="1" maxlength="255"></input>
            <button class="blog_edit_button" type="submit" style="margin-left: 10px;flex-grow: 2;" @click="generateTitleRequest">Generate</button>
        </div>
        <div class="blog_edit_hbox" style="height: 80%;">

            <textarea class="blog_content_editor" v-model="content" name="content" placeholder="content" style="width: 0px;"></textarea>

            <p v-if="isPolishing" style="margin: 5px;">-></p>
            <div v-if="isPolishing" class="blog_polish_vbox">
                <div style="display: flex;flex-direction: row;">
                    <input  class="blog_edit_prompt" v-model="polishingPrompt" placeholder="prompt"></input>
                    <button class="blog_edit_button" style="flex-grow: 1;margin-left: 10px;" @click="polishContentRequest">Polish</button>
                </div>
                <textarea class="blog_content_editor" v-model="polishingContent" placeholder="polished content"></textarea>
            </div>
        </div>
        <div class="blog_edit_hbox">
            <button class="blog_edit_button" type="submit" style="flex-grow: 2;" @click="submitPost">{{ operationType }}</button>
            <button v-if="!isPolishing" class="blog_edit_button" type="submit" style="margin-left: 10px;flex-grow: 2;" @click="polishContent">Polish Blog</button>
            <div v-if="isPolishing" style="margin-left: 20px;"></div>
            <button v-if="isPolishing" class="blog_edit_button" style="margin-left: 10px;flex-grow: 1;" @click="confirmPolishing">Confirm</button>
            <button v-if="isPolishing" class="blog_edit_button" style="margin-left: 10px;flex-grow: 1;" @click="cancelPolishing">Cancel</button>
        </div>
        
    </form>
</template>

<style>
.blog_edit_prompt {
    font-size: large;
    flex-grow: 2;
    box-sizing: border-box;
}
.blog_polish_vbox {
    display: flex;
    flex-direction: column;
    height: 100%;
    width: 0px;
    flex-grow: 1;
    margin-top: 10px;
    box-sizing: border-box;
}
.blog_edit_hbox {
    display: flex;
    flex-direction: row;
    width: 100%;
    align-items: center;
    margin: 5px;
}

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
    width: 80%;
}

.blog_content_editor {
    height: 100%;
    /* width: 100%; */
    flex-grow: 1;
    margin-top: 10px;
    margin-bottom: 10px;
    box-sizing: border-box;
}

.blog_edit_button {
    background-color: white;
    color: black;
    border-color: black;
    border-width: 1px;
    border-radius: 0px;
    border-style: solid;
    text-align: center;
    /* width: 100%; */
    width: 0px;
    height: 35px;
}

.blog_edit_button:hover{
    background-color: black;
    color: white;
}

</style>