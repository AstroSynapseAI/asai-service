<script setup>
import { onMounted, ref, toRef, computed } from 'vue';
import { useAgentStore } from '@/stores/agent.store';
import { useUserStore } from '@/stores/user.store';
import { useAvatarStore } from '@/stores/avatar.store';

const user = useUserStore();

const agent = useAgentStore();
const agentsRecords = toRef(agent, 'records');

const avatar = useAvatarStore();
const activeAgents = toRef(avatar, 'activeAgents');

const agentIsActive = (agentID) => {
  const activeAgent = activeAgents.value.find(activeAgent => {
    return activeAgent.agent.ID === agentID;
  })

  return activeAgent ? activeAgent.is_active : false;
}

const toggleActive = async (agentID) => {
  const activeAgent = activeAgents.value.find(activeAgent => {
    return activeAgent.agent.ID == agentID;
  });

  if (activeAgent) {
    activeAgent.is_active = !activeAgent.is_active;
  }

  const formData = {
    is_active: activeAgent ? activeAgent.is_active : false,
    avatar_id: user.avatar.ID
  }

  try {
    await agent.toggleActiveAgent(agentID, formData);
  }
  catch (error) {
    console.log(error);
  }
}

const getActiveAgentID = (agentID) => {
  const activeAgent = activeAgents.value.find(a => a.agent.ID === agentID);
  return activeAgent ? activeAgent.ID : null;
}

const chunkedAgentsRecords = computed(() => {
  const result = [];
  const items = [...agentsRecords.value]; // Clone the original array to avoid direct modification
  while (items.length) {
    result.push(items.splice(0, 2)); // Split the array into chunks of 2
  }
  return result;
});

onMounted(async () => {
  await agent.getAgents();
  try {
    await avatar.getActiveAgents(user.avatar.ID);
  }
  catch (error) {
    console.log(error);
  }

  feather.replace();
});

</script>

<template>

  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Agents</h1>
    <div class="row">
      <div class="col-12">
        <div class="container">

          <div v-for="(chunk, index) in chunkedAgentsRecords" :key="index" class="row">
            <div class="col-6" v-for="agent in chunk" :key="agent.ID">
              <div class="card">
                <div class="card-header">
                  <div class="row">
                    <div class="col">
                      <h5 class="card-title">{{ agent.name }}</h5>
                    </div>
                    <div class="col-auto">
                      <div class="form-check form-switch d-flex align-items-center" v-if="getActiveAgentID(agent.ID)">
                        <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault"
                          :checked="agentIsActive(agent.ID)" @click="toggleActive(agent.ID)">
                        <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="card-body">
                  <p>{{ agent.description }}</p>
                  <div>
                    <router-link
                      :to="{ name: agent.slug, params: { avatar_id: user.avatar.ID, agent_id: agent.ID, active_agent_id: getActiveAgentID(agent.ID) } }"
                      class="btn btn-primary">Configure</router-link>
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
  min-height: 250px;
}
</style>
