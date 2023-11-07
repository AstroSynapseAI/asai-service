<script setup>

import PromptInput from './components/PromptInput.vue';
import { ref, onMounted, watch, nextTick } from 'vue';
import { storeToRefs } from 'pinia';
import  MarkdownIt  from 'markdown-it';

import { useChatStore } from './stores/chat.store.js';
import { useUsersStore } from './stores/user.store.js';

const chatStore = useChatStore();
const { messages } = storeToRefs(chatStore);

const usersStore = useUsersStore();
const conversationContainer = ref(null);

const md = new MarkdownIt({
    breaks: true
  }
);

usersStore.getSession();

async function scrollToBottom() {
  console.log('Scrolling to bottom');
  requestAnimationFrame(() => {
    if (conversationContainer.value) {
      conversationContainer.value.scrollTop = conversationContainer.value.scrollHeight;
    }
  });
}

onMounted(() => {
  chatStore.connectWebSocket();
})

watch(messages, () => {
  scrollToBottom();
});

</script>

<template>
  <div class="container border-start border-end border-white border-5 min-vh-100 d-flex flex-column">
    <div ref="conversationContainer" class="conversation-container container-fluid flex-grow-1 overflow-auto">
      
      <template v-if="messages.length > 0">
      
        <div class="conversation-item row" v-for="(message, index) in messages" :key="index">
        
          <div class="col-1 col-xs-4">
            <img src="./assets/asai-icon.png" class="logo" alt="Asai Icon" v-if="message.sender === 'ai'"/>
            <img src="./assets/user-icon.png" class="logo" alt="User Icon" v-if="message.sender === 'human'"/>
          </div>
        
          <div class="col-11 col-xs-8">
            <div v-if="message.content !== 'loader'" class="message-content" v-html="md.render(message.content.trim())"></div>
            <div v-else>
              <span class="spinner mb-2 me-2"></span> I'm thinking... 
            </div>
            
            <!-- <div class="message-content" v-html="md.render(message.content.trim())"></div> -->
          </div>
          
          <hr class="separator opacity-100" v-if="messages.length > 1 && index !== messages.length - 1">

        </div>

      </template>

    </div>
    
    <div class="prompt-container container">
      <hr class="border border-3 opacity-100">
      <PromptInput />
    </div>
  </div>
</template>

<style scoped>
.conversation-container {
  max-height: 90vh;
  padding: 1.25rem;
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE and Edge */
}
.conversation-container ::-webkit-scrollbar {
  display: none; /* Chrome, Safari */
  padding: 1.5rem;
}

.separator {
  width: 95%;
  margin: 10px auto;
}

.conversation-item img {
  max-width: 35px;
  max-height: 50px;
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

@keyframes rotate {
  0%    { transform: rotate(0deg); }
  25%   { transform: rotate(90deg); }
  50%   { transform: rotate(180deg); }
  75%   { transform: rotate(270deg); }
  100%  { transform: rotate(360deg); }
}
.spinner::after {
  content: '|';
  display: inline-block;
  vertical-align: middle;
  transform-origin: 50% 50%;
  animation: rotate 0.5s linear infinite;
}


</style>
