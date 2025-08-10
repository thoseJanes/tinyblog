<script setup>
import axios from 'axios';
import { onMounted, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import PageJumper from './PageJumper.vue';
import PostItemView from './PostItemView.vue';

const pageJumperRef = ref(null)
const postItemViewRef = ref(null)

const router = useRouter()
const route = useRoute()

onMounted(()=>{
    const searchValue = route.query.text
    if(searchValue){
        searchPostRequest(searchValue)
    }
})

watch(
    ()=>route.query,
    (newValue)=>{
        const searchValue = route.query.text
        searchPostRequest(searchValue)
    }
)

watch(
    ()=>pageJumperRef.value?.offset,
    (newValue)=>{
        searchPostRequest(route.query.text)
    }
)

function searchPostRequest(text){
    const url = "/api/v1/posts/search?text=" + text + "&limit=" + pageJumperRef.value.limit + "&offset=" + pageJumperRef.value.offset
    axios.get(url)
    .then((response)=>{
        postItemViewRef.value.posts = response.data.posts
        pageJumperRef.value.totalPage = parseInt((response.data.totalCount-1)/pageJumperRef.value.limit + 1)
    }).catch(err=>{
        console.log(err)
        alert(err.response.data.message)
        router.push("/")
    })
}

</script>

<template>
    <div class="search_post_content">
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
</style>