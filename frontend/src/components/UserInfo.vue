<script setup>
import {ref, onMounted} from 'vue'
import axios from 'axios';
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/user';

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

const router = useRouter()
const authStore = useAuthStore()


onMounted(() => {
    console.log("isAuthenticated")
    getRequest()
    if(!authStore.isAuthenticated){
        router.push("/")
    }
})

function getRequest(){
    axios.get("/api/v1/users/" + authStore.username)
    .then(function(response){
        if(response.status == axios.HttpStatusCode.Ok){
            username.value = response.data.username
            nickname.value = response.data.nickname
            phone.value = response.data.phone
            email.value = response.data.email
            createdTime.value = response.data.createdAt
            updatedTime.value = response.data.updatedAt
        }else{
            authStore.clearUserInfo()
            alert("网络错误，请重新登陆")
            useRouter().push("/login")
        }
    })
    .catch(function(err){
        console.log(err)
        if(err.response.data.code == "AuthFailure.TokenInvalid"){
            authStore.clearUserInfo()
            alert("网络错误，请重新登陆")
            useRouter().push("/login")
        }
    })
}
function updateRequest(){
    const form = document.getElementById("userProfileForm")
    if(!form.checkValidity()){
        return
    }
    if(editNickname.value == nickname.value
        && editEmail.value == email.value
        && editPhone.value == phone.value
    ){
        isEditing.value = false;
        return
    }
    const data = new FormData(form)
    axios.put("/api/v1/users/" + authStore.username, Object.fromEntries(data.entries()))
    .then(function(response){
        if(isEditing.value == true){
            isEditing.value = false;
        }
        getRequest()
    })
    .catch(function(err){
        console.log(err)
        alert(err.response.data.message)
        location.reload()
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
    <div style="height: 10%;"></div>
    <form id = "userProfileForm">
        <div>
            <label class="label_key">username:</label>
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
            <button v-if="!isEditing" style="width: 50;flex-grow: 1;" @click="editProfile">[edit profile]</button>
            <button v-if="isEditing" style="width: 50;flex-grow: 1;" @click="saveProfile" type="submit">[save profile]</button>
            <button style="width: 50;flex-grow: 1;">[change password]</button>
        </div>
    </form>
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

#userProfileForm div {
    height: 50px;
    width: 400px;
    display: flex;
    flex-direction: row;
}

#userProfileForm input {
    height: 25px;
    min-width: none;
    width: 50px;
    margin-bottom: 20px;
    font-size: large;
    flex-grow: 1;
}

</style>