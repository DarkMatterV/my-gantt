import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      redirect: '/gantt'
    },
    {
      path: '/gantt',
      name: 'gantt',
      component: () => import('../views/GanttTableView.vue')
    }
  ],
})

export default router
