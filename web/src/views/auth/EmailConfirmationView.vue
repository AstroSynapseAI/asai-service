<script setup>
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth.store';
import { onMounted } from 'vue';
import { useToast } from 'vue-toastification';

const auth = useAuthStore();
const toast = useToast();

onMounted(async () => {
    try {
        const urlParams = new URLSearchParams(window.location.search);
        const email = urlParams.get('email');
        const token = urlParams.get('token');

        console.log('Email:', email);
        console.log('Token:', token);

        const formData = await auth.validateConfirmationEmail({
            email: email,
            token: token
        });
        
      }
      catch (error) {
        console.log("err-", error)
        toast.error(error)
    }
});

</script>
<template>
  <div class="container d-flex flex-column vh-60">
    <nav class="navbar navbar-expand-md bg-dark bg-transparent">
      <div class="container-fluid">
        <div class="row w-100">
          <div class="col-6">
            <span class="navbar-brand text-white">AI Avatars Platform</span>
          </div>

          <div class="navbar-menu col-6 d-flex justify-content-end">

            <div class="row">
              <div class="col-auto">
                <router-link :to="{ name: 'home' }" class="btn text-white">
                  <i class="align-middle feather-icon" data-feather="home"></i>
                </router-link>
              </div>
            </div>

          </div>
        </div>
      </div>
    </nav>

      <div class="col-md-4">
        <img class="logo" src="@/assets/ASAILogotype.svg" alt="">
        <div class="card">
            <button class="send-button btn btn-light" @click="''">Go to profile</button>
        </div>
      </div>
  </div>
</template>

<style scoped>
a {
  color: white !important;
}

nav {
  margin-top: 50px;
  margin-bottom: 15em;
}

.navbar-brand {
  color: white !important;
}

h1,
h2,
h3,
h4,
h5,
h6 {
  color: white;
}


.feather-icon {
  width: 24px !important;
  height: 24px !important;
}

.logo {
  width: 100%;
  margin-bottom: 3em;
}

.card {
  background-color: black;
  border: 1px solid white;
  border-radius: 0;
  width: 100%;
}

.card-body {
  color: white;
  /* To make text visible in dark background */
}

.form-control {
  background-color: black;
  border: none;
  width: 100%;
  color: white;
  border-radius: 0;
}

.email-input,
.pass-input {
  width: 100%;
  margin-bottom: 2em;
  height: 3em;
  background-color: transparent;
  border: 1px solid white;
  color: white;
  border-radius: 0;
}

.btn {
  background-color: black;
  border-color: white;
  border-radius: 0;
  text-decoration: none;
  color: white;
}

.btn:hover {
  color: black !important;
  background-color: white;
}

.loader {
  width: 18px;
  height: 18px;
  border: 2px solid #FFF;
  border-bottom-color: transparent;
  border-radius: 50%;
  display: inline-block;
  box-sizing: border-box;
  animation: rotation 1s linear infinite;
}

@keyframes rotation {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

.send-button {
  min-width: 150px;
}

@media only screen and (max-width: 576px) {
  nav {
    margin-top: 10px;
    margin-bottom: 2em;
  }
}
</style>
