<script>
import DropDown from '../components/DropDown.vue'
import AbsButton from '../components/AbsButton.vue'
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'

export default {
  components: {
    DropDown,
    AbsButton,
    Header,
    Footer
  },
  data(){
    return{
        title:"事前欠席登録"
    }
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
      <Header :title="title"/>
    <body>
      <div>
        <DropDown />
      </div>
      <div>
        <AbsButton />
      </div>
    </body>
    <Footer />
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