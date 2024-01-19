import { defineStore } from 'pinia';
import { fetchWrapper } from '../helpers/fetch-wrapper.js';

const apiUrl = import.meta.env.VITE_API_URL;
const avatarsURL = `${apiUrl}/avatars`;

export const useAvatarStore = defineStore({
  id: 'avatar',
  state: () => ({
    currentUser: JSON.parse(localStorage.getItem('user')),
    userAvatar: {},
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
          avatar.ID = formData.ID;
        }

        const userAvatar = await fetchWrapper.post(`${avatarsURL}/save`, avatar);
        this.userAvatar = userAvatar;

      } catch (error) {
        console.error(error);
      }
    }
  }
})