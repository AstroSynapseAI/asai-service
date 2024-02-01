<script setup>
import { onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useToolStore } from '@/stores/tool.store';
import { useAvatarStore } from '@/stores/avatar.store';

const route = useRoute();
const router = useRouter();
const avatar = useAvatarStore();
const tool = useToolStore();

const toolToken = ref('');
const isActiveTool = ref(false);
const isPublicTool = ref(false);
const toolName = ref(''); 

const toggleActive = () => {
  isActiveTool.value = !isActiveTool.value
}

const togglePublic = () => {
  isPublicTool.value = !isPublicTool.value
}

const saveTool = async () => {
  try {
    await tool.saveAvatarTool({
      ID: parseInt(route.params.active_tool_id),
      avatar_id: parseInt(route.params.avatar_id),
      tool_id: parseInt(route.params.tool_id),
      token: toolToken.value,
      is_active: isActiveTool.value,
      is_public: isPublicTool.value
    })
    router.push({name: 'tools', params: {avatar_id: route.params.avatar_id}});
  }
  catch (error) {
    console.log(error);
  }
}

onMounted( async () => {
  try {
    await tool.getTool(route.params.tool_id);
    toolName.value = tool.record.name;
    if (route.params.active_tool_id) {
      await avatar.getActiveTool(route.params.avatar_id, route.params.tool_id);
      if (avatar.activeTool) {
        isActiveTool.value = avatar.activeTool.is_active;
        isPublicTool.value = avatar.activeTool.is_public;
        toolToken.value = avatar.activeTool.token;
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
    <h1 class="h3 mb-3">Configure: {{ toolName }}
      <div class="form-check form-switch float-end me-5">
        <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isActiveTool" @click="toggleActive">
      </div>
      <div class="form-check form-switch float-end me-3">
        <label class="form-check-label" for="flexSwitchCheckDefault">Public</label>
        <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isPublicTool" @click="togglePublic">
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
                    <input v-model="toolToken" type="text" class="form-control" id="floatingInput" placeholder="Token...">
                    <label for="floatingInput">Token</label>
                  </div>
                </div>
              </div>
              
              <div class="row mt-3">
                <div class="col-12">
                  <button type="button" class="btn btn-secondary" @click="saveTool">Save</button>
                </div>    
              </div>

            </div>

          </div>

        </div>

      </div>
    </div>
  </div>

</template>