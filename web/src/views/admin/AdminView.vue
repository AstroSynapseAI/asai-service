<script setup>
import Sidebar from '@/components/admin/Sidebar.vue';
import Navbar  from '@/components/admin/Navbar.vue';
import '@/assets/adminkit/dist/css/dark.css';
import '@/assets/adminkit/dist/js/app.js';
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