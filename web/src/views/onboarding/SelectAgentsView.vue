<script setup>
import { onMounted, ref, toRef } from 'vue';
import { useRouter } from 'vue-router';
import { Form, Field, ErrorMessage } from 'vee-validate';

import { useLLMStore } from '@/stores/llm.store';
import { useAgentStore } from '@/stores/agent.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';
import { useToast } from 'vue-toastification';

const router = useRouter();
const toast = useToast();

const selectedAgents = ref([]);
const skipSMTP = ref(false);
const SMTPRequired = ref(false);
const agentConfig = ref({});

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

const searchConfig = ref({
  "ddg_is_active": true,
  "google_api_token": "",
  "google_is_active": false,
  "exa_api_token": "",
  "exa_is_active": false
})

const user = useUserStore();
const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

const agent = useAgentStore();
const agentRecords = toRef(agent, 'records');

const avatar = useAvatarStore()

const selectedLLM = ref('');
const selectedAgent = ref('');

const createAvatar = async () => {
  let onbaordingData = JSON.parse(localStorage.getItem('onboarding_data'));
  if (onbaordingData) {
    try {
      await llm.getLLMs();

      if (onbaordingData.avatar_llm == 'gpt') {
        // if GPT family of models is selected set GPT-4 as default
        selectedLLM.value = llmRecords.value.find(llm => llm.slug === 'gpt-4');
      }

      if (onbaordingData.avatar_llm === 'mistral') {
        selectedLLM.value = llmRecords.value.find(llm => llm.slug === 'open-mixtral-8x7b');
      }

      if (selectedLLM.value != '') {
        const avatarData = {
          "user_id": user.current.ID,
          "name": onbaordingData.avatar_name,
          "primer": onbaordingData.avatar_primer,
          "llm": selectedLLM.value.ID,
        }

        await avatar.saveAvatar(avatarData);

        const activeLLMData = {
          "token": onbaordingData.openai_token,
          "is_active": true,
          "is_public": false,
          "llm_id": selectedLLM.value.ID,
          "avatar_id": avatar.userAvatar.ID
        }

        await llm.saveLLM(activeLLMData);
      }

      await agent.getAgents();

      if (onbaordingData.avatar_agents) {
        for (let i = 0; i < onbaordingData.avatar_agents.length; i++) {
          selectedAgent.value = agentRecords.value.find(agent => agent.slug === onbaordingData.avatar_agents[i]);
          if (selectedAgent.value.slug === 'email-agent') {
            agentConfig.value = config.value;
          }

          if (selectedAgent.value.slug === 'search-agent') {
            agentConfig.value = searchConfig.value;
          }


          let agent_data = {
            "is_active": true,
            "is_public": false,
            "primer": selectedAgent.value.primer,
            "llm_id": selectedLLM.value.ID,
            "avatar_id": avatar.userAvatar.ID,
            "agent_id": selectedAgent.value.ID
          }

          if (agentConfig) {
            agent_data['config'] = JSON.stringify(agentConfig.value);
          }
          await agent.saveActiveAgent(agent_data);
        }
      }

      localStorage.removeItem('onboarding_data');
    } catch (error) {
      console.error(error);
      toast.error("Something went worng, we couldn't create your avatar");
    }
  }
}

const toggleAgent = (agent) => {
  let onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  if (selectedAgents.value.includes(agent)) {
    if (agent === 'email-agent') {
      SMTPRequired.value = false;
    }
    selectedAgents.value = selectedAgents.value.filter(item => item !== agent);
  }
  else {
    if (agent === 'email-agent') {
      SMTPRequired.value = true;
    }
    selectedAgents.value.push(agent);
  }

  onboardingData['avatar_agents'] = selectedAgents.value;
  localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
}

const agentSelected = (agent) => {
  return selectedAgents.value.includes(agent);
}

const skip = () => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  SMTPRequired.value = false;
  skipSMTP.value = true;

  onboardingData['skip_smtp'] = skipSMTP.value;
  localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
}

const next = async () => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  onboardingData['email_config'] = config.value;
  try {
    await createAvatar();
    localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
    router.push({ name: 'avatar-created' });
  } catch (error) {
    console.error(error);
  }
}

onMounted(() => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  selectedAgents.value = onboardingData?.avatar_agents || [];
  skipSMTP.value = onboardingData?.skip_smtp || false;
  config.value = onboardingData?.email_config || {};


  if (!skipSMTP.value) {
    selectedAgents.value.includes('email-agent') && (SMTPRequired.value = true);
  }
  feather.replace();
})
</script>


