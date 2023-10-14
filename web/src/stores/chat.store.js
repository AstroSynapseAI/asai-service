import { defineStore } from 'pinia';

import { fetchWrapper }  from '../helpers/fetch-wrapper.js';
import { useUsersStore } from './user.store.js';

const chatURL = `${import.meta.env.VITE_API_URL}/chat`;

export const useChatStore = defineStore({
  id: 'chat',
  state: () => ({
    messages: [
      {
        sender: "ai",
        content: "Hello there... I'm Asai, How can I help you?"
      }
    ],
    user: useUsersStore().user
  }),
  actions: {
    async loadHistory() {
      console.log("Loading history...")
      const session_id = this.user.session_id;
      try {
        const response = await fetchWrapper.get(`${chatURL}/history/${session_id}`);
        if (response.length > 0) {
          this.messages = response
        }
        console.log("History:", this.messages);
      } catch (error) {
        console.error(error);
      } 
    },

    async sendPrompt(content) {      
      console.log("Sending prompt...")
      var msg = {
        sender: "human",
        content: content
      }

      this.messages = [...this.messages, msg];
      
      const data = {
        session_id: this.user.session_id,
        user_prompt: content
      }
      try {
        const response = await fetchWrapper.post(`${chatURL}/msg`, data);
        msg = {
          sender: "ai",
          content: response.content
        }
        this.messages = [...this.messages, msg];
      } catch (error) {
        console.error(error);
      }
    }
  }
})