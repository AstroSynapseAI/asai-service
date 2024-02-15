import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const avatarsURL = `${apiUrl}/avatars`;

export const useAvatarStore = defineStore({
  id: 'avatar',
  state: () => ({
    currentUser: JSON.parse(localStorage.getItem('user')),
    userAvatar: {},
    activeAgents: [],
    activeAgent: {},
    activePlugins: [],
    activePlugin: {},
    activeLLMs: [],
    activeLLM: {},
    activeTools: [],
    activeTool: {},
  }),
  actions: {
    async saveAvatar(formData) {
      try {
        const avatar = {
          user_id: this.currentUser.ID,
          avatar_name: formData.name,
          avatar_primer: formData.primer,
          avatar_llm_id: formData.llm,
          avatar_id: formData.ID
        }

        const userAvatar = await fetchWrapper.post(`${avatarsURL}/save`, avatar);
        console.log("userAvatar", userAvatar);
        localStorage.setItem('user', JSON.stringify(userAvatar));
        this.userAvatar = userAvatar;

      } catch (error) {
        console.error(error);
      }
    },

    async getActiveAgents(avatar_id) {
      try {
        const agents = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/agents`);
        this.activeAgents = agents;
      } catch (error) {
        console.error(error);
      }
    },

    async getActiveAgent(agent_id, avatar_id) {
      try {
        const agent = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/agents/${agent_id}`);
        this.activeAgent = agent;
      } catch (error) {
        console.error(error);
      }
    },

    async getActivePlugins(avatar_id) {
      try {
        const plugins = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/plugins`);
        this.activePlugins = plugins;
      } catch (error) {
        console.error(error);
      }

    },

    async getActivePlugin(avatar_id, plugin_id) {
      try {
        const plugin = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/plugins/${plugin_id}`);
        this.activePlugin = plugin;
      } catch (error) {
        console.error(error);
      }
    },

    async getActiveLLMs(avatar_id) {
      try {
        const llms = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/llms`);
        this.activeLLMs = llms;
      } catch (error) {
        console.error(error);
      }
    },

    async getActiveLLM(avatar_id, llm_id) {
      try {
        const llm = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/llms/${llm_id}`);
        this.activeLLM = llm;
      } catch (error) {
        console.error(error);
      }
    }, 

    async getActiveTools(avatar_id) {
      try {
        const tools = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/tools`);
        this.activeTools = tools;
      } catch (error) {
        console.error(error);
      }
    },

    async getActiveTool(avatar_id, tool_id) {
      try {
        const tool = await fetchWrapper.get(`${avatarsURL}/${avatar_id}/tools/${tool_id}`);
        this.activeTool = tool;
      } catch (error) {
        console.error(error);
      }
    }
  }
})