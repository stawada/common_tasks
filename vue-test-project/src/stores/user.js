import { defineStore } from "pinia";

export const useUserStore = defineStore("user", {
  state: () => ({
    user_id: null,
  }),
  getters: {
    //ローカルにuser_id情報があるかを確認、あったらstoreにローカルからもらってきたuser_idを代入しtrueを返還、なかったらfalseを返還
    isLoggedIn() {
        let local_user_id = this.getUserId;
        if(local_user_id != null){
            this.user_id = local_user_id;
            return true;
        }else{
            return false;
        }
    },
    //ローカルからuser_idを持ってくる
    getUserId() {
      return localStorage.getItem("user_id")
    }
},
actions: {
    //ログインできたらuser_idをローカルにsave
    setUserId(student_id) {
      localStorage.setItem("user_id", student_id)
    },
    // 後で消す予定!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    clear() {
        localStorage.clear();
    }
    
  },
});