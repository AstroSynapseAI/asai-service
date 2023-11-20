import { defineStore } from 'pinia';
import infoFigureLeft from '@/assets/AIAvatar-figure-left.png';
import infoFigureRight from '@/assets/AIAvatar-figure-right.png';

export const useHomeStore = defineStore({
  id: 'home',
  state: () => ({
    contents: {
      "about": {
        title: "Info",
        description:`
        <div class="container">
          <p class="mb-4 mt-3">Asai is an AI Avatar platform, an interconnected system driven by several types of AI models (LLM, A/V) that entirely replace all your interactions with software and the internet using simple text and voice commands.</p>
          <p>Whereas products around LLM’s and other models like GPTChat, Bard chat, Midjourney etc, are apps for users to add to their already large arsenal of apps, Asai is a system designed to help you easily interact with your cluster of apps, thereby simplifying your private or professional workflow.</p> 
          <div class='row'>
            <div class='col-md-6'><img width='90%' class='d-block mx-auto d-md-inline-block' src='${infoFigureLeft}' /></div>
            <div class='col-md-6'><img width='90%' class='d-block mx-auto d-md-inline-block float-md-end mt-3 mt-md-0' src='${infoFigureRight}' /></div>
          </div>
          <p class="mt-3 mt-md-5">Out of the box, your AI Avatar comes with a toolkit of autonomous agents that can perform various operations. Your avatar can search the web for you, read webpages, and the documents you provide it or give access to on GDrive. It connects to your Gmail or private Mailbox for easier email curation, and can even manage your social media identity if connected to your social media accounts.</p>
          <p>Access it via a standard web-based UI, or use plugins to integrate a chatbot to Slack, Discord, or equivalent. Through its dashboard, you can create a completely custom persona that manages your workflows and controls your agents. Choose which agents you want to use and modify their actions, select specific foundation models for your agents to use. If all of that is too overwhelming, use our simplified creation wizard.</p>
        </div>`,
      },
      "open-source": {
        title: "Open Source",
        description: `<div class="container">
        <p class="mt-3">Asai is an open-source project, centered around transparency and community collaboration. Our platform is based on open source tools like Langchain-Go and Ollama, among other popular libraries widely utilized in software development. We value trust and transparency, which is why we have chosen to build Asai in the open-source domain.</p>

        <p>In the early stages of development, the source code will be momentarily closed for public contributions. However, this is only so that we can perfect the initial stages of the Asai cloud platform. Immediately after this phase, the source code will be published on GitHub, becoming openly available to the user community.</p>
        </div>`,
      }, 
      "transparency": {
        title: "Transparent",
        description: `<div class="container">
        <p class="mt-3">By choosing to work in the open-source world, we aim to build a strong sense of trust with our user base. We understand the value of transparency in showing how user data and workflows are managed. When users can see and understand the internal workings of the system they use, it fosters trust and encourages continuous platform improvement.</p>

        Through our forthcoming GitHub public contributions, we will not only be sharing our platform with the world but also learning from global insights. This allows users to be confident in the integrity of their data, its management, and their overall interaction with the Asai platform. Our ultimate goal is to provide transparency, ensuring users can see how their contributions are influencing the continual enhancement and growth of Asai.</p>

        <p>In addition, Asai provides you with exceptional control over your Avatar. You can decide on the type of models you want to use and actively manage the costs associated with using these models. One of the major advantages is your ability to manage where your data and documentation is stored. This control extends to how your AI avatar manages these elements. The design of Asai is rooted in a user-first approach, and we provide every tool possible for users to have direct control over their experience. We aim for you to feel confident in using Asai, knowing that you have control over your avatar's learning, cost, and data handling process.</p>
        <div class="container">`,
      },
      "customization": {
        title: "Customizable",
        description: `<div class="container">
        <p class="mt-3">We've built a user-friendly avatar creation wizard to get you started with your Avatar swiftly, but if you'd like to delve deep and micromanage your Avatar, there's a plethora of possibilities available. </p>

        <p>You can custom-create the basic instructions and a persona that your Avatar will embody, thereby crafting its unique personality. Decide on the foundational models your avatar has access to, allowing you to balance between cost and capabilities.</p>
        
        <p>Once your avatar is set up, there's a marketplace of agents waiting to be integrated into your Avatar. These agents can be utilised via a range of plugins, allowing seamless communication with your creation through popular chat apps like Discord or Slack, or integrated as an assistant in product management tools such as Trello and Confluence. The expansive customization options at your disposal make it a cinch to create an AI Avatar that fits your unique needs</p>
        </div>`,
      },
      "personalization": {
        title: "Personalized",
        description: `<div class="container">
        <p class="mt-3">Personalized experiences are at the heart of what makes Asai special. Our platform fosters the creation of lifelike artificial personalities, enabling you to communicate and collaborate with an AI that truly 'gets' you. With the integration of the <a href="https://en.wikipedia.org/wiki/Big_Five_personality_traits" target="_blank">OCEAN model</a>, also known as the Big Five personality traits, your AI Avatar goes beyond just understanding your commands - it understands your nuances, embodied in a quantified version of the persona you envision.</p>

        <p>This advanced implementation couples with our memory-centric <a href="https://github.com/cpacker/MemGPT" target="_blank">MemGPT-based</a> architecture. Your AI Avatar doesn't just 'know' - it 'remembers'. From short-term interactions to long-term trends, your Avatar develops a coherent and consistent persona. It doesn't just respond; it evolves with you.</p>
        
        <p>With Asai, you're not just opting for automation. You're choosing a personalized experience - a unique, intuitive, and evolving assistant that's custom-made for you.</p>
        </div>`,
      },
      "privacy": {
        title: "Private",
        description: `<div class="container">
        <p class="mt-3">With Asai, we prioritize privacy and transparency. In an AI-driven world with justified concerns about personal data, we provide a clear view of our processes through an open-source approach. At the same time, we allow for self-hosting or operating your Avatar locally, offering an additional level of data security and control based on your needs.</p>

        <p>More than just technology, Asai represents an initiative to seamlessly integrate AI into the fabric of your life, your workflows, and your world. It's about empowering individuals with customized AI technologies without compromising privacy. With us, you get the privilege of a personalized AI experience, keeping you in control while staying true to our commitment to data protection.</p>
        </div>`,
      },
      "integrations": {
        title: "Integrations",
        description: `<div class="container">
        <p class="mt-3">Our AI Avatar platform is designed with adaptability at its core. Understanding that every business has unique needs, we offer specifically tailored Avatars that can be seamlessly integrated into your business and internal infrastructure. This level of customization empowers you to shape your Avatar to fit perfectly within your specific context, further enhancing your workflows and productivity.

        <p>Moreover, this bespoke setup allows us to establish a custom ecosystem catered to your requirements. This ecosystem can be further extended with specific agents, each one custom-tailored, thereby truly personalizing your AI technology experience.</p>
        
        <p>For more information or to discuss custom integrations, please reach out to us at <a href="mailto:contact@astrosynapse.ai">contact@astrosynapse.ai</a>. We're committed to helping you create an AI solution that's just right for you.</p>
        </div>`,
      } 
    },
  }),
  actions: {

  },
})