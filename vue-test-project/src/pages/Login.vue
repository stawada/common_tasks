<script setup>
  import { useUserStore } from '../stores/user'
  const user = useUserStore();

</script>

<script>
//各種インポート
import TextField from '../components/TextField.vue'
import LoginButton from '../components/LoginButton.vue'
import axios from 'axios'

axios.defaults.withCredentials = true;

export default {
    components: {
        TextField,
        LoginButton
    },

    methods: {

      getVal(id, pw){
        this.student_id = id;
        this.password = pw;
      },


      async onclick(){
        // alert("click");
        //console.log("親コンポーネント id: " + this.student_id + " pw: " +  await this.sha256(this.password));
        this.postJson()
      },

      async postJson(){
        axios.post('http://localhost:8080/attendance/login',{
        "student_id": this.student_id,
        "hashed_password": await this.sha256(this.password),
      })
      .then(
        response => {
          console.log(response)
          if (response.data.match_flag==1){
            axios.post('http://localhost:8080/setCookie', {
            "student_id": this.student_id,
            }).then(res => {
                console.log(res);
                this.$router.go(0)
            }).catch( err => {console.log(err)} )
          }
          else {
            alert("学生ID または パスワードが違います");
          }
          }).catch(
          error=>{
                console.log(error)
                switch(error.code){
                    case 'ERR_NETWORK': alert("サーバーに問題が発生しております。担当者にお問い合わせください。"); this.$router.go(0); break; //reload apiにはrouter不要
                    default: alert("エラーが発生しました。もう一度お試しください。"); this.$router.go(0); break;　//reload apiにはrouter不要
                }
            }
          )},

      async sha256(message) {
        const msgUint8 = new TextEncoder().encode(message);                           // encode as (utf-8) Uint8Array
        const hashBuffer = await crypto.subtle.digest('SHA-256', msgUint8);                 // hash the message
        const hashArray = Array.from(new Uint8Array(hashBuffer));                     // convert buffer to byte array
        const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join(''); // convert bytes to hex string
        //console.log("hashpassword: " + this.password + " = " + hashHex);
        return hashHex;
      },


    data() {
      return {
        student_id: "",
        password: "",
      }
    },
  }
};

</script>

<template>
  <div class="login">
    <div class="logo-container">
      <header>
        <img src="@/assets/logo.png" alt="logo" class="logo">
      </header>
    </div>
    <div class="textfield">
      <TextField @appendVal="getVal"/>
    </div>
    <LoginButton @handleAbsent="onclick"/>
  </div>
</template>

<style>
  .login{
  background-color: #F3AF2B;
  height: 100%;
  }
  .login div.logo-container{
    padding-top: 100px;
    padding-bottom: 40px;
  }
  .login div.textfield{
    padding-bottom: 45px;
  }
</style>