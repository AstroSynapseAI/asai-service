<script setup>

import PromptInput from '../components/PromptInput.vue';
import { ref, onMounted, watch, nextTick, computed } from 'vue';
import { storeToRefs } from 'pinia';
import  MarkdownIt  from 'markdown-it';

import { useChatStore } from '../stores/chat.store.js';
import { useUsersStore } from '../stores/user.store.js';

const chatStore = useChatStore();
const { messages, connectionErr } = storeToRefs(chatStore);
const usersStore = useUsersStore();
const conversationContainer = ref(null);
const promptContainer = ref(null);
const md = new MarkdownIt({breaks: true});
const lastMessageIndex = computed(() => messages.value.length - 1)
const lastMessageText = computed(() => messages.value[lastMessageIndex.value]?.content || "")

usersStore.getSession();

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

onMounted(() => {
  scrollToBottom();
  feather.replace();
  chatStore.connectWebSocket();
})

</script>

<template>
  <div class="container border-start border-end border-white border-5 min-vh-100 d-flex flex-column ">

    <div class="error-container text-white border border-1 border-white px-3" v-if="connectionErr.active">
      <div class="row h-100 align-items-center justify-content-center text-center">
        <div class="col">
          <p class="mb-0">An has occurred! Please refresh the page and try again.</p>
        </div>
        <div class="col-auto">
          <button class="btn-close btn-close-white" v-on:click="chatStore.closeError"></button>
        </div>
      </div>
    </div>

    <div ref="conversationContainer" class="conversation-container container-fluid flex-grow-1 overflow-auto">
      
      <template v-if="messages.length > 0">
      
        <div class="conversation-item row" v-for="(message, index) in messages" :key="index">
          <div class="col-12">

            <!-- <div class="row">
              <div class="col-12">
                <button class="msg-btn btn btn-dark btn-sm float-end me-3 mb-1" v-if="message.sender === 'human'">
                  <i class="msg-btn-icon d-block" data-feather="refresh-cw"></i>
                </button>
                <button class="msg-btn btn btn-dark btn-sm float-end me-3 mb-1" v-if="message.sender === 'ai'">
                  <i class="msg-btn-icon d-block" data-feather="clipboard"></i>
                </button>
              </div>
            </div> -->

            <div class="row">
              <div class="col-1 col-xs-4">
                <img src="../assets/asai-icon.png" class="logo" alt="Asai Icon" v-if="message.sender === 'ai'"/>
                <img src="../assets/user-icon.png" class="logo" alt="User Icon" v-if="message.sender === 'human'"/>
              </div>
            
              <div class="col-11 col-xs-8">

                <div v-if="message.isLoading">
                  <p><span class="me-3">I'm thinking...  </span><span class="spinner mb-2 me-2"><img v-if="!message.isAgentRunnig" src="../assets/loader.png" alt=""></span></p>
                  <p v-if="message.isAgentRunnig"><span class="me-3 fst-italic">Activating agent: {{ message.agentName }}...  </span><span class="spinner mb-2 me-2"><img src="../assets/loader.png" alt=""></span></p>
                </div>

                <div v-else class="message-content pe-3" v-html="md.render(message.content.trim())"></div>

              </div>
            </div>

          </div>
          <hr class="separator opacity-100" v-if="messages.length > 1 && index !== messages.length - 1">
        </div>

      </template>

    </div>
    
    <div ref="promptContainer" class="prompt-container container">
      <hr class="border border-3 opacity-100">
      <PromptInput />
    </div>
  </div>
</template>

<style scoped>
.error-container {
  position: fixed;
  top: 40px;
  left: 50%;
  transform: translateX(-50%);
  width: 600px;
  height: 60px;
  background-color: black;
}

.conversation-container {
  max-height: calc(90vh - 30px);
  padding: 1.25rem;
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE and Edge */
}
.conversation-container ::-webkit-scrollbar {
  display: none; /* Chrome, Safari */
  padding: 1.5rem;
}

.msg-btn {
  padding-right: 23px;
  padding-bottom: 23px !important;
  width: 25px;
  height: 25px;
}

.msg-btn-icon {
  width: 16px;
  height: 16px;
}

.prompt-container {
  background-color: black;
  bottom: 30px;
  position: sticky;
}

.separator {
  width: 95%;
  margin: 10px auto;
}

.conversation-item img {
  max-width: 35px;
  max-height: 50px;
}

@keyframes rotate {
  0%    { transform: rotate(0deg); }
  25%   { transform: rotate(90deg); }
  50%   { transform: rotate(180deg); }
  75%   { transform: rotate(270deg); }
  100%  { transform: rotate(360deg); }
}

.spinner img {
  display: inline-block;
  vertical-align: middle;
  transform-origin: 50% 50%;
  animation: rotate 0.5s linear infinite;
  height: 18px;
  margin-bottom: 5px;
}

@media only screen and (max-width: 600px) {

  .conversation-container {
    padding: 0.25rem;
  }
  .conversation-item img {
    max-width: 25px;
    max-height: 35px;
    padding-top: 4px;
  }
}
</style>
