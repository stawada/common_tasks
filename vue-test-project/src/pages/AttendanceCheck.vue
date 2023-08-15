<template>
  <header>
    <div class="container">
      <img src="@/assets/check.png" alt="check" class="check-img"/>
    </div>
  </header>
  <body>
    <p class="massage">出席登録できました。</p>
  </body>
</template>

<style>
html {
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
}

.check-img {
  width: 100px;
  height: 100px;
}

.massage{
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
  margin-top: 50px;
  color: white;
}

</style>