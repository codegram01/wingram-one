<script setup lang="ts">
import NavbarLayout from '@/components/layouts/NavbarLayout.vue';

</script>

<template>
  <div class="main">
    <div class="navbar-ctn">
      <NavbarLayout/>
    </div>
    <div class="page-ctn">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
.main {
  width: 100%;
  height: 100%;
  position: relative;
}

.navbar-ctn {
  width: 100%;
  height: 72px;
  position: absolute;
}

.page-ctn {
  width: 100%;
  height: calc(100% - 72px);
  top: 72px;
  position: absolute;
}
</style>
