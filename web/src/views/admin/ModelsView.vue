<script setup>
import { onMounted, toRef, ref } from 'vue';
import { useLLMStore } from '@/stores/llm.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';

const user = useUserStore();
const avatar = useAvatarStore();
const activeModels = toRef(avatar, 'activeLLMs');

const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

const isActive = (ID) => {
  const activeModel = activeModels.value.find(activeModel => {
    return activeModel.llm.ID == ID;
  });
  
  return activeModel ? activeModel.is_active : false; 
}

const toggleActive = async (ID) => {
  const activeModel = activeModels.value.find(activeModel => {
    return activeModel.llm.ID == ID;
  });
  
  if(activeModel){
    activeModel.is_active = !activeModel.is_active;
  }

  const formData = {
    is_active: activeModel ? activeModel.is_active : false,
    avatar_id: user.avatar.ID
  }
  try {    
    await llm.toggleActiveLLM(ID, formData)
  }
  catch (error) {
    console.log(error);
  }
}

const getActiveLLMID = (llmID) => {
  const activeLLM = activeModels.value.find(m => m.llm.ID === llmID);
  return activeLLM ? activeLLM.ID : null;
}

onMounted(async () => {
  try {
    await llm.getLLMs();
    await avatar.getActiveLLMs(user.avatar.ID);
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
});
</script>

<template>
          
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Models</h1>
    
    <div class="card" v-for="(llm, index) in llmRecords" :key="'row' + index">
      <div class="card-header">
        <div class="row">
          <div class="col">
            <h5 class="card-title">{{ llm.name }}</h5>
          </div>
          <div class="col-auto">
            <div class="form-check form-switch d-flex align-items-center" v-if="getActiveLLMID(llm.ID)">
              <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault" :checked="isActive(llm.ID)" @click="toggleActive(llm.ID)">
              <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label>
            </div>
          </div>
        </div>
      </div>

      <div class="card-body">
        <p>{{ llm.description }}</p>
        <div>
          <router-link 
          :to="{name: 'model-config', params: {avatar_id: user.avatar.ID, model_id: llm.ID, active_model_id: getActiveLLMID(llm.ID)}}" 
          class="btn 
          btn-primary">
            Configure
          </router-link>
        </div>
      </div>
    </div>

  </div>
        
  </template>