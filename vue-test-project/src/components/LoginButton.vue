<template>
    <div class="LoginButton">
        <button @click='$emit("handleAbsent")'>本人確認</button>
    </div>
</template>

<script>
    export default{
        name:'LoginButton',
        methods: {
        }
    }
</script>

<style>
.LoginButton button{
    all: unset;
    height: 55px;
    width: 170px;

    background-color: #696969;
	border: 3px solid white;
	border-radius: 100px;
    box-shadow: 0 2px 7px rgba(0, 0, 0, .3);

    color: white;
    font-size: 20px;
    font-weight: 600;
}
.LoginButton button:active{
	color: #fff;
	background: #5a5a5a;
	border-bottom: 3px solid #fff;
	transform: translateY(5px);
}
</style>