import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const llmsURL = `${apiUrl}/llms`;

export const useLLMStore = defineStore({
  id: 'llm',
  state: () => ({
    records: {},
  }),
  actions: {
    async getLLMs() {
      try {
        const llms = await fetchWrapper.get(`${llmsURL}`);
        this.records = llms;
      } catch (error) {
        console.error(error);
      }
    }
  }
})