<template>
    <div class="dropdown-container">
        <div class="dropdown">
        <select v-model="selectSubject" @change="bothFanction">
            <option value="" selected>講義名</option>
            <option v-for="subject_name in subjectName" :key="subject_name" >{{ subject_name }}</option>
            </select>
        </div>
        <div class="dropdown">
        <select v-model="lectureDateAndTime" @change="onEmit">
            <option v-for="subject_time in dateTime" :key="subject_time" >{{ subject_time }}</option>
        </select>
        </div>
    </div>
</template>

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
</style>

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
        }
    },

    data (){
        return {
            selectSubject:'',
            lecture_date_and_time:'',
            key:'',
            dateTime:[],
            subTimeKey:[],
            subTimeValues:[]
        }
    },
    methods: {
        bothFanction(){
            this.onEmit(),
            this.arrayKey()
        },

        onEmit(){
            this.$emit("onChange",this.selectSubject,this.lectureDateAndTime)
        },
//subjectNameのkeyに対応したsubjectTimeのみを表示する
        arrayKey(){
            this.key = this.selectSubject;
            this.subTimeKey = Object.keys(this.subjectTime)
            this.subTimeValues = Object.values(this.subjectTime)
            console.log(this.key,this.subTimeKey)
            for(let i=0; i<this.subjectName.length; i++){
                if(this.key == this.subTimeKey[i]){
                    var date = this.subTimeValues[i]
                    this.dateTime[i] = new Date(date*1000).toLocaleString() ;
                }
            }
        }
    }
}
</script>
