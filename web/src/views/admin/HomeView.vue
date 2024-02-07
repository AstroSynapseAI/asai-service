<script setup>
import PromptInput from '@/components/admin/PromptInput.vue';
import { onMounted, watch, ref, computed } from 'vue';
import { useRoute } from 'vue-router';
import { storeToRefs } from 'pinia';
import { useChatStore } from '@/stores/chat.store';
import { useUserStore } from '@/stores/user.store';
import MarkdownIt from 'markdown-it';

const route = useRoute();

const chatStore = useChatStore();
const { messages } = storeToRefs(chatStore, "messages");

const user = useUserStore();

const conversationContainer = ref(null);
const promptContainer = ref(null);

const md = new MarkdownIt({breaks: true});
const lastMessageIndex = computed(() => messages.value.length - 1)
const lastMessageText = computed(() => messages.value[lastMessageIndex.value]?.content || "")

async function scrollToBottom() {
  requestAnimationFrame(() => {
    if (conversationContainer.value) {
      var offset = promptContainer.value.offsetHeight + 30;
      var scrollTo = conversationContainer.value.scrollHeight + offset
      conversationContainer.value.scrollTop = scrollTo
    }
  });
}

watch(messages, () => {
  scrollToBottom();
});

watch(lastMessageText, () => {
  scrollToBottom();
});

onMounted(async () => {
  try {
    await user.getPrivateSession(route.params.avatar_id);
  }
  catch (error) {
    console.log(error);
  }
  feather.replace();
});

</script>

<template>  
  <div class="container-fluid p-0">
    <h1 class="h3 mb-3">Avatar</h1>
    <div class="row">
      <div class="col-12">
        <div class="card">
          <div class="card-body">
            <div class="container p-4 d-flex flex-column">
              
              <div ref="conversationContainer" class="conversation-container container-fluid flex-grow-1 overflow-auto">
                
                <template v-if="messages.length > 0">
                
                  <div class="conversation-item row" v-for="(message, index) in messages" :key="index">
                  
                    <div class="col-1">
                      <img src="../assets/avatar.png" class="logo" alt="Avatar Icon" width="35" v-if="message.type === 'ai'"/>
                      <img src="../assets/user.svg" class="logo" alt="User Icon" width="35" height="50" v-if="message.type === 'user'"/>
                    </div>
                  
                    <div class="col-11">
                      <div v-if="message.isLoading">
                        <p><span class="me-3">I'm thinking...  </span><span class="spinner mb-2 me-2"><img v-if="!message.isAgentRunnig" src="../assets/loader.png" alt=""></span></p>
                        <p v-if="message.isAgentRunnig"><span class="me-3 fst-italic">Activating agent: {{ message.agentName }}...  </span><span class="spinner mb-2 me-2"><img src="../assets/loader.png" alt=""></span></p>
                      </div>
                      
                      <div v-else class="message-content pe-3" v-html="md.render(message.content.trim())"></div>
                    </div>
                    
                    <hr class="separator opacity-100" v-if="messages.length > 1 && index !== messages.length - 1">

                  </div>

                </template>

              </div>
            
              <div ref="promptContainer" class="prompt-container">
                <hr class="border border-3 opacity-100">
                <PromptInput />
              </div>

            
            </div>
          </div>
        </div>
      </div>

      
    
    </div>
  </div>
</template>

<style scoped>
main {
  display: block !important;
}
.container-fluid {
  color: white;
}

.card-body .container {
  min-height: 85vh;
}
.card-body {
  background-color: #19232E !important;
  color: white !important;
}
</style>