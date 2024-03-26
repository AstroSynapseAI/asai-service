import { createRouter, createWebHistory } from "vue-router";

const onboardingRoutes = [
  {
    path: '',
    name: 'welcome',
    component: () => import('@/views/admin/onboarding/WelcomeView.vue')
  },
  {
    path: 'create',
    name: 'create-avatar',
    component: () => import('@/views/admin/onboarding/CreateAvatarView.vue')
  },
  {
    path: 'model',
    name: 'choose-model',
    component: () => import('@/views/admin/onboarding/ChooseModelView.vue')
  },
  {
    path: 'agents',
    name: 'select-agents',
    component: () => import('@/views/admin/onboarding/SelectAgentsView.vue')
  },
  {
    path: 'created',
    name: 'avatar-created',
    component: () => import('@/views/admin/onboarding/AvatarCreatedView.vue')
  }
]

export default onboardingRoutes
