import { createRouter, createWebHistory } from 'vue-router'
import AsaiChatView from '@/views/AsaiChatView.vue'
import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/Home/AboutView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'asai',
      component: AsaiChatView
    },
    {
      path: '/login',
      name: 'login',
      component: HomeView
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView,
      children: [
        {
          path: 'about/:slug?',
          name: 'about',
          component: AboutView
        },
      ]
    },
  ],
  scrollBehavior(to, from, savedPosition) {
    if (to.hash) {
      return new Promise((resolve) => {
        setTimeout(() => {
          resolve({ el: to.hash, behavior: 'smooth' })
        }, 500)
      })
    }
    else if (savedPosition) {
      return savedPosition;
    }
    else {
      return { top: 0 };
    }
  }
})

export default router