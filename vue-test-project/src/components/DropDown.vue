<template>
    <div class="dropdown-container">
        <div class="dropdown">
        <select v-model="selectSubject" @change="onEmit">
            <option value="" selected>講義名</option>
            <option v-for="subject_name in subjectName" :key="subject_name" >{{ subject_name }}</option>
            </select>
        </div>
        <div class="dropdown">
        <select v-model="lecture_date_and_time">
            <option v-for="subject_time in subjectTime" :key="subject_time" >{{ subject_time }}</option>
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
    align-items: center;
}

.dropdown {
    margin: 10px;
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
        }
    },

    data (){
        return {
            selectSubject:'',
            subject_name:[],
        }
    },
    methods: {
        onEmit(){
            this.$emit("onChange")
        },

    }
}
</script>
