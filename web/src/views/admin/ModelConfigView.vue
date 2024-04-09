<script setup>
import { useRoute, useRouter } from 'vue-router';
import { onMounted, ref, toRef } from 'vue';
import { useLLMStore } from '@/stores/llm.store';
import { useAvatarStore } from '@/stores/avatar.store';
import { useToast } from 'vue-toastification';
import { Form, Field, ErrorMessage } from 'vee-validate';
import * as yup from 'yup';

const schema = yup.object({
  ApiToken: yup.string().required(),
});

const toast = useToast();

const route = useRoute();
const router = useRouter();
const llm = useLLMStore();
const avatar = useAvatarStore();

const modelName = ref('');
const modelToken = ref('');
const llmRecords = toRef(llm, 'records');
const activeLLMs = toRef(avatar, 'activeLLMs');
const selectedModels = ref([]);
const filteredLLMs = ref();

const llms = () => {
  const provider = route.params.provider.toLowerCase();
  filteredLLMs.value = llmRecords.value.filter(llm => llm.provider.toLowerCase() === provider);

  return filteredLLMs.value
}


const isActive = (ID) => {
  const activeLLM = activeLLMs.value.find(activeLLM => {
    return activeLLM.llm.ID === ID;
  })
  return activeLLM ? activeLLM.is_active : false
}

const isSelected = (ID) => {
  return selectedModels.value.includes(ID);
}

const toggleActive = async (ID) => {
  if (selectedModels.value.includes(ID)) {
    selectedModels.value = selectedModels.value.filter(m => m !== ID);
  } else {
    selectedModels.value.push(ID);
  }
}

const saveModel = async () => {
  try {
    for (const model of llms()) {
      const activeLLM = activeLLMs.value.find(activeLLM => {
        return activeLLM.llm.ID === model.ID;
      })

      await llm.saveLLM({
        ID: activeLLM ? activeLLM.ID : undefined,
        token: modelToken.value,
        avatar_id: parseInt(route.params.avatar_id),
        llm_id: model.ID,
        is_active: isSelected(model.ID)
      })
    }

    await avatar.getActiveLLMs(route.params.avatar_id);
    toast.success('Model configuration saved');
  }
  catch (error) {
    console.log(error);
    toast.error('Error while saving model configuration');
  }
}

onMounted(async () => {
  if (route.params.provider === 'openai') {
    modelName.value = 'GPT by OpenAI'
  }

  if (route.params.provider === 'mistral') {
    modelName.value = 'Mistral by MistralAI'
  }

  try {
    await llm.getLLMs();
    for (const model of llms()) {
      if (isActive(model.ID)) {
        selectedModels.value.push(model.ID);
      }
    }

    await avatar.getActiveLLMs(route.params.avatar_id);
    for (const activeLLM of activeLLMs.value) {
      if (activeLLM.llm.provider.toLowerCase() === route.params.provider && activeLLM.token) {
        modelToken.value = activeLLM.token;
        break
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
    <h1 class="h3 mb-3">Configure: {{ modelName }}</h1>

    <div class="row">
      <div class="col-12 card">
        <div class="card-body">
          <div class="container">
            <div class="row">

              <div class="col-6">
                <div class="form-floating mb-3">
                  <input v-model="modelToken" type="text" name="ApiToken" class="form-control" id="floatingInput"
                    placeholder="Token...">
                  <ErrorMessage name="ApiToken" class="text-danger" />
                  <label for="floatingInput">Token*</label>
                </div>
              </div>

              <!-- This list needs to be properly loaded by checking with openai or other llm providers to see what llms the token has available -->
              <div class="col-6">
                <div class="row" v-for="(llm, index) in filteredLLMs">
                  <div class="col-4">
                    <p>{{ llm.name }}</p>
                  </div>
                  <div class="col-4">
                    <div class="form-check form-switch float-end me-5">
                      <label class="form-check-label" :for="llm.name">Active</label>
                      <input class="form-check-input" type="checkbox" :id="llm.name" :checked="isActive(llm.ID)"
                        @click="toggleActive(llm.ID)">
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="row mt-3">
              <div class="col-12">
                <button type="button" class="btn btn-primary" @click="saveModel">Save</button>
              </div>
            </div>
          </div>
        </div>

      </div>
    </div>
  </div>

</template>

<style scoped>
.form-control {
  background-color: #374151;
}
</style>
