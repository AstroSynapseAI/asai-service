<script setup>
import {useRoute} from 'vue-router';
import { onMounted, ref } from 'vue';
import Sidebar from '@/components/admin/Sidebar.vue';
import Navbar  from '@/components/admin/Navbar.vue';

import '@/assets/adminkit/static/css/dark.css';
import '@/assets/adminkit/static/js/app.js';

const route = useRoute();
const routeMainId = ref('default_view');

onMounted(async () => {
  if (route.path.search('admin/avatar/')) {
    routeMainId.value = 'chat_view';
  }
});
</script>
<template>
  
  <div class="wrapper">
    <Sidebar />
    
    <div class="main">
      <Navbar />
      <main class="content" v-bind:id="routeMainId">
        <router-view />
      </main>
    </div>
  
  </div>
</template> 
<style scoped>
  main#chat_view {
  padding-bottom: 0 !important;
}
</style>
