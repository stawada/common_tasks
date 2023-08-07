import { createRouter, createWebHistory } from 'vue-router'
import Top from '../pages/Top.vue'


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Top
    },
    {
      path: '/attend',
      name: 'attend',
      component: () => import('../pages/Attend.vue')
    },
    {
      path: '/absent',
      name: 'absent',
      component: () => import('../pages/Absent.vue')
    },
    {
      path: '/absent_check',
      name: 'absent_check',
      component: () => import('../pages/AbsentCheck.vue')
    },
    {
      path: '/attendance_check',
      name: 'attendance_check',
      component: () => import('../pages/AttendanceCheck.vue')
    }
  
  ]
})

export default router
