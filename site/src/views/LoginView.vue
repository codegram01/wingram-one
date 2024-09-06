<script setup lang="ts">
import { setAuthLogin } from '@/auth';
import { login_api, type Login_req } from '@/services/auth';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

const loginReq = ref<Login_req>({
    email: "",
    password: ""
})

const login = async () => {
    try {
        await login_api(loginReq.value).then(res => {
            setAuthLogin(res.data)

            alert("Login success")
            router.push({name: "home"})
        })
    } catch (error) {
        console.log(error)
    }
}
</script>

<template>
  <main class="page">
    <h1>Login</h1>

    <form @submit.prevent="login">
        <label for="email">Email</label>
        <input type="text" v-model="loginReq.email">

        <label for="password">Password</label>
        <input type="password" v-model="loginReq.password">

        <button type="submit">Login</button>
    </form>
  </main>
</template>
