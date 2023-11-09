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
    // messages: [
    //   {
    //       "sender": "human",
    //       "content": "My vue routing is not working when hosted from a golang server \n```"
    //   },
    //   {
    //       "sender": "ai",
    //       "content": "    Hosting a Vue app from a Go server is a common deployment plan, but one frequent issue is that the Go server may not correctly serve the Vue app's routes. \n\nGo's default behavior is to try and match the exact path/URL to its routing table. However, a Vue app typically uses Vue Router with HTML5 history mode enabled, which means that the URL paths are managed on the client-side by Vue and the browser, not on the server-side. As a result, the Go server might not understand how to handle these paths.\n\nTo resolve this, we have to make sure that all routes (except for the specific API routes or other server-side routes that Go handles) fall back to the Vue app's `index.html` file. This is because, in a Vue app, the `index.html` file is the entry point for the Vue Router. \n\nHere's a simple example of how your Go server setup might look to accommodate the Vue Router:\n\n```go\npackage main\n\nimport (\n\t\"net/http\"\n\t\"os\"\n\t\"path/filepath\"\n)\n\nfunc main() {\n\thttp.Handle(\"/\", fallbackToIndex(http.FileServer(http.Dir(\"path/to/your/vue/dist\"))))\n\n\thttp.ListenAndServe(\":8080\", nil)\n}\n\nfunc fallbackToIndex(h http.Handler) http.HandlerFunc {\n\treturn func(w http.ResponseWriter, r *http.Request) {\n\t\t_, err := os.Stat(filepath.Join(\"path/to/your/vue/dist\", r.URL.Path))\n\t\t\n\t\tif os.IsNotExist(err) {\n\t\t\t// if the requested url does not match a file, fallback to index.html\n\t\t\thttp.ServeFile(w, r, \"path/to/your/vue/dist/index.html\")\n\t\t} else {\n\t\t\t// else, handle the request as usual\n\t\t\th.ServeHTTP(w, r)\n\t\t}\n\t}\n}\n```\n\nReplace `\"path/to/your/vue/dist\"` with the path to your Vue app's built dist directory, the output of `npm run build`.\n\nThis setup serves the Vue app and, in the case of a URL not matching any file in the Vue `dist` directory (which means Vue Router should handle the path), it serves the `index.html` file.\n\nI would suggest trying this setup in your Go server code and see if your Vue routing starts to work correctly. However, I recommend you thoroughly test your application afterward to verify if any other issues might pop up due to this change. Let me know if you need any more help!"
    //   },
    //   {
    //       "sender": "human",
    //       "content": "ok here's my setup:\n```\nimport { createRouter, createWebHistory } from 'vue-router'\nimport HomeView from '../views/HomeView.vue'\nimport AsaiChatView from '../views/AsaiChatView.vue'\n\nconst router = createRouter({\n  history: createWebHistory(import.meta.env.BASE_URL),\n  routes: [\n    {\n      path: '/',\n      name: 'asai',\n      component: AsaiChatView\n    },\n    {\n      path: '/home',\n      name: 'home',\n      component: HomeView\n    },\n```\n\n```\n\nfunc NewAsaiServer() *AsaiServer {\n\tvar err error\n\tserver := &AsaiServer{}\n\n\tserver.asaiChain, err = chains.NewAsaiChain()\n\tif err != nil {\n\t\tfmt.Println(\"Failed to create AsaiChain:\", err)\n\t}\n\n\ttoken := os.Getenv(\"DISCORD_API_KEY\")\n\tserver.discordClient, err = discordgo.New(\"Bot \" + token)\n\tif err != nil {\n\t\tfmt.Println(\"Failed to create Discord session:\", err)\n\t}\n\n\tserver.wsManager = ws.NewManager(context.Background())\n\n\treturn server\n}\n\nfunc (server *AsaiServer) Run() error {\n\t// Initialize the Discord client\n\tserver.discordClient.AddHandler(controllers.DiscordMsgHandler)\n\tserver.discordClient.Identify.Intents = discordgo.IntentsGuildMessages\n\n\terr := server.discordClient.Open()\n\tif err != nil {\n\t\tfmt.Println(\"Failed to open Discord connection:\", err)\n\t\treturn err\n\t}\n\n\trouter := rest.NewRouter()\n\trouter.Mux.StrictSlash(true)\n\n\t// Web client server\n\tstatic := http.FileServer(http.Dir(\"./web/static\"))\n\tassets := http.FileServer(http.Dir(\"./web/static/assets\"))\n\n\trouter.Mux.PathPrefix(\"/assets/\").Handler(http.StripPrefix(\"/assets/\", assets))\n\trouter.Mux.Handle(\"/\", static)\n\n\t// API server\n\tctrl := rest.NewController(router)\n\tctrl.Get(\"/api/chat/history/{session_id}\", controllers.GetHistory)\n\tctrl.Post(\"/api/chat/msg\", controllers.PostHandler)\n\tctrl.Get(\"/api/users/session\", controllers.GetSession)\n\n\t// Websocket server\n\trouter.Mux.HandleFunc(\"/api/chat/socket\", server.wsManager.Handler)\n\n\tport := os.Getenv(\"PORT\")\n\tif port == \"\" {\n\t\terr := router.Listen(\":8082\")\n\t\tif err != nil {\n\t\t\tfmt.Println(\"Failed to listen:\", err)\n\t\t\treturn err\n\t\t}\n\t\treturn nil\n\t}\n\n\tlistener, err := net.Listen(\"tcp\", \":\"+port)\n\tif err != nil {\n\t\tfmt.Println(\"Failed to listen:\", err)\n\t\treturn err\n\t}\n\n\t// Start the HTTP server using the router and the listener\n\terr = http.Serve(listener, router.Mux)\n\tif err != nil {\n\t\tfmt.Println(\"Failed to serve:\", err)\n\t\treturn err\n\t}\n\n\t// Setup signal capturing\n\tstop := make(chan os.Signal, 1)\n\tsignal.Notify(stop, syscall.SIGTERM)\n\n\t// Wait for SIGTERM signal\n\t<-stop\n\n\t// Cleanly close down the Discord session.\n\tserver.discordClient.Close()\n\n\treturn nil\n\n}\n```"
    //   },
    //   {
    //       "sender": "ai",
    //       "content": "     That's a good point. With this configuration, static files that are not available under the '/assets/' prefix may not be served correctly.\n\nOne common solution is to set a distinct route prefix for the API and anything else that should be server managed (for the purposes of the demonstration, let's use '/api'). Then use a wildcard (`*`) route to catch all other requests and serve your `index.html` file at that route. However, within this catch-all route, you would first need to check if the requested file exists and in that case serve the file instead. That way, static assets like CSS or JS that are located outside the '/assets/' path are also served. \n\nHere's an amended version of your Go code:\n\n```go\nrouter := rest.NewRouter()\nrouter.Mux.StrictSlash(true)\n\n// Web client server\nstaticDir := \"./web/static\"\nstatic := http.FileServer(http.Dir(staticDir))\nassets := http.FileServer(http.Dir(staticDir + \"/assets\"))\n\nrouter.Mux.PathPrefix(\"/assets/\").Handler(http.StripPrefix(\"/assets/\", assets))\n\n// API server\nctrl := rest.NewController(router)\nctrl.Get(\"/api/chat/history/{session_id}\", controllers.GetHistory)\nctrl.Post(\"/api/chat/msg\", controllers.PostHandler)\nctrl.Get(\"/api/users/session\", controllers.GetSession)\n\n// Websocket server\nrouter.Mux.HandleFunc(\"/api/chat/socket\", server.wsManager.Handler)\n\n// Fallback to index.html for Vue Router\nrouter.Mux.PathPrefix(\"/\").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n    path:= filepath.Join(staticDir + r.URL.Path)\n    _, err := os.Stat(path)\n\n    if os.IsNotExist(err) {\n        http.ServeFile(w,r, staticDir + \"/index.html\")\n        return\n    }\n\n    // If request is not for a directory, serve with the static file server as normal\n    static.ServeHTTP(w, r)\n    return\n})\n```\n\nThe `os.Stat` call checks if the requested file exists. If it does not, it serves the index.html file; if it does exist, it serves the file using the static file server.\n\nGive this a try and let me know if it resolves the routing issue. Please do not hesitate to reach out for more assistance if needed."
    //   },
    //   {
    //       "sender": "human",
    //       "content": "In my chat view I have a watch function that tracks messages and scrolls to bottom to show the newly added msg, that works, but I alos have message.content that is updated with an incoming stream of tokens via websocket from the backend. I need the scroll to keep scroling as the content is added to the message adn a new line is added. I've used `watchEffect` by simply replacing the watch function, but the problem is that the function is repeteadly for some reason and the scroll is not happening as expected. Here's my AsaiChatView:\n\n ```\n<script setup>\n\nimport PromptInput from '../components/PromptInput.vue';\nimport { ref, onMounted, watch, nextTick, watchEffect } from 'vue';\nimport { storeToRefs } from 'pinia';\nimport  MarkdownIt  from 'markdown-it';\n\nimport { useChatStore } from '../stores/chat.store.js';\nimport { useUsersStore } from '../stores/user.store.js';\n\nconst chatStore = useChatStore();\nconst { messages } = storeToRefs(chatStore);\n\nconst usersStore = useUsersStore();\nconst conversationContainer = ref(null);\n\nconst md = new MarkdownIt({\n    breaks: true\n  }\n);\n\nusersStore.getSession();\n\nasync function scrollToBottom() {\n  console.log('Scrolling to bottom');\n  requestAnimationFrame(() => {\n    if (conversationContainer.value) {\n      var promptContainerHeight = document.querySelector('.prompt-container').offsetHeight;\n      var scrollTo = conversationContainer.value.scrollHeight + 30;\n      conversationContainer.value.scrollTop = scrollTo\n    }\n  });\n}\n\nonMounted(() => {\n  chatStore.connectWebSocket();\n})\n\nwatch(messages, () => {\n  scrollToBottom();\n  nextTick(() => {\n    feather.replace();\n  });\n});\n\n</script>\n\n<template>\n  <div class=\"container border-start border-end border-white border-5 min-vh-100 d-flex flex-column\">\n    <div ref=\"conversationContainer\" class=\"conversation-container container-fluid flex-grow-1 overflow-auto\">\n      \n      <template v-if=\"messages.length > 0\">\n      \n        <div class=\"conversation-item row\" v-for=\"(message, index) in messages\" :key=\"index\">\n          <div class=\"col-12\">\n\n            <!-- <div class=\"row\">\n              <div class=\"col-12\">\n                <button class=\"msg-btn btn btn-dark btn-sm float-end me-3 mb-1\" v-if=\"message.sender === 'human'\"><i class=\"msg-btn-icon d-block\" data-feather=\"refresh-cw\"></i></button>\n                <button class=\"msg-btn btn btn-dark btn-sm float-end me-3 mb-1\" v-if=\"message.sender === 'ai'\"><i class=\"msg-btn-icon d-block\" data-feather=\"clipboard\"></i></button>\n              </div>\n            </div> -->\n\n            <div class=\"row\">\n              <div class=\"col-1 col-xs-4\">\n                <img src=\"../assets/asai-icon.png\" class=\"logo\" alt=\"Asai Icon\" v-if=\"message.sender === 'ai'\"/>\n                <img src=\"../assets/user-icon.png\" class=\"logo\" alt=\"User Icon\" v-if=\"message.sender === 'human'\"/>\n              </div>\n            \n              <div class=\"col-11 col-xs-8\">\n                <div v-if=\"message.content !== 'loader'\" class=\"message-content pe-3\" v-html=\"md.render(message.content.trim())\"></div>\n                <div v-else>\n                  <span class=\"me-3\">I'm thinking...  </span><span class=\"spinner mb-2 me-2\"><img src=\"../assets/loader.png\" alt=\"\"></span>\n                </div>\n              </div>\n            </div>\n\n          </div>\n          <hr class=\"separator opacity-100\" v-if=\"messages.length > 1 && index !== messages.length - 1\">\n        </div>\n\n      </template>\n\n    </div>\n    \n    <div class=\"prompt-container container\">\n      <hr class=\"border border-3 opacity-100\">\n      <PromptInput />\n    </div>\n  </div>\n</template>\n\n<style scoped>\n.conversation-container {\n  max-height: calc(90vh - 30px);\n  padding: 1.25rem;\n  scrollbar-width: none; /* Firefox */\n  -ms-overflow-style: none; /* IE and Edge */\n}\n.conversation-container ::-webkit-scrollbar {\n  display: none; /* Chrome, Safari */\n  padding: 1.5rem;\n}\n\n.msg-btn {\n  padding-right: 23px;\n  padding-bottom: 23px !important;\n  width: 25px;\n  height: 25px;\n}\n\n.msg-btn-icon {\n  width: 16px;\n  height: 16px;\n}\n\n.prompt-container {\n  bottom: 30px;\n  position: sticky;\n}\n\n.separator {\n  width: 95%;\n  margin: 10px auto;\n}\n\n.conversation-item img {\n  max-width: 35px;\n  max-height: 50px;\n}\n\n@keyframes rotate {\n  0%    { transform: rotate(0deg); }\n  25%   { transform: rotate(90deg); }\n  50%   { transform: rotate(180deg); }\n  75%   { transform: rotate(270deg); }\n  100%  { transform: rotate(360deg); }\n}\n\n.spinner img {\n  display: inline-block;\n  vertical-align: middle;\n  transform-origin: 50% 50%;\n  animation: rotate 0.5s linear infinite;\n  height: 18px;\n  margin-bottom: 5px;\n}\n\n@media only screen and (max-width: 600px) {\n\n  .conversation-container {\n    padding: 0.25rem;\n  }\n  .conversation-item img {\n    max-width: 25px;\n    max-height: 35px;\n    padding-top: 4px;\n  }\n}\n</style>\n```\nand the included component PromptInput\n```\n<script setup>\n// Tu ide send logika \nimport { ref, onMounted } from 'vue';\nimport { Form, Field, useForm } from 'vee-validate';\nimport { useChatStore } from '../stores/chat.store.js';\n\nconst chatStore = useChatStore();\nconst prompt = ref(\"\");\n\nlet { resetForm } = useForm();\n\nfunction resizeTextArea(event) {\n  const lineHeight = parseInt(window.getComputedStyle(event.target).getPropertyValue(\"line-height\"));\n  const currentRows = event.target.scrollHeight / lineHeight;\n  const maxRows = 10;\n\n  if (currentRows <= maxRows) {\n    event.target.style.height = 'auto';\n    event.target.style.height = event.target.scrollHeight + 'px';\n  }\n\n  // Scroll to bottom\n  var promptContainerHeight = document.querySelector('.prompt-container').offsetHeight;\n  var scrollTo = document.body.scrollHeight + promptContainerHeight;\n  window.scrollTo(0, scrollTo);\n}\n\nfunction submitPrompt(event, resetForm) {\n  if (event.shiftKey && event.key == 'Enter') {\n    event.preventDefault();\n    let cursorPos = event.target.selectionStart;\n    let textBeforeCursor = prompt.value.substring(0, cursorPos);\n    let textAfterCursor = prompt.value.substring(cursorPos);\n    prompt.value = textBeforeCursor + '\\n' + textAfterCursor;\n  }\n  else if (event.key == 'Enter') {\n    event.preventDefault();\n    if (prompt.value != '') {\n      chatStore.sendPrompt(prompt.value);\n      resetForm();\n    }\n  }\n}\n\nonMounted(() => {\n  feather.replace();\n});\n\n</script>\n\n<template>\n  <Form v-slot=\"{ resetForm }\">\n    <div class=\"textarea-container\">\n      <Field\n        v-on:input=\"resizeTextArea\"\n        @keydown.enter=\"submitPrompt($event, resetForm)\"\n        name=\"prompt\"\n        v-model=\"prompt\"\n        type=\"text\"\n        as=\"textarea\"\n        class=\"form-control\"\n        rows=\"2\"\n        placeholder=\"Send a message...\"\n        :disabled=\"chatStore.isLoading\"\n      ></Field>\n      <button class=\"send-button btn btn-light\" @click=\"submitPrompt($event, resetForm)\">\n        <i class=\"align-middle\" data-feather=\"send\"></i>\n      </button>\n    </div>\n  </Form>\n  \n</template>\n\n<style scoped>\n.textarea-container {\n  position: relative;\n}\n\n.message-input {\n  padding: 1rem 3rem 1rem 1rem;  /* make space for the button */\n}\n\n.send-button {\n  position: absolute;\n  right: 0.5rem;\n  bottom: 0.7rem;\n  background-color: black;\n  color: white;\n  border: 1px solid white;\n}\n\ntextarea {\n  width: 100%;\n  background-color: transparent;\n  color: white !important;\n}\n\ntextarea:focus {\n  background-color: transparent;\n  color: white;\n}\n\ntextarea:disabled {\n  background-color: grey;\n  color: white;\n}\n\n</style>\n```"
    //   },
    //   {
    //       "sender": "ai",
    //       "content": " It looks like you've run into some issues with your scrolling functionality when receiving a stream of tokens from a websocket and adding new lines to your chat messages.\n\nThere are a couple of things that could be causing the scroll function to behave unexpectedly. \n\nFirstly, when you call `scrollToBottom()`, it is asynchronously requesting the scroll action via `requestAnimationFrame()`, which means the scroll action might not execute immediately.\n\nSecondly, the scroll event is only triggered when `messages` change. If the websocket is adding tokens to an existing message (instead of adding a new message to `messages`), the scroll action will not get triggered because your `watch(messages, ...)` is only watching the `messages` list itself, not the content of each individual message.\n\nTo make your scroll work as expected, you might have to adjust your approach. Here's a suggestion:\n\nSince your Vue component receives a stream of tokens via websocket and adds new lines to your chat messages, you could create a ref like `const lastUpdatedMessage = ref(null)`. Now, every time you receive a token from the websocket, update `lastUpdatedMessage.value` to be the message that's being updated. \n\nNow you have a reactive dependency that changes every time a message gets new content from the websocket. You can use a `watch` or `watchEffect` to monitor `lastUpdatedMessage`, and run the scroll function when it changes:\n\n```js\nwatch(lastUpdatedMessage , () => {\n  scrollToBottom();\n});\n```\n\nReplace `messages` with `lastUpdatedMessage` in your `watch` hook from the existing code. This way, your `scrollToBottom` function is triggered not only when a new message is added to `messages`, but also when an existing message gets updated with new content from the websocket.\n\nTry making these changes and see if it resolves your issue. Please let me know if it works or if you have any further questions, as there could be other factors in your setup that might be contributing to the issue."
    //   }
    // ],
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
        this.isLoading = true;
      } catch (error) {
        console.error("Failed to send prompt:", error);
      }

    }
  }
})