<script setup>
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref, toRef } from 'vue';
import { usePluginStore } from '@/stores/plugin.store';
import { useAvatarStore } from '@/stores/avatar.store';

const route = useRoute();
const router = useRouter();
const plugin = usePluginStore();
const avatar = useAvatarStore();

const pluginToken = ref('');
const isActivePlugin = ref(false);
const isPublicPlugin = ref(false);
const pluginName = ref('');

const toggleActive = () => {
  isActivePlugin.value = !isActivePlugin.value
}

const togglePublic = () => {
  isPublicPlugin.value = !isPublicPlugin.value
}

const savePlugin = async () => {
  try {
    await plugin.saveActivePlugin({
      ID: parseInt(route.params.active_plugin_id),
      avatar_id: parseInt(route.params.avatar_id),
      plugin_id: parseInt(route.params.plugin_id),
      token: pluginToken.value,
      is_active: isActivePlugin.value,
      is_public: isPublicPlugin.value
    });

    router.push({name: 'plugins', params: {avatar_id: route.params.avatar_id}});
  }
  catch (error) {
    console.log(error);
  }
  
}

onMounted(async () => {
  try {
    await plugin.getPlugin(route.params.plugin_id);
    pluginName.value = plugin.record.name;
    if (route.params.active_plugin_id) {
      await avatar.getActivePlugin(route.params.avatar_id, route.params.plugin_id);
      if (avatar.activePlugin) {
        isActivePlugin.value = avatar.activePlugin.is_active;
        isPublicPlugin.value = avatar.activePlugin.is_public;
        pluginToken.value = avatar.activePlugin.token;
      }
    }
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
});


</script>

<template>

  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Configure: {{ pluginName }}
      <div class="form-check form-switch float-end me-5">
        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isActivePlugin" @click="toggleActive">
      </div>
      <div class="form-check form-switch float-end me-3">
        <label class="form-check-label" for="flexSwitchCheckDefault">Public</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isPublicPlugin" @click="togglePublic">
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
                    <input v-model="pluginToken" type="text" class="form-control" id="floatingInput" placeholder="Token...">
                    <label for="floatingInput">Token</label>
                  </div>
                </div>
              </div>
              
              <div class="row mt-3">
                <div class="col-12">
                  <button type="button" class="btn btn-secondary" @click="savePlugin">Save</button>
                </div>    
              </div>

            </div>

          </div>

        </div>

      </div>
    </div>
  </div>

  </template>