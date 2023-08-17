<!--
<template>
    <div class="dropdown" style="width:100%; display:flex; justify-content: center; align-items: center;">
        <div>
            <select>
            <option value="math">数学</option>
            <option value="japanese">国語</option>
            <option value="english">英語</option>
        </select>
        </div>
        <div>
            <select>
            <option value="option1">7/20 11:00</option>
            <option value="option2">7/20 14:20</option>
            <option value="option3">7/22 13:30</option>
        </select>
        </div>
    </div>
</template>
-->

<template>
    <div class="dropdown-container">
        <label class="headline">講義を選択してください</label>
        <div class="dropdown">
        <select>
            <option value=""></option>
            <option value="math">数学</option>
            <option value="japanese">国語</option>
            <option value="english">英語</option>
        </select>
        </div>
        <label class="headline">日付を選択してください</label>
        <div class="dropdown">
        <select>
            <option value=""></option>
            <option value="option1">7/20 11:00</option>
            <option value="option2">7/20 14:20</option>
            <option value="option3">7/22 13:30</option>
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

