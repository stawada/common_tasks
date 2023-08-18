<template>
    <div class="dropdown-container">
        <label class="headline">講義選択をしてください</label>
        <div class="dropdown">
        <select v-model="selectSubject" @change="bothFanction">
            <option value=""></option>
            <option v-for="subject_name in subjectName" :key="subject_name" >{{ subject_name }}</option>
            </select>
        </div>
        <label class="headline">日付をしてください</label>
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
            type: Array,
            default (){
                return []
            }
        },
        subjectTime: {
            type: Array,
            default (){
                return []
            }
        },
        responseJSON: {
            type: Array,
            default (){
                return []
            }
        },
    },

    data (){
        return {
            selectSubject:'',
            lecture_date_and_time:'',
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
            console.log(this.selectSubject,this.lectureDateAndTime,this.dic)
            this.$emit("onChange",this.selectSubject,this.lectureDateAndTime,this.dic)
        },

        arrayKey(){
            this.dateTime = []
            this.date = []
            this.key = this.selectSubject;
            this.subTimeKey = Object.keys(this.subjectTime)
            this.subTimeValues = Object.values(this.subjectTime)
            console.log(this.subTimeKey,this.subTimeValues)

            for(let i=0; i<this.subjectName.length; i++){
                if(this.key == this.subTimeKey[i]){
                    this.date.push(this.subTimeValues[i])
                    console.log(this.date)
                    //this.dateTime.push(new Date(this.date*1000).toLocaleString());
                    // console.log(this.dateTime)
                }
            }
            this.date[0].sort();
            console.log(this.date);
            for(let i=0; i<this.date[0].length; i++){
                var newDate = this.date[0][i]
                console.log(newDate)
                this.dateTime.push(new Date(newDate*1000).toLocaleString())
            }

            //科目名と科目時間から一意に定まる辞書型配列
            for(let i=0;i<this.subTimeKey.length;i++){
                this.dic[this.subTimeKey[i]] = {};
            }
            console.log(this.dic)
            console.log(this.responseJSON)
            for(let i=0;i<this.responseJSON.data.length;i++){
                var subjectname = this.responseJSON.data[i].subject_name;
                var subjectid = this.responseJSON.data[i].subject_id;
                var subjecttime = this.responseJSON.data[i].subject_time;
                this.dic[subjectname][subjecttime] = subjectid;
            console.log(this.dic)
            }
        }
    }
}
</script>

<style>
.dropdown-container {
    width: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    /* align-items: center; */

}

.headline {
    margin-top: 5%;
    text-align: left;
    margin-left:13%;
}
select{
    -webkit-appearance: none;
    appearance: none; /* デフォルトの矢印を非表示 */
    background: #D9D9D9;
    width: 80%;
    font-size: 30px;
    margin-bottom: 5%;
    background-image: url("@/assets/check.png"); /*ここに画像を入れてね */
    background-position: right 10px center;
    background-repeat: no-repeat;
    background-size: 13px 13px;
}

select::-ms-expand {
  display: none; /* デフォルトの矢印を非表示(IE用) */
}

.dropdown {
    color: #D9D9D9;
    margin: 3%;
}
</style>