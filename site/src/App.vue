<script setup lang="ts">
import NavbarLayout from '@/components/layouts/NavbarLayout.vue';

</script>

<template>
  <div class="app">
    <div class="navbar-ctn">
      <NavbarLayout/>
    </div>
    <div class="page-ctn">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
.app {

}

.navbar-ctn {

}

.page-ctn {
  
}
</style>
