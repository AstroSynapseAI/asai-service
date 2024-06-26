<script setup>
import { reactive, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { onMounted } from 'vue';
import { useAuthStore } from '@/stores/auth.store';
import { useUserStore } from '@/stores/user.store.js';
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';

import AsaiAlert from '@/views/ui/AsaiAlert.vue';

const schema = yup.object({
  Username: yup.string().required(),
  Password: yup.string().required().min(8),
  ConfirmPassword: yup.string().required('Please confirm your password').min(8)
    .oneOf([yup.ref('Password'), null], 'Passwords must match')
});

const user = useUserStore();
const router = useRouter();
const route = useRoute();
const auth = useAuthStore();
const username = ref('');
const password = ref('');
const confirmPassword = ref('');
const isAsaiAlertActive = ref(false);
const apiErrorText = ref('');

const formState = reactive({
  isSubmitting: false,
});

const closeAsaiAlert = () => {
  isAsaiAlertActive.value = false;
};

const register = async () => {
  formState.isSubmitting = true;
  try {
    const loggedIn = await auth.registerInvite({
      username: username.value,
      password: password.value,
      invite_token: route.params.invite_token
    });

    if (loggedIn) {
      user.current = auth.user;
      router.push({ name: 'welcome' });
    }
  }
  catch (error) {
    apiErrorText.value = error;
    isAsaiAlertActive.value = true;
  }
  finally {
    formState.isSubmitting  = false
  }
};

onMounted(async () => {
  feather.replace();
  if (route.params.invite_token) {
    try {
      let user = await auth.getInvitedUser(route.params.invite_token);
      username.value = user.username;
    }
    catch (error) {
      apiErrorText.value = error;
      isAsaiAlertActive.value = true;
    }
  }
});
</script>

<template>
  <div class="container d-flex flex-column vh-100">
    <AsaiAlert :is-active="isAsaiAlertActive" :errorText="apiErrorText" @close="closeAsaiAlert" />
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
    <div class="row">
      <div class="col-md-6">
        <img class="logo" src="@/assets/ASAILogotype.svg" alt="">
        <div class="card">
          <div class="card-body">
            <Form class="form-control" @submit="register" :validation-schema="schema">
              <ErrorMessage name="Username" />
              <Field v-model="username" name="Username" type="email" class="email-input d-block"
                placeholder="Username" />
              <ErrorMessage name="Password" />
              <Field v-model="password" name="Password" type="password" class="pass-input d-block"
                placeholder="Password" />
              <ErrorMessage name="ConfirmPassword" />
              <Field v-model="confirmPassword" name="ConfirmPassword" type="password" class="pass-input d-block"
                placeholder="Confirm Password" />
              <button class="send-button btn btn-light" :disabled="formState.isSubmitting">
                <span v-if="formState.isSubmitting">
                  <span class="loader"></span>
                </span>
                <span v-else>REGISTER</span>
              </button>
            </Form>
          </div>
        </div>
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

@media only screen and (max-width: 576px) {
  nav {
    margin-top: 10px;
    margin-bottom: 2em;
  }
}
</style>
