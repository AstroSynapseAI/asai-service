import { defineStore } from 'pinia';

import { fetchWrapper }  from '../helpers/fetch-wrapper.js';
import { useUsersStore } from './user.store.js';

const chatURL = `${import.meta.env.API_URL}/chat`;

export const useChatStore = defineStore({
  id: 'chat',
  state: () => ({
    messages: [],
    user: useUsersStore().user
  }),
  actions: {
    loadHistory() {
      console.log("Loading history...")
      
      // const session_id = this.user.session_id;
      // try {
      //   this.messages = fetchWrapper.get(`${chatURL}/history/${session_id}`);
      // } catch (error) {
      //   console.error(error);
      // } 
      

      this.messages = [
        {
          sender: "Asai",
          content: "I'm Asai, the AI Avatar of Astro Synapse. Embarking on a journey through the realm of AI research and integration is more engaging with me by your side. I'm here to answer your questions and guide you through our groundbreaking solutions. Welcome to our exploration of AI!"
        },
        {
          sender: "User",
          content: "Oh, hello Asai! It's amazing being greeted by an AI Avatar. Looking forward to exploring the world of AI with Astro Synapse. I'm particularly interested in learning about AI integration in businesses. Where should I start?"
        }
      ]
    },

    async sendPrompt(content) {      
      console.log("Sending prompt...")
      var msg = {
        sender: "User",
        content: content
      }

      this.messages.push(msg); 
      
      const data = {
        session_id: this.user.session_id,
        user_prompt: content
      }
      try {
        const response =  await fetchWrapper.post(`${chatURL}/msg`, data);
        msg = {
          sender: "Asai",
          content: response.content
        }
        this.messages.push(msg);
      } catch (error) {
        console.error(error);
      }
    }
  }
})