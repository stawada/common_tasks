<template>
  <header>
    <div class="container">
      <br><br><br><br><br><img src="@/assets/check.png" alt="check" class="check-img"/><br><br><br>
    </div>
  </header>
  <body>
    <p class="massage">　欠席登録できました。</p>
  </body>
</template>

<style>
@media screen and (min-width: 400px){ /*PC用 */
html {
  background-color: white;
  margin-left: 22%;
  margin-right: 22%;
  padding: 0;
}

body {
background-color: gray;
}
}

@media screen and (max-width: 400px){ /*スマホ用 */
html {
  background-color: gray;
}

body {
background-color: gray;
}
}

header {
  background-color: gray;
}

.container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
  background-color: gray;

}

.check-img {
  text-align: center;
  width: auto;
  height: auto;
}

.massage{
  font-size: 25px;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: auto;
  /*margin-top: 50px;*/
  color: white;
  background-color: gray;
}

</style>