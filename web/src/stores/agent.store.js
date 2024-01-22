import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const avatarsURL = `${apiUrl}/avatars`;

export const useAgentStore = defineStore({
  id: 'agent',
  state: () => ({
    data: {},
    allAgents: [],
    activeAgents: [],
  }),
  actions: {
    async getAgents() {
      try {
        const agents = await fetchWrapper.get(`${apiUrl}/agents`);
        this.allAgents = agents;
      } catch (error) {
        console.error(error);
      }
    }
  }
})