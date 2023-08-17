
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
const user_id = user.getUserId;

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
        console.log('reload API 実行')
        const nowTime = Math.floor((new Date()).getTime() / 1000)
        // console.log("현재시간 타임스탬프 : "+ nowTime)
        console.log("user_id : " + user_id)

        // ページロード時、実行されるmethod
        // 学生id、現在時間をPOSTで送り、subject_name, subject_id, http_statusをもらう
        axios.post(this.BASE_URL+"reload", { "student_id":"G000000000", "now_time":1691553600}) //!!!!!!!!!!!!!!!!!テスト
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
                    if(object.http_status == 400){
                        alert("BadRequest")
                    }else if(object.http_status == 401){
                        alert("Unauthorized")
                    }else if(object.http_status == 404){
                        alert("Not found")
                    }else if(object.http_status == 412){
                        alert("Precondition Failed")
                    }else if(object.http_status == 429){
                        alert("To Many Requests")
                    }else if(object.http_status == 503){
                        alert("Service Unavailable")
                    }else if(object.http_status == 504){
                        alert("Gateway Timeout")
                    }else{
                        alert("Unknown Error")
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

                console.log(error.response)
                const status = error.response.status
                const message = error.response.data.message
                alert(status + "ERROR : " + message)
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
            const nowTime = Math.floor((new Date()).getTime() / 1000)

            axios.post(
                this.BASE_URL + 'attend',
                {
                    attend_flag : 1,
                    now_time : 1699999900,      //!!!!!!!!!!nowTime
                    student_id : "G000000000",  //!!!!!!!!!!user_id
                    subject_id : this.subject_id
                }
            ).then( response => {
                console.log(response)
                const object = response.data
                if(object.http_status < 400){
                    if(object.check_flag>0){
                        setTimeout(() => {this.$router.push('/attendance_check')}, 500);
                    }else if(object.http_status == 400){
                        alert("BadRequest")
                        this.$router.go(0)
                    }else if(object.http_status == 401){
                        alert("Unauthorized")
                        this.$router.go(0)
                    }else if(object.http_status == 404){
                        alert("Not found")
                        this.$router.go(0)
                    }else if(object.http_status == 412){
                        alert("Precondition Failed")
                        this.$router.go(0)
                    }else if(object.http_status == 429){
                        alert("To Many Requests")
                        this.$router.go(0)
                    }else if(object.http_status == 503){
                        alert("Service Unavailable")
                        this.$router.go(0)
                    }else if(object.http_status == 504){
                        alert("Gateway Timeout")
                        this.$router.go(0)
                    }else{
                        alert("Unknown Error")
                        this.$router.go(0)
                    }
                }else{
                    alert("ネットワークエラーが発生しました。")
                    router.go(0)
                }
            }).catch(
                error=>{
                    this.can_attend = false
                    this.button_clicked = true

                    console.log(error.response)
                    const status = error.response.status
                    const message = error.response.data.message
                    alert(status + "ERROR : " + message)
                })
        }
    }
}

</script>

<style></style>