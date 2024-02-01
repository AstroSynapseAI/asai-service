import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const toolsURL = `${apiUrl}/tools`;

export const useToolStore = defineStore({
  id: 'tool',
  state: () => ({
    records: {},
    record: {},
  }),
  actions: {
    async getTools() {
      try {
        const tools = await fetchWrapper.get(`${toolsURL}`);
        this.records = tools;
      } catch (error) {
        console.error(error);
      }
    },
    async getTool(tool_id) {
      try {
        const tool = await fetchWrapper.get(`${toolsURL}/${tool_id}`);
        this.record = tool;
      } catch (error) {
        console.error(error);
      }
    }, 

    async saveAvatarTool(formData) {
      try {
        const tool = await fetchWrapper.post(`${toolsURL}/save/avatar`, formData);
        this.record = tool;
      } catch (error) {
        console.error(error);
      }
    }, 

    async saveAgentTool(formData) {
      try {
        const tool = await fetchWrapper.post(`${toolsURL}/save/agent`, formData);
        this.record = tool;
      } catch (error) {
        console.error(error);
      }
    }, 

    async toggleAvatarTool(ID, formData) {
      try {
        await fetchWrapper.post(`${toolsURL}/${ID}/toggle/avatar`, formData);
      } catch (error) {
        console.error(error);
      }
    }, 

    async toggleAgentTool(ID, formData) {
      try {
        await fetchWrapper.post(`${toolsURL}/${ID}/toggle/agent`, formData);
      } catch (error) {
        console.error(error);
      }
    }
  }
})