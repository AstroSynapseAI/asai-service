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
  "imap_server": "",
  "smtp_server": "",
  "imap_port": "",
  "smtp_port": "",
  "username": "",
  "password": "",
  "encryption": "",
  "sender": "",
  "reply_to": "",
});

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
    toast.success('Agent configuration saved');
  }
  catch (error) {
    console.log(error);
    toast.error('Error while saving agent configuration');
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
                <h3 class="my-3">Email Config</h3>

                <div class="row">

                  <div class="col-6">
                    <div class="form-floating mb-3">
                      <input type="text" class="form-control" id="smtp-server" placeholder="mail.example.com"
                        v-model="config.smtp_server">
                      <label for="smtp-server">SMTP Server</label>
                    </div>
                  </div>

                  <div class="col-4">
                    <div class="form-floating mb-3">
                      <select v-model="config.encryption" class="form-select model-select" aria-label="Select Model">
                        <option value="" disabled selected>Select Encryption Type</option>
                        <option value="ssl">SSL</option>
                        <option value="tls">TLS</option>
                        <option value="starttls">STARTTLS</option>
                        <option value="ssltls">SSL/TLS</option>
                        <option value="none">None</option>
                      </select>
                      <label for="smtp-encryption">Encryption</label>
                    </div>
                  </div>

                  <div class="col-2">
                    <div class="form-floating mb-3">
                      <input type="text" class="form-control" id="smtp-port" placeholder="587"
                        v-model="config.smtp_port">
                      <label for="smtp-port">Port</label>
                    </div>
                  </div>

                </div>

                <div class="row">

                  <div class="col-6">
                    <div class="form-floating mb-3">
                      <input type="text" class="form-control" id="smtp-username" placeholder="username"
                        v-model="config.username">
                      <label for="smtp-username">Username</label>
                    </div>
                  </div>

                  <div class="col-6">
                    <div class="form-floating mb-3">
                      <input type="password" class="form-control" id="smtp-password" placeholder="password"
                        v-model="config.password">
                      <label for="smtp-password">Password</label>
                    </div>
                  </div>

                </div>

                <div class="row">

                  <div class="col-6">
                    <div class="form-floating mb-3">
                      <input type="text" class="form-control" id="sender" placeholder="sender" v-model="config.sender">
                      <label for="sender">Sender</label>
                    </div>
                  </div>

                  <div class="col-6">
                    <div class="form-floating mb-3">
                      <input type="text" class="form-control" id="reply-to" placeholder="reply-to"
                        v-model="config.reply_to">
                      <label for="reply-to">Reply To</label>
                    </div>
                  </div>

                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="row mt-3">
          <div class="col-12">
            <button type="button" class="btn btn-secondary" @click="submitForm">Save</button>
          </div>
        </div>

      </div>
    </div>
  </div>

</template>

<style scoped>
.form-control,
.form-select {
  background-color: #374151;
}
</style>
