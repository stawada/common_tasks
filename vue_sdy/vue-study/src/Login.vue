<script>
import TextField from './components/TextField.vue'
import LoginButton from './components/LoginButton.vue'
export default {
    components: {
        TextField,
        LoginButton
    }
};
</script>

<template>
    <html lang="ja">
        <head>
            <meta charset='UTF-8'>
            <meta name='viewpoint' content="width=device-width" initial-scale=1.0>
            <title>本人確認画面</title>
        </head>
        <header>
            <h1>KANEKO UNIVERSITY</h1>
        </header>
        <body>
            <div>
                <TextField />
                <LoginButton />
            </div>
        </body>
    </html>
</template>

<style>
h1 {
    margin: 100px auto;
    text-align: center;
font-size: 50px;
};

body {
    /*display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;*/
    margin: center;
};


</style>