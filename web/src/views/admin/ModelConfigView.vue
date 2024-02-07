<script setup>
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref, toRef } from 'vue';
import { useLLMStore } from '@/stores/llm.store';
import { useAvatarStore } from '@/stores/avatar.store';

const route = useRoute();
const router = useRouter();
const llm = useLLMStore();
const avatar = useAvatarStore();

const modelName = ref('');
const modelToken = ref('');
const isActive = ref(false);

const toggleActive = () => {
  isActive.value = !isActive.value
}

const saveModel = async () => {
  try {
    await llm.saveLLM({
      ID: parseInt(route.params.active_model_id),
      avatar_id: parseInt(route.params.avatar_id),
      llm_id: parseInt(route.params.model_id),
      is_active: isActive.value,
      token: modelToken.value
    })  
    router.push({name: 'models', params: {avatar_id: route.params.avatar_id}});
  }
  catch (error) {
    console.log(error);
  }
  
}

onMounted(async () => {
  try {
    await llm.getLLM(route.params.model_id);
    modelName.value = llm.record.name;
    if (route.params.active_model_id) {
      await avatar.getActiveLLM( route.params.avatar_id, route.params.model_id);
      if (avatar.activeLLM) {
        isActive.value = avatar.activeLLM.is_active;
        modelToken.value = avatar.activeLLM.token;
      }
    }
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
})
</script>

<template>

  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Configure: {{ modelName }}
      <div class="form-check form-switch float-end me-5">
        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isActive" @click="toggleActive">
      </div>
    </h1>

    <div class="row">
      <div class="col-12">

        <div class="card">

          <div class="card-body">
            <div class="container">

              <div class="row">
                <div class="col-12">
                  <div class="form-floating mb-3">
                    <input v-model="modelToken" type="text" class="form-control" id="floatingInput" placeholder="Token...">
                    <label for="floatingInput">Token</label>
                  </div>
                </div>
              </div>
              
              <div class="row mt-3">
                <div class="col-12">
                  <button type="button" class="btn btn-secondary" @click="saveModel">Save</button>
                </div>    
              </div>

            </div>

          </div>

        </div>

      </div>
    </div>
  </div>

  </template>