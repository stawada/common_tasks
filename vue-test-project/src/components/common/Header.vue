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

<style></style>
