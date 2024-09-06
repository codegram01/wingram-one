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
        </div>
        <div class="nav-action">
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
    </div>
</template>

<style scoped>
.navbar {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
}

.nav-action {
    margin-left: auto;
}
</style>