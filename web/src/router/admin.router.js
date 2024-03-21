import { createRouter, createWebHistory } from 'vue-router'
import AgentConfigView from '@/views/admin/AgentConfigView.vue'
import agentConfigRoutes from "./agents.router";
import onboardingRoutes from "./onboarding.router";

const adminRoutes = [
  {
    path:'avatar/onboarding',
    component: () => import('@/views/admin/OnboardingView.vue'),
    children: onboardingRoutes
  },
  {
    path: 'avatar/create',
    name: 'create-avatar',
    component: () => import('@/views/admin/CreateAvatarView.vue')
  },
  {
    path: 'avatar/:avatar_id',
    name: 'admin',
    component: () => import('@/views/admin/HomeView.vue')
  },
  {
    path: 'avatar/:avatar_id/personality',
    name: 'personality',
    component: () => import('@/views/admin/PersonalityView.vue')
  },
  {
    path: 'avatar/:avatar_id/agents',
    name: 'agents',
    component: () => import('@/views/admin/AgentsView.vue')
  },
  {
    path: 'avatar/:avatar_id/agents',
    component: AgentConfigView,
    children: agentConfigRoutes
  },
  {
    path: 'avatar/:avatar_id/plugins',
    name: 'plugins',
    component: () => import('@/views/admin/PluginsView.vue')
  },
  {
    path: 'avatar/:avatar_id/plugins/:plugin_id/config/:active_plugin_id?',
    name: 'plugin-config',
    component: () => import('@/views/admin/PluginConfigView.vue')
  },
  {
    path: 'avatar/:avatar_id/models',
    name: 'models',
    component: () => import('@/views/admin/ModelsView.vue')
  },
  {
    path: 'avatar/:avatar_id/models/:model_id/config/:active_model_id?',
    name: 'model-config',
    component: () => import('@/views/admin/ModelConfigView.vue')
  },
  {
    path: 'account',
    name: 'account',
    component: () => import('@/views/admin/AccountView.vue')
  },
  {
    path: 'avatar/:avatar_id/tools',
    name: 'tools',
    component: () => import('@/views/admin/ToolsView.vue')
  },
  {
    path: 'avatar/:avatar_id/tools/:tool_id/config/:active_tool_id?',
    name: 'tool-config',
    component: () => import('@/views/admin/ToolConfigView.vue')
  },
  {
    path: 'users',
    name: 'users',
    component: () => import('@/views/admin/UsersView.vue')
  },
  {
    path: 'profile',
    name: 'profile',
    component: () => import('@/views/admin/ProfileView.vue')
  },
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

export default adminRoutes
