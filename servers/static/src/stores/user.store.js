import { defineStore } from 'pinia';

import { fetchWrapper } from '../helpers/fetch-wrapper.js';
import { useChatStore } from './chat.store.js';

const usersURL = `${import.meta.env.API_URL}/users`;

export const useUsersStore = defineStore({
  id: 'users',
  state: () => ({
    user: {},
  }),
  actions: {
    async getSession() {
      console.log("Creating session...")
      const chatStore = useChatStore(); 
      this.user = JSON.parse(localStorage.getItem('user'));
      
      if (!this.user) {
        try {
          this.user = await fetchWrapper.get(`${usersURL}/session`);
          localStorage.setItem('user', JSON.stringify(this.user));
        } catch (error) {
          console.error(error);
        }    
      }

      chatStore.loadHistory();
    }  
  }
})