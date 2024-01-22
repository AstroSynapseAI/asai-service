<script setup>
import { ref, onMounted, toRef } from 'vue';
import { Form, Field } from 'vee-validate';
import { useAvatarStore } from '@/stores/avatar.store.js';
import { useLLMStore } from '@/stores/llm.store.js';
import { useUserStore } from '@/stores/user.store.js';

// Initiate stores
const llm = useLLMStore();
const avatar = useAvatarStore();
const user = useUserStore();

// Form data
const avatarName = ref('');
const avatarLLMID = ref('');
const avatarPrimer = ref('');

const llms = toRef(llm, 'llms');

const submitForm = () => {
  const formData = {
    name: avatarName.value,
    llm: avatarLLMID.value,
    primer: avatarPrimer.value,
  }

  if (user.avatar.ID) {
    formData.ID = user.avatar.ID;
  }

  console.log(formData);
  avatar.saveAvatar(formData);
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