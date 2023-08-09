<script>
//各種インポート
import TextField from '../components/TextField.vue'
import LoginButton from '../components/LoginButton.vue'
import axios from 'axios'


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
      onclick(){
        alert("click");
        console.log("親コンポーネント　id: " + this.student_id + "　pw: " + this.password)
        this.getJson()
      },

      getJson(){
        axios.post('http://localhost:8080/attendance/login',{
        "student_id": this.student_id,
        "hashed_password": this.password,
      })
      .then(
        response => console.log(response),
        this.$router.push('/attendance_check')
      ).catch(error => console.log(error))
      },
    },

    data() {
      return {
        student_id: "",
        password: ""
      }
    },


};
</script>

<template>
  <div class="all">
    <div class="logo-container">
      <header>
        <img src="@/assets/logo.png" alt="logo" class="logo">
      </header>
    </div>
    <div class='app'>
      <div class="components">
        <TextField @appendVal="getVal"/>
        <LoginButton @handleAbsent="onclick"/>
      </div>
    </div>
  </div>
</template>

<style>
body{
  background-color: white;
  height: 100%;
}

.logo-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
}

.app{
  height: 100%;
}
.all {
  /*display: flex;
  flex-direction: column;
  justify-content: center;
  padding:50px 150px 50px 150px;
  margin: center;
  align-content: center;
  width: 300px;
  */
  height: 100%;
  padding: 30% 0px 20% 0px;
  background-color: #F3AF2B;

}


</style>