import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const agentsURL = `${apiUrl}/agents`;

export const useAgentStore = defineStore({
  id: 'agent',
  state: () => ({
    data: {},
    allAgents: [],
  }),
  actions: {
    async getAgents() {
      try {
        const agents = await fetchWrapper.get(`${agentsURL}`);
        this.allAgents = agents;
      } catch (error) {
        console.error(error);
      }
    },
  }
})