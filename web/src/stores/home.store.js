import { defineStore } from 'pinia';


export const useHomeStore = defineStore({
  id: 'home',
  state: () => ({
    contents: {
      "about": {
        title: "Info",
        description: "Asai is open source and you can see the source code on GitHub.",
      },
      "open-source": {
        title: "Open Source",
        description: "Asai is open source and you can see the source code on GitHub.",
      }, 
      "transparency": {
        title: "Transparent",
        description: "Asai is an open source project, you can see the source code on GitHub.",
      },
      "customization": {
        title: "Customizable",
        description: "Asai is an open source project, you can see the source code on GitHub.",
      },
      "personalization": {
        title: "Personalized",
        description: "Asai is an open source project, you can see the source code on GitHub.",
      },
      "privacy": {
        title: "Private",
        description: "Asai is an open source project, you can see the source code on GitHub.",
      },
      "integrations": {
        title: "Integrations",
        description: "Asai is an open source project, you can see the source code on GitHub.",
      } 
    },
  }),
  actions: {

  },
})