import { defineStore } from "pinia";
import axios from 'axios'

axios.defaults.withCredentials = true;

export const useUserStore = defineStore("user", {
  state: () => ({
    user_id: null
  }),
  getters: {
    //Cookieからユーザーidを読み込み、情報があったらuser_idに保存してからtrueを、なかったらfalseを返還
    isLoggedIn(state) {
      state.user_id = $cookies.get("student_id");
      return state.user_id != null
    }
},
  actions: {
  }

});