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
    aiMsg: null,
    socket: null,
    isLoading: false
  }),
  actions: {
    async connectWebSocket() {
      this.socket = new WebSocket(wsURL);

      this.socket.addEventListener('open', (event) => {
        console.log('WebSocket connected', event);
      });

      this.socket.addEventListener('message', (event) => {
        if (event.data === "[chain start]") {
          this.aiMsg = {
            sender: "ai",
            content: "loader"
          };
          this.messages = [...this.messages, this.aiMsg];
        } else if (event.data === "[chain end]") {
          this.aiMsg = null;
          this.isLoading = false;
        } else if (this.aiMsg) {
          if (this.aiMsg.content == "loader") {
            this.aiMsg.content = ""
          }
          this.aiMsg.content += event.data;
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
        }
        // console.log("Messages:", this.messages);
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
        this.isLoading = true;
      } catch (error) {
        console.error("Failed to send prompt:", error);
      }

    }
  }
})