<script setup>
import { ref } from 'vue'

const searchContent = ref("")
const hasLogin = ref(false)
</script>

<template>
<div style="min-width: 100%;width: max-content;min-height: 100%;height: max-content;display:flex;flex-direction: column;align-items: center;">
    <div class="top_bar">
        <div class="top_bar_search">
            <input v-model="searchContent" class="top_bar_search_input"></input>
            <button class="top_bar_search_button">search</button>
        </div>

        <!-- <router-link to="/login"><button class="login">login</button></router-link>
        <router-link to="/user"><button class="login">userInfo</button></router-link> -->
        <a href="/login.html"><button class="login">login</button></a>
        <a href="/userpage.html"><button class="login">userInfo</button></a>
    </div>
    </br>
    <div class="content">
    </div>
</div>
</template>

<style>
.top_bar {
    /* width: 100%;
    vertical-align: top; */
    top: 0;
    height: fit-content;
    padding: 5px;
    align-content: center;
    width: max-content;
    min-width: 100%;
    
    background-color: rgb(255, 255, 255);

    font-size: large;
}

.content {
    width: 100%;
    vertical-align: bottom;
    background-color: rgb(199, 199, 199);
    display: flex;
    flex-grow: 1;
}

.top_bar_search {
    display: inline;
}

.top_bar_search_input {
    border-radius: 3px;
    background-color: inherit;
    font-size: large;
}

.top_bar_search_button {
    margin-left: 5px;
}

.top_bar button {
    background-color:inherit;
    color: black;
    border-color: black;
    border-width: 2px;
    border-style: solid;
    border-radius: 5px;
    padding: 2px;
}

.top_bar button:hover {
    background-color: black;
    color: white;
}

.login {
    float: right;
    margin-right: 10px;
    height: inherit;
    box-sizing: border-box;
}
</style>