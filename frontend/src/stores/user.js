import { defineStore } from "pinia";
import { ref } from 'vue'
import axios from 'axios'

export const useAuthStore = defineStore('auth', () => {
  const username = ref(localStorage.getItem('username') || null)
  const token = ref(localStorage.getItem('token') || null)
  const isAuthenticated = ref(!!token.value)
  
  const setUserInfo = (new_token, new_username) => {
    token.value = new_token
    username.value = new_username
    isAuthenticated.value = true
    localStorage.setItem('token', new_token)
    localStorage.setItem('username', new_username)
    axios.defaults.headers.common['Authorization'] = `Bearer ${token.value}`
  }
  
  const clearUserInfo = () => {
    token.value = null
    username.value = null
    isAuthenticated.value = false
    localStorage.removeItem('token')
    delete axios.defaults.headers.common['Authorization']
  }
  
  return { token, isAuthenticated, username, setUserInfo, clearUserInfo }
})
