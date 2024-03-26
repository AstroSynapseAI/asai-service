<script setup>
import { onMounted, ref, toRef } from 'vue';
import { useLLMStore } from '@/stores/llm.store';
import { useAgentStore } from '@/stores/agent.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';
import { useToast } from 'vue-toastification';

const user = useUserStore();
const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

const agent = useAgentStore();
const agentRecords = toRef(agent, 'records');


const avatar = useAvatarStore()

var avatarCreated = ref(false);
var agentsCreated = ref(false);

const selectedLLM = ref('');
const selectedAgent = ref('');


const createAvatar = async () => {
  let onbaordingData = JSON.parse(localStorage.getItem('onboarding_data'));
  if (onbaordingData) {
    try {
      await llm.getLLMs();

      console.log(llmRecords.value);
      console.log(onbaordingData);
      if (onbaordingData.avatar_llm == 'gpt') {
        // if GPT family of models is selected set GPT-4 as default
        console.log('setting GPT-4 as default LLM');
        selectedLLM.value = llmRecords.value.find(llm => llm.slug === 'gpt-4');
        console.log(selectedLLM.value);
      }
      else {
        console.log('missing LLM');
      }

      const avatarData = {
        "user_id": user.current.ID,
        "name": onbaordingData.avatar_name,
        "primer": onbaordingData.avatar_primer,
        "llm": selectedLLM.value.ID,
      }

      await avatar.saveAvatar(avatarData); 

      avatarCreated.value = true;
      
      await agent.getAgents();

      for (let i = 0; i < onbaordingData.avatar_agents.length; i++) {
        selectedAgent.value = agentRecords.value.find(agent => agent.slug === onbaordingData.avatar_agents[i]);
        let agent_data = {
          "is_active": true,
          "is_public": false,
          "primer": selectedAgent.value.primer,
          "llm_id": selectedLLM.value.ID,
          "avatar_id": avatar.userAvatar.ID,
          "agent_id": selectedAgent.value.ID
        }
        await agent.saveActiveAgent(agent_data);
      }

      agentsCreated.value = true;
      localStorage.removeItem('onboarding_data');
    } catch (error) {
      console.error(error);
    }
  }
}

onMounted( async () => {
  try {
    await createAvatar();
  }
  catch (error) {
    console.error(error);
  }
  feather.replace();
})
</script>


<template>
  <div class="container">
    <div class="row mb-5">
      <div class="col-8 text-center offset-2">
        <h1 class="display-6 mb-3">Your Avatar is getting ready</h1>
        <h2 class="subtitle mb-5">Your agents are being created and your avatar will be ready in a few moments.</h2>
      </div>
    </div>
    <div class="row">
      <div class="col-12 text-center">
        <img src="@/assets/avatar-rdy.svg" class="img-fluid mb-5" alt="Avatar created">
      </div>
    </div>

    <div class="row">
      <div class="col-4 offset-5">
        
        <div class="row mb-3" if="avatarCreated">
          <div class="col-1">
            <div class="checkmark-icon d-flex align-items-center justify-content-center">
              <i data-feather="check" style="width: 24px; height: 24px; color: #1c64f2"></i>
            </div>
          </div>
          <div class="col-8">
            <p class="ml-2">Your AI Avatar is created</p>
          </div>
        </div>

        <div class="row mb-3" if="agentsCreated">
          <div class="col-1">
            <div class="checkmark-icon d-flex align-items-center justify-content-center">
              <i data-feather="check" style="width: 24px; height: 24px; color: #1c64f2"></i>
            </div>
          </div>
          <div class="col-8">
            <p class="ml-2">Agents are created</p>
          </div>
        </div>
      </div>
    </div>
    <div class="row" style="margin-top: 80px;" if="agentsCreated && avatarCreated">
      <div class="col-4 offset-2 text-center d-grid">
        <router-link :to="{name: 'select-agents'}" class="btn btn-primary btn-lg btn-back">Back</router-link>
      </div>
      <div class="col-4 text-center d-grid">
        <router-link :to="{name: 'admin', params: { avatar_id: user.avatar.ID }}" class="btn btn-primary btn-lg">Continue</router-link>
      </div>
    </div>
  </div>
</template>

<style scoped>
.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.btn-back {
  background-color: transparent;
}

.checkmark-icon {
  background-color: white;
  border-radius: 50%;
  padding: 5px;
  height: 24px;
  width: 24px;
  display: flex; 
  align-items: center;
  justify-content: center;
}

</style>
