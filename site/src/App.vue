<script setup lang="ts">
import NavbarLayout from '@/components/layouts/NavbarLayout.vue';

</script>

<template>
  <div class="app">
    <div class="navbar">
      <NavbarLayout/>
    </div>
    <div class="main">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>

</style>
