<script setup>
import {ref} from 'vue'
import axios from 'axios';
import { useRoute, useRouter } from 'vue-router'

const showContent = ref("")
const token = localStorage.getItem("jwt_token")
const username=ref("")
const nickname=ref("")
const phone=ref("")
const email=ref("")
const createdTime=ref("")
const updatedTime=ref("")

const editNickname=ref("")
const editPhone=ref("")
const editEmail=ref("")
const isEditing=ref(false)

function getRequest(){
    axios.get("/", token)
    .then(function(response){
        if(response.status == axios.HttpStatusCode.Ok){
            username.value = response.data.username
            nickname.value = response.data.nickname
            phone.value = response.data.phone
            email.value = response.data.email
        }else{
            localStorage.removeItem("jwt_token")
            alert("网络错误，请重新登陆")
            useRouter().push("/login")
        }
    })
    .catch(function(err){
        console.log(err)
    })
}
function updateRequest(){
    const form = document.getElementById("userProfileForm")
    if(!form.checkValidity()){
        return
    }
    const data = new FormData()
    axios.put("/", data)
    .then(function(response){
        if(isEditing.value == true){
            isEditing.value = false;
        }
    })
    .catch(function(err){
        console.log(err)
    })
}
function editProfile(){
    document.getElementById("userProfileForm").addEventListener("submit", (event) => {
        event.preventDefault()
    })

    console.log("edit")
    if(isEditing.value == false){
        isEditing.value = true;
        editEmail.value = email.value
        editNickname.value = nickname.value
        editPhone.value = phone.value
    }
}
function saveProfile(){
    updateRequest();
}
</script>

<template>
    <div class = "user_content">
        <form id = "userProfileForm">
            <div>
                <label class="label_key">username</label>
                <label class="label_value">{{ username }}</label>
                <!-- <input v-if="isEditing" v-model="username" name = "username" type = "text"></input> -->
            </div>
            <div>
                <label class="label_key">nickname:</label>
                <label v-if="!isEditing" class="label_value">{{ nickname }}</label>
                <input v-if="isEditing" v-model="editNickname" name = "nickname" type = "text" required></input>
            </div>
            <div>
                <label class="label_key">phone:</label>
                <label v-if="!isEditing" class="label_value">{{ phone }}</label>
                <input v-if="isEditing" v-model="editPhone" name = "phone" type = "text" required pattern="([0-9]){11}"></input>
            </div>
            <div>
                <label class="label_key">email:</label>
                <label v-if="!isEditing" class="label_value">{{ email }}</label>
                <input v-if="isEditing" v-model="editEmail" name = "email" type = "email" required></input>
            </div>
            <div>
                <label class="label_key">created at:</label>
                <label class="label_value">{{ createdTime }}</label>
            </div>
            <div>
                <label class="label_key">updated at:</label>
                <label class="label_value">{{ updatedTime }}</label>
            </div>
            <div>
                <button v-if="!isEditing" style="width: 50;flex-grow: 1;" @click="editProfile">edit</button>
                <button v-if="isEditing" style="width: 50;flex-grow: 1;" @click="saveProfile" type="submit">save</button>
                <button style="width: 50;flex-grow: 1;">changePassword</button>
            </div>
        </form>
    </div>

</template>


<style>
.label_key {
    height: 50px;
    width: 150px;
    text-align: left;
}

.label_value {
    flex-grow: 1;
    padding-left: 5px;
    text-align: left;
}

.user_content input {
    height: 25px;
    min-width: none;
    width: 50px;
    margin-bottom: 20px;
    font-size: large;
    flex-grow: 1;
}
.user_content {
    display: flex;
    flex-direction: column;
    flex-grow: 5;
    align-items: center;
    background-color: rgb(178, 255, 188);
    padding: 20px;
    
}

.user_content div {
    height: 50px;
    width: 400px;
    display: flex;
    flex-direction: row;
}
</style>