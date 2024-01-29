<script setup>
import { onMounted, toRef } from 'vue';
import { useLLMStore } from '@/stores/llm.store';
import { useAvatarStore } from '../../stores/avatar.store';
import { useUserStore } from '../../stores/user.store';

const user = useUserStore();
const avatar = useAvatarStore();
const activeModels = toRef(avatar, 'activeLLMs');

const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

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
    <h1 class="h3 mb-3">Models
      <!-- <button class="btn btn-primary float-end">Add Models <i data-feather="plus-square"></i></button> -->
    </h1>
    
    
    <div class="card" v-for="(llm, index) in llmRecords" :key="'row' + index">
      <div class="card-header">
        <div class="row">
          <div class="col">
            <h5 class="card-title">{{ llm.name }}</h5>
          </div>
          <div class="col-auto">
            <div class="form-check form-switch d-flex align-items-center">
              <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault">
              <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label>
            </div>
          </div>
        </div>
      </div>

      <div class="card-body">
        <p>{{ llm.description }}</p>
        <div>
          <router-link 
          :to="{name: 'model-config', params: {avatar_id: user.avatar.ID, model_id: llm.ID}}" 
          class="btn 
          btn-primary">
            Configure
          </router-link>
        </div>
      </div>
    </div>

  </div>
        
  </template>