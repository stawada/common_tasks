<template>
    <div class="dropdown-container">
        <div class="headline_container">
            <label class="headline">講義選択をしてください</label>
        </div>
        <div class="dropdown">
            <select v-model="selectSubject" @change="bothFanction">
                <option value=""></option>
                <option v-for="subject_name in subjectName" :key="subject_name" >{{ subject_name }}</option>
            </select>
        </div>
        <div class="headline_container">
            <label class="headline">日付をしてください</label>
        </div>
        <div class="dropdown">
            <select v-model="lectureDateAndTime" @change="onEmit">
                <option value=""></option>
                <option v-for="subject_time in dateTime" :key="subject_time" >{{ subject_time }}</option>
            </select>
        </div>
    </div>
</template>

<script>

export default {
    props:{
        subjectName: {
            type: Object
        },
        subjectTime: {
            type: Object
        },
        responseJSON: {
            type: Object
        }
    },

    data (){
        return {
            selectSubject:'',
            lectureDateAndTime:'',
            key:'',
            dateTime:[],
            subTimeKey:[],
            subTimeValues:[],
            date:[],
            dic:{},
        }
    },


    methods: {
        bothFanction(){
            this.onEmit(),
            this.arrayKey()
        },

        onEmit(){
            this.$emit("onChange",this.selectSubject,this.lectureDateAndTime,this.dic)
        },

        arrayKey(){
            this.dateTime = []
            this.date = []
            this.key = this.selectSubject;
            this.subTimeKey = Object.keys(this.subjectTime)
            this.subTimeValues = Object.values(this.subjectTime)

            for(let i=0; i<this.subjectName.length; i++){
                if(this.key == this.subTimeKey[i]){
                    this.date.push(this.subTimeValues[i])
                }
            }
            this.date[0].sort();
            for(let i=0; i<this.date[0].length; i++){
                var newDate = this.date[0][i]
                this.dateTime.push(new Date(newDate*1000).toLocaleString())
            }

            //科目名と科目時間から一意に定まる辞書型配列
            for(let i=0;i<this.subTimeKey.length;i++){
                this.dic[this.subTimeKey[i]] = {};
            }
            for(let i=0;i<this.responseJSON.data.length;i++){
                var subjectname = this.responseJSON.data[i].subject_name;
                var subjectid = this.responseJSON.data[i].subject_id;
                var subjecttime = this.responseJSON.data[i].subject_time;
                this.dic[subjectname][subjecttime] = subjectid;
            }
        }
    }
}
</script>

<style>
select{
  -webkit-appearance: none;
  appearance: none; /* デフォルトの矢印を非表示 */
  background: #D9D9D9;
  border: 0;
  width: 304px;
  height: 62px;

  font-size: 17px;
  font-weight: 550;
  color: #494848;
  text-align: center;

  background-image: url("@/assets/pulldown.png"); /*ここに画像を入れてね */
  background-position: right 10px center;
  background-repeat: no-repeat;
  background-size: 34px 29px;
}

.headline {
  font-size: 20px;
  color: #494848;
  font-weight: 600;
}
.headline_container{
  width: 304px;
  padding-bottom: 5px;
  margin: 0 auto;
  text-align: left;
}

select::-ms-expand {
display: none; /* デフォルトの矢印を非表示(IE用) */
}

.dropdown {
  padding-bottom: 30px;
}

</style>

