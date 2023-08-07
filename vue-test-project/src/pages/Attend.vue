
<template>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
    </head>
    
    <div class="app">
        <Header :title='title'/>
        <div class="body">
            <div id="title">
                <!-- 出席可能な講義がある場合講義名を、ない場合メッセージを表示 -->
                <div id="can_attend" v-if="can_attend">
                    <h1 id="subject">{{subject}}</h1>
                    <h3>に出席されますか</h3>
                </div>
                <div id="can_attend" v-else>
                    <h2>出席可能な講義がありません</h2>
                </div>
            </div>
            <button :class = "{can_click_btn:!button_clicked, cant_click_btn:button_clicked}" @click="onclick()" v-bind:disabled="button_clicked==true">出席</button>
            <br>
            <router-link to="/absent">事前欠席登録はこちら</router-link>
            <Footer />
        </div>

    
    </div>
</template>

<script>
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'

// window.onload = function(){
//     console.log("start")
// };

export default {
    created() {
        console.log('load API 実行')
    },
    data (){
        return{
            title:'出席情報登録',
            can_attend: true,
            subject:'講義名',
            button_clicked: false
        }
    },
    components: {
        Header,
        Footer
    },
    methods: {
        //出席ボタンを押した時のイベント
        onclick(){
            this.button_clicked= true
            //[POST]/api/attendence/attend
        }
    }
}
</script>

<style></style>