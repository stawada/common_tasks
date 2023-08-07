<template>
    <form name="login">
        <div class="textfield-container">
            <div class="id">
                <input type="id" name="student_id" placeholder="学生番号">
            </div>
            <div class="password">
                <input type="password" name="password" placeholder="パスワード">
            </div>
        </div>
    </form>
</template>

<style>

form {
    /*
    width: 400px;
    padding: 40px 10px;
    */
    margin: center;
}

form input[type ="id"],form input[type ="password"]{
border: 0;
background: none;
display: block;
margin: 20px auto;
border: 2px solid #C0C0C0;
padding: 15px 10px;
width: 200px;
outline: none;
border-radius: 0.25px;
transition: 0.25s;
text-align: center;
}
</style>