<template>
  <div class="container">

    <div class="row">
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5">
          <h2 class="circle-text"><i data-feather="check" style="width: 36px; height: 36px;"></i></h2>
        </div>
        <h3 class="mb-3">Create Avatar</h3>
        <p class="lead mb-5">Give your AI avatar a name and describe how it should behave.</p>
      </div>
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5">
          <h2 class="circle-text"><i data-feather="check" style="width: 36px; height: 36px;"></i></h2>
        </div>
        <h3 class="mb-3">Choose models</h3>
        <p class="lead mb-5">Select one or more LLM models yor avatar will be using.</p>
      </div>
      <div class="col-4 text-center d-flex flex-column align-items-center">
        <div class="circle current mb-5">
          <h2 class="circle-text">3</h2>
        </div>
        <h3 class="mb-3">Select Agents</h3>
        <p class="lead mb-5">Your AI Avatar can browse the internet, answer emails, post on social media nad more!</p>
      </div>
    </div>

    <div class="row">

      <div class="col-4 d-flex">
        <div class="card h-100">
          <div class="card-body">
            <h5 class="card-title">Search Agent <i data-feather="info" class="float-end"></i></h5>
            <p class="card-text">Utilizes search engines such as Google, DuckDuckGo, and Metaphor for automated web
              searches.</p>
            <div class="card-action mt-4">
              <div class="form-check form-switch">
                <input class="form-check-input float-end" type="checkbox" role="switch" id="search-agent"
                  @click="toggleAgent('search-agent')" :checked="agentSelected('search-agent')">
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="col-4 d-flex">
        <div class="card h-100">
          <div class="card-body">
            <h5 class="card-title">Email Agent <i data-feather="info" class="float-end"></i></h5>
            <p class="card-text">Email agent connects to your email server, and then comopses and sends email for you.
            </p>
            <div class="card-action mt-4">
              <div class="form-check form-switch">
                <input class="form-check-input float-end" type="checkbox" role="switch" id="search-agent "
                  @click="toggleAgent('email-agent')" :checked="agentSelected('email-agent')">
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="col-4 d-flex">
        <div class="card h-100">
          <div class="card-body">
            <h5 class="card-title">Library Agent <i data-feather="info" class="float-end"></i></h5>
            <p class="card-text">Utilizes search engines such as Google, DuckDuckGo, and Metaphor for automated web
              searches.</p>
            <div class="card-action mt-4">
              <div class="form-check form-switch">
                <input class="form-check-input float-end" type="checkbox" role="switch" id="search-agent">
              </div>
            </div>
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>

    </div>

    <div class="row mt-3" v-if="SMTPRequired">
      <div class="col-8 offset-2">
        <div class="card">
          <div class="card-body">

            <div class="row">
              <div class="col-1">
                <i class="fas fa-gear fa-2x float-end"></i>
              </div>
              <div class="col-10">
                <h2>Email agent setup.</h2>
                <p class="lead">In order to dispatch emails yor email agent needs access to your preferd SMTP server.
                </p>
              </div>
            </div>

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
                  <input type="text" class="form-control" id="smtp-port" placeholder="587" v-model="config.smtp_port">
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

            <div class="row mt-1">
              <div class="col-10 offset-1">
                <p class="text-info fs-4">You can setup your email agent later. <a href="#" @click="skip">Skip for
                    now</a></p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row" style="margin-top: 80px;">
      <div class="col-4 offset-2 text-center d-grid">
        <router-link :to="{ name: 'choose-model' }" class="btn btn-primary btn-lg btn-back">Back</router-link>
      </div>
      <div class="col-4 text-center d-grid">
        <button class="btn btn-primary btn-lg" @click="next">Next</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1,
h3 {
  font-weight: 700;
}

.form-control {
  background-color: #374151;
}

.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.btn-back {
  background-color: transparent;
}

.subtitle,
.lead {
  color: grey;
}

.circle {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 85px;
  height: 85px;
  border: 2px solid #1c64f2;
  border-radius: 50%;
}

.current {
  background-color: #1c64f2;
}

.circle-text {
  text-align: center;
  line-height: 50px;
  /* match this with the height of .circle */
  margin: 0;
}

.card-icon {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 85px;
  height: 85px;
  background-color: #374151;
  border-radius: 5px;
}

.card-overlay {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 10;
}
</style>
