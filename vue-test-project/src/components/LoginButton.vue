<template>
    <div class="button-container">
        <div class="LoginButton">
            <button @click="handleAbsent(); getLoginInfo()">本人確認</button>
        </div>
    </div>
</template>

<script>
    import TextField from './TextField.vue';
    import axios from 'axios';

    export default{
        name:'LoginButton',
        methods: {
            handleAbsent() {
                // console.log("ボタンを押しました");
                var [std_id, password] = TextField.getLoginInfo(); // 入力した要素を取得
                console.log(std_id, password)

                const APIServer = 'http://localhost:8080/attendance/login'
                const RequestLoginJson = axios.post(APIServer, {
                    student_id: std_id,
                    hashed_password: password,
                })
                console.log(RequestLoginJson);
            }
        }
    }
</script>

<style>
.button-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
}

.AbsButton {
    margin: 10px;
}
</style>