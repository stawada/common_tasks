<template>
    <div class="LoginButton">
        <button @click='$emit("handleAbsent")'>本人確認</button>
    </div>
</template>

<script>
    export default{
        name:'LoginButton',
        methods: {
            
            // handleAbsent() {
            //     console.log("ボタンを押しました");
            // }
        }
    }
</script>

<style>
.LoginButton button{
    border: 0;
    height: 47px;
    width: 160px;

    color: white;
    font-size: 32px;
    background-color: #787878;
}
</style>