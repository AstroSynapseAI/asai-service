<script setup>
import { onMounted, ref } from 'vue';
import { useAuthStore } from '@/stores/auth.store.js';

const authStore = useAuthStore();
const username = ref('');

const logout = () => {
  authStore.logout();
}

onMounted(() => {
  username.value = authStore.user?.username;
  feather.replace();
})
</script>

<template>
  <nav class="navbar navbar-expand navbar-light navbar-bg">
    <h2 v-if="!authStore.isLoggedIn">AI Avatar</h2>
    <!-- <a class="sidebar-toggle js-sidebar-toggle"> -->
    <!--   <i class="hamburger align-self-center"></i> -->
    <!-- </a> -->

    <div class="navbar-collapse collapse">
      <ul class="navbar-nav navbar-align">

        <li class="nav-item dropdown">
          <!-- <a class="nav-icon dropdown-toggle d-inline-block d-sm-none" href="#" data-bs-toggle="dropdown"> -->
          <!--   <i class="align-middle" data-feather="settings"></i> -->
          <!-- </a> -->

          <a class="nav-link dropdown-toggle d-none d-sm-inline-block" href="#" data-bs-toggle="dropdown">
            <span class="text-dark me-2">{{ username }}</span>
          </a>

          <div class="dropdown-menu dropdown-menu-end">
            <router-link :to="{ name: 'profile' }" class="dropdown-item"><i class="align-middle me-1"
                data-feather="user"></i> Profile</router-link>

            <div class="dropdown-divider"></div>
            <a class="dropdown-item" href="#" @click="logout"><i class="align-middle me-1"
                data-feather="log-out"></i>Log out</a>
          </div>
        </li>

      </ul>

    </div>
  </nav>
</template>

<style scoped>
.navbar-bg {
  background-color: #222e3c !important;
}

.navbar {
  height: 65px;
}

.navbar-brand {
  color: #ffffff !important;
}
</style>
