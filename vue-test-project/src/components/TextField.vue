<template>
    <form name="login">
        <div class="textfield-container">
            <div class="id">
                <p class="label">学生番号</p>
                <input type="id" name="student_id" placeholder="学生番号" v-model="student_id" @change="onEmit">
            </div>
            <div class="password">
                <p class="label">パスワード</p>
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

    .label{
        text-align: left;
        margin: auto;
        padding-bottom: 5px;
        width: 277px;
        font-size: 20px;
        color: white;
        font-weight: 600;
    }

</style>