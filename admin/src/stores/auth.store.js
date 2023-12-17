import { defineStore } from 'pinia'

export const useAuthStore = defineStore({
  id: 'auth',
  state: () => ({
    currentUser: JSON.parse(localStorage.getItem('user')),
    isLogedIn: false,
    apiToken: null
  }),
  actions: {
    async login(username, password) {
      req = {
        username: username,
        password: password
      }

      try {
        const user = await fetchWrapper.post(`${usersURL}/login`, req);
        localStorage.setItem('user', JSON.stringify(user));
        if (this.user.apiToken) {
          this.apiToken = this.user.apiToken
          this.isLogedIn = true
        }
        return true
      } catch (error) {
        console.error(error);
        localStorage.removeItem('user');
        this.isLogedIn = false
        this.apiToken = null
        return false
      }
    },

    async logout() {
      localStorage.removeItem('user');
      this.isLogedIn = false
      this.apiToken = null
    }
  }
})