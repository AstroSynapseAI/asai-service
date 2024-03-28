<script setup>
import { onMounted, toRef, ref } from 'vue';
import { useLLMStore } from '@/stores/llm.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useUserStore } from '@/stores/user.store';
import { useToast } from 'vue-toastification';

const toast = useToast();

const user = useUserStore();
const avatar = useAvatarStore();
const activeModels = toRef(avatar, 'activeLLMs');
const selectedModels = ref([]);

const llm = useLLMStore();
const llmRecords = toRef(llm, 'records');

const isModelSelected = (model) => {
  return selectedModels.value.includes(model);
}

const isModelConfigured = (model) => {
  var modelConfigured = false
  for (const activeModel of activeModels.value) {
    if (mode === 'gpt' && activeModel.llm.provider === 'OpenAI') {
      modelConfigured = activeModel.token ? true : false
    }
  }

  return modelConfigured
}

const selectModel = (model) => {
  if (!isModelConfigured(model)) {
    toast.error('Please configure ' + model + ' first')
    return
  }

  if (isModelSelected(model)) {
    selectedModels.value = selectedModels.value.filter(m => m !== model);
  } else {
    selectedModels.value.push(model);
  }
}

const isActive = (ID) => {
  const activeModel = activeModels.value.find(activeModel => {
    return activeModel.llm.ID == ID;
  });

  return activeModel ? activeModel.is_active : false;
}

const toggleActive = async (ID) => {
  const activeModel = activeModels.value.find(activeModel => {
    return activeModel.llm.ID == ID;
  });

  if (activeModel) {
    activeModel.is_active = !activeModel.is_active;
  }

  const formData = {
    is_active: activeModel ? activeModel.is_active : false,
    avatar_id: user.avatar.ID
  }
  try {
    await llm.toggleActiveLLM(ID, formData)
  }
  catch (error) {
    console.log(error);
  }
}

const getActiveLLMID = (llmID) => {
  const activeLLM = activeModels.value.find(m => m.llm.ID === llmID);
  return activeLLM ? activeLLM.ID : null;
}

onMounted(async () => {
  try {
    await llm.getLLMs();
    await avatar.getActiveLLMs(user.avatar.ID);
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
});
</script>

<template>

  <div class="contanier">
    <div class="row">
      <div class="col-3 card llm-card">

        <div class="card-body">
          <div class="card-checkmark d-flex flex-column justify-content-center align-items-center"
            v-if="isModelSelected('gpt')">
            <i class="fas fa-check"></i>
          </div>

          <div class="row">
            <div class="col-12 d-flex flex-column justify-content-center align-items-center">
              <!-- <div class="card-icon" @click="selectModel('gpt')"> -->
              <div class="card-icon">
                <i class="fas fa-puzzle-piece fa-3x"></i>
              </div>
              <div class="card-text mt-3">
                <h2>GPT</h2>
              </div>
            </div>
          </div>

          <div class="row mt-3">
            <div class="col-12 text-center d-grid">
              <router-link :to="{ name: 'models-config', params: { provider: 'openai' } }"
                class=" btn btn-primary btn-lg">
                Configure
              </router-link>
            </div>
          </div>

        </div>
      </div>

    </div>
  </div>

  <!-- <div class="container-fluid p-0"> -->
  <!--   <h1 class="h3 mb-3">Models</h1> -->
  <!---->
  <!--   <div class="card" v-for="(llm, index) in llmRecords" :key="'row' + index"> -->
  <!--     <div class="card-header"> -->
  <!--       <div class="row"> -->
  <!--         <div class="col"> -->
  <!--           <h5 class="card-title">{{ llm.name }}</h5> -->
  <!--         </div> -->
  <!--         <div class="col-auto"> -->
  <!--           <div class="form-check form-switch d-flex align-items-center" v-if="getActiveLLMID(llm.ID)"> -->
  <!--             <input class="form-check-input me-2" type="checkbox" id="flexSwitchCheckDefault" -->
  <!--               :checked="isActive(llm.ID)" @click="toggleActive(llm.ID)"> -->
  <!--             <label style="margin-bottom: -5px;" for="flexSwitchCheckDefault">Active</label> -->
  <!--           </div> -->
  <!--         </div> -->
  <!--       </div> -->
  <!--     </div> -->
  <!---->
  <!--     <div class="card-body"> -->
  <!--       <p>{{ llm.description }}</p> -->
  <!--       <div> -->
  <!--         <router-link -->
  <!--           :to="{ name: 'model-config', params: { avatar_id: user.avatar.ID, model_id: llm.ID, active_model_id: getActiveLLMID(llm.ID) } }" -->
  <!--           class="btn  -->
  <!--       btn-primary"> -->
  <!--           Configure -->
  <!--         </router-link> -->
  <!--       </div> -->
  <!--     </div> -->
  <!--   </div> -->
  <!---->
  <!-- </div> -->

</template>

<style scoped>
.card-icon {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 85px;
  height: 85px;
  background-color: #374151;
  border-radius: 5px;
}

/* .card-icon:hover { */
/*   border: 2px solid #1c64f2; */
/*   cursor: pointer; */
/* } */

.llm-card {
  border: 2px solid transparent;
}

/* .llm-card:hover { */
/*   border: 2px solid #1c64f2; */
/*   cursor: pointer; */
/* } */

.card-overlay {
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  background-color: rgba(0, 0, 0, 0.5);
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
