<script setup>
import { onMounted} from 'vue';

onMounted(() => {
  feather.replace();
})

</script>

<template>
  <div class="container d-flex flex-column vh-100">
    <nav class="navbar navbar-dark bg-transparent">
    <router-link :to="{name: 'asai'}" class="navbar-brand" href="#">
      <img src="../assets/asai-icon.png" height="24" class="d-inline-block" alt="">
      <span class="ms-3">TALK TO ASAI</span>
    </router-link>
    </nav>
    <div class="header row">
      <div class="col-12">
        <h3>AI Avatars Platform</h3>
      </div>
    </div>
    
    <div class="row">
      <div class="col-6">
        <img src="../assets/Logo.v1.white.png" alt="">
      </div>
      <div class="col-6">
        <h5 class="mb-3">Open source</h5>
        <h5 class="mb-3">Transparent</h5>
        <h5 class="mb-3">Customizable</h5>
        <h5 class="mb-3">Personalized</h5>
        <h5 class="mb-3">Private</h5>
      </div>
    </div>
    
    <hr class="mt-5 opacity-100">
    
    <div class="cta row">
      <div class="col-4">
        <button class="btn btn-primary w-100 btn-lg">CLOUD</button>
      </div>
      <div class="col-4">
        <button class="btn btn-primary w-100 btn-lg">SELF-HOSTED</button>
      </div>
      <div class="col-4">
        <button class="btn btn-primary w-100 btn-lg">LOCAL</button>
      </div>
    </div>
    
    <hr class="mt-5 opacity-100">
    
    <div class="contacts row">
      <div class="col-4 offset-4">
        
        <div class="row">
          <div class="col-4 d-flex justify-content-center align-items-center">
            <button class="btn d-flex justify-content-center align-items-center"><i class="d-block" data-feather="github"></i> </button>
          </div>
          <div class="col-4 d-flex justify-content-center align-items-center">
            <button class="btn"><img src="../assets/discord-mark-white.png" alt="discord" height="24"></button>
          </div>
          <div class="col-4 d-flex justify-content-center align-items-center">
            <button class="btn d-flex justify-content-center align-items-center"><i class="d-block" data-feather="twitter"></i> </button>
          </div>
        </div>

      </div>
    </div>

    <div class="footer row mt-auto">
      <div class="col-12">
        <p class="text-center"><a href="https://astrosynapse.ai" target="_blank">astrosynapse.ai © 2023</a></p>
      </div>
    </div>
  </div>
  
</template>

<style scoped>
a {
  color: white;
  text-decoration: none;
}

a:hover {
  color: white;
  text-decoration: underline;
}

.header {
  margin-top: 150px;
  margin-bottom: 30px;
}

.cta {
  margin-top: 100px;  
  margin-bottom: 120px;
}
.cta .btn {
  background-color: black;
  border-color: white;
}

.contacts {
  margin-top: 30px;
}

.contacts .btn {
  background-color: black;
  color: white;
  border: 1px solid white;
  width: 60px;
}

</style>