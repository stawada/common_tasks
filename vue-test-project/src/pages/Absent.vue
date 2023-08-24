<script>
import DropDown from '../components/DropDown.vue'
import AbsButton from '../components/AbsButton.vue'
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'
import axios from 'axios'
import { useUserStore } from '../stores/user'

const user = useUserStore();
const user_id = user.user_id;
const nowTime = Math.floor((new Date()).getTime() / 1000);

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
            this.ans = this.dic[this.subject][this.dateUTC]
            this.button_clicked = true;
            // console.log(this.ans)

            if(this.ans == undefined){
              alert("入力内容が不正です。もう一度やり直して下さい。")
              this.$router.go(0)
            }else{
              axios.post(
                  this.BASE_URL + 'attend',
                  {
                      attend_flag : -1,
                      now_time : nowTime,
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
                    case 400: alert("BadRequest");  this.$router.go(0); break;
                    case 401: alert("Unauthorized");  this.$router.go(0); break;
                    case 404: alert("Not found");  this.$router.go(0); break;
                    case 412: alert("Precondition Failed");  this.$router.go(0); break;
                    case 429: alert("To Many Requests");  this.$router.go(0); break;
                    case 503: alert("Service Unavailable");  this.$router.go(0); break;
                    case 504: alert("Gateway Timeout");  this.$router.go(0); break;
                    default: alert("Unknown Error"); this.$router.go(0); break;
                  }
                }
              }).catch(
                  error=>{
                      this.button_clicked = true
                      console.log(error)
                      alert(error.message)
                  })
            }
        }
    },
    created(){
      axios.post(this.BASE_URL + 'reload'
            ,{
                student_id: user_id,
                now_time : nowTime,
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
                    case 400: alert("BadRequest");  this.$router.go(0); break;
                    case 401: alert("Unauthorized");  this.$router.go(0); break;
                    case 404: alert("Not found");  this.$router.go(0); break;
                    case 412: alert("Precondition Failed");  this.$router.go(0); break;
                    case 429: alert("To Many Requests");  this.$router.go(0); break;
                    case 503: alert("Service Unavailable");  this.$router.go(0); break;
                    case 504: alert("Gateway Timeout");  this.$router.go(0); break;
                    default: alert("Unknown Error"); this.$router.go(0); break;
                  }
                }
              }).catch(error=>{
                    this.button_clicked = true
                    console.log(error)
                    alert(error.message)
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