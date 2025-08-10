<script setup>
import axios from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';

import PageJumper from './PageJumper.vue';
import PostItemView from './PostItemView.vue';

const pageJumperRef = ref(null)
const postItemViewRef = ref(null)

const allPosts = ref(null)

const router = useRouter()
const route = useRoute()

const evaluation = ref("")
onMounted(()=>{
    const searchValue = route.query.prompt
    if(searchValue){
        aiSearchPostRequest(searchValue)
    }
})

watch(
    ()=>route.query,
    (newValue)=>{
        const searchValue = route.query.prompt
        aiSearchPostRequest(searchValue)
    }
)

watch(
    ()=>pageJumperRef.value?.offset,
    (newValue)=>{
        if(allPosts.value){
            const end = Math.min(pageJumperRef.value.offset+pageJumperRef.value.limit, allPosts.value.length)
            postItemViewRef.value.posts = allPosts.value.slice(pageJumperRef.value.offset, end)
        }
    }
)


function aiSearchPostRequest(prompt){
    const url = "/api/v1/posts/aisearch?prompt=" + prompt
    evaluation.value = "搜索中，请耐心等待..."
    axios.get(url)
    .then((response)=>{
        allPosts.value = response.data.posts
        evaluation.value = response.data.evaluation
        pageJumperRef.value.totalPage = parseInt((allPosts.value.length-1)/pageJumperRef.value.limit + 1)

        const end = Math.min(pageJumperRef.value.offset+pageJumperRef.value.limit, allPosts.value.length)
        postItemViewRef.value.posts = allPosts.value.slice(pageJumperRef.value.offset, end)
    }).catch(err=>{
        console.log(err)
        alert(err.response.data.message)
        router.push("/")
    })
}

</script>

<template>
    <div class="search_post_content">
        <div class="aisearch_evaluation">
            <p>{{ evaluation }}</p>
        </div>
        <post-item-view ref="postItemViewRef"></post-item-view>
        <page-jumper ref="pageJumperRef"></page-jumper>
    </div>
</template>


<style>
.search_post_content {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
}

.aisearch_evaluation {
    width: 80%;
}
</style>