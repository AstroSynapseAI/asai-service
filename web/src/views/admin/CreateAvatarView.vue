<script setup>
import { ref, onMounted, toRef } from 'vue';
import { useRouter } from 'vue-router';
import { Form, Field } from 'vee-validate';
import { useLLMStore } from '@/stores/llm.store.js';
import { useAvatarStore } from '@/stores/avatar.store.js';

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
    console.log(error)
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
            <div class="container p-4">

              <div class="row mb-5">

                <div class="col-6">
                  <div class="form-floating mb-3">
                    <input v-model="avatarName" type="text" class="form-control" id="floatingInput" placeholder="Name your Avatar...">
                    <label for="floatingInput">Avatar name</label>
                  </div>
                </div>

                <div class="col-6">
                  <select v-model="avatarLLMID" class="form-select model-select" aria-label="Select Model">
                    <option value="" disabled selected>Select a LLM</option>
                    <option v-for="(llm, index) in llms" :value="llm.ID" :key="index">
                      {{ llm.name }}
                    </option>
                  </select>
                </div>

              </div>

              <div class="row mb-5">
                <div class="col-12">
                  <h3>Primer</h3>
                  <Field
                    v-model="avatarPrimer"
                    name="avatar-primer"
                    type="text"
                    as="textarea"
                    class="form-control"
                    rows="8"
                    placeholder=""
                  ></Field>
                </div>
              </div>

              <button type="button" class="btn btn-secondary" @click="submitForm">Save</button>

            </div>
          </div>

        </div>

      </div>
    </div>
  </div>

</template>