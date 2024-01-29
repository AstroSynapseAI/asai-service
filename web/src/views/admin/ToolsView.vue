<script setup>
import { onMounted, ref, toRef } from 'vue';
import { useToolStore } from '@/stores/tool.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';

const user = useUserStore();
const tool = useToolStore();
const toolsRecords = toRef(tool, 'records');

const avatar = useAvatarStore();
const activeTools = toRef(avatar, 'activeTools');

const toolIsActive = ref(false);

onMounted( async () => {
  try {
    await tool.getTools();
    await avatar.getActiveTools(user.avatar.ID);
  }
  catch (error) {
    console.log(error);
  }

  feather.replace();
})
</script>

<template>
          
  <div class="container-fluid p-0">
    
    <h1 class="h3 mb-3">Tools</h1>
    <div class="row">
      <div class="col-12">
        <div class="container">
          
          <div class="row">
            <div class="row" v-for="(tool, index) in toolsRecords" :key="'row' + index">
              
              
              <div class="card">
                  
                <div class="card-header">
                  <div class="row">
                    <div class="col">
                      <h5 class="card-title">{{ tool.name }}</h5>
                    </div>
                    <div class="col-auto">
                      <div class="form-check form-switch d-flex align-items-center">
                        <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault" :checked="toolIsActive">
                        <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="card-body">
                  <p>{{ tool.description }}</p>
                  <div>
                    <router-link 
                    :to="{name: 'tool-config', params: {avatar_id: user.avatar.ID, tool_id: tool.ID}}" 
                    class="btn 
                    btn-primary">
                      Configure
                    </router-link>
                  </div>
                </div>
              
              </div>
              
            </div>
          </div>


        </div>
      </div>
    </div>
  </div>
      
</template>