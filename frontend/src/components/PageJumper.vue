
<script setup>
import { ref, watch } from 'vue'

const presentPage = ref(1)

const totalPage = ref(1)
const limit = ref(10)
const offset = ref(0)

watch(presentPage, (newValue, oldValue, cleanUp) => {
    if(presentPage.value > totalPage.value){
        presentPage.value = oldValue
    }else if(presentPage.value < 1) {
        presentPage.value = 1
    }else{
        offset.value = limit.value*(presentPage.value-1)
        
    }
})

defineExpose({
    totalPage,
    offset,
    limit
})

function nextPage(){
    if(presentPage.value < totalPage.value){
        presentPage.value += 1
    }
}

function beforePage(){
    if(presentPage.value > 1){
        presentPage.value -= 1
    }
}

function firstPage(){
    if(presentPage.value != 1){
        presentPage.value = 1
    }
    
}

function lastPage(){
    if(presentPage.value != totalPage.value){
        presentPage.value = totalPage.value
    }
    
}

</script>

<template>
    <div class="list_post_form">
        <button class="set_page_button" @click="firstPage"><<</button>
        <button class="set_page_button" @click="beforePage"><</button>
        <p style="margin-left: 20px;">page: </p>
        <input name="page" style="width: 25px;flex-grow: 0;height: 18px;text-align: center;font-size: large;" v-model="presentPage" type="number" min="1" :max="totalPage"></input>
        <p style="margin-right: 20px;">/{{ totalPage }}</p>
        <button class="set_page_button" @click="nextPage">></button>
        <button class="set_page_button" @click="lastPage">>></button>
    </div>
</template>


<style>
.list_post_form {
    width: 90%;
    height: 50px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    font-size: large;
}

.set_page_button {
    border: 1px solid black;
    border-radius: 1px;
    width: 25px;
    margin: 1px;
}
</style>