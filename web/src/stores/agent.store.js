import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const agentsURL = `${apiUrl}/agents`;

export const useAgentStore = defineStore({
  id: 'agent',
  state: () => ({
    record: {},
    records: [],
    activeAgents: [],
    activeAgent: {},
    activeTools: [],
    activeTool: {},
  }),
  actions: {
    async getAgents() {
      try {
        const agents = await fetchWrapper.get(`${agentsURL}`);
        this.records = agents;
      } catch (error) {
        console.error(error);
      }
    },

    async getAgent(agent_id) {
      try {
        const agent = await fetchWrapper.get(`${agentsURL}/${agent_id}`);
        this.record = agent;
      } catch (error) {
        console.error(error);
      }
    },
    async saveActiveAgent(formData) {
      try {
        console.log("Saving active agent...", formData)
        const activeAgent = await fetchWrapper.post(`${agentsURL}/save/active`, formData);
        this.activeAgent = activeAgent;
      } catch (error) {
        console.error(error);
      }


    },
    async getActiveTools(agent_id) {
      try {
        const tools = await fetchWrapper.get(`${agentsURL}/${agent_id}/tools`);
        this.activeTools = tools;
      } catch (error) {
        console.error(error);
      }
    },
  }
})