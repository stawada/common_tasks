<script setup>
import { useUserStore } from '../stores/user'
  const user = useUserStore();

</script>

<script>
import { useUserStore } from '../stores/user'
const user = useUserStore();

export default {
  created() {
    console.log("test")
    var subjectDict = {}
    var timeDict = {}
    timeDict['time1']='english-1'
    subjectDict['english']
  }
  ,data() {
    return{
        student_id: ""
    }
  },methods : {
    reload() {
 
    },
    login(student_id){
      user.setUserId(student_id)
    }
  }
}
</script>

<template>
    <h1>Login Test</h1>
    <input type="text" name="" id="" v-model="student_id">
    <input type="button" value="login" @click="login(student_id)">
    <input type="button" value="reload" @click="reload">
    <input type="button" value="storage clear" @click="user.clear">
    <p> localStorage.user_id : {{ user.getUserId }}</p>
    <p> localStorage.login_date : {{ new Date(user.getLogin_date) }}</p>
    <p> isLoggedIn : {{ user.isLoggedIn }}</p>
</template>