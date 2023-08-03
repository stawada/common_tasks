<script>
import DropDown from './components/DropDown.vue'
import AbsButton from './components/AbsButton.vue'
export default {
  components: {
    DropDown,
    AbsButton
  }
};
</script>

<template>
  <html lang="ja">
    <head>
        <meta charset='UTF-8'>
        <meta name='viewpoint' content="width=device-width" initial-scale=1.0>
        <title>出欠選択画面</title>
    </head>
    <header>
      <h1>事前欠席登録</h1>
    </header>
    <body>
      <div>
        <DropDown />
      </div>
      <div>
        <AbsButton />
      </div>
    </body>
</html>
</template>

<style>
body {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
};

head {
  margin: 50px;
  padding: 100px;
};
</style>