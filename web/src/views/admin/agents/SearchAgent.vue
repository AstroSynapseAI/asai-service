<script setup>
import { onMounted, ref, toRef } from 'vue';
import { Form, Field } from 'vee-validate';
import { useAgentStore } from '@/stores/agent.store';
import { useLLMStore } from '@/stores/llm.store';
import { useRoute } from 'vue-router';
import { useAvatarStore } from '@/stores/avatar.store';
import { useToast } from 'vue-toastification';

const toast = useToast();

const avatar = useAvatarStore();
const agent = useAgentStore();
const agentRecord = toRef(agent, 'record');
const route = useRoute();
const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

const agentPrimer = ref('');
const agentName = ref('');
const activeAgentLLMID = ref('');
const isActiveAgent = ref(false);
const isPublicAgent = ref(false);
const config = ref({
  "ddg_is_active": true,
  "google_api_token": "",
  "google_is_active": false,
  "exa_api_token": "",
  "exa_is_active": false
});

const toggleSearchEngine = (id) => {
  console.log(id)

  if (id === 'ddg') {
    config.value.ddg_is_active = !config.value.ddg_is_active
  }

  if (id === 'google') {
    config.value.google_is_active = !config.value.google_is_active
  }

  if (id === 'exa') {
    config.value.exa_is_active = !config.value.exa_is_active
  }
}

const submitForm = async () => {
  const formData = {
    ID: parseInt(route.params.active_agent_id),
    avatar_id: parseInt(route.params.avatar_id),
    agent_id: parseInt(route.params.agent_id),
    llm_id: activeAgentLLMID.value,
    primer: agentPrimer.value,
    is_active: isActiveAgent.value,
    is_public: isPublicAgent.value,
    config: JSON.stringify(config.value),
  }

  if (agentRecord.value.ID) {
    formData.ID = agentRecord.value.ID;
  }

  try {
    await agent.saveActiveAgent(formData);
    toast.success("Agent saved successfully");
  }
  catch (error) {
    toast.error("Error while saving agent");
    console.log(error);
  }
}

onMounted(async () => {
  try {
    await llm.getLLMs();
    await agent.getAgent(route.params.agent_id);

    agentPrimer.value = agentRecord.value.primer;
    agentName.value = agentRecord.value.name;

    if (route.params.active_agent_id) {
      await avatar.getActiveAgent(route.params.agent_id, route.params.avatar_id);
      activeAgentLLMID.value = avatar.activeAgent.llm_id;
      agentPrimer.value = avatar.activeAgent.primer;
      isActiveAgent.value = avatar.activeAgent.is_active;
      isPublicAgent.value = avatar.activeAgent.is_public;

      if (avatar.activeAgent.config) {
        config.value = JSON.parse(avatar.activeAgent.config);
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
  <h1 class="h3 mb-3">Configure: {{ agentRecord.name }}
    <div class="form-check form-switch float-end me-5">
      <label class="form-check-label" for="flexSwitchCheckDefault">Active</label>
      <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isActiveAgent"
        @click="isActiveAgent = !isActiveAgent">
    </div>
    <div class="form-check form-switch float-end me-3">
      <label class="form-check-label" for="flexSwitchCheckDefault">Public</label>
      <input class="form-check-input" type="checkbox" id="flexSwitchCheckDefault" :checked="isPublicAgent"
        @click="isPublicAgent = !isPublicAgent">
    </div>
  </h1>


  <div class="card">
    <div class="card-body">
      <div class="container-fluid">

        <div class="row">

          <div class="col-6">
            <h3 class="mb-3">Primer</h3>
            <Field v-model="agentPrimer" name="agent_primer" type="text" as="textarea" class="form-control" rows="30"
              placeholder="How should the agent behave..."></Field>
          </div>

          <div class="col-6">

            <div class="row">
              <div class="col-12">
                <h3 class="mb-3">Agent Model</h3>
                <h5> LLM </h5>
                <select v-model="activeAgentLLMID" class="form-select model-select" aria-label="Select Model">
                  <option value="" disabled selected>Select a LLM</option>
                  <option v-for="(llm, index) in llmRecords" :value="llm.ID" :key="index">
                    {{ llm.name }}
                  </option>
                </select>
              </div>
            </div>

            <div class="row">
              <div clas="col-12">
                <h3 class="my-3">Search Tools</h3>
                <div class="form-check form-switch mb-3">
                  <span>Duck Duck Go</span>
                  <input class="form-check-input" type="checkbox" id="ddg-checkbox" @click="toggleSearchEngine('ddg')"
                    :checked="config.ddg_is_active">
                </div>

                <div class="form-check form-switch">
                  <input class="form-check-input" type="checkbox" id="google-checkbox"
                    @click="toggleSearchEngine('google')" :checked="config.google_is_active">
                  <span> Google Search</span>
                </div>
                <div class="form-floating mb-3">
                  <input type="text" class="form-control" id="apiToken" placeholder="Tool Token"
                    v-model="config.google_api_token">
                  <label for="apiToken">Api Token</label>
                </div>

                <div class="form-check form-switch">
                  <input class="form-check-input" type="checkbox" id="exa-checkbox" @click="toggleSearchEngine('exa')"
                    :checked="config.exa_is_active">
                  <span> Exa Search</span>
                </div>
                <div class="form-floating mb-3">
                  <input type="text" class="form-control" id="apiToken" placeholder="Tool Token"
                    v-model="config.exa_api_token">
                  <label for="apiToken">Api Token</label>
                </div>
              </div>
            </div>

          </div>

        </div>

        <div class="row mt-3">
          <div class="col-12">
            <button type="button" class="btn btn-primary" @click="submitForm">Save All</button>
          </div>
        </div>

      </div>
    </div>
  </div>

</template>

<style scoped>
.form-control, .form-select {
  background-color: #374151;
}
</style>
