
<template>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
    </head>

    <Header :title='title'/>
    <div id="attend_content">
        <!-- 出席可能な講義がある場合講義名を、ない場合メッセージを表示 -->
        <div id="title">
            <div id="can_attend" v-if="can_attend">
                <h1>{{subject}}</h1>
                <h3>に出席されますか</h3>
            </div>
            <div id="can_attend" v-else>
                <h2>只今出席可能な</h2>
                <h2>講義がありません</h2>
            </div>
        </div>
        <div id="button">
            <button class = "radius_btn" :class = "{can_click_btn:!button_clicked, cant_click_btn:button_clicked}" @click="onclick" v-bind:disabled="button_clicked==true">出席</button>
        </div>
        <br>
        <router-link to="/absent">事前欠席登録はこちら</router-link>
    </div>
    <Footer />

</template>

<script>
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'
import axios from 'axios'

import { useUserStore } from '../stores/user'
const user = useUserStore();
const user_id = user.user_id;

export default {
    data (){
        return{
            title:'出席情報登録',
            can_attend: true,
            subject:'講義名',
            subject_id: '',
            button_clicked : false
        }
    },
    mounted() {
        // ページロード時、実行されるmethod
        // 学生id、現在時間をPOSTで送り、subject_name, subject_id, http_statusをもらう
        const reloadNowTime = Math.floor((new Date()).getTime() / 1000);
        axios.post(this.BASE_URL+"reload", { "student_id":user_id, "now_time":reloadNowTime})
        .then(response=> {
            const object = response.data[0]
            console.log(response)

            if(object != null){
                if(object.http_status < 400){
                    this.can_attend = true;
                    this.subject = object.subject_name
                    this.subject_id = object.subject_id
                }else{
                    this.can_attend = false
                    this.button_clicked = true
                    switch(object.http_status){
                        case 400: alert("URLに誤りがあります。URLを確かめた上でもう一度入り直してください。"); break;
                        case 401: alert("閲覧権限がありません。URLを確かめた上でもう一度入り直してください。"); break;
                        case 404: alert("該当するページが見つかりません。URLに誤りがないか確認してください。"); break;
                        case 429: alert("サーバーへのアクセスが集中しています。少し時間を空けてください。"); break;
                        case 503: alert("一時的にサービスが利用できません。少し時間を空けてください。"); break;
                        case 504: alert("正しく処理が行われませんでした。もう一度お試しください。"); break;
                        default: alert("エラーが発生しました。もう一度お試しください。"); break;
                    }
                }
            }else{
                this.can_attend = false
                this.button_clicked = true
            }
        })
        .catch(
            error=>{
                this.can_attend = false
                this.button_clicked = true

                console.log(error)
                switch(error.code){
                    case 'ERR_NETWORK': alert("サーバーに問題が発生しております。担当者にお問い合わせください。"); break;
                    default: alert("エラーが発生しました。もう一度お試しください。"); break;
                }
            }
        )

    },
    components: {
        Header,
        Footer
    },
    methods: {
        //出席ボタンを押した時、出席登録を行うイベント
        onclick(){
            this.button_clicked = true;
            const attendNowTime = Math.floor((new Date()).getTime() / 1000)
            axios.post(
                this.BASE_URL + 'attend',
                {
                    attend_flag : 1,
                    now_time : attendNowTime,
                    student_id : user_id,
                    subject_id : this.subject_id
                }
            ).then( response => {
                console.log(response)
                const object = response.data
                if(object.http_status < 400){
                    if(object.check_flag>0){
                        user.can_back = true;
                        setTimeout(() => {this.$router.push('/attendance_check')}, 500);
                    }else{
                        alert("出席登録できませんでした。再度お試しください。")
                    }
                }else{
                    switch(object.http_status){
                        case 400: alert("URLに誤りがあります。URLを確かめた上でもう一度入り直してください。");  this.$router.go(0); break;
                        case 401: alert("閲覧権限がありません。URLを確かめた上でもう一度入り直してください。");  this.$router.go(0); break;
                        case 404: alert("該当するページが見つかりません。URLに誤りがないか確認してください。");  this.$router.go(0); break;
                        case 429: alert("サーバーへのアクセスが集中しています。少し時間を空けてください。");  this.$router.go(0); break;
                        case 503: alert("一時的にサービスが利用できません。少し時間を空けてください。");  this.$router.go(0); break;
                        case 504: alert("正しく処理が行われませんでした。もう一度お試しください。");  this.$router.go(0); break;
                        default: alert("エラーが発生しました。もう一度お試しください。"); this.$router.go(0); break;
                  }
                }
            }).catch( error=>{
                this.can_attend = false
                this.button_clicked = true
                console.log(error)
                switch(error.code){
                    case 'ERR_NETWORK': alert("サーバーに問題が発生しております。担当者にお問い合わせください。"); this.$router.go(0); break;
                    default: alert("エラーが発生しました。もう一度お試しください。"); this.$router.go(0); break;
                }
            })
        }
    }
}

</script>

<style></style>