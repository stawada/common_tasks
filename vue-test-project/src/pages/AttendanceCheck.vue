<template>
    <div class="check" id="attendance_check" @click="onclick">
      <div class="img">
        <img src="@/assets/check.png" alt="check" class="check-img"/>
      </div>
      <p class="massage">出席登録できました。</p>
      <h5>画面をタッチしたらホームに戻ります。</h5>
    </div>
</template>
<script>
  export default{
    methods: {
      onclick(){
        this.$router.push("/attend");
      }
    }
  }
</script>
<style>
</style>

