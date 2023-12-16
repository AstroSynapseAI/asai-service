import { defineStore } from 'pinia';

import { fetchWrapper } from '../helpers/fetch-wrapper.js';
import { useChatStore } from './chat.store.js';

const apiUrl = import.meta.env.VITE_API_URL;
const usersURL = `${apiUrl}/users`;

export const useUsersStore = defineStore({
  id: 'users',
  state: () => ({
    // user: JSON.parse(localStorage.getItem('user')),
    user: {
      id: 'UDID-01',
      avatar: 'asai',
    }
  }),
  actions: {
    async getUsers() {
      
    },

    async getUser() {
      
    },
    
    async getUserAvatar() {
      
    },

    async getUserAgents() {
      
    },

    async getUserPlugins() {
      
    },

    async getUserModels() {
      
    },

    async getSession() {
      console.log("Creating session...")
      const chatStore = useChatStore(); 
      this.user = JSON.parse(localStorage.getItem('user'));
      
      if (!this.user) {
        try {
          chatStore.messages = [{
              sender: "ai",
              content: "Hello there... I'm Asai, How can I help you?"
          }];
          const user = await fetchWrapper.get(`${usersURL}/session`);
          localStorage.setItem('user', JSON.stringify(user));
          this.user = user;
        } catch (error) {
          console.error(error);
        }    
      }

      if (this.user) {
        console.log("User:", this.user);
        chatStore.loadHistory();
      }
    }
  }

})