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

const submitForm =  async () => {
  formState.isSubmitting = true; 
  try {
    const formData = {
      name: avatarName.value,
      llm: avatarLLMID.value,
      primer: avatarPrimer.value,
    }
    await avatar.saveAvatar(formData);
    await router.replace({name: 'personality', params: {avatar_id: avatar.userAvatar.ID}});
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

  <div class="container-fluid p-0">

    <h1 class="h3 mb-3">Create your AI Avatar</h1>

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
</style>