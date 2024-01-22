import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const avatarsURL = `${apiUrl}/llms`;

export const useLLMStore = defineStore({
  id: 'llm',
  state: () => ({
    llms: {},
  }),
  actions: {
    async getLLMs() {
      try {
        const llms = await fetchWrapper.get(`${apiUrl}/llms`);
        this.llms = llms;
      } catch (error) {
        console.error(error);
      }
    }
  }
})