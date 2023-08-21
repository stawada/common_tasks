<template>
    <div class="button-container">
        <div class="AbsButton">
            <button class="radius_btn" @click='$emit("handleAbsent")' :class = "{can_click_btn:!button_clicked, cant_click_btn:button_clicked}">欠席</button>
        </div>
    </div>
</template>

<script>
    export default{
        name:'AbsButton',
        props: {
            button_clicked: {
                type: Boolean,
                default (){
                    return []
                }
            }
        }
    }
</script>
