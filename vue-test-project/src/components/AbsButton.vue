<template>
    <div class="button-container">
        <div class="AbsButton">
            <button @click="handleAbsent">欠席</button>
        </div>
    </div>
</template>

<script>
    export default{
        name:'AbsButton',
        methods: {
            handleAbsent() {
                console.log("欠席ボタンがクリックされました");
            }
        }
    }
</script>
