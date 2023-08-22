<template>
    <form name="login">
        <div class="textfield-container">
            <div class="id">
                <input type="id" name="student_id" placeholder="学生番号" v-model="student_id" @change="onEmit">
            </div>
            <div class="password">
                <input type="password" name="password" placeholder="パスワード" v-model="password" @change="onEmit">
            </div>
        </div>
    </form>
</template>

<script>

export default {
    data (){
        return{
            student_id:'',
            password:''
        }
    },

    methods: {
        onEmit()
        {
        this.$emit("appendVal", this.student_id, this.password)
        }
    }
}
</script>

<style>
    .id input, .password input{
    border: 0;
    background-color: white;

    text-align: center;
    font-size: 30px;
    color: #494848;

    width: 270px ;
    height: 47px;
    }
    .id{
    padding-bottom: 35px;
    }

</style>