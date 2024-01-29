import { defineStore } from 'pinia';
import { fetchWrapper }  from '@/helpers/fetch-wrapper.js';

const pluginsURL = `${import.meta.env.VITE_API_URL}/plugins`;

export const usePluginStore = defineStore({
  id: 'plugin',
  state: () => ({
    records: [],
    record: {},
    activePlugin: {},
  }),
  actions: {
    async getPlugins() {
      try {
        const plugins = await fetchWrapper.get(`${pluginsURL}`);
        this.records = plugins;
      } catch (error) {
        console.error(error);
      }
    },
    async getPlugin(plugin_id) {
      try {
        const plugin = await fetchWrapper.get(`${pluginsURL}/${plugin_id}`);
        this.record = plugin;
      } catch (error) {
        console.error(error);
      }
    }, 
    async saveActivePlugin(formData) {
      try {
        const activePlugin = await fetchWrapper.post(`${pluginsURL}/save/active`, formData);
        this.activePlugin = activePlugin;
      } catch (error) {
        console.error(error);
      }
    }
  }
})