
import { createRouter, createWebHistory } from "vue-router";

const agentConfigRoutes = [
  {
    path: ':agent_id/config/:active_agent_id?',
    name: 'search-agent',
    component: () => import('@/views/admin/agents/SearchAgent.vue')
  },
  {
    path: ':agent_id/config/:active_agent_id?',
    name: 'browser-agent',
    component: () => import('@/views/admin/agents/BrowserAgent.vue')
  },
  {
    path: ':agent_id/config/:active_agent_id?',
    name: 'email-agent',
    component: () => import('@/views/admin/agents/EmailAgent.vue')
  }
] 

export default agentConfigRoutes
