<script setup>
import { onMounted, ref, toRef } from 'vue';
import { useToolStore } from '@/stores/tool.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';

const user = useUserStore();
const avatar = useAvatarStore();
const activeTools = toRef(avatar, 'activeTools');

const tool = useToolStore();
const toolsRecords = toRef(tool, 'records');

const isActive = (ID) => {
  const activeTool = activeTools.value.find(activeTool => {
    return activeTool.tool.ID == ID;
  });

  return activeTool ? activeTool.is_active : false;
}

const toggleActive = async (ID) => {
  const activeTool = activeTools.value.find(activeTool => {
    return activeTool.tool.ID == ID;
  });

  if (activeTool) {
    activeTool.is_active = !activeTool.is_active;
  }

  const formData = {
    is_active: activeTool ? activeTool.is_active : false,
    avatar_id: user.avatar.ID
  }
  try {
    await tool.toggleAvatarTool(ID, formData)
  }
  catch (error) {
    console.log(error);
  }
}

const getActiveToolID = (ToolID) => {
  const activeTool = activeTools.value.find(m => m.tool.ID === ToolID);
  return activeTool ? activeTool.ID : null;
}

onMounted(async () => {
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
                      <h5 class="card-title">
                        {{ tool.name }}
                        <span class="badge bg-success ms-5">Releasing in v.0.2.0</span>
                      </h5>
                    </div>
                    <div class="col-auto">
                      <div class="form-check form-switch d-flex align-items-center" v-if="getActiveToolID(tool.ID)">
                        <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault"
                          :checked="isActive(tool.ID)" @click="toggleActive(tool.ID)">
                        <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="card-body">
                  <p>{{ tool.description }}</p>
                  <div>
                    <button class="btn btn-primary">Configure</button>
                    <!-- <router-link  -->
                    <!-- :to="{name: 'tool-config', params: {avatar_id: user.avatar.ID, tool_id: tool.ID, active_tool_id: getActiveToolID(tool.ID)}}"  -->
                    <!-- class="btn  -->
                    <!-- btn-primary"> -->
                    <!--   Configure -->
                    <!-- </router-link> -->
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
