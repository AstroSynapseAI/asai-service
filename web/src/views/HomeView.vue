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
        <img src="@/assets/asai-icon.png" height="24" class="d-inline-block" alt="">
        <span class="ms-3">TALK TO ASAI</span>
      </router-link>
    </nav>
    
    <div class="header row">
      <div class="col-12">
        <h3>AI Avatars Platform</h3>
      </div>
    </div>
    
    <div class="row">
      <div class="col-sm-6 col-md-9 col-xl-6">
        <img class="logo" src="@/assets/Logo.v1.white.png" alt="">
      </div>
      <div class="col-sm-6 col-md-3 col-xl-6">
        <router-link :to="{name: 'about', params: {slug: 'open-source'}}">
          <h5 class="mb-3">Open source <i data-feather="arrow-right"></i></h5>
        </router-link>

        <router-link :to="{name: 'about', params: {slug: 'transparency'}}">
          <h5 class="mb-3">Transparent<i data-feather="arrow-right"></i></h5>
        </router-link>

        <router-link :to="{name: 'about', params: {slug: 'customization'}}">
          <h5 class="mb-3">Customizable<i data-feather="arrow-right"></i></h5>
        </router-link>
        
        <router-link :to="{name: 'about', params: {slug: 'personalization'}}">
          <h5 class="mb-3">Personalized<i data-feather="arrow-right"></i></h5>
        </router-link>
        <router-link :to="{name: 'about', params: {slug: 'privacy'}}">
          <h5 class="mb-3">Private<i data-feather="arrow-right"></i></h5>
        </router-link>
      </div>
    </div>
    
    <hr class="opacity-100">
    
    <router-view> </router-view>
    
    <hr class="opacity-100">
    
    <div class="row">
      <div class="col-12">
        <h5 class="mb-3 mt-3 text-center">WOULD YOU LIKE TO KNOW MORE?</h5>
        <router-link :to="{name: 'about'}">
          <h5 class="mb-3 mt-3 text-center">MORE INFO </h5>
        </router-link>
        <router-link :to="{name: 'about', params: {slug: 'integrations'}}">
          <h5 class="mb-3 mt-3 text-center">CUSTOM INTEGRATIONS </h5>
        </router-link>
      </div>

    </div>

    <div class="row mt-auto">
      <div class="col-12">
        <div class="contacts row">
          
            <div class="col-4 offset-4">

              <div class="row">
                <div class="col-4 d-flex justify-content-center align-items-center">
                  <button class="btn d-flex justify-content-center align-items-center"><i class="d-block" data-feather="github"></i></button>
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

        <div class="row">
          <div class="col-12">
            <p class="text-center"><a href="https://astrosynapse.ai" target="_blank">astrosynapse.ai © 2023</a></p>
          </div>
        </div>
      </div>
    </div>
    
  </div>
  
</template>

<style scoped>
a {
  color: white;
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
  margin-bottom: 100px;
}
.cta .btn {
  background-color: black;
  border-color: white;
}

.contacts {
  margin-bottom: 15px;
  margin-top: 30px; 
}

.contacts .btn {
  background-color: black;
  color: white;
  border: 1px solid white;
  width: 60px;
}

.btn:hover {
  color: black;
  background-color: white;
}

@media only screen and (max-width: 576px) {
  hr {
    margin-top: 10px;
    margin-bottom: 10px;
  }

  .header {
    margin-top: 20px;
  }

  .logo {
    margin-bottom: 25px;
  }

  .cta {
    margin-top: 25px;
    margin-bottom: 25px;
  }
  .cta .btn {
    margin-bottom: 15px;
  }
  
}

</style>