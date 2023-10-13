<script setup>
import PromptInput from './components/PromptInput.vue';

import { storeToRefs } from 'pinia';
import { useChatStore } from './stores/chat.store.js';
import { useUsersStore } from './stores/user.store.js';

const chatStore = useChatStore();
const { messages } = storeToRefs(chatStore);

const usersStore = useUsersStore();

usersStore.getSession();
</script>

<template>
  <div class="container p-4 border-start border-end border-white border-5 min-vh-100 d-flex flex-column">
    <div class="conversation-container flex-grow-1">
      
      <template v-if="messages.length > 0">
      
        <div class="conversation-item row" v-for="(message, index) in messages" :key="index">
        
          <div class="col-1">
            <img src="./assets/asai-icon.png" class="logo" alt="Asai Icon" width="35" height="50" v-if="message.sender === 'Asai'"/>
            <img src="./assets/user-icon.png" class="logo" alt="User Icon" width="35" height="50" v-if="message.sender === 'User'"/>
          </div>
        
          <div class="col-11">  
            <p class="message-content">{{ message.content }}</p>
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
.separator {
  width: 95%;
  margin: 10px auto;
}
</style>
