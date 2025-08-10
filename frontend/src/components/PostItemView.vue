<script setup>
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';


const posts = ref(null)
const editable = ref(false)
const tobeFreshed = ref(false)

const router = useRouter()

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
        tobeFreshed.value = true
    })).catch((error)=>{
        alert("删除失败")
        console.log(error)
    })
}

defineExpose({
    posts,
    editable,
    tobeFreshed
})

</script>

<template>
    <div class="post_item_view" v-for="post in posts">
        <div class="post_item_head">
            <div style="display: flex;flex-direction: row;flex-grow: 1;flex-shrink: 1;width: 80%;">
                <div class="overflow_hidden_element" style="font-size:x-large;font-weight: bold;align-self: center;justify-self: center;flex-shrink: 1;">
                    {{post.title}}
                </div>
                <div class="overflow_hidden_element" style="align-self:self-end;flex-shrink: 1;flex-grow: 1;">
                    （updated at {{ post.updatedAt }}）
                </div>
            </div>
            <div style="margin-left: auto;display: flex;flex-direction: row;">
                <button v-if="editable" class="edit_post_button" @click="()=>onDeletePost(post)">delete</button>
                <button v-if="editable" class="edit_post_button" @click="()=>updatePost(post.postId)">edit</button>
            </div>
        </div>
        <p class="overflow_hidden_element">{{ post.content }}</p>
    </div>
</template>


<style>
.post_item_head {
    display: flex;
    flex-direction: row;
    padding: 2px;
    flex-grow: 1;
}

.post_item_view {
    width: 1000px;
    max-width: 83%;
    height: 80px;
    border: 2px solid white;
    box-sizing: border-box;
    display: flex;
    flex-direction: column;
}

.post_item_view:hover {
    border: 0.5px solid black;
    box-sizing: border-box;
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
    flex-shrink: 0;
    
}

.edit_post_button:hover {
    background-color: black;
    color: white;
}


</style>