<script setup>
import { ref, onMounted, toRef, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { useToast } from 'vue-toastification';
import * as yup from 'yup';

const router = useRouter();

const toast = useToast();
const schema = yup.object({
  AvatarName: yup.string().required("The Avatar Name is required"),
  AvatarPrimer: yup.string().required("The Avatar Primer is required"),
});

// Form data
const avatarName = ref('');
const avatarPrimer = ref('');

const submitForm = () => {
  let onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));

  if (!onboardingData) {
    onboardingData = {};
  }

  onboardingData['avatar_name'] = avatarName.value;
  onboardingData['avatar_primer'] = avatarPrimer.value;

  localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
  router.push({ name: 'choose-model'});
}

onMounted(async () => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  if (onboardingData) {
    avatarName.value = onboardingData.avatar_name;
    avatarPrimer.value = onboardingData.avatar_primer;
  }
  feather.replace();
});

</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5 current"><h2 class="circle-text">1</h2></div>
        <h3 class="mb-3">Create Avatar</h3>
        <p class="lead mb-5">Give your AI avatar a name and describe how it should behave.</p>
      </div>
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5"><h2 class="circle-text">2</h2></div>
        <h3 class="mb-3">Choose models</h3>
        <p class="lead mb-5">Select one or more LLM models yor avatar will be using.</p>
      </div>
      <div class="col-4 text-center d-flex flex-column align-items-center">
        <div class="circle mb-5"><h2 class="circle-text">3</h2></div>
        <h3 class="mb-3">Select Agents</h3>
        <p class="lead mb-5">Your AI Avatar can browse the internet, answer emails, post on social media nad more!</p>
      </div>
    </div>

    <div class="row">
      <div class="col-8 offset-1">

        <div class="card">

          <div class="card-body">
            <Form class="container p-4" @submit="submitForm" :validation-schema="schema">

              <div class="row mb-5">

                <div class="col-12">
                  <div class="form-floating mb-1">
                    <Field v-model="avatarName" name="AvatarName" type="text" class="form-control" id="floatingInput"
                      placeholder="Name your Avatar..." />
                    <label for="floatingInput">Avatar name</label>
                  </div>
                  <ErrorMessage name="AvatarName" />
                </div>

              </div>

              <div class="row mb-5">
                <div class="col-12">
                  <h3>Primer</h3>

                  <Field v-model="avatarPrimer" name="AvatarPrimer" type="text" as="textarea" class="form-control mb-4"
                    rows="18" placeholder="" />
                  <ErrorMessage name="AvatarPrimer" />
                </div>
              </div>

              <div class="row">
                <div class="col-6 text-center d-grid">
                  <router-link :to="{ name: 'welcome' }" class="btn btn-primary btn-lg btn-back">Back</router-link>
                </div>
                <div class="col-6 text-center d-grid mx-auto">
                  <button type="submit" class="btn btn-primary btn-lg">Next</button>
                </div>
              </div>
            </Form>
          </div>

        </div>

      </div>
    </div>
  </div>
</template>

<style scoped>
h1, h3 {
  font-weight: 700;
}

.form-control  {
  background-color: #374151;
}

.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.btn-back {
  background-color: transparent;
}
.subtitle, .lead {
  color: grey;
}

.circle {
   display: flex;
   justify-content: center;
   align-items: center;
   width: 85px;
   height: 85px;
   border: 2px solid #1c64f2;
   border-radius: 50%;
}

.current {
  background-color: #1c64f2;
}

.circle-text {
   text-align: center;
   line-height: 50px; /* match this with the height of .circle */
   margin: 0;
}

</style>
