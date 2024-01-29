<script setup>
import { onMounted, toRef } from 'vue';
import { usePluginStore } from '@/stores/plugin.store';
import { useUserStore } from '@/stores/user.store';
import { useAvatarStore } from '@/stores/avatar.store';

const user = useUserStore();

const plugin = usePluginStore();
const pluginsRecords = toRef(plugin, 'records');

const avatar = useAvatarStore();
const activePlugins = toRef(avatar, 'activePlugins');

const pluginIsActive = (pluginID) => {
  return pluginsRecords.value.some(plugin => plugin.ID === pluginID && plugin.isActive);
}

const getActivePluginID = (pluginID) => {
  const activePlugin = activePlugins.value.find(p => p.tool.ID === pluginID);
  return activePlugin ? activePlugin.ID : null;
}

onMounted(async () => {
  try {
    await plugin.getPlugins();
    await avatar.getActivePlugins(user.avatar.ID);
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
});
</script>

<template>
          
  <div class="container-fluid p-0">
    
    <h1 class="h3 mb-3">Plugins</h1>
    <div class="row">
      <div class="col-12">
        <div class="container">
          
          <div class="row">
            <div class="row" v-for="(_, index) in pluginsRecords.filter((a, i) => i % 2 === 0)" :key="'row' + index">
              <!-- Render the current and next agent (if it exists) within the same row -->
              <div class="col-6" v-for="plugin in pluginsRecords.slice(index, index + 2)" :key="plugin.ID">
                <div class="card">
                  
                  <div class="card-header">
                    <div class="row">
                      <div class="col">
                        <h5 class="card-title">{{ plugin.name }}</h5>
                      </div>
                      <div class="col-auto">
                        <div class="form-check form-switch d-flex align-items-center" v-if="getActivePluginID(plugin.ID)">
                          <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault" :checked="pluginIsActive(plugin.ID)">
                          <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label>
                        </div>
                      </div>
                    </div>
                  </div>

                  <div class="card-body">
                    <p>{{ plugin.description }}</p>
                    <div>
                      <router-link 
                      :to="{name: 'plugin-config', params: {avatar_id: user.avatar.ID, plugin_id: plugin.ID, active_plugin_id: getActivePluginID(plugin.ID)}}" 
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
  </div>
      
</template>

<style scoped>
.card {
    min-height: 220px;
  }
</style>