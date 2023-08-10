<template>
    <div class="button-container">
        <div class="LoginButton">
            <button @click='$emit("handleAbsent")'>本人確認</button>
        </div>
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