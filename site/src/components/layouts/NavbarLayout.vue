<script setup lang="ts">
import { getAuthInfo, isLoggedIn } from '@/stores/auth';
import { onMounted } from 'vue';

onMounted(()=> {
    console.log(getAuthInfo.value)
})
</script>

<template>
    <div class="navbar">
        <div>
            <RouterLink :to="{name: 'home'}">Home</RouterLink>
            <RouterLink :to="{name: 'about'}">About</RouterLink>
        </div> -
        <div v-if="!isLoggedIn">
            <RouterLink :to="{name: 'login'}">Login</RouterLink>
            <RouterLink :to="{name: 'register'}">Register</RouterLink>
        </div>
        <div v-else>
            <div v-if="getAuthInfo">
                {{ getAuthInfo.name }}
            </div>
        </div>
    </div>
</template>

<style scoped>

</style>