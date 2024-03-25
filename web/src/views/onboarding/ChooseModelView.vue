<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { Form, Field, ErrorMessage } from 'vee-validate';

const router = useRouter();

const apiToken = ref('');
const tokenRequired = ref(false);
const skipToken = ref(false);

const selectedModel = ref('');
const chooseModel = (model) => {
  // Retrieve existing onboarding data
  let onboardingData = JSON.parse(localStorage.getItem('onboarding_data')) || {};

  // if the selected model is already chosen, unselect it
  if (selectedModel.value === model) {
    if (model === 'gpt') {
      tokenRequired.value = false;
    }
    
    selectedModel.value = '';
    delete onboardingData['avatar_llm'];
  } else { // else select the model
    selectedModel.value = model;
    onboardingData['avatar_llm'] = model;
    
    if (model === 'gpt') {
      tokenRequired.value = true;
    }  
  }

  // Store the updated onboarding data
  localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
}

const skip = () => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  
  tokenRequired.value = false;
  skipToken.value = true;
  
  onboardingData['skip_token'] = skipToken.value;
  localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
  
}

const next = () => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  onboardingData['api_token'] = apiToken.value;
  localStorage.setItem('onboarding_data', JSON.stringify(onboardingData));
  router.push({ name: 'select-agents' });
}
onMounted(() => {
  const onboardingData = JSON.parse(localStorage.getItem('onboarding_data'));
  selectedModel.value = onboardingData?.avatar_llm || '';
  skipToken.value = onboardingData?.skip_token || false;

  if (!skipToken.value) {
    selectedModel.value === 'gpt' && (tokenRequired.value = true);
}
  
  feather.replace();
})
</script>


<template>
  <div class="container">
    <div class="row">
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle mb-5"><h2 class="circle-text"><i data-feather="check"style="width: 36px; height: 36px;"></i></h2></div>
        <h3 class="mb-3">Create Avatar</h3>
        <p class="lead mb-5" style="color: white">Give your AI avatar a name and describe how it should behave.</p>
      </div>
      <div class="col-4 d-flex flex-column align-items-center text-center">
        <div class="circle current mb-5"><h2 class="circle-text">2</h2></div>
        <h3 class="mb-3">Choose models</h3>
        <p class="lead mb-5" style="color: white">Select one or more LLM models yor avatar will be using.</p>
      </div>
      <div class="col-4 text-center d-flex flex-column align-items-center">
        <div class="circle mb-5"><h2 class="circle-text">3</h2></div>
        <h3 class="mb-3">Select Agents</h3>
        <p class="lead mb-5">Your AI Avatar can browse the internet, answer emails, post on social media nad more!</p>
      </div>
    </div>
    <div class="row">

      <div class="col-3">
        <div class="card llm-card" @click="chooseModel('gpt')">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
            <div class="card-checkmark d-flex flex-column justify-content-center align-items-center" v-if="selectedModel === 'gpt'">
              <i class="fas fa-check"></i>
            </div>
            
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>GPT</h2>
            </div>
            
          </div>
        </div>
      </div>

      <div class="col-3">
        <div class="card llm-card" data-bs-toggle="tooltip" data-bs-placement="top" title="Coming Soon">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>LLama 2</h2>
            </div>
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>
      
      <div class="col-3">
        <div class="card llm-card" data-bs-toggle="tooltip" data-bs-placement="top" title="Coming Soon">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>Mistral</h2>
            </div>
            
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>

      <div class="col-3">
        <div class="card llm-card" data-bs-toggle="tooltip" data-bs-placement="top" title="Coming Soon">
          <div class="card-body d-flex flex-column justify-content-center align-items-center">
           
            <div class="card-icon">
              <i class="fas fa-puzzle-piece fa-3x"></i>
            </div>

            <div class="card-text mt-3">
              <h2>Falcon</h2>
            </div>
            
            <div class="card-overlay position-absolute"></div>
          </div>
        </div>
      </div>
    </div>

    <div class="row mt-3" v-if="tokenRequired">
      <div class="col-8 offset-2">  
        <div class="card">
          <div class="card-body">
            <div class="row">
              <div class="col-1">
                <i class="fas fa-gear fa-2x float-end"></i>
              </div>
              <div class="col-10">
                <h2>OpenAI API Token required.</h2>
                <p class="lead">GPT family of models is provided via OpenAI API and is a paid service, to aquire your API token plese visit OpenAI playground <a href="https://platform.openai.com" target="_blank">here</a>.</p>
              </div>
            </div>

            <div class="row">
              <div class="col-10 offset-1">

                  <div class="form-floating mb-1">
                    <Field 
                      v-model="apiToken"
                      name="ApiToken"
                      type="text"
                      class="form-control mb-2"
                      id="floatingInput"
                      placeholder="Provide API Token..."  
                    />
                    <label for="floatingInput">Provide API Token</label>
                  </div>
              </div>
            </div>

            <div class="row mt-1">
              <div class="col-10 offset-1">
                <p class="text-info fs-4">You can provide this token later. <a href="#" @click="skip">Skip for now</a></p>
              </div>
            </div>
            
          </div>
        </div>
      </div>
    </div>

    <div class="row" style="margin-top: 80px;"> 
      <div class="col-4 offset-2 text-center d-grid">
        <router-link :to="{ name: 'create-avatar' }" class="btn btn-primary btn-lg btn-back">Back</router-link>
      </div>
      <div class="col-4 text-center d-grid">
        <button type="submit" class="btn btn-primary btn-lg" @click="next">Next</button>
        <!-- <router-link :to="{ name: 'select-agents' }" class="btn btn-primary btn-lg">Next</router-link> -->
      </div>
    </div>

  </div>
</template>

<style scoped>
h1, h3 {
  font-weight: 700;
}

.btn {
  background-color: #1c64f2;
  border-color: #1c64f2;
}

.form-control  {
  background-color: #374151;
}

.btn-back {
  background-color: transparent;
}

.subtitle, .lead {
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
   line-height: 50px; /* match this with the height of .circle */
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

.llm-card {
  border: 2px solid transparent;
}

.llm-card:hover  {
  border: 2px solid #1c64f2;
  cursor: pointer;
}

.card-overlay {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0,0,0,0.5);
  z-index: 10;
}

.card-checkmark {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;

  /* Add these lines for colors */
  color: white;
  background-color: #1c64f2;
  border-radius: 50%;
  width: 30px;
  height: 30px;
}
</style>
