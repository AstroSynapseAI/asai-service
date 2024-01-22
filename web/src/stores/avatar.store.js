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
  }),
  actions: {
    async saveAvatar(formData) {
      try {
        const avatar = {
          user_id: this.currentUser.ID,
          avatar_name: formData.name,
          avatar_primer: formData.primer,
          avatar_llm_id: formData.llm
        }

        if (formData.ID) {
          avatar.avatar_id = formData.ID;
        }

        const userAvatar = await fetchWrapper.post(`${avatarsURL}/save`, avatar);
        this.userAvatar = userAvatar;

      } catch (error) {
        console.error(error);
      }
    }, 

    async getActiveAgents(avatar_id) {
      try {
        const agents = await fetchWrapper.get(`${avatarsURLl}/${avatar_id}/agents`);
        this.activeAgents = agents;
      } catch (error) {
        console.error(error);
      }
    }, 
    
    async getActiveAgent(agent_id, avatar_id) {
      try {
        const agent = await fetchWrapper.get(`${avatarsURLl}/${avatar_id}/agents/${agent_id}`);
        this.activeAgent = agent;
      } catch (error) {
        console.error(error);
      }
    }
  }
})