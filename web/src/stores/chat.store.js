import { defineStore } from 'pinia';

import { fetchWrapper }  from '../helpers/fetch-wrapper.js';
import { useUsersStore } from './user.store.js';

const chatURL = `${import.meta.env.VITE_API_URL}/chat`;
const wsURL = `${import.meta.env.VITE_WS_URL}/chat/socket`;

export const useChatStore = defineStore({
  id: 'chat',
  state: () => ({
    messages: [{
      sender: "ai",
      content: "Hello there... I'm Asai, How can I help you?"
    }],
    currentMsg: null,
    socket: null
  }),
  actions: {
    async connectWebSocket() {
      this.socket = new WebSocket(wsURL);

      this.socket.addEventListener('open', (event) => {
        console.log('WebSocket connected', event);
      });

      this.socket.addEventListener('message', (event) => {
        if (this.currentMsg) {
          this.currentMsg.content += event.data;
        }
          
      });

      this.socket.addEventListener('close', (event) => {
        console.log('WebSocket closed', event);
      });
    },
    async loadHistory() {
      const userStore = useUsersStore();
      const session_id = userStore.user.session_id;
      try {
        console.log("Loading history...", session_id);
        const response = await fetchWrapper.get(`${chatURL}/history/${session_id}`);

        console.log("Response:", response);
        var responseMsgs = []

        if (response) {
          for (var i = 0; i < response.length; i++) {
            var msg = {
              sender: response[i].type,
              content: response[i].text
            }
            responseMsgs.push(msg)
          }
          this.messages = responseMsgs;
          console.log("Fetched History:", response);
        }
        console.log("Messages:", this.messages);
      } catch (error) {
        console.error(error);
      } 
    },

    async sendPrompt(content) {      
      console.log("Sending prompt...")
      const userStore = useUsersStore();

      var msg = {
        sender: "human",
        content: content
      }

      this.messages = [...this.messages, msg];
      
      const data = {
        session_id: userStore.user.session_id,
        user_prompt: content
      }

      try {
        this.socket.send(JSON.stringify(data));
        console.log("Sent prompt...");
        this.currentMsg = {
          sender: "ai",
          content: ""
        };
        this.messages = [...this.messages, this.currentMsg];
      } catch (error) {
        console.error("Failed to send prompt:", error);
      }

    }
  }
})