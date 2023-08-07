<script>
import TextField from '../components/TextField.vue'
import LoginButton from '../components/LoginButton.vue'
export default {
    components: {
        TextField,
        LoginButton
    }
};
</script>

<template>
  <head>
    <meta name='viewpoint' content="width=device-width" initial-scale=1.0>
    <title>本人確認画面</title>
  </head>
  <div class="logo-container">
    <header>
      <img src="@/assets/logo.png" alt="logo" class="logo">
    </header>
  </div>
  <body>
    <div class="components">
      <TextField />
      <LoginButton />
    </div>
  </body>
</template>

<style>
h1 {
    margin: 100px auto;
    text-align: center;
font-size: 50px;
};

.logo-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
};

body {
    /*display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;*/
    margin: center;
};


</style>