<script setup>
import { ref, onMounted, toRef, reactive } from 'vue';
import { useRouter } from 'vue-router';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { useLLMStore } from '@/stores/llm.store.js';
import { useAvatarStore } from '@/stores/avatar.store.js';
import { useToast } from 'vue-toastification';
import * as yup from 'yup';

const toast = useToast();
const schema = yup.object({
  AvatarName: yup.string().required("The Avatar Name is required"),
  AvatarLLMID: yup.string().required("The LLM is required"),
  AvatarPrimer: yup.string()
});

const formState = reactive({
  isSubmitting: false,
});

// Initiate stores
const llm = useLLMStore();
const avatar = useAvatarStore();
const router = useRouter();

// Form data
const avatarName = ref('');
const avatarLLMID = ref('');
const avatarPrimer = ref('');

const llms = toRef(llm, 'records')

const submitForm = async () => {
  formState.isSubmitting = true;
  try {
    const formData = {
      name: avatarName.value,
      llm: avatarLLMID.value,
      primer: avatarPrimer.value,
    }
    await avatar.saveAvatar(formData);
    await router.replace({ name: 'personality', params: { avatar_id: avatar.userAvatar.ID } });
    toast.success('Avatar saved successfully');
    formState.isSubmitting = false;
    window.location.reload();
  }
  catch (error) {
    toast.error(error)
    formState.isSubmitting = false;
  }
}

onMounted(async () => {
  try {
    await llm.getLLMs();
  } catch (error) {
    console.log(error);
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
                  <router-link :to="{ name: 'choose-model' }" class="btn btn-primary btn-lg">Next</router-link>
                </div>
              </div>
              <!-- <button type="submit" class="btn btn-secondary" :disabled="formState.isSubmitting"> -->
              <!--   <span v-if="formState.isSubmitting"> -->
              <!--     <span class="loader"></span> -->
              <!--   </span> -->
              <!--   <span v-else>Save</span> -->
              <!-- </button> -->
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

.form-control {
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
