<script>
import DropDown from '../components/DropDown.vue'
import AbsButton from '../components/AbsButton.vue'
import Header from '../components/common/Header.vue'
import Footer from '../components/common/Footer.vue'
import axios from 'axios'

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
            button_clicked : false,
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
        },

        onClick(){ //欠席ボタン押下時の処理
          this.ans = this.dic[this.subject][this.dateUTC]
            this.button_clicked = true;
            const nowTime = Math.floor((new Date()).getTime() / 1000)

            axios.post(
                this.BASE_URL + 'attend',
                {
                    attend_flag : -1,
                    now_time : nowTime,      //!!!!!!!!!!nowTime
                    student_id : "G000000000",  //!!!!!!!!!!user_id ←保持しているuser_idの表記方法が決まり次第修正
                    subject_id : this.ans
                }
            ).then( response => {
                console.log(response)
                const object = response.data
                if(object.http_status < 400){
                  if(object.check_flag>-1){
                      this.button_clicked = true
                      setTimeout(() => {this.$router.push('/absent_check')}, 500);
                }else{
                  if(object.http_status == 400){
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
                }
              }
            }).catch(
                error=>{
                    this.button_clicked = true
                    console.log(error.response)
                    const status = error.response.status
                    const message = error.response.data.message
                    alert(status + "ERROR : " + message)
                })
        }
    },

    created(){
      axios.post(this.BASE_URL + 'reload'
            ,{
                student_id: 'G000000000',
                //now_time : nowTime //AttendのnowTimeを引っ張ってくる
            }
            ).then(
              response => {
                const object = response.data
                if(object.http_status < 400){
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
                }else{
                  if(object.http_status == 400){
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
                }
              }
            ).catch(error=>{
                    this.button_clicked = true
                    console.log(error.response)
                    const status = error.response.status
                    const message = error.response.data.message
                    alert(status + "ERROR : " + message)})
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
    <body>
      <div class="absent_dropdown">
        <DropDown :subjectName="subjectName" :subjectTime="subjectTime" :responseJSON="responseJSON" @onChange="getVal"/>
      </div>
      <div>
        <AbsButton @handleAbsent="onClick"/>
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
