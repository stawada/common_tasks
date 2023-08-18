<template>
    <div class="button-container">
        <div class="AbsButton">
            <button class="radius_btn" :class = "{can_click_btn:!button_clicked, cant_click_btn:button_clicked}" @click="handleAbsent">欠席</button>
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
        },
        data() {
            return{
                button_clicked: false
            }
        }
    }
</script>

<style>

</style>