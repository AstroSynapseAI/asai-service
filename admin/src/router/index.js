import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/:avatar',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    },
    {
      path: '/:avatar/personality',
      name: 'personality',
      component: () => import('../views/PersonalityView.vue')
    },
    {
      path: '/:avatar/agents',
      name: 'agents',
      component: () => import('../views/AgentsView.vue')
    },
    {
      path: '/:avatar/agents/:slug/config',
      name: 'agent-config',
      component: () => import('../views/AgentConfigView.vue')
    },
    {
      path: '/:avatar/plugins',
      name: 'plugins',
      component: () => import('../views/PluginsView.vue')
    },
    {
      path: '/:avatar/models',
      name: 'models',
      component: () => import('../views/ModelsView.vue')
    },
    {
      path: '/account',
      name: 'account',
      component: () => import('../views/AccountView.vue')
    },
  ]
})

export default router

// import { createRouter, createWebHistory } from 'vue-router';

// import { useAuthStore, useAlertStore } from '@/stores';
// import { Home } from '@/views';
// import accountRoutes from './account.routes';
// import usersRoutes from './users.routes';

// export const router = createRouter({
//     history: createWebHistory(import.meta.env.BASE_URL),
//     linkActiveClass: 'active',
//     routes: [
//         { path: '/', component: Home },
//         { ...accountRoutes },
//         { ...usersRoutes },
//         // catch all redirect to home page
//         { path: '/:pathMatch(.*)*', redirect: '/' }
//     ]
// });

// router.beforeEach(async (to) => {
//     // clear alert on route change
//     const alertStore = useAlertStore();
//     alertStore.clear();

//     // redirect to login page if not logged in and trying to access a restricted page 
//     const publicPages = ['/account/login', '/account/register'];
//     const authRequired = !publicPages.includes(to.path);
//     const authStore = useAuthStore();

//     if (authRequired && !authStore.user) {
//         authStore.returnUrl = to.fullPath;
//         return '/account/login';
//     }
// });

