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
            can_attend: true,
            subject:'数学I',
            subject_id: '',
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

    },

    created(){
      axios.post(this.BASE_URL + 'reload' //?????
            ,{
                student_id: 'G000000000',  //!!!!!저장소에서 학생아이디 꺼내기
                //now_time : nowTime
            }
            ).then(
              response => {
                console.log(response)
                for(let i=0; i<response.data.length; i++){
                  console.log(i),
                this.subjectName[i] = response.data[i].subject_name
                }

                for(let i=0; i<response.data.length; i++){
                  // console.log(this.subjectTime[i][response.data[i].subject_name] , response.data[i].subject_time)
                  // this.subjectTime[response.data[i].subject_name] = response.data[i].subject_time
                  this.subjectTime[this.subjectName[i]] = []
                }
                for(let i=0; i<response.data.length; i++){
                  this.subjectTime[response.data[i].subject_name].push(response.data[i].subject_time)
                }
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
        <DropDown :subjectName="subjectName" :subjectTime="subjectTime"/>
      </div>
      <div>
        <AbsButton />
      </div>
      <div>
        <!-- <p>{{ this.subjectTime }}</p> -->
      </div>
    </body>
    <Footer />
</html>
</template>
