import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const llmsURL = `${apiUrl}/llms`;

export const useLLMStore = defineStore({
  id: 'llm',
  state: () => ({
    records: {},
    record: {},
  }),
  actions: {
    async getLLMs() {
      try {
        const llms = await fetchWrapper.get(`${llmsURL}`);
        this.records = llms;
      } catch (error) {
        console.error(error);
      }
    }, 

    async getLLM(llm_id) {
      try {
        const llm = await fetchWrapper.get(`${llmsURL}/${llm_id}`);
        this.record = llm;
      } catch (error) {
        console.error(error);
      }
    },

    async saveLLM(formData) {
      try {
        const llm = await fetchWrapper.post(`${llmsURL}/save/active`, formData);
        this.record = llm;
      } catch (error) {
        console.error(error);
      }
    }, 
    async toggleActiveLLM(ID, formData) {
      try {
        await fetchWrapper.post(`${llmsURL}/${ID}/toggle/active`, formData);
      } catch (error) {
        console.error(error);
      }
      
    }
  }
})