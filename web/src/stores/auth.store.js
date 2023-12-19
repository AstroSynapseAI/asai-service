import { defineStore } from 'pinia'
import { fetchWrapper }  from '../helpers/fetch-wrapper.js';

const usersURL = `${import.meta.env.VITE_API_URL}/users`;

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    currentUser: JSON.parse(localStorage.getItem('user')),
    isLogedIn: false,
    apiToken: null
  }),
  actions: {
   async login(username, password) {
      const reqBody = {
        username: username,
        password: password
      }

      try {
        const user = await fetchWrapper.post(`${usersURL}/login`, reqBody);
        if (user) {
          localStorage.setItem('user', JSON.stringify(user));
          if (user.apiToken) {
            this.apiToken = user.apiToken
          }
          this.isLogedIn = true
          return true
        }

      } catch (error) {
        console.error(error);
        localStorage.removeItem('user');
        this.isLogedIn = false
        this.apiToken = null
        return false
      }
    },

    logout() {
      localStorage.removeItem('user');
      this.isLogedIn = false
      this.apiToken = null
    }
  }
})