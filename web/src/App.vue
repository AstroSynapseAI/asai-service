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
  // await nextTick()
  // if (conversationContainer.value) {
  //   conversationContainer.value.scrollTop = conversationContainer.value.scrollHeight;
  // }
  requestAnimationFrame(() => {
    if (conversationContainer.value) {
      conversationContainer.value.scrollTop = conversationContainer.value.scrollHeight;
    }
  });
}

// onMounted(() => {
//   console.log('App mounted');
//   scrollToBottom();
// });

watch(messages, () => {
  console.log('Messages changed');
  scrollToBottom();
});


</script>

<template>
  <div class="container p-4 border-start border-end border-white border-5 min-vh-100 d-flex flex-column">
    <div ref="conversationContainer" class="conversation-container flex-grow-1 overflow-auto">
      
      <template v-if="messages.length > 0">
      
        <div class="conversation-item row" v-for="(message, index) in messages" :key="index">
        
          <div class="col-1">
            <img src="./assets/asai-icon.png" class="logo" alt="Asai Icon" width="35" height="50" v-if="message.sender === 'ai'"/>
            <img src="./assets/user-icon.png" class="logo" alt="User Icon" width="35" height="50" v-if="message.sender === 'human'"/>
          </div>
        
          <div class="col-11">  
            <div class="message-content" v-html="md.render(message.content.trim())"></div>
          </div>
          
          <hr class="separator opacity-100" v-if="messages.length > 1 && index !== messages.length - 1">

        </div>

      </template>

    </div>
    
    <div class="prompt-container">
      <hr class="border border-3 opacity-100">
      <PromptInput />
    </div>
  </div>
</template>

<style scoped>
.conversation-container {
  max-height: 90vh;
}
.separator {
  width: 95%;
  margin: 10px auto;
}
</style>
