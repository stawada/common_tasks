<script setup>
  import { useStoreCounter } from '../stores/counter'
  const counter = useStoreCounter();

  import { useUserStore } from '../stores/user'
  const user = useUserStore();
</script>

<script>
import TextField from '../components/TextField.vue'
import LoginButton from '../components/LoginButton.vue'

export default {
    components: {
        TextField,
        LoginButton
    },
    data() {
      return {
        student_id: "",
        password: ""
      }
    },
    methods:{
      getVal(id, pw){
        this.student_id = id;
        this.password = pw;
        // console.log("親コンポーネント　id: " + this.student_id + "　pw: " + this.password)
      },
      onclick(){
        alert("click");
        console.log("id: " +this.student_id + " pw: " + this.password)
        // localStorage.setItem("student_id", student_id)
      }
    }
};
</script>

<template>
  <head>
    <meta name='viewpoint' content="width=device-width" initial-scale=1.0>
    <title>本人確認画面</title>
  </head>
  <div class="logo-container">
    <header>
      <img src="@/assets/logo.png" alt="logo" class="logo">
    </header>
  </div>
  <body>
    <div class="components">
      <p>{{counter.count}}</p>
      <p>{{user.user_id}}</p>
      <p>{{user.isLoggedIn}}</p>
      <TextField @appendVal="getVal"/>
      <LoginButton @handleAbsent="onclick"/> <!-- emit -->
    </div>
  </body>
</template>

<style>
/*
h1 {
    margin: 100px auto;
    text-align: center;
font-size: 50px;
};

.logo-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
};
*/
/*
body {
    /*display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin: center;
}; */


</style>