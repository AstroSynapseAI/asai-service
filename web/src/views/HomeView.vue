<script setup>
import { onMounted } from 'vue';


onMounted(() => {
  feather.replace();
});
</script>

<template>
  <div class="container d-flex flex-column vh-100">
    
    <!-- Navigation  -->

    <nav class="navbar navbar-expand-md bg-dark bg-transparent">
      <div class="container-fluid">
        <div class="row w-100">
          <div class="col-xl-6 col-md-4">
            <span class="navbar-brand text-white">AI Avatars Platform</span>
          </div>
          <div class="navbar-menu col-xl-6 col-md-8 d-flex justify-content-md-end justify-content-sm-start">

            <div class="row">
              <div class="col-auto">
                <router-link :to="{name: 'asai'}" class="btn text-white px-md-4 px-lg-5 px-sm-2">
                  <img src="@/assets/asai-icon.png" height="24" alt="">
                  <span class="ms-3">Talk to Asai</span>
                </router-link>
              </div>
              <div class="col-auto">
                <router-link :to="{name: 'about', hash: '#content'}" class="btn text-white px-md-5 px-sm-0">
                  <span>About</span>
                </router-link>
              </div>

              <div class="col-auto">
                <router-link :to="{name: 'login'}" class="btn text-white">
                  <i class="align-middle" data-feather="log-in"></i>
                </router-link>
              </div>
            </div>

          </div>
        </div>
      </div>
    </nav>

    <!-- Navigation End -->

    <hr class="top-hr opacity-100">

    <!-- Content -->
    
    <div class="sub-menu row">
      <div class="col-lg-6">
        <img class="logo" src="@/assets/Logo.v1.white.png" alt="">
      </div>

      <div class="col-lg-6">
        <div class="row">

          <div class="col-md-6">
            <table>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'open-source'}}">
                    <h5 class="mb-3">Open source</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3 ms-3" data-feather="arrow-right"></i>
                </td>
              </tr>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'transparency'}}">
                    <h5 class="mb-3">Transparent</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'personalization'}}">
                    <h5 class="mb-3">Personalized</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'privacy'}}">
                    <h5>Private</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>
            </table>
          </div>

          <div class="col-md-6">
            <table>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'open-source'}}">
                    <h5 class="mb-3">Cloud</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3 ms-3" data-feather="arrow-right"></i>
                </td>
              </tr>
              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'transparency'}}">
                    <h5 class="mb-3">Self-Hosted</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'integrations'}}">
                    <h5 class="mb-3">Integrations</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>

              <tr>
                <td>
                  <router-link :to="{name: 'about', hash: '#content', params: {slug: 'integrations'}}">
                    <h5 class="mb-3">Local</h5>
                  </router-link>
                </td>
                <td class="d-flex justify-content-end">
                  <i class="mb-3" data-feather="arrow-right"></i>
                </td>
              </tr>
            </table>
          </div>

        </div>
      </div>

    </div>
    
    <hr class="bottom-hr opacity-100">
    
    <div id="content"></div>
    <router-view> </router-view>

    <!-- Content End -->

    <div class="row mt-auto">
      <div class="col-12 mt-lg-5">
        <div class="contacts row">
          
            <div class="col-md-4 offset-md-4 col-sm-6 offset-sm-2">

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

nav {
  margin-top: 50px;
  margin-bottom: 6em;
}

nav .btn {
  background-color: black;
  border-color: white;
  border-radius: 0;
  text-decoration: none;
}

nav .btn:hover {
  color: black !important;
  background-color: white;
}

.sub-menu h5 {
  font-size: 20px;
  font-weight: 400;
}

table i {
  margin-left: 5em;
}

.header {
  margin-top: 150px;
  margin-bottom: 30px;
}

.top-hr {
  margin-bottom: 6em;
}

.bottom-hr {
  margin-top: 6em;
  margin-bottom: 6em;
}

.contacts {
  margin-bottom: 15px;
}

.contacts .btn {
  background-color: black;
  color: white;
  border: 1px solid white;
  width: 60px;
  border-radius: 0;
}

.contacts .btn:hover {
  color: black;
  background-color: white;
}

@media screen and (max-width: 1024px) {
  .logo {
    width: 100%;
    margin-bottom: 3em;
  }
}

@media only screen and (max-width: 576px) {

  nav {
    margin-top: 10px;
    margin-bottom: 2em;
  }
  .header {
    margin-top: 20px;
  }

  .navbar-menu {
    margin-top: 2em;
  }

  .navbar-menu .row > * {
    padding-right: calc(var(--bs-gutter-x) * .40) !important;
    padding-left: calc(var(--bs-gutter-x) * .40) !important;
  }

  .sub-menu table {
    width: 100%;
  }

  .top-hr {
    margin-bottom: 3em;
  }

  .bottom-hr {
    margin-top: 3em;
    margin-bottom: 3em;
  }


}

</style>