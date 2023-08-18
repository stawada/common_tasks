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
  setup(){
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

        getVal(selectSubject, lectureDateAndTime, dic){
          this.subject = selectSubject,
          this.dateTime = lectureDateAndTime,
          this.dic = dic,
          this.dateUTC = (new Date(this.dateTime).getTime())/1000,
          console.log(this.subject,this.dateUTC,this.dic);
          //this.subject_id = this.responseJSON.find((this.subject,this.dateUTC))
        },

        onClick(){
          this.ans = this.dic[this.subject][this.dateUTC]
          console.log('検索結果:'+this.ans)
          // this.ans = this.ans[this.dateUTC]
          // console.log(this.ans)
            this.button_clicked = true;
            const nowTime = Math.floor((new Date()).getTime() / 1000)
            axios.post(
                this.BASE_URL + 'attend',
                {
                    attend_flag : -1,
                    now_time : nowTime,      //!!!!!!!!!!nowTime
                    student_id : "G000000000",  //!!!!!!!!!!user_id
                    subject_id : this.ans
                }
            ).then( response => {
                console.log(response)
                const object = response.data
                if(object.http_status < 400){
                    if(object.check_flag>0){
                        setTimeout(() => {this.$router.push('/absent_check')}, 500);
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
                    this.$router.go(0)
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
    },

    created(){
      axios.post(this.BASE_URL + 'reload' //?????
            ,{
                student_id: 'G000000000',  //!!!!!저장소에서 학생아이디 꺼내기
                //now_time : nowTime
            }
            ).then(
              response => {
                this.responseJSON = response
                console.log(this.responseJSON)
                var set = new Set();
                for(let i=0; i<response.data.length; i++){
                  set.add(response.data[i].subject_name)
                }
                this.subjectName = Array.from(set);
                //console.log(this.subjectName);

                for(let i=0; i<response.data.length; i++){
                  // console.log(this.subjectTime[i][response.data[i].subject_name] , response.data[i].subject_time)
                  // this.subjectTime[response.data[i].subject_name] = response.data[i].subject_time
                  this.subjectTime[this.subjectName[i]] = []
                }
                console.log(this.subjectTime)
                for(let i=0; i<response.data.length; i++){
                  this.subjectTime[response.data[i].subject_name].push(response.data[i].subject_time)
                }
                console.log(this.subjectTime)
            },
            //response => this.subjectName = response.data,
            //         if(response.status>0){
            //             this.can_attend = true;
            //             this.can_attend_subject = "math"
            //             this.subject_id = '111111'
            //         }else{
            //             this.can_attend = false
            //         }
            ).catch(console => console.log('error'))

    }

    // console.log(this.selectSubject,this.lectureDateAndTime)
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
      <div>
        <DropDown :subjectName="subjectName" :subjectTime="subjectTime" :responseJSON="responseJSON" @onChange="getVal"/>
      </div>
      <div>
        <AbsButton @handleAbsent="onClick"/>
      </div>
      <div>
        <!-- <p>{{ this.subjectTime }}</p> -->
      </div>
    </body>
    <Footer />
</html>
</template>
