<script setup>
import axios from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRouter } from 'vue-router';
import PostItemView from './PostItemView.vue';
import PageJumper from './PageJumper.vue';

const pageJumperRef = ref(null)
const postItemViewRef = ref(null)

const deletedPostTitle = ref("")
let btn = document.getElementById("delete_post_message")

const router = useRouter()

onMounted(()=>{
    postItemViewRef.value.editable = true
})

//这里要加?，可能是因为在页面刚加载还没有把pageJumperRef和子组件联系在一起时就开始监听了。
//这也导致，即使不onMounted，也能监听到pageJumperRef一开始的变化行为，从而获取posts
watch(
    ()=>pageJumperRef.value?.offset,
    (newValue)=>{
        listPostRequest()
    }
)

watch(
    ()=>postItemViewRef.value?.tobeFreshed,
    (newValue)=>{
        if(postItemViewRef.value.tobeFreshed){
            listPostRequest()
            postItemViewRef.value.tobeFreshed = false
        }
    }
)

function listPostRequest(){
    axios.get("/api/v1/posts?" + "limit=" + pageJumperRef.value.limit + "&offset=" + pageJumperRef.value.offset)
    .then((response)=>{
        postItemViewRef.value.posts = response.data.posts
        pageJumperRef.value.totalPage = parseInt((response.data.totalCount-1)/pageJumperRef.value.limit + 1)
    })
}



</script>

<template>
    <post-item-view ref="postItemViewRef"></post-item-view>
    <page-jumper ref="pageJumperRef"></page-jumper>

    <!-- <div id="delete_post_message">
        <p>确定要删除{{deletedPostTitle}}吗？</p>
        <button class="delete_post_message_button">yes</button>
        <button class="delete_post_message_button">no</button>
    </div> -->
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
</style>