<script setup>
import {handleError, onMounted, ref} from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/user';

const isLogin = ref(true)
import axios from 'axios';

const loginText = ref("登陆")
const signupText = ref("注册>>")

const authStore = useAuthStore()
const router = useRouter()

onMounted(() => {
    document.getElementById("userInfo").addEventListener("submit", function(event) {
        event.preventDefault();
    });
    clearTip()
})

function login(){
    clearTip()
    
    if(isLogin.value == false){
        isLogin.value = true;
        loginText.value = "登陆"
        signupText.value = "注册>>"
        document.getElementById("legend").innerText = "Login"
    }else{
        loginRequest();
    }
}

function signup(){
    clearTip()

    if(isLogin.value == true){
        isLogin.value = false;
        loginText.value = "<<登陆"
        signupText.value = "注册"
        document.getElementById("legend").innerText = "Signup"
    }else{
        signupRequest();
    }
}

function loginRequest(){
    const form = document.getElementById("userInfo")
    const data = new FormData(form)
    
    if(!form.checkValidity()){
        return
    }
    axios.post('/api/login/', Object.fromEntries(data.entries()))
    .then(function (response) {
        const token = response?.data?.token;
        if(!token){
            throw error("未获得token!")
        }
        authStore.setUserInfo(token, data.get("username"))
        router.push("/userpage")
        console.log(response);
        isLogin.value = true;
    })
    .catch(function (error) {
        const errno = error.response?.data?.code;
        handleLoginErrno(errno)
        console.log(error);
    });
}

function handleLoginErrno(errno){
    if(!errno){
        alert("未知错误,");
    }else if(errno == "ResourceNotFound.UserNotFound"){
        setTip('username', '用户不存在!')
    }else if(errno == "AuthFailure.PasswordIncorrect"){
        setTip('password', '密码错误！')
    }else{
        alert("服务器错误!");
    }
}

function signupRequest(){
    const form = document.getElementById("userInfo")
    const data = new FormData(form)
    if(!form.checkValidity()){
        return
    }
    if(data.get("password") != data.get("confirmPassword")){
        setTip('confirmPassword', '密码不匹配！')
        return
    }

    data.delete("confirmPassword")
    axios.post('/api/v1/users/', Object.fromEntries(data.entries()))
    .then(function(response){
        console.log(response);
    })
    .catch(function (error) {
        const code = error.response?.data?.code;
        if(code == "FailedOperation.UserAlreadyExist"){
            setTip("username", "用户名已存在")
        }
        console.log(error);
    });
}

function setTip(name, tip){
    clearTip()
    const form = document.getElementById("userInfo")
    const input = document.getElementsByName(name)
    const p = findNextElement(input[0], 'p')
    if(p){
        p.innerText = tip
    }
}

function clearTip(){
    const xpath = "//form[@id='userInfo']//p";  // 查找所有 class="item" 的 div
    const result = document.evaluate(
        xpath,
        document,
        null,
        XPathResult.ORDERED_NODE_SNAPSHOT_TYPE,
        null
    );
    console.log("clear %d tips", result.snapshotLength)
    for (let i = 0; i < result.snapshotLength; i++) {
        const item = result.snapshotItem(i);
        item.innerText = ""
    }
}

function findNextElement(startElement, targetTagName) {
    let next = startElement.nextElementSibling;
    while (next) {
        if (next.tagName.toLowerCase() == targetTagName.toLowerCase()) {
            return next; // 找到匹配的元素
        }
        next = next.nextElementSibling; // 继续查找下一个兄弟节点
    }
    return null; // 没找到
}

</script>

<template>
<div style="min-width: 100%;width: max-content;display:flex;flex-direction: column;align-items: center;">
    <div class="login_box">
        <fieldset style="border-style: dashed;border-width: 5px; border-color: rgb(109, 109, 109);padding: 30px;">
            <legend style="font-size: 30px;color: black;" id="legend">Login</legend>
            <form id="userInfo">
                <div class="userinput">
                    <template v-if="!isLogin">
                        <input type="text" name="nickname" placeholder="昵称" required maxlength="20"></input>
                        <p></p>
                        <input type="phone" name="phone" placeholder="电话" required pattern="([0-9]){11}">
                        <p></p>
                        <input type="email" name="email" placeholder="邮箱" required minlength="6" maxlength="20">
                        <p></p>
                    </template>
                        <input type="text" name="username" placeholder="用户名" pattern="([0-9]|[a-z])+" required maxlength="20">
                        <p></p>
                        <input type="password" name="password" placeholder="密码" required minlength="6" maxlength="20">
                        <p></p>
                    <template v-if="!isLogin">
                        <input name="confirmPassword" type="password" placeholder="确认密码" required minlength="6" maxlength="20">
                        <p></p>
                    </template>
                </div>
                <div class="startbuttons">
                    <button name="action" value="login" type="submit" @click="login">{{ loginText }}</button>
                    <div style="flex-grow: 3;"></div>
                    <button name="action" value="signup" type="submit" @click="signup">{{ signupText }}</button>
                </div>
            </form>
        </fieldset>
    </div>
    <div style="flex-grow: 5;height: 20%;">
        <p></p>
    </div>
</div>
</template>

<style>
.login_box {
    width: max-content;
    height: 100%;
    display: flex;
    flex-direction: row;
    align-items: center;
    flex-grow: 2;
}

.userinput p {
    margin: 0px;
    margin-left: 5px;
    color: red;
    text-align: start;
    height: 25px;
}

.userinput input{
	height: 40px;
	background-color: black;
	color: rgb(173, 167, 167);
	font-size: large;
	border-radius: 5px;
	transition: all 0.4s;
}

.userinput input:focus{
	background-color: rgb(66, 66, 66);
	color: rgb(226, 222, 222);
}

/*重置button样式*/
button {
    /* 移除默认边框和背景 */
    border: none;
    background: none;
	background-color: none;
    
    /* 移除浏览器默认的内边距和外观 */
    padding: 0;
    margin: 0;
    outline: none; /* 移除点击时的焦点边框（可选） */
    
    /* 确保字体继承父元素（避免浏览器默认字体） */
    font-family: inherit;
    font-size: inherit;
    color: inherit;
    
    /* 防止按钮被浏览器默认样式干扰 */
    -webkit-appearance: none; /* Safari/Chrome */
    -moz-appearance: none;    /* Firefox */
    appearance: none;
}

.startbuttons{
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    
}
.startbuttons button{
	background-color: #ffffff;
	background: #ffffff;
	color: #000000;
	font-style: italic;
	font-size: 18px;
	width: 70px;
    height: 40px;
    vertical-align: middle;
	/* padding: 3px; */
	cursor:pointer;
	transition: all 0.4s;

	border-radius: 5px;
    border-width: 2px;
    border-style: solid;
    border-color: #000000;

    flex-grow: 1;
}
.startbuttons button:hover {
    background-color: #000000;
	background: #000000;
	color: #d3d3d3;
}
</style>