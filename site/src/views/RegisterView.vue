<script setup lang="ts">
import { register_api, type Register_req } from '@/services/auth';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter()

const registerReq = ref<Register_req>({
    email: "",
    password: "",
    name: ""
})

const register = async () => {
    try {
        await register_api(registerReq.value).then(res => {
            alert("Register success")

            router.push({name: "login"})
        })
    } catch (error) {
        console.log(error)
    }
}
</script>

<template>
  <main>
    <h1>Register</h1>

    <form @submit.prevent="register">
        <label for="email">Email</label>
        <input type="text" v-model="registerReq.email">

        <label for="name">Name</label>
        <input type="text" v-model="registerReq.name">

        <label for="password">Password</label>
        <input type="password" v-model="registerReq.password">

        <button type="submit">Register</button>
    </form>
  </main>
</template>
