<template>
    <div class="button-container">
        <div class="AbsButton">
            <button class="radius_btn" @click='$emit("handleAbsent")'>欠席</button>
        </div>
    </div>
</template>

<script>
    export default{
        name:'AbsButton',
        methods: {
        }
    }
</script>
