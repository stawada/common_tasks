<script>
import DropDown from '../components/DropDown.vue'
import AbsButton from '../components/AbsButton.vue'
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'
import axios from 'axios'
import { useUserStore } from '../stores/user'

const user = useUserStore();
const user_id = user.user_id;

export default {
  components: {
    DropDown,
    AbsButton,
    Header,
    Footer
  },

    data (){

        return{
            title:'事前欠席登録',
            subject:'',
            dateTime:'',
            dateUTC:'',
            subject_id: '',
            ans:'',
            dic:{},
            responseJSON:{},
            button_clicked : true,
            subjectName :[],
            subjectTime :{}
        }
    },
    methods: {
        getVal(selectSubject, lectureDateAndTime, dic){ //DropDownからのデータ受けとり
          this.subject = selectSubject,
          this.dateTime = lectureDateAndTime,
          this.dic = dic,
          this.dateUTC = (new Date(this.dateTime).getTime())/1000
          console.log(this.dateTime)
          if(this.dateTime != ''){
            this.button_clicked = false
          }
        },

        onClick(){ //欠席ボタン押下時の処理
            const attendNowTime = Math.floor((new Date()).getTime() / 1000);
            this.ans = this.dic[this.subject][this.dateUTC]
            this.button_clicked = true;
            // console.log(this.ans)

            if(this.ans == undefined){
              alert("入力内容が不正です。もう一度やり直してください。")
              this.$router.go(0)
            }else{
              axios.post(
                  this.BASE_URL + 'attend',
                  {
                      attend_flag : -1,
                      now_time : attendNowTime,
                      student_id : user_id,
                      subject_id : this.ans
                  }
              ).then( response => {
                  console.log(response)
                  const object = response.data
                  if(object.http_status < 400){
                    if(object.check_flag>0){
                        this.button_clicked = true;
                        user.can_back = true;
                        setTimeout(() => {this.$router.push('/absent_check')}, 500);
                    }else{
                      alert("欠席登録できませんでした。もう一度やり直してください。")
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
              }).catch(
                error=>{
                this.button_clicked = true //login不要

                console.log(error)
                switch(error.code){
                    case 'ERR_NETWORK': alert("サーバーに問題が発生しております。担当者にお問い合わせください。"); this.$router.go(0); break; //reload apiにはrouter不要
                    default: alert("エラーが発生しました。もう一度お試しください。"); this.$router.go(0); break;　//reload apiにはrouter不要
                }
            })
            }
        }
    },
    created(){
      const reloadNowTime = Math.floor((new Date()).getTime() / 1000);
      axios.post(this.BASE_URL + 'reload'
            ,{
                student_id: user_id,
                now_time : reloadNowTime,
                attend_flag: -1
            }
            ).then(
              response => {
                console.log(response)
                const object = response.data
                let http_status = 0;
                if(response.data.length<1){
                  http_status = response.data.http_status
                }else{
                  http_status = response.data[0].http_status;
                }

                if( http_status < 400 ){
                  this.responseJSON = response
                  var set = new Set();
                  for(let i=0; i<response.data.length; i++){
                    set.add(response.data[i].subject_name)
                  }
                  this.subjectName = Array.from(set);

                  for(let i=0; i<response.data.length; i++){
                    this.subjectTime[this.subjectName[i]] = []
                  }
                  for(let i=0; i<response.data.length; i++){
                    this.subjectTime[response.data[i].subject_name].push(response.data[i].subject_time)
                  }
                }else {
                  switch(http_status){
                    case 400: alert("URLに誤りがあります。URLを確かめた上でもう一度入り直してください。"); break;
                    case 401: alert("閲覧権限がありません。URLを確かめた上でもう一度入り直してください。"); break;
                    case 404: alert("該当するページが見つかりません。URLに誤りがないか確認してください。"); break;
                    case 429: alert("サーバーへのアクセスが集中しています。少し時間を空けてください。"); break;
                    case 503: alert("一時的にサービスが利用できません。少し時間を空けてください。"); break;
                    case 504: alert("正しく処理が行われませんでした。もう一度お試しください。"); break;
                    default: alert("エラーが発生しました。もう一度お試しください。"); break;
                  }
                }
              }).catch(error=>{
                this.button_clicked = true//login不要

                console.log(error)
                switch(error.code){
                    case 'ERR_NETWORK': alert("サーバーに問題が発生しております。担当者にお問い合わせください。"); break; //reload apiにはrouter不要
                    default: alert("エラーが発生しました。もう一度お試しください。"); break; //reload apiにはrouter不要
                }
            })
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
      <Header :title="title" />
      <div class="absent_dropdown">
        <DropDown :subjectName="subjectName" :subjectTime="subjectTime" :responseJSON="responseJSON" @onChange="getVal"/>
      </div>
      <div>
        <AbsButton @handleAbsent="onClick" :button_clicked="button_clicked"/>
      </div>
    <Footer />
</html>
</template>

<style>
.absent_dropdown{
  padding-top: 90px;
  padding-bottom: 10px;
}
</style>