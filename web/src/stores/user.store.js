import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';
import { useChatStore } from './chat.store.js';

const apiUrl = import.meta.env.VITE_API_URL;
const usersURL = `${apiUrl}/users`;

export const useUserStore = defineStore({
  id: 'users',
  state: () => ({
    current: JSON.parse(localStorage.getItem('user')),
    avatar: JSON.parse(localStorage.getItem('avatar')),
    record: {},
    records: [],
  }),
  actions: {
    async getUsers() {

    },

    async getUser() {

    },

    async getUserAvatar(user_id) {
      try {
        const avatar = await fetchWrapper.get(`${usersURL}/${user_id}/avatars`);
        localStorage.setItem('avatar', JSON.stringify(avatar));
      } catch (error) {
        console.error(error);
      }
    },

    async getUserAccounts(user_id) {
      try {
        const accounts = await fetchWrapper.get(`${usersURL}/${user_id}/accounts`);
        user = this.user;
        user.accounts = accounts;
        localStorage.setItem('user', JSON.stringify(user));
      } catch (error) {
        console.error(error);
      }
    },

    async getUserRole() {

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