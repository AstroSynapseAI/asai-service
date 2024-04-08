import { defineStore } from "pinia";
import { useStorage } from "@vueuse/core";
import { fetchWrapper } from "../helpers/fetch-wrapper.js";
import router from "@/router/index.js";

const usersURL = `${import.meta.env.VITE_API_URL}/users`;

export const useAuthStore = defineStore({
  id: "auth",
  state: () => ({
    isLoggedIn: useStorage("isLoggedIn", false),
    apiToken: useStorage("apiToken", null),
    user: useStorage("user", {}),
  }),
  actions: {
    async login(username, password) {
      const reqBody = {
        username: username,
        password: password,
      };
      try {
        const user = await fetchWrapper.post(`${usersURL}/login`, reqBody);
        if (user) {
          this.apiToken = user.api_token || null;
          this.isLoggedIn = true;
          this.user = user;
          return true;
        }
        return false;
      } catch (error) {
        this.user = {};
        this.isLoggedIn = false;
        this.apiToken = null;
        throw error.Error;
      }
    },

    async inviteUser(username) {
      const reqBody = {
        username: username,
      };
      try {
        const user = await fetchWrapper.post(`${usersURL}/invite`, reqBody);
        return user;
      } catch (error) {
        console.error(error);
        return false;
      }
    },

    async sendRecoverPasswordLink(email) {
      const reqBody = {
        email: email,
      };
      try {
        await fetchWrapper.post(`${usersURL}/password_recovery`, reqBody);
      } catch (error) {
        throw error.Error 
      }
    },

    async validateRecoveryToken(token) {
      try {
        const user = await fetchWrapper.get(
          `${usersURL}/password_recovery/${token}`
        );
        return user
      } catch (error) {
        throw error.Error 
      }
    },

    async validateConfirmationEmail(formData) {
      try {
        const user = await fetchWrapper.post(
          `${usersURL}/confirm_email`, formData
        );
        return user
      } catch (error) {
        throw error.Error 
      }
    },

    async getInvitedUser(inviteToken) {
      try {
        const user = await fetchWrapper.get(
          `${usersURL}/invited/${inviteToken}`
        );
        return user;
      } catch (error) {
        console.error(error);
        return false;
      }
    },

    async registerInvite(formData) {
      try {
        const user = await fetchWrapper.post(`${usersURL}/register/invite`, formData);
        if (user) {
          this.apiToken = user.apiToken || null;
          this.isLoggedIn = true;
          this.user = user;
          return true;
        }
        return false;
      } catch (error) {
        console.error(error);
        this.user = {};
        this.isLoggedIn = false;
        this.apiToken = null;
        throw error.Error;
      }
    },

    logout() {
      localStorage.removeItem("user");
      localStorage.removeItem("avatar");
      localStorage.removeItem("onboarding_data");
      localStorage.removeItem("session_id");
      this.isLoggedIn = false;
      this.apiToken = null;
      router.push("/login");
    },
  },
});
