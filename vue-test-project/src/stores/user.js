import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: () => ({
    user_id: null,
  }),
  getters: {
    //ローカルにuser_id情報があるかを確認、あったらstoreにローカルからもらってきたuser_idを代入しtrueを返還、なかったらfalseを返還
    isLoggedIn() {
      let local_user_id = localStorage.getItem("user_id")
      if(local_user_id != null){
        const nowTime = new Date().getTime();
        const login_date = parseInt(localStorage.getItem("login_date"))
        
        //test 180000 3分
        //ログインした時から24時間(86400000)が経った場合||24時間内で合っても日付が変わった場合、ログアウトさせる
        if(nowTime-login_date >= 86400000 || new Date(nowTime).getDate() != new Date(login_date).getDate()){　// !!!!!!!!!!!!!!
          console.log("nowTime : " + new Date(nowTime));
          console.log("login_date : " + new Date(login_date));
          alert("ログアウトします");
          localStorage.clear();
          return false
        }else{
          this.user_id = local_user_id;
          return true;
        }
      }else{
          return false;
      }
    },
    //ローカルからuser_idを持ってくる
    getUserId() {
      return localStorage.getItem("user_id")
    },
    getLogin_date() {
      return parseInt(localStorage.getItem("login_date"))
    }
},
actions: {
    //ログインできたらuser_id,login_dateをローカルにsave
    setUserId(student_id) {
      const nowTime = new Date().getTime();

      localStorage.setItem("user_id", student_id)
      localStorage.setItem("login_date", nowTime)
    },
    // test!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    clear() {
        localStorage.clear();
    }
    
  },
});