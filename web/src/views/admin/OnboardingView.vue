<script setup>
import Navbar  from '@/components/admin/Navbar.vue';

import '@/assets/adminkit/static/css/dark.css';
import '@/assets/adminkit/static/js/app.js';

</script>

<template>

  <div class="wrapper">
    <Navbar />
    <div class="main">
      <main class="content">

        <div class="container-fluid p-0">
          <div class="row">
            <div class="col-12">
              <router-view />
            </div>
          </div>
        </div>

      </main>
    </div>
  </div>

</template>
