
<template>
    <head>
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
    </head>
    
    <Header :title='title'/>
        <div id="attend_content">
            <!-- 出席可能な講義がある場合講義名を、ない場合メッセージを表示 -->
            <div id="can_attend" v-if="can_attend">
                <h1>{{subject}}</h1>
                <h3>に出席されますか</h3>
            </div>
            <div id="can_attend" v-else>
                <h2>出席可能な講義がありません</h2>
            </div>
            <button class = "radius_btn" :class = "{can_click_btn:!button_clicked, cant_click_btn:button_clicked}" @click="onclick()" v-bind:disabled="button_clicked==true">出席</button>
            <br>
            <router-link to="/absent">事前欠席登録はこちら</router-link>
        </div>
       <Footer />

</template>

<script>
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'
import axios from 'axios'

export default {
    created() {
        console.log('load API 実行')
        const nowTime = Date.now();
        console.log("현재시간 타임스탬프 : "+ nowTime)
        console.log("baseurl "+ this.BASE_URL)


        //ページロード時、実行されるmethod：学生id、現在時間をPOSTで送り、受講できるかについてのstatus、受講できる講義名、講義idをもらう
        axios.post(this.BASE_URL + 'reload' //?????
            ,{
                student_id: 'G000000000',  //!!!!!저장소에서 학생아이디 꺼내기
                //now_time : nowTime
            }
            // ).then( response => console.log(response)
            //         if(response.status>0){
            //             this.can_attend = true;
            //             this.can_attend_subject = "math"
            //             this.subject_id = '111111'
            //         }else{
            //             this.can_attend = false
            //         }
            // ).catch(console => console.log(error))
        )
    },
    data (){
        return{
            title:'出席情報登録',
            can_attend: true,
            subject:'数学I',
            subject_id: '',
            button_clicked : false
        }
    },
    components: {
        Header,
        Footer
    },
    methods: {
        //出席ボタンを押した時のイベント
        // onclick(){
        //     this.button_clicked = true;
        //     alert("버튼 클릭");
        //     setTimeout(()=>{
        //         this.$router.push('/attendance_check')
        //     },1000);
        //     axios.post(this.BASE_URL + 'attend',
        //     {
        //         attend_flag : 1,
        //         attend_day : Date.now(),
        //         student_id : "111111",
        //         subject_id : this.subject_id
        //     })
        //     .then(
        //         response=>console.log(response)
        //         this.$router.push('/attendance_check')
        //     ).catch(error => console.log(error))
        // }
    }
}
</script>

<style></style>