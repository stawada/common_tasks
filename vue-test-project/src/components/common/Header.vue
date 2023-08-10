<template>
    <header class="Header">
        <p>{{title}}</p>
    </header>
</template>

<script>
let str = '';
let url = window.location.href;
//console.log(url);

export default {
    props:{
        title:String
    },
}

</script>

<style>
.Header{
  background-color: #F3AF2B;
  height: 150px;
}
.Header p{
  color: white;
  font-size: 45px;
  text-align: center;
  font-weight: 550;
  padding: 20% 0 0 0;
}

/* スマホ以外の場合、適用されるCSS */
@media(min-width: 576px){
  .Header p{
  }

}
</style>
