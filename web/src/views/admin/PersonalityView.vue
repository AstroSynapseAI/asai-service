<script setup>
import { ref, onMounted, toRef } from 'vue';
import { Form, Field } from 'vee-validate';
import { useAvatarStore } from '@/stores/avatar.store.js';
import { useLLMStore } from '@/stores/llm.store.js';
import { useUserStore } from '@/stores/user.store.js';


const llmStore = useLLMStore();
const avatarStore = useAvatarStore();
const userStore = useUserStore();


// Form data
const avatarName = ref('');
const avatarLLMID = ref('');
const avatarPrimer = ref('');

const llms = toRef(llmStore, 'llms');
llmStore.getLLMs();

console.log(JSON.parse(localStorage.getItem('user')))

console.log(userStore.currentUser);

const userAvatar = userStore.getUserAvatar(userStore.currentUser.ID);
if (userAvatar) {
  avatarName.value = userAvatar.name;
  avatarLLMID.value = userAvatar.llm_id;
  avatarPrimer.value = userAvatar.primer;
}

const submitForm = () => {
  const formData = {
    name: avatarName.value,
    llm: avatarLLMID.value,
    primer: avatarPrimer.value,
  }

  if (userAvatar) {
    formData.ID = userAvatar.ID;
  }

  console.log(formData);

  avatarStore.saveAvatar(formData);
}

onMounted(() => {
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

<style scoped>
.model-select {
  height: 58px;
}
</style>