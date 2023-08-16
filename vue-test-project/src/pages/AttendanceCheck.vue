<template>
  <header>
    <div class="container">
      <br><br><br><br><br><img src="@/assets/check.png" alt="check" class="check-img"/><br><br><br>
    </div>
  </header>
  <body>
    <p class="massage">　出席登録できました。</p>
  </body>
</template>

<style>
html {
  background-color: orange;
}

body {
  background-color: orange;
}

header {
  margin-top: 100px;
}

.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
  background-color: orange;
}

.check-img {
  text-align: center;
  width: auto;
  height: auto;
  background-color: orange;
}

.massage{
  font-size: 25px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
  /*margin-top: 100px;*/
  color: white;
  background-color: orange;
}



</style>

