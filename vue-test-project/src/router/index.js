import { createRouter, createWebHistory } from 'vue-router'
import Top from '../pages/Top.vue'
import { useUserStore } from '../stores/user'

const router = new createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
      component: Top,
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
    },
    // 後で消す!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    {
      path: '/logout',
      name: 'logout',
      component: () => import('../pages/Logout.vue')
    }
  ]
})
    //!!!!!!!!!!!!!!test後修正する
    router.beforeEach((to,next) => {
      const user = useUserStore();
      if(to.name != "logout" && !user.isLoggedIn){
        alert("user.isLoggedIn : " + user.isLoggedIn)
        return "/logout"
      }else if(to.name == 'login' && user.isLoggedIn){
        alert("既にログイン情報があります。")
        return "/attend"
      }
      // //ログインしていないユーザがログイン外のページに移動する場合、ログインページに遷移する
      // if(to.name != "login" && !user.isLoggedIn){
      //   alert("ログインしてください。")
      //   return "/"
      // //ログインしているユーザがログインページに移動する場合、出席登録ページに遷移する
      // }else if(to.name == 'login' && user.isLoggedIn){
      //   alert("既にログイン情報があります。")
      //   return "/attend"
      // }
    })
    
    export default router