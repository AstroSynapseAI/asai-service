<script setup>
import { ref, onMounted, toRef, reactive } from 'vue';
import { Form, Field, ErrorMessage } from 'vee-validate';
import { useAvatarStore } from '@/stores/avatar.store.js';
import { useLLMStore } from '@/stores/llm.store.js';
import { useUserStore } from '@/stores/user.store.js';
import { useToast } from 'vue-toastification';
import * as yup from 'yup';

const toast = useToast();
const schema = yup.object({
  AvatarName: yup.string().required("The Avatar Name is required"),
  AvatarLLMID: yup.string().required("The LLM is required"),
  AvatarPrimer: yup.string().required("The Primer is required"),
});

const formState = reactive({
  isSubmitting: false, 
});

// Initiate stores
const llm = useLLMStore();
const avatar = useAvatarStore();
const user = useUserStore();

// Form data
const avatarName = ref('');
const avatarLLMID = ref('');
const avatarPrimer = ref('');

const llms = toRef(llm, 'records');

const submitForm = async () => {
  formState.isSubmitting = true; 
  try {
    const formData = {
      name: avatarName.value,
      llm: avatarLLMID.value,
      primer: avatarPrimer.value,
    }
  
    if (user.avatar.ID) {
      formData.ID = user.avatar.ID;
    }
    await avatar.saveAvatar(formData);
    toast.success('Avatar saved successfully');
    formState.isSubmitting = false;
  } catch (error) {
    toast.error(error)
    formState.isSubmitting = false; 
  }
}

onMounted(async () => {
  try {
    await llm.getLLMs();
    await user.getUserAvatar(user.current.ID);

    if (user.avatar) {
      avatarName.value = user.avatar.name;
      avatarLLMID.value = user.avatar.llm_id;
      avatarPrimer.value = user.avatar.primer;
    }
  } catch (error) {
    console.log(error);
  }

  feather.replace();
});

</script>

<template>

  <div class="container-fluid p-0">

    <h1 class="h3 mb-3">Personality</h1>

    <div class="row">
      <div class="col-12">

        <div class="card">

          <div class="card-body">
            <Form class="container p-4" @submit="submitForm" :validation-schema="schema">
              <div class="row mb-5">

                <div class="col-6">
                  <div class="form-floating mb-1">
                    <Field v-model="avatarName" name="AvatarName" type="text" class="form-control" id="floatingInput" placeholder="Name your Avatar..."/>
                    <label for="floatingInput">Avatar name</label>
                  </div>
                  <ErrorMessage name="AvatarName" />
                </div>

                <div class="col-6">
                  <Field name="AvatarLLMID" as="select" v-model="avatarLLMID"  class="form-select model-select mb-1" aria-label="Select Model">
                    <option value="" disabled selected>Select a LLM</option>
                    <option v-for="(llm, index) in llms" :value="llm.ID" :key="index">
                      {{ llm.name }}
                    </option>
                  </Field>
                  <ErrorMessage name="AvatarLLMID" />
                </div>

              </div>

              <div class="row mb-5">
                <div class="col-12">
                  <h3>Primer</h3>

                  <Field
                    v-model="avatarPrimer"
                    name="AvatarPrimer"
                    type="text"
                    as="textarea"
                    class="form-control mb-4"
                    rows="8"
                    placeholder=""
                  />
                  <ErrorMessage name="AvatarPrimer"/>
                </div>
              </div>

              <button type="submit" class="btn btn-secondary" :disabled="formState.isSubmitting">
                <span v-if="formState.isSubmitting">
                  <span class="loader"></span>
                </span>
                <span v-else>Save</span>
              </button>
            </Form>
          </div>

        </div>

      </div>
    </div>
  </div>

</template>

<style scoped>
.model-select {
  height: 58px;
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
</style>