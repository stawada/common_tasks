<script setup>
    import { useUserStore } from '../stores/user'
    const user = useUserStore();


</script>

<script>
export default {
  data() {
    return{
        student_id: ""
    }
  }
}
</script>

<template>
    <h1>LOGOUT</h1>
    <input type="text" name="" id="" v-model="student_id">
    <input type="button" value="login" @click="user.setUserId(student_id)">
    <input type="button" value="logout">
    <input type="button" value="storage clear" @click="user.clear">
    <p> user storage.user_id : {{ user.user_id }}</p>
    <p> localStorage.user_id : {{ user.getUserId }}</p>
    <p> isLoggedIn : {{ user.isLoggedIn }}</p>
</template>