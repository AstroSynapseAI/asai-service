import { createRouter, createWebHistory } from 'vue-router'

const adminRoutes = [
  {
    path: 'avatar/:avatar',
    name: 'admin',
    component: () => import('@/views/admin/HomeView.vue')
  },
  {
    path: 'avatar/:avatar/personality',
    name: 'personality',
    component: () => import('@/views/admin/PersonalityView.vue')
  },
  {
    path: 'avatar/:avatar/agents',
    name: 'agents',
    component: () => import('@/views/admin/AgentsView.vue')
  },
  {
    path: 'avatar/:avatar/agents/:slug/config',
    name: 'agent-config',
    component: () => import('@/views/admin/AgentConfigView.vue')
  },
  {
    path: 'avatar/:avatar/plugins',
    name: 'plugins',
    component: () => import('@/views/admin/PluginsView.vue')
  },
  {
    path: 'avatar/:avatar/models',
    name: 'models',
    component: () => import('@/views/admin/ModelsView.vue')
  },
  {
    path: 'account',
    name: 'account',
    component: () => import('@/views/admin//AccountView.vue')
  },
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

export default adminRoutes
