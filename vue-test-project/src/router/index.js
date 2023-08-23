import { createRouter, createWebHistory } from 'vue-router'
import Top from '../pages/Top.vue'
import { useUserStore } from '../stores/user'

const router = new createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'login',
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

    router.beforeEach((to,from) => {
        const user = useUserStore();
        console.log("canclick" + user.can_back)
  
        //ログインしていないユーザがログイン外のページに移動する場合、ログインページに遷移する
        if(to.name != "login" && !user.isLoggedIn){
          console.log("isLoggedIn: " + user.isLoggedIn) //!!!!!!!!!!!!
          alert("ログインしてください。")
          return "/"
        //ログインしているユーザがログインページに移動する場合、出席登録ページに遷移する
        }else if(to.name == 'login' && user.isLoggedIn){
          console.log("isLoggedIn: " + user.isLoggedIn) //!!!!!!!!!!!
          return "/attend"
        }

        //確認ページに戻れなくする
        if( (user.can_back==false && to.name=="absent_check") || (user.can_back==false && to.name=="attendance_check") ){
          console.log("back block")
          return "/attend"
        }else if(from.name == "absent_check" || from.name == "attendance_check"){
          user.can_back = false;
          console.log("can_back change to false");
        }
    })
    
    export default router
