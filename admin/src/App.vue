<script setup>
import Sidebar from './components/Sidebar.vue';
import Navbar  from './components/Navbar.vue';
</script>

<template>
  
  <div class="wrapper">
    <Sidebar />
    
    <div class="main">
      <Navbar />
      <main class="content">
        <router-view />
      </main>
    </div>
  
  </div>

</template>

<style scoped>
h1, h2, h3, h4, h5, h6 {
  color: white;
}
.content {
  background-color: #141A21;
  color: white;
}
</style>
