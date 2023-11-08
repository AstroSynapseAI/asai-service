import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import AsaiChatView from '../views/AsaiChatView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'asai',
      component: AsaiChatView
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView
    },
    // {
    //   path: '/about',
    //   name: 'about',
    //   // route level code-splitting
    //   // this generates a separate chunk (About.[hash].js) for this route
    //   // which is lazy-loaded when the route is visited.
    //   component: () => import('../views/AboutView.vue')
    // },
    // {
    //   path: '/:avatar/personality',
    //   name: 'personality',
    //   component: () => import('../views/PersonalityView.vue')
    // },
    // {
    //   path: '/:avatar/agents',
    //   name: 'agents',
    //   component: () => import('../views/AgentsView.vue')
    // },
    // {
    //   path: '/:avatar/agents/:slug/config',
    //   name: 'agent-config',
    //   component: () => import('../views/AgentConfigView.vue')
    // },
    // {
    //   path: '/:avatar/plugins',
    //   name: 'plugins',
    //   component: () => import('../views/PluginsView.vue')
    // },
    // {
    //   path: '/:avatar/models',
    //   name: 'models',
    //   component: () => import('../views/ModelsView.vue')
    // },
    // {
    //   path: '/account',
    //   name: 'account',
    //   component: () => import('../views/AccountView.vue')
    // },
  ]
})

export default router