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
    account: {},
    record: {},
    records: [],
  }),
  actions: {
    isAdmin() {
      // return this.current.roles.some(role => role.permission === 'admin');
      return true;
    },

    async getUsers() {
      try {
        const users = await fetchWrapper.get(`${usersURL}`);
        this.records = users;
      } catch (error) {
        console.error(error);
      }
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
        this.account = accounts[0];
      } catch (error) {
        console.error(error);
      }
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
    }, 

    async saveProfile(formData) {
      if (formData.ID) {
        try {
          const user = await fetchWrapper.put(`${usersURL}/${formData.ID}/update`, formData);
          localStorage.setItem('user', JSON.stringify(user));
        } catch (error) {
          console.error(error);
        }
      }

      if (formData.account_id) {
        const accountData = {
          ID: formData.account_id,
          first_name: formData.first_name,
          last_name: formData.last_name,
          email: formData.email
        } 
        try {
          const account = await fetchWrapper.put(`${apiUrl}/accounts/${formData.account_id}`, accountData);
          this.account = account;
        } catch (error) {
          console.error(error);
        }
      }
    }
  }

})