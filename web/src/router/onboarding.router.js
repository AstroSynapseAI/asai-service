import { createRouter, createWebHistory } from "vue-router";

const onboardingRoutes = [
  {
    path: '',
    name: 'welcome',
    component: () => import('@/views/onboarding/WelcomeView.vue')
  },
  {
    path: 'create',
    name: 'create-avatar',
    component: () => import('@/views/onboarding/CreateAvatarView.vue')
  },
  {
    path: 'model',
    name: 'choose-model',
    component: () => import('@/views/onboarding/ChooseModelView.vue')
  },
  {
    path: 'agents',
    name: 'select-agents',
    component: () => import('@/views/onboarding/SelectAgentsView.vue')
  },
  {
    path: 'created',
    name: 'avatar-created',
    component: () => import('@/views/onboarding/AvatarCreatedView.vue')
  }
]

export default onboardingRoutes